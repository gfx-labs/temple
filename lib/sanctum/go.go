package sanctum

import (
	"fmt"
	"go/format"
	"path/filepath"
)

type GoPrayer struct {
	// Input specifies which template to use.
	Input string
	// Obj is the data that will be passed into the template.
	Obj any
	// Args will be accessible within the template as arg0, arg1, argN...
	Args []any

	// Package defines the name of the package for this prayer.
	Package string
	// Output defines the file path and name to place the output of this prayer.
	Output string
}

func (g *GoPrayer) Template() string {
	return g.Input
}

func (g *GoPrayer) Object() any {
	return g.Obj
}

func (g *GoPrayer) Arguments() []any {
	return g.Args
}

func (g *GoPrayer) Format(bytes []byte) ([]byte, error) {
	pkg := g.Package
	if pkg == "" {
		pkg = filepath.Base(filepath.Dir(g.Output))
	}
	src := fmt.Sprintf("package %s\n\n%s", pkg, string(bytes))
	return format.Source([]byte(
		src,
	))
}

func (g *GoPrayer) FileName() string {
	return g.Output
}

var _ Prayer = (*GoPrayer)(nil)
