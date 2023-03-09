package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func DoCmd(cmdName string, args []string) ([]byte, error) {
	fmt.Printf("DoCmd:%s with %v\n", cmdName, args)

	cmd := exec.Command(cmdName, args...)

	var stdoutBuf bytes.Buffer //TODO: check ram usage
	cmd.Stdout = &stdoutBuf

	e := cmd.Run()
	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	return stdoutBuf.Bytes(), nil
}

func StartProcess(processName string, args []string) {
	os.StartProcess(processName, args, nil)
}
