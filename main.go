package main

import "fmt"

func main() {
	state, e := ClusterStateExport()
	if e != nil {
		fmt.Printf("ClusterStateExport err:%s\n", e.Error())
		return
	}

	pinset, e := IpfsPinLs()
	if e != nil {
		fmt.Printf("IpfsPinLs err:%s\n", e.Error())
		return
	}

	fmt.Println("cluster state:")
	for _, pin := range state {
		fmt.Printf("%s:%v\n", pin.Cid, pin.Allocations)
	}

	fmt.Println("ipfs pinset:")
	for _, cid := range pinset {
		fmt.Println(cid)
	}
}
