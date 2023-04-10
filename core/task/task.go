package task

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/breezedup/goserver/core"
	"github.com/breezedup/goserver/core/basic"
	"github.com/breezedup/goserver/core/container"
	"github.com/breezedup/goserver/core/container/recycler"
	"github.com/breezedup/goserver/core/logger"
	"github.com/breezedup/goserver/core/profile"
)

type Callable interface {
	Call(*basic.Object) interface{}
}

type CompleteNotify interface {
	Done(interface{}, Task)
}

type Task interface {
	AddRefCnt(cnt int32) int32
	GetRefCnt() int32
	Get() interface{}
	GetWithTimeout(timeout time.Duration) interface{}
	GetEnv(k interface{}) interface{}
	PutEnv(k, v interface{}) bool
	SetAlertTime(alertt time.Duration)
	GetCostTime() time.Duration
	GetRunTime() time.Duration
	Start()
	StartByExecutor(name string) bool
	StartByFixExecutor(name string) bool
	BroadcastToAllExecutor() bool
	StartByGroupExecutor(gname string, name string) bool
	StartByGroupFixExecutor(name, gname string) bool
	//inner
	clone(name string) Task
	run(o *basic.Object) (e error)
	done(n CompleteNotify)
	sendRsp()
	setAfterQueCnt(n int)
	setBeforeQueCnt(n int)
	getS() *basic.Object
	getC() Callable
	getN() CompleteNotify
}

type CallableWrapper func(o *basic.Object) interface{}

func (cw CallableWrapper) Call(o *basic.Object) interface{} {
	return cw(o)
}

type CompleteNotifyWrapper func(interface{}, Task)

func (cnw CompleteNotifyWrapper) Done(i interface{}, t Task) {
	cnw(i, t)
}

type baseTask struct {
	imp          Task
	s            *basic.Object
	c            Callable
	n            CompleteNotify
	r            chan interface{}
	v            interface{}
	env          *container.SynchronizedMap
	tCreate      time.Time
	tStart       time.Time
	alertTime    time.Duration
	name         string
	refTaskCnt   int32
	beforeQueCnt int //入队列前，等待中的任务数量
	afterQueCnt  int //出队列后，等待中的任务数量
}

func New(s *basic.Object, c Callable, n CompleteNotify, name ...string) Task {
	return newBaseTask(s, c, n, name...)
}

func newBaseTask(s *basic.Object, c Callable, n CompleteNotify, name ...string) *baseTask {
	t := &baseTask{
		s:       s,
		c:       c,
		n:       n,
		r:       make(chan interface{}, 1),
		tCreate: time.Now(),
	}
	if len(name) != 0 {
		t.name = name[0]
	}
	if s == nil {
		t.s = core.CoreObject()
	}
	t.imp = t
	return t
}

func (t *baseTask) clone(name string) Task {
	fullname := t.name
	if name != "" {
		fullname += "-" + name
	}
	return New(t.s, t.c, t.n, fullname)
}

func (t *baseTask) setAfterQueCnt(n int) {
	t.afterQueCnt = n
}

func (t *baseTask) setBeforeQueCnt(n int) {
	t.beforeQueCnt = n
}

func (t *baseTask) getS() *basic.Object {
	return t.s
}

func (t *baseTask) getC() Callable {
	return t.c
}

func (t *baseTask) getN() CompleteNotify {
	return t.n
}

func (t *baseTask) AddRefCnt(cnt int32) int32 {
	return atomic.AddInt32(&t.refTaskCnt, cnt)
}

func (t *baseTask) GetRefCnt() int32 {
	return atomic.LoadInt32(&t.refTaskCnt)
}

func (t *baseTask) Get() interface{} {
	if t.n != nil {
		panic("Task result by CompleteNotify return")
	}

	return <-t.r
}

func (t *baseTask) GetWithTimeout(timeout time.Duration) interface{} {
	if timeout == 0 {
		return t.Get()
	} else {
		timer := recycler.GetTimer(timeout)
		defer recycler.GiveTimer(timer)
		select {
		case r, ok := <-t.r:
			if ok {
				return r
			} else {
				return nil
			}
		case <-timer.C:
			return nil
		}
	}
	return nil
}

func (t *baseTask) GetEnv(k interface{}) interface{} {
	if t.env == nil {
		return nil
	}
	return t.env.Get(k)
}

func (t *baseTask) PutEnv(k, v interface{}) bool {
	if t.env == nil {
		t.env = container.NewSynchronizedMap()
	}
	if t.env != nil {
		t.env.Set(k, v)
	}

	return true
}

func (t *baseTask) run(o *basic.Object) (e error) {
	watch := profile.TimeStatisticMgr.WatchStart(fmt.Sprintf("/task/%v/run", t.name), profile.TIME_ELEMENT_TASK)
	defer func() {
		if watch != nil {
			watch.Stop()
		}

		if err := recover(); err != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			logger.Logger.Error("Task::run stack--->", string(buf[:n]))
		}
	}()

	t.tStart = time.Now()
	wait := t.tStart.Sub(t.tCreate)
	t.v = t.c.Call(o)
	dura := t.GetRunTime()

	if t.r != nil {
		t.r <- t.v
	}

	t.imp.sendRsp()

	if t.alertTime != 0 && t.name != "" {
		cost := t.GetCostTime()
		if cost > t.alertTime {
			logger.Logger.Warn("task [", t.name, "] since createTime(",
				cost, ") since startTime(", dura, "), in quene wait(", wait, ")", " beforeQueCnt(", t.beforeQueCnt, ") afterQueCnt(", t.afterQueCnt, ")")
		}
	}
	return nil
}

func (t *baseTask) done(n CompleteNotify) {
	if n != nil {
		n.Done(t.v, t)
	}
}

func (t *baseTask) sendRsp() {
	if t.n != nil {
		SendTaskRes(t.s, t, t.n)
	}
}

func (t *baseTask) Start() {
	go t.imp.run(nil)
}

func (t *baseTask) SetAlertTime(alertt time.Duration) {
	t.alertTime = alertt
}

func (t *baseTask) GetCostTime() time.Duration {
	return time.Now().Sub(t.tCreate)
}

func (t *baseTask) GetRunTime() time.Duration {
	return time.Now().Sub(t.tStart)
}

func (t *baseTask) StartByExecutor(name string) bool {
	return sendTaskReqToExecutor(t, name, "")
}

func (t *baseTask) StartByFixExecutor(name string) bool {
	return sendTaskReqToFixExecutor(t, name, "")
}

func (t *baseTask) BroadcastToAllExecutor() bool {
	return sendTaskReqToAllExecutor(t)
}

func (t *baseTask) StartByGroupExecutor(gname string, name string) bool {
	return sendTaskReqToExecutor(t, name, gname)
}

func (t *baseTask) StartByGroupFixExecutor(name, gname string) bool {
	return sendTaskReqToFixExecutor(t, name, gname)
}
