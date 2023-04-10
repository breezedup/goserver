package signal

import (
	"os"

	"github.com/breezedup/goserver/core/logger"
	"github.com/breezedup/goserver/core/module"
)

type KillSignalHandler struct {
}

func (ish *KillSignalHandler) Process(s os.Signal, ud interface{}) error {
	logger.Logger.Warn("Receive Kill signal, process be close")
	module.Stop()
	return nil
}

func init() {
	SignalHandlerModule.RegisteHandler(os.Kill, &KillSignalHandler{}, nil)
}
