// +build dev

package main

import (
	"net/http"
	"github.com/shurcooL/vfsgen"
	"log"
)


//go:generate go run -tags=dev assets_generate.go
func main() {
	var fs http.FileSystem = http.Dir("./asset")

	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:    "web/data_go.go",
		PackageName:  "web",
		BuildTags:    "vfs",
		VariableName: "assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
