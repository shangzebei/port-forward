package main

import (
	"port-info/web"
)

func main() {
	web.InitConfig().Run(":5555")
}
