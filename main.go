package main

import (
	"port-info/web"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	web.InitConfig().Run(":5555")
}
