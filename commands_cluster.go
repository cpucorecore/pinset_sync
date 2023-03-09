package main

import (
	"encoding/json"
)

func ClusterStateExport() (ClusterState, error) {
	cmd := "ipfs-cluster-service"
	args := []string{"state", "export"}

	d, e := DoCmd(cmd, args)
	if e != nil {
		return nil, e
	}

	var state ClusterState

	lines := bytes2Lines(d)
	for line := range lines {
		var pin ClusterPin

		e = json.Unmarshal(line, &pin)
		if e != nil {
			// TODO: log error
			// TODO: continue or return err?
			continue
		}

		state = append(state, pin)
	}

	return state, nil
}
