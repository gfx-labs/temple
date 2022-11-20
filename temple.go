package temple

import (
	"gfx.cafe/util/temple/lib/preset"
	"gfx.cafe/util/temple/lib/sanctum"
)

var t = sanctum.New(".")
var Sanctum = t

var curPreset preset.Preset

func SelectPreset(p preset.Preset) {
	curPreset = p
}

func RegisterTemplateFile(name string) {
	t.RegisterTemplateFile(name)
}
func RegisterTemplate(name string, content string) {
	t.RegisterTemplate(name, content)
}

func RegisterFunc(s string, fn any) {
	t.RegisterFunc(s, fn)
}

func RegisterFuncVar(s string, val any) {
	t.RegisterFuncVar(s, val)
}

func Prepare(p *sanctum.Prayer) {
	t.Prepare(p)
}

func Pray() error {
	if curPreset != nil {
		curPreset.Bind(t)
		if !curPreset.Initialized() {
			err := curPreset.Initialize()
			if err != nil {
				return err
			}
		}
		return curPreset.Run()
	}
	return t.Pray()
}
