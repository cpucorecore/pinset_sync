package main

import (
	"testing"
)

func TestIpfsPinLs(t *testing.T) {
	ps, e := IpfsPinLs()
	if e != nil {
		t.Error(e)
	}
	t.Log(ps)
}

func TestClusterStateExport(t *testing.T) {
	state, e := ClusterStateExport()

	if e != nil {
		t.Error(e)
	}

	t.Log(state)
}
