package main

//go:generate go run .
import (
	"log"

	"gfx.cafe/util/temple"
	"gfx.cafe/util/temple/lib/sanctum"
)

func main() {
	temple.RegisterTemplateFile("mapn.tmpl")
	temple.Prepare(&sanctum.Prayer{
		Input: "mapn",
		Obj: map[string]any{
			"Count": 3,
		},
		Args:        nil,
		PackagePath: "maps",
		PackageName: "maps",
		FileName:    "map2.go",
	})
	err := temple.Pray()
	if err != nil {
		log.Println(err)
	}
}
