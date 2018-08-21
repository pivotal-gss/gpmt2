package bash

import (
	"fmt"
	"os/exec"
	"bytes"
	"time"
	log "github.com/sirupsen/logrus"
)

// Wrapper to execute bash command, this will take in the below arguments.
// timeout: timeout value in ( sections ) which is the maximum time it waits before terminating the process ( 0 to disable )
// command: the bash command that it needs to execute
// args: the arguments the needs to be used along with the command.
// eg.s RunBashCmd(5, "ping", "-c25", "8.8.8.8")
// this will return the result of the output and a error if it falls into any
func RunBashCmd(timeout int, command string, args ...string) (string, error) {

	// Check if the command executable exists
	_, err :=  exec.LookPath(command)
	if err != nil {
		return "", err
	}

	// instantiate new command
	cmd := exec.Command(command, args...)

	// get pipe to standard output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "" , fmt.Errorf("cmd.StdoutPipe() error: %s", err.Error())
	}

	// start process via command
	if err := cmd.Start(); err != nil {
		return "" , fmt.Errorf("cmd.Start() error: %s", err.Error())
	}

	// now we have the command started
	// the process number for this process is
	process := cmd.Process
	pid := process.Pid
	log.Debugf("Executing the command: %s %v, Process: %d", command, args, pid)

	// Start a ticker to capture the time that it waiting
	// only good for debugging.
	ticker := time.NewTicker(time.Second)
	go func(ticker *time.Ticker) {
		now := time.Now()
		for _ = range ticker.C {
			log.Debugf("Waiting on process(%d): %s", pid, []byte(fmt.Sprintf("%s", time.Since(now))))
		}
	}(ticker)

	// setup a buffer to capture standard output
	var buf bytes.Buffer

	// create a channel to capture any errors from wait
	done := make(chan error)
	go func() {
		if _, err := buf.ReadFrom(stdout); err != nil {
			log.Panicf("buf.Read(stdout) error: %s", err.Error())
		}
		done <- cmd.Wait()
	}()

	// If timeout is disabled then wait for the command to finish.
	if timeout == 0 {
		err := <-done
		if err != nil {
			close(done)
			return "", fmt.Errorf("process(%d) done, with error: %s", pid, err.Error())
		}
		// Done, send the information back to user
		return buf.String(), nil
	} else {
		// block on select, and switch based on actions received
		select {
		// Terminate the process since the timeout has reached.
		case <-time.After(time.Duration(timeout) * time.Second):
			if err := process.Kill(); err != nil {
				return "", fmt.Errorf("failed to kill the process(%d): %s", pid, err.Error())
			}
			return "", fmt.Errorf("timeout reached, process(%d) killed", pid)
		// The command is done, send the output back to the user
		case err := <-done:
			if err != nil {
				close(done)
				return "",fmt.Errorf("process(%d) done, with error: %s", pid, err.Error())
			}
			return buf.String(), nil
		}
	}

	// Should never reach here.
	return "", nil
}