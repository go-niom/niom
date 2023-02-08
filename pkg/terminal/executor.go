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
)

var (
	cmd           *exec.Cmd
	IsCodeUpdated bool
	isUserExit    bool
)

func KillFunc() {
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err == nil {
		syscall.Kill(-pgid, 15) // note the minus sign
	}
}

func CmdExecute(dir, app string, args []string) string {

	//to user cancel event cmd+c
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		KillFunc()
		isUserExit = true
	}()

	//check command has args
	if len(args) > 1 {
		cmd = exec.Command(app, args...)
	} else {
		cmd = exec.Command(app)
	}
	cmd.Dir = dir
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		log.Printf("Error executing command: %s......\n", err.Error())
		return err.Error()
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
		if (err.Error() == "signal: terminated") && IsCodeUpdated {
			log.Println("Server Restarting ", IsCodeUpdated)
			IsCodeUpdated = false
			CmdExecute("myapp", "go", []string{"run", "main.go"})
		} else if (err.Error() == "signal: terminated") && isUserExit {
			log.Println("Server Stopped")
			isUserExit = false
		} else {
			log.Printf("Error waiting for command execution: %s......\n", err.Error())
			return err.Error()
		}
	}
	return ""
}
