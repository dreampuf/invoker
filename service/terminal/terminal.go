package terminal

import (
	"context"
	"github.com/dreampuf/invoker/service/log"
	"github.com/kr/pty"
	"io"
	"os/exec"
)

type Remote struct {

}

func NewRemote() *Remote {
	return &Remote{

	}
}

func Connect() (io.ReadWriteCloser,context.CancelFunc) {
	cmdCtx, cancel := context.WithCancel(context.TODO())
	cmd := exec.CommandContext(cmdCtx, "bash")
	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.WithError(err).Error("starting pty failure")
	}
	if err := cmd.Start(); err != nil {
		log.WithError(err).Error()
	}
	go func() {
		if err := cmd.Wait(); err != nil && err.Error() != "signal: killed" {
			log.WithError(err).Error()
		}
		ptmx.Close()
	}()
	return ptmx, cancel
}
