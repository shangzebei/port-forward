package main

import (
	"net/http"
	"github.com/shurcooL/vfsgen"
	"log"
)

func main() {
	var fs http.FileSystem = http.Dir("assets")

	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "main",
		BuildTags:    "vfs",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
