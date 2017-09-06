package main

import (
	_ "port-info/web"
	"port-info/web"
	"net/http"
	"github.com/shurcooL/vfsgen"
	"log"
)

func main() {
	var fs http.FileSystem = http.Dir("assets")

	err := vfsgen.Generate(fs, vfsgen.Options{})
	if err != nil {
		log.Fatalln(err)
	}
	web.InitConfig().Run(":5555")
}
