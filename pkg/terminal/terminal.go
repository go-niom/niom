package terminal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
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

func KillFunc() error {
	PID, err := syscall.Getpgid(cmd.Process.Pid)
	if err == nil {
		logger.Warn("Killing App runner")
		if errKill := syscall.Kill(-PID, 15); errKill != nil {
			return errKill
		}
	}
	if PID < 0 && !isUserExit {
		go TermCmd.Execute()
	}
	return nil
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

// CMD Exit check
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

func (term *TerminalCmd) Execute() error {
	//check command has args
	if len(term.Args) > 0 {
		cmd = exec.Command(term.App, term.Args...)
	} else {
		cmd = exec.Command(term.App)
	}

	if cmd != nil {
		cmd.Dir = term.Dir
	}
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		log.Printf("Error executing command: %s\n", err.Error())
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		io.Copy(os.Stdout, stdout)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		io.Copy(os.Stderr, stderr)
		reader := bufio.NewReader(stdout)
		line, err := reader.ReadString('\n')
		for err == nil {
			fmt.Println(line)
			line, err = reader.ReadString('\n')
		}
	}()

	wg.Wait()
	if err := cmd.Wait(); err != nil {
		return term.cmdErrors(err)
	}
	if term.ShowMessage {
		logger.Info("Command Executed Successfully")
	}
	return nil
}

func CmdExecute(dir, app string, args []string, showMessage bool) {
	initSignal()
	TermCmd = TerminalCmd{
		Dir: dir, App: app, Args: args,
		ShowMessage: showMessage,
	}
	TermCmd.Execute()
}
