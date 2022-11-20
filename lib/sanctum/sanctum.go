package sanctum

import (
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/iancoleman/strcase"
	"github.com/spf13/afero"
)

type Sanctum struct {
	template map[string]string
	fm       template.FuncMap

	foyer Foyer
	fs    afero.Fs
}

type Foyer struct {
	Prayers []Prayer
}

type Prayer struct {
	Input string
	Obj   any
	Args  []any

	PackagePath string
	PackageName string
	FileName    string
}

var defaultFuncs = template.FuncMap{
	// list packs multiple arguments into one object (a slice)
	// arguments can be re obtained by calling index (e.g. `index s 0`)
	"list": func(v ...any) []any {
		return v
	},
	// funcs to convert between cases
	"lowerSnake": strcase.ToSnake,
	"upperSnake": strcase.ToScreamingSnake,
	"lowerCamel": strcase.ToLowerCamel,
	"upperCamel": strcase.ToCamel,
	// some is a helper function to detect whether a value exists.
	// by default, doing something like `if v` where v is an int that is 0 will return false
	// if you do `if some v`, this will always return true if v is not nil
	// useful for getting detecting if a map key exists: `if some map.Foo`
	"some": func(v any) bool {
		return v != nil
	},
	"repeat":     strings.Repeat,
	"trimPrefix": strings.TrimPrefix,
	"trimSuffix": strings.TrimSuffix,
	"hasPrefix":  strings.HasPrefix,
	"hasSuffix":  strings.HasSuffix,
	"trim":       strings.Trim,
	"toLower":    strings.ToLower,
	"toUpper":    strings.ToUpper,
	"equalFold":  strings.EqualFold,
}

func New(path string) *Sanctum {
	return &Sanctum{
		template: map[string]string{},
		fm:       defaultFuncs,
		fs:       afero.NewBasePathFs(afero.NewOsFs(), path),
	}
}

func (t *Sanctum) FS() afero.Fs {
	return t.fs
}

func (t *Sanctum) RegisterTemplateFile(name string) {
	bts, err := afero.ReadFile(t.fs, name)
	if err != nil {
		return
	}
	sbts := string(bts)
	sbts = strings.TrimSuffix(sbts, ".tmpl")
	sbts = strings.TrimSuffix(sbts, ".gotmpl")
	t.template[name] = sbts
}
func (t *Sanctum) RegisterTemplate(name string, content string) {
	t.template[name] = content
}

func (t *Sanctum) RegisterFunc(s string, fn any) {
	t.fm[s] = fn
}
func (t *Sanctum) RegisterFuncVar(s string, val any) {
	t.fm[s] = func() any {
		return val
	}
}
func (t *Sanctum) Prepare(p *Prayer) {
	if p != nil {
		t.foyer.Prayers = append(t.foyer.Prayers, *p)
	}
}
func (t *Sanctum) Pray() error {
	for _, v := range t.foyer.Prayers {
		output, err := t.execute(v.Input, v.Obj, v.Args...)
		if err != nil {
			return fmt.Errorf("exec tmpl=%s obj=%+v args=%v err=%w", v.Input, v.Obj, v.Args, err)
		}
		t.fs.MkdirAll(v.PackagePath, 0777)
		file, err := t.fs.Create(path.Join(v.PackagePath, v.FileName))
		defer file.Close()
		if err != nil {
			return fmt.Errorf("openfile tmpl=%s obj=%+v args=%v err=%w", v.Input, v.Obj, v.Args, err)
		}
		err = file.Truncate(0)
		if err != nil {
			return fmt.Errorf("truncfile tmpl=%s obj=%+v args=%v err=%w", v.Input, v.Obj, v.Args, err)
		}
		file.WriteString(fmt.Sprintf("package %s\n\n", v.PackageName))
		file.WriteString(output)
	}
	return nil
}

func (t *Sanctum) execute(s string, obj any, args ...any) (string, error) {
	tfm := make(template.FuncMap)
	for k, v := range args {
		tfm[fmt.Sprintf("arg%d", k)] = func() any {
			return v
		}
	}
	tmp := template.New(s).
		Funcs(sprig.FuncMap()).
		Funcs(t.fm).
		Funcs(tfm)
	for _, v := range t.template {
		tmp = template.Must(tmp.Parse(v))
	}
	tmp = template.Must(tmp.Parse(t.template[s]))
	sb := new(strings.Builder)
	err := tmp.Execute(sb, obj)
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
