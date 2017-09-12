package main

import (
	"runtime"
	"port-forward/web"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	web.InitConfig().Run(":5555")
}
