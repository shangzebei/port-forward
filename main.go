package main

import (
	"port-info/port"
	"fmt"
	"time"
)

func main() {
	a := port.StartPortForward("0.0.0.0:8700", "192.168.0.88:13000")

	b := port.StartPortForward("0.0.0.0:8701", "192.168.0.88:13000")
	for true {
		time.Sleep(time.Second)
		fmt.Println("aaa= ", a.TotalByte.String())
		fmt.Println("bbb=", b.TotalByte.String())
	}
}
