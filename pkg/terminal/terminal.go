package terminal

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/go-niom/niom/pkg/logger"
)

var (
	cmd             *exec.Cmd
	IsCodeUpdated   bool
	isUserExit      bool
	TermCmd         TerminalCmd
	TerminalChannel = make(chan string, 1000)
)

type TerminalCmd struct {
	Dir         string
	App         string
	Args        []string
	ShowMessage bool
}

// to check user cancel event cmd+c
func initSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		defer close(c)
		sig := <-c
		if sig == os.Interrupt {
			isUserExit = true
			KillFunc()
			close(TerminalChannel)
		}
	}()
}

// cmdErrors handles the thrown by the shell
func (term *TerminalCmd) cmdErrors(err error) error {
	if (err.Error() == "signal: terminated") && IsCodeUpdated {
		logger.Info("Restarting....")
		IsCodeUpdated = false
		TermCmd.Execute()
	} else if (err.Error() == "signal: terminated") && isUserExit {
		logger.Warn("Niom App Terminated ")
		isUserExit = false
	} else {
		logger.Error("Stopped ", err.Error())
		logger.Warn("Waiting for file change....")
		return err
	}
	return err
}

// CmdExecute trigger the Execute function
// Also init the signal to check signal notification
func CmdExecute(dir, app string, args []string, showMessage bool) {
	initSignal()
	TermCmd = TerminalCmd{
		Dir: dir, App: app, Args: args,
		ShowMessage: showMessage,
	}
	TermCmd.Execute()
}
