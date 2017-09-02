package main

import (
	"time"
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"port-info/util"
)

func main() {
	//a := port.StartPortForward("0.0.0.0:8700", "192.168.0.88:13000")
	//a.AddStatics(&port.Limit{})
	//port.StartPortForward("0.0.0.0:8701", "192.168.0.88:13000")
	v, _ := mem.VirtualMemory()


	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", util.GetBytes(float64(v.Total)), util.GetBytes(float64(v.Available)), v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
	for true {
		time.Sleep(time.Second)
		//fmt.Println("aaa= ", a.TotalByte.String())
		//fmt.Println("bbb=", b.TotalByte.String())
		//fmt.Println("ccc=", port.ForwardPoll)
	}
}
