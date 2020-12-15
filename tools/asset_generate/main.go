package main

import (
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
	"os"
)

func main() {
	var fs http.FileSystem = http.Dir("assets")
	_ = os.MkdirAll("internal/pkg/templates", os.ModeDir)
	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:     "internal/pkg/templates/templates.go",
		PackageName:  "templates",
		VariableName: "Templates",
	})
	if err != nil {
		log.Fatal(err)
	}
}
