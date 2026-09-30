# temple

go codegen from `text/template`s. you register templates, queue up "prayers" (template + data + output file), then `Pray()` renders them all and runs a formatter on the output.

## install

```
go get github.com/gfx-labs/temple
```

## usage

stick your templates (`.tmpl` or `.gotmpl`) next to a `main.go` and wire it up with `go:generate`:

```go
package main

//go:generate go run .

import (
	"log"

	"github.com/gfx-labs/temple"
	"github.com/gfx-labs/temple/lib/prayer"
)

func main() {
	temple.RegisterTemplateDir(".") // template name = file name w/o extension
	temple.Prepare(&prayer.Go{
		Input:  "mapn",
		Obj:    map[string]any{"Count": 10},
		Output: "./maps/maps.go",
	})
	if err := temple.Pray(); err != nil {
		log.Fatal(err)
	}
}
```

then `go generate ./...`

the top level `temple.*` funcs use a global sanctum rooted at `./`. if you want your own, use `sanctum.New(path)` or `sanctum.NewWithFS(afero.Fs)`.

## prayers

- `prayer.Go` - prepends `package <name>` (defaults to the output dir name, override with `Package`) and runs `go/format`
- `prayer.Rust` - pipes output through `rustfmt` (needs it on your path)
- `prayer.Raw` - no formatting unless you pass a `Formatter`

or implement `prayer.Prayer` yourself.

## template stuff

- all templates are loaded together, so you can `{{template "other" .}}` between files
- all of [sprig](https://masterminds.github.io/sprig/) is available
- extra builtins: `list`, `some`, `iterate`, `iterateFromN`, `iterateReverseFromN`, `lowerSnake`, `upperSnake`, `lowerCamel`, `upperCamel`, plus a few `strings` helpers
- `Args` on a prayer show up as `arg0`, `arg1`, ... inside the template
- add your own with `RegisterFunc(name, fn)`
- `ReadObjectFile(&v, "file.yaml")` loads yaml/json into whatever you pass

## examples

see `examples/`:

- `mapn` - generates typed n-key nested sync maps
- `enums` - generates go enums (String, Parse, text marshaling) from a yaml file. good starting point
