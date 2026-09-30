package main

//go:generate go run .

import (
	"log"

	"github.com/gfx-labs/temple/lib/prayer"
	"github.com/gfx-labs/temple/lib/sanctum"
)

type Spec struct {
	Enums []struct {
		Name   string
		Values []string
	}
}

func main() {
	t := sanctum.New(".")
	if err := t.RegisterTemplateDir("."); err != nil {
		log.Fatal(err)
	}
	// custom func available inside templates
	t.RegisterFunc("quote", func(s string) string { return `"` + s + `"` })

	var spec Spec
	if err := t.ReadObjectFile(&spec, "enums.yaml"); err != nil {
		log.Fatal(err)
	}
	t.Prepare(&prayer.Go{
		Input:  "enums",
		Obj:    spec,
		Args:   []any{"enums.yaml"},
		Output: "./enums/enums.go",
	})
	if err := t.Pray(); err != nil {
		log.Fatal(err)
	}
}
