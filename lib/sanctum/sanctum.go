package sanctum

import (
	"fmt"
	"strings"
	"text/template"
)

type Sanctum struct {
	template map[string]string
	fm       template.FuncMap
}

func New() *Sanctum {
	return &Sanctum{
		template: map[string]string{},
		fm:       template.FuncMap{},
	}
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

func (t *Sanctum) Execute(s string, obj any, args ...any) (string, error) {
	tfm := make(template.FuncMap)
	for k, v := range args {
		tfm[fmt.Sprintf("arg%d", k)] = func() any {
			return v
		}
	}
	tmp := template.Must(template.New(s).
		Funcs(t.fm).
		Funcs(tfm).
		Parse(t.template[s]))
	sb := new(strings.Builder)
	err := tmp.Execute(sb, obj)
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
