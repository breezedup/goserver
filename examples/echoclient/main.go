// main
package main

import (
	"github.com/breezedup/goserver/core"
	"github.com/breezedup/goserver/core/module"
)

func main() {
	defer core.ClosePackages()
	core.LoadPackages("config.json")

	waiter := module.Start()
	waiter.Wait("main")
}
