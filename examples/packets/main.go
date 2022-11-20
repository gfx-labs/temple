package main

import (
	"embed"
	"fmt"
	"gfx.cafe/util/temple/lib/sanctum"
	"github.com/iancoleman/strcase"
	"io/fs"
	"path/filepath"
	"sigs.k8s.io/yaml"
)

const (
	OUTPUT = "./out/packets"
)

//go:embed packets.tmpl
var packetsTmpl string

//go:embed types.tmpl
var typesTmpl string

//go:embed spec
var input embed.FS

func removeExt(v string) string {
	ext := filepath.Ext(v)
	return v[:len(v)-len(ext)]
}

func main() {
	t := sanctum.New(OUTPUT)
	t.RegisterTemplate("packets", packetsTmpl)
	t.RegisterTemplate("types", typesTmpl)
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
	typesFile, err := input.ReadFile("spec/types.yaml")
	if err != nil {
		panic(err)
	}
	var types map[string]any
	err = yaml.Unmarshal(typesFile, &types)
	t.Prepare(&sanctum.Prayer{
		Input: "types",
		Obj:   types,

		PackageName: "packets",
		FileName:    "types.go",
	})

	var states []fs.DirEntry
	states, err = input.ReadDir("spec")
	if err != nil {
		panic(err)
	}

	for _, state := range states {
		if !state.IsDir() {
			continue
		}
		var directions []fs.DirEntry
		directions, err = input.ReadDir(filepath.Join("spec", state.Name()))
		for _, direction := range directions {
			directionName := removeExt(direction.Name())

			var packetsFile []byte
			packetsFile, err = input.ReadFile(filepath.Join("spec", state.Name(), direction.Name()))
			if err != nil {
				panic(err)
			}

			var packets map[string]any
			err = yaml.Unmarshal(packetsFile, &packets)
			if err != nil {
				panic(err)
			}

			packets["Types"] = types["Types"]
			packets["Name"] = strcase.ToCamel(fmt.Sprintf("%s_%s", state.Name(), directionName))

			t.Prepare(&sanctum.Prayer{
				Input: "packets",
				Obj:   packets,

				PackagePath: filepath.Join(state.Name(), directionName),
				PackageName: state.Name() + "_" + directionName,
				FileName:    "packets.go",
			})
		}
	}

	err = t.Pray()
	if err != nil {
		panic(err)
	}
}
