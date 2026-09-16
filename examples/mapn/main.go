package main

//go:generate go run .
//go:generate cat ./maps/maps.go
import (
	"github.com/gfx-labs/temple/lib/prayer"
	"log"

	"github.com/gfx-labs/temple"
)

func main() {
	temple.RegisterTemplateDir(".")
	temple.Prepare(&prayer.Go{
		Input: "mapn",
		Obj: map[string]any{
			"Count": 10,
		},
		Args:   nil,
		Output: "./maps/maps.go",
	})
	temple.Prepare(&prayer.Go{
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
