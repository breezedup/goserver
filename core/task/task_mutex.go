package task

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/breezedup/goserver/core/basic"
)

var taskMutexLock sync.Mutex
var taskMutexPool = make(map[string]Task)
var ErrTaskIsRunning = errors.New("mutex task is running")

// 互斥任务，相同key的任务，只有一个CompleteNotify，后边再触发的自动忽略，例如：客户端多次点击导致的请求，只有第一次给反馈
type mutexTask struct {
	*baseTask
	running  int32  //是否正在运行
	mutexKey string //互斥任务key
}

func NewMutexTask(s *basic.Object, c Callable, n CompleteNotify, key, name string) (t Task, done bool) {
	mutexKey := name + key
	taskMutexLock.Lock()
	if t, ok := taskMutexPool[mutexKey]; ok {
		taskMutexLock.Unlock()
		return t, true
	}

	base := newBaseTask(s, c, n, name)
	t = &mutexTask{
		baseTask: base,
		mutexKey: mutexKey,
	}
	base.imp = t
	taskMutexPool[mutexKey] = t
	taskMutexLock.Unlock()
	return t, false
}

// 不支持
func (t *mutexTask) clone(name string) Task {
	return nil
}

func (t *mutexTask) run(o *basic.Object) (e error) {
	// process mutex task
	if !atomic.CompareAndSwapInt32(&t.running, 0, 1) {
		return ErrTaskIsRunning
	}

	e = t.baseTask.run(o)

	taskMutexLock.Lock()
	delete(taskMutexPool, t.mutexKey)
	taskMutexLock.Unlock()

	return nil
}
