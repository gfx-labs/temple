package sanctum

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/iancoleman/strcase"
	"github.com/spf13/afero"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v2"
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
	"iterate": func(counta any) []int {
		count := cast.ToInt(counta)
		var i int
		var Items []int
		for i = 0; i < (count); i++ {
			Items = append(Items, i)
		}
		return Items
	},
	"iterateFromN": func(na any, counta any) []int {
		n := cast.ToInt(na)
		count := cast.ToInt(counta)
		var i int
		var Items []int
		for i = n; i < (count + n); i++ {
			Items = append(Items, i)
		}
		return Items
	},
	"iterateReverseFromN": func(na any, counta any) []int {
		n := cast.ToInt(na)
		count := cast.ToInt(counta)
		var i int
		var Items []int
		for i = n; i < (count + n); i++ {
			Items = append(Items, i)
		}
		for i2, j := 0, len(Items)-1; i2 < j; i2, j = i2+1, j-1 {
			Items[i2], Items[j] = Items[j], Items[i2]
		}
		return Items
	},
}

func New(path string) *Sanctum {
	s := &Sanctum{
		template: map[string]string{},
		fm:       defaultFuncs,
		fs:       afero.NewBasePathFs(afero.NewOsFs(), path),
	}
	if filepath.Clean(path) == "." {
		s.fs = afero.NewOsFs()
	}
	return s
}

func (t *Sanctum) FS() afero.Fs {
	return t.fs
}

func (t *Sanctum) ReadObjectFile(item any, path ...string) error {
	bts, err := afero.ReadFile(t.fs, filepath.Join(path...))
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bts, item)
}

func (t *Sanctum) RegisterTemplateFile(name string) {
	bts, err := afero.ReadFile(t.fs, name)
	if err != nil {
		return
	}
	name = filepath.Base(name)
	name = strings.TrimSuffix(name, ".tmpl")
	name = strings.TrimSuffix(name, ".gotmpl")
	sbts := string(bts)
	t.RegisterTemplate(name, sbts)
}
func (t *Sanctum) RegisterTemplateDir(path string) {
	files, err := afero.ReadDir(t.fs, path)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			// TODO should this be recursive?
			continue
		}

		switch filepath.Ext(file.Name()) {
		case ".gotmpl", ".tmpl":
			t.RegisterTemplateFile(filepath.Join(path, file.Name()))
		default:
			continue
		}
	}
}
func (t *Sanctum) RegisterTemplate(name string, content string) {
	t.template[name] = strings.Trim(strings.TrimSpace(content), "\n")
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
		err = t.fs.MkdirAll(v.PackagePath, 0o744)
		if err != nil {
			log.Printf("WARNING: mkdirall failed (%s)", err)
		}
		file, err := t.fs.Create(filepath.Join(v.PackagePath, v.FileName))
		if err != nil {
			return fmt.Errorf("openfile tmpl=%s obj=%+v args=%v err=%w", v.Input, v.Obj, v.Args, err)
		}
		defer file.Close()
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
	tmp := template.New("current").
		Funcs(sprig.FuncMap()).
		Funcs(t.fm).
		Funcs(tfm)
	for name, v := range t.template {
		tmp = template.Must(tmp.Parse(fmt.Sprintf(`{{define "%s"}}%s{{end}}`, name, v)))
	}
	sb := new(strings.Builder)
	err := tmp.ExecuteTemplate(sb, s, obj)
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
