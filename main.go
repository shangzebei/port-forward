package main

import (
	_ "port-info/web"
	"port-info/web"
)

func main() {

	web.InitConfig().Run(":5555")
}
