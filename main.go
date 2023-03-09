package main

import (
	"fmt"
	"os"
)

func main() {
	//e := StopCluster()
	//if e != nil {
	//	fmt.Printf("StopCluster err:%s\n", e.Error())
	//	return
	//}
	//
	//state, e := ClusterStateExport()
	//if e != nil {
	//	fmt.Printf("ClusterStateExport err:%s\n", e.Error())
	//	return
	//}
	//
	//pinset, e := IpfsPinLs()
	//if e != nil {
	//	fmt.Printf("IpfsPinLs err:%s\n", e.Error())
	//	return
	//}
	//
	//fmt.Println("cluster state:")
	//for _, pin := range state {
	//	fmt.Printf("%s:%v\n", pin.Cid, pin.Allocations)
	//}
	//
	//fmt.Println("ipfs pinset:")
	//for _, cid := range pinset {
	//	fmt.Println(cid)
	//}

	//attr := syscall.SysProcAttr{Setpgid: true}
	proc, e := os.StartProcess("/bin/bash", []string{"./scripts/test.sh"}, &os.ProcAttr{Files: []*os.File{nil, nil, nil}})
	if e != nil {
		fmt.Printf("StartCluster err:%s\n", e.Error())
		return
	}

	proc.Release()

	//proc.Wait()
}
