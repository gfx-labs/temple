package main

//go:generate go run .
//go:generate cat ./maps/maps.go
import (
	"go/format"
	"log"

	"gfx.cafe/util/temple"
	"gfx.cafe/util/temple/lib/sanctum"
)

func main() {
	temple.RegisterTemplateDir(".")
	temple.Prepare(&sanctum.Prayer{
		Input: "mapn",
		Obj: map[string]any{
			"Count": 10,
		},
		Args:        nil,
		Formatter:   format.Source,
		PackagePath: "./maps",
		PackageName: "maps",
		FileName:    "maps.go",
	})
	temple.Prepare(&sanctum.Prayer{
		Input:       "sync_map",
		Obj:         nil,
		Args:        nil,
		Formatter:   format.Source,
		PackagePath: "./maps",
		PackageName: "maps",
		FileName:    "sync_map.go",
	})
	err := temple.Pray()
	if err != nil {
		log.Println(err)
	}
}
