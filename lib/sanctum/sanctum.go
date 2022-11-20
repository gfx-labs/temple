package sanctum

import (
	"strings"
	"text/template"
)

type Sanctum struct {
	template map[string]string
	fm       template.FuncMap
}

func (t *Sanctum) RegisterFunc(s string, fn any) {
	t.fm[s] = fn
}

func (t *Sanctum) Execute(s string, obj any) (string, error) {
	tmp := template.Must(template.New(s).Funcs(t.fm).Parse(t.template[s]))
	sb := new(strings.Builder)
	err := tmp.Execute(sb, obj)
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}
