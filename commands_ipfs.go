package main

import (
	"strings"
)

func IpfsPinLs() (PinSet, error) {
	cmd := "ipfs"
	args := []string{"pin", "ls", "--type", "recursive"}

	d, e := DoCmd(cmd, args)

	if e != nil {
		return nil, e
	}

	var pinSet PinSet
	lines := bytes2Lines(d)
	for line := range lines {
		columns := strings.Split(string(line), " ")
		pinSet = append(pinSet, columns[0])
	}

	return pinSet, nil
}
