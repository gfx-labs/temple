package main

//go:generate go run .
//go:generate cat ./maps/maps.go
import (
	"log"

	"gfx.cafe/util/temple"
	"gfx.cafe/util/temple/lib/sanctum"
)

func main() {
	temple.RegisterTemplateDir(".")
	temple.Prepare(&sanctum.GoPrayer{
		Input: "mapn",
		Obj: map[string]any{
			"Count": 10,
		},
		Args:   nil,
		Output: "./maps/maps.go",
	})
	temple.Prepare(&sanctum.GoPrayer{
		Input:  "sync_map",
		Obj:    nil,
		Args:   nil,
		Output: "./maps/sync_map.go",
	})
	err := temple.Pray()
	if err != nil {
		log.Println(err)
	}
}
