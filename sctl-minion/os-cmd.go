package main

import (
	"bytes"
	"os/exec"
	"strings"
)

//RunCommand executes a command against the host shell
func RunCommand(command string) (string, error) {
	parts := strings.Fields(command)
	command = parts[0]
	args := parts[1:len(parts)]
	cmd := exec.Command(command, args...)
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
