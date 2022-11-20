package temple

import (
	"gfx.cafe/util/temple/lib/preset"
	"gfx.cafe/util/temple/lib/sanctum"
)

var Sanctum = t
var t *sanctum.Sanctum

func init() {
	Sanctum = sanctum.New("./")
	t = Sanctum
}

var curPreset preset.Preset

func SelectPreset(p preset.Preset) {
	curPreset = p
}

func ReadObjectFile(item any, path ...string) {
	t.ReadObjectFile(item, path...)
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
