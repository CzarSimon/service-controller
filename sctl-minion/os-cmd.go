package main

import (
	"bytes"
	"os/exec"
)

//RunCommand executes a command against the host shell
func RunCommand(command string) (string, error) {
	cmd := exec.Command(command)
	var out bytes.Buffer
	cmd.Stdout = &out
	var errOut bytes.Buffer
	cmd.Stderr = &errOut
	err := cmd.Run()
	if err != nil {
		return errOut.String(), err
	}
	return out.String(), nil
}
