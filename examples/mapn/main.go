package main

//go:generate go run .
//go:generate cat ./maps/maps.go
import (
	"log"

	"gfx.cafe/util/temple"
	"gfx.cafe/util/temple/lib/sanctum"
)

func main() {
	temple.RegisterTemplateFile("mapn.tmpl")
	temple.RegisterTemplateFile("typedef.tmpl")
	temple.RegisterTemplateFile("sync_map.gotmpl")
	temple.Prepare(&sanctum.Prayer{
		Input: "mapn",
		Obj: map[string]any{
			"Count": 2,
		},
		Args:        nil,
		PackagePath: "./maps",
		PackageName: "maps",
		FileName:    "maps.go",
	})
	temple.Prepare(&sanctum.Prayer{
		Input:       "sync_map",
		Obj:         nil,
		Args:        nil,
		PackagePath: "./maps",
		PackageName: "maps",
		FileName:    "sync_map.go",
	})
	err := temple.Pray()
	if err != nil {
		log.Println(err)
	}
}
