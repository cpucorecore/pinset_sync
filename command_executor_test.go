package main

import (
	"encoding/json"
	"os/exec"
	"testing"
)

func TestIpfsPinLs(t *testing.T) {
	args := []string{"pin", "ls", "--type", "recursive"}
	d, e := exec.Command("ipfs", args...).Output()

	if e != nil {
		t.Error(e)
	}

	t.Log(string(d))
}

func TestIpfsClusterStateExport(t *testing.T) {
	args := []string{"state", "export"}
	d, e := exec.Command("ipfs-cluster-service", args...).Output()

	if e != nil {
		t.Error(e)
	}

	var pins ClusterPins

	lines := bytes2Lines(d)
	for line := range lines {
		var pin ClusterPin
		e = json.Unmarshal(line, &pin)
		if e != nil {
			t.Error(e)
		}

		t.Log(pin)
		pins = append(pins, pin)
	}
}
