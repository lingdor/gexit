package gexit

import (
	"os"
	"os/signal"
	"syscall"
)

func signalExitNotify(ch chan os.Signal) {
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
}