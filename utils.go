package main

import (
	"bufio"
	"bytes"
)

func bytes2Lines(d []byte) chan []byte {
	scanner := bufio.NewScanner(bytes.NewReader(d))
	scanner.Split(bufio.ScanLines)

	linesChan := make(chan []byte, 1000)
	defer func() {
		close(linesChan)
	}()

	for scanner.Scan() {
		linesChan <- scanner.Bytes()
	}

	return linesChan
}
