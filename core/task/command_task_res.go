package task

import (
	"github.com/breezedup/goserver/core/basic"
	"github.com/breezedup/goserver/core/utils"
)

type taskResCommand struct {
	t Task
	n CompleteNotify
}

func (trc *taskResCommand) Done(o *basic.Object) error {
	defer o.ProcessSeqnum()
	defer utils.DumpStackIfPanic("taskExeCommand")
	trc.t.done(trc.n)
	return nil
}

func SendTaskRes(o *basic.Object, t Task, n CompleteNotify) bool {
	if o == nil {
		return false
	}
	return o.SendCommand(&taskResCommand{t: t, n: n}, true)
}
