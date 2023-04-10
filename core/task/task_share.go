package task

import (
	"github.com/breezedup/goserver/core/basic"
	"sync"
	"sync/atomic"
)

var taskShareLock sync.Mutex
var taskSharePool = make(map[string]*shareTask)

// 共享任务，多次请求共享一个Callable；返回多个CompleteNotify；例如：多个用户查询同一份榜单数据，避免缓存击穿
type shareTaskNotify struct {
	s *basic.Object
	n CompleteNotify
}

type shareTask struct {
	*baseTask
	sync.RWMutex
	notifies []*shareTaskNotify
	running  int32  //是否正在运行
	shareKey string //共享任务key
}

func RunShareTask(s *basic.Object, c Callable, n CompleteNotify, key, name string) (t Task, done bool) {
	mutexKey := name + key
	taskShareLock.Lock()
	if t, ok := taskSharePool[mutexKey]; ok {
		taskShareLock.Unlock()
		if t.v != nil {
			SendTaskRes(t.s, t, t.n)
		} else {
			t.Lock()
			t.notifies = append(t.notifies, &shareTaskNotify{s: s, n: n})
			t.Unlock()
		}
		return t, true
	}

	bt := newBaseTask(s, c, n, name)
	st := &shareTask{
		baseTask: bt,
		shareKey: mutexKey,
	}
	t = st
	bt.imp = t
	taskSharePool[mutexKey] = st
	taskShareLock.Unlock()

	go st.run(nil)

	return t, false
}

// 不支持
func (t *shareTask) clone(name string) Task {
	return nil
}

func (t *shareTask) run(o *basic.Object) (e error) {
	// process mutex task
	if !atomic.CompareAndSwapInt32(&t.running, 0, 1) {
		return ErrTaskIsRunning
	}

	e = t.baseTask.run(o)

	taskShareLock.Lock()
	delete(taskSharePool, t.shareKey)
	taskShareLock.Unlock()

	return nil
}

func (t *shareTask) sendRsp() {
	if t.n != nil {
		SendTaskRes(t.s, t, t.n)
	}

	if len(t.notifies) != 0 {
		t.RLock()
		defer t.RUnlock()
		for _, s := range t.notifies {
			SendTaskRes(s.s, t, s.n)
		}
	}
}
