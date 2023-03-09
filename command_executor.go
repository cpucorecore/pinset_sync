package main

import (
	"os/exec"
)

func DoCmd(cmd string, args []string) ([]byte, error) {
	return exec.Command(cmd, args...).Output()
}
