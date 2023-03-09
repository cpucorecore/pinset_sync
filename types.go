package main

type ClusterPin struct {
	Cid         string
	Allocations []string
}

type ClusterState []ClusterPin

type PinSet []string
