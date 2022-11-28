package sanctum

import (
	"bytes"
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
	// Input specifies which template to use.
	Input string
	// Obj is the data that will be passed into the template.
	Obj any
	// Args will be accessible within the template as arg0, arg1, argN...
	Args []any
	// Formatter defines which formatter to use (for go output use format.Source aka gofmt).
	// Setting to nil will use no formatter.
	Formatter func([]byte) ([]byte, error)

	// PackagePath defines which directory to place the output of this prayer.
	PackagePath string
	// PackageName defines the name of the package for this prayer.
	PackageName string
	// FileName defines the file to place the output of this prayer.
	FileName string
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

func NewWithFS(fs afero.Fs) *Sanctum {
	return &Sanctum{
		template: map[string]string{},
		fm:       defaultFuncs,
		fs:       fs,
	}
}

func New(path string) *Sanctum {
	if filepath.Clean(path) == "." {
		return NewWithFS(afero.NewOsFs())
	}
	return NewWithFS(afero.NewBasePathFs(afero.NewOsFs(), path))
}

func (t *Sanctum) FS() afero.Fs {
	return t.fs
}

// ReadObjectFile reads the file at path as yaml and unmarshal it into item
func (t *Sanctum) ReadObjectFile(item any, path ...string) error {
	bts, err := afero.ReadFile(t.fs, filepath.Join(path...))
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bts, item)
}

// RegisterTemplateFile reads the file at name and registers it as a template using the file's base name (without the extension)
func (t *Sanctum) RegisterTemplateFile(name string) {
	bts, err := afero.ReadFile(t.fs, name)
	if err != nil {
		return
	}
	name = filepath.Base(name)
	// trim extension
	name = name[:len(name)-len(filepath.Ext(name))]
	sbts := string(bts)
	t.RegisterTemplate(name, sbts)
}

// RegisterTemplateDir scans for files ending in .gotmpl or .tmpl and registers them as templates
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

// RegisterTemplate registers the template using name
func (t *Sanctum) RegisterTemplate(name string, content string) {
	t.template[name] = strings.Trim(strings.TrimSpace(content), "\n")
}

func (t *Sanctum) RegisterFunc(s string, fn any) {
	t.fm[s] = fn
}

func (t *Sanctum) RegisterFuncs(m template.FuncMap) {
	for s, fn := range m {
		t.RegisterFunc(s, fn)
	}
}

func (t *Sanctum) RegisterFuncVar(s string, val any) {
	t.fm[s] = func() any {
		return val
	}
}

// Prepare queues a Prayer for execution
func (t *Sanctum) Prepare(p *Prayer) {
	if p != nil {
		t.foyer.Prayers = append(t.foyer.Prayers, *p)
	}
}

// Curse discards all current Prayer without executing them
func (t *Sanctum) Curse() {
	t.foyer.Prayers = t.foyer.Prayers[:0]
}

// Pray executes all current Prayer
func (t *Sanctum) Pray() error {
	var buf bytes.Buffer
	for _, v := range t.foyer.Prayers {
		output, err := t.execute(v.Input, v.Obj, v.Args...)
		if err != nil {
			return fmt.Errorf("exec tmpl=%s obj=%+v args=%v err=%w", v.Input, v.Obj, v.Args, err)
		}
		err = t.fs.MkdirAll(v.PackagePath, 0o744)
		if err != nil {
			log.Printf("WARNING: mkdirall failed (%s)", err)
		}
		err = func() error {
			file, err := t.fs.Create(filepath.Join(v.PackagePath, v.FileName))
			if err != nil {
				return fmt.Errorf("openfile tmpl=%s obj=%+v args=%v err=%w", v.Input, v.Obj, v.Args, err)
			}
			defer file.Close()
			err = file.Truncate(0)
			if err != nil {
				return fmt.Errorf("truncfile tmpl=%s obj=%+v args=%v err=%w", v.Input, v.Obj, v.Args, err)
			}

			buf.Reset()
			buf.WriteString(fmt.Sprintf("package %s\n\n", v.PackageName))
			buf.WriteString(output)
			var fmtd []byte
			if v.Formatter != nil {
				fmtd, err = v.Formatter(buf.Bytes())
				if err != nil {
					return fmt.Errorf("gofmt tmpl=%s obj=%+v args=%v err=%w", v.Input, v.Obj, v.Args, err)
				}
			} else {
				fmtd = buf.Bytes()
			}
			_, err = file.Write(fmtd)
			if err != nil {
				return fmt.Errorf("write tmpl=%s obj=%+v args=%v err=%w", v.Input, v.Obj, v.Args, err)
			}
			return nil
		}()
		if err != nil {
			return err
		}
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
