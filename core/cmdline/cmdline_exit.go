package cmdline

import (
	"fmt"

	"github.com/breezedup/goserver/core/module"
)

type exitExecuter struct {
}

func (this exitExecuter) Execute(args []string) {
	module.Stop()
}

func (this exitExecuter) ShowUsage() {
	fmt.Println("usage: exit")
}

func init() {
	RegisteCmd("exit", &exitExecuter{})
}
