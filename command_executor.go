package main

import (
	"os/exec"
)

func IpfsPinLs(cmd string, args []string) (PinSet, error) {
	//_ := []string{"pin", "ls", "--type", "recursive"}
	_, e := exec.Command(cmd, args...).Output()

	if e != nil {
		return nil, e
	}

	return nil, nil
}
