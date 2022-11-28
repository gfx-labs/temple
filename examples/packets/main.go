package main

import (
	_ "embed"
	"fmt"
	"gfx.cafe/util/temple/lib/sanctum"
	"github.com/iancoleman/strcase"
	"github.com/spf13/afero"
	"go/format"
	"io/fs"
	"path/filepath"
)

//go:generate go run .

//go:embed packets.tmpl
var packetsTmpl string

//go:embed types.tmpl
var typesTmpl string

func removeExt(v string) string {
	ext := filepath.Ext(v)
	return v[:len(v)-len(ext)]
}

func main() {
	t := sanctum.New(".")
	t.RegisterTemplateDir(".")
	t.RegisterFunc(
		"isCustomType", func(v string) bool {
			switch v {
			case "VarInt", "VarLong", "NBT", "UUID":
				return true
			default:
				return false
			}
		})
	t.RegisterFunc(
		"isGeneratedType", func(v string) bool {
			switch v {
			case "byte", "uint8", "uint16", "uint32", "uint64", "uint",
				"int8", "int16", "int32", "int64", "int",
				"float32", "float64", "string", "bool",
				"VarInt", "VarLong", "NBT", "UUID":
				return false
			default:
				return true
			}
		})
	t.RegisterFunc(
		"emptyMap", func() map[any]any {
			return make(map[any]any)
		})
	t.RegisterFunc(
		"cloneMap", func(m map[any]any) map[any]any {
			out := make(map[any]any)
			for k, v := range m {
				out[k] = v
			}
			return out
		})
	t.RegisterFunc(
		"setMap", func(m map[any]any, k any, v any) struct{} {
			m[k] = v
			return struct{}{}
		})
	t.RegisterFunc(
		"getMap", func(m map[any]any, k any) any {
			return m[k]
		})
	t.RegisterFunc(
		"isString", func(v any) bool {
			_, ok := v.(string)
			return ok
		})
	t.RegisterFunc(
		"add", func(a, b int) int {
			return a + b
		})

	var ty map[string]any
	err := t.ReadObjectFile(&ty, "spec/types.yaml")
	if err != nil {
		panic(err)
	}
	t.Prepare(&sanctum.Prayer{
		Input:     "types",
		Obj:       ty,
		Formatter: format.Source,

		PackagePath: "out",
		PackageName: "packets",
		FileName:    "types.go",
	})

	err = afero.Walk(t.FS(), "spec", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Base(path) == "types.yaml" {
			return nil
		}

		stateName := filepath.Base(filepath.Dir(path))
		directionName := removeExt(filepath.Base(path))

		var v map[string]any
		err = t.ReadObjectFile(&v, path)
		if err != nil {
			return err
		}

		v["Types"] = ty["Types"]
		v["Name"] = strcase.ToCamel(fmt.Sprintf("%s_%s", stateName, directionName))

		t.Prepare(&sanctum.Prayer{
			Input:     "packets",
			Obj:       v,
			Formatter: format.Source,

			PackagePath: filepath.Join("out", stateName, directionName),
			PackageName: stateName + "_" + directionName,
			FileName:    "packets.go",
		})

		return nil
	})
	if err != nil {
		panic(err)
	}

	err = t.Pray()
	if err != nil {
		panic(err)
	}
}
