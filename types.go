package main

type ClusterPin struct {
	Cid         string
	Allocations []string
}

type ClusterPins []ClusterPin

type PinSet []string
