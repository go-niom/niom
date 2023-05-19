//go:build windows
// +build windows

package terminal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"

	"github.com/go-niom/niom/pkg/logger"
)

// KillFunc kills the shell started by the Execute
func KillFunc() error {

	if cmd.ProcessState == nil {
		logger.Warn("Killing App runner")
		if errKill := cmd.Process.Kill(); errKill != nil {
			return errKill
		}
	}

	if cmd.ProcessState != nil && !isUserExit {
		go TermCmd.Execute()
	}

	return nil
}

// Execute executes shell command
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
