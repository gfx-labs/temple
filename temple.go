package temple

import (
	"gfx.cafe/util/temple/lib/prayer"
	"gfx.cafe/util/temple/lib/preset"
	"gfx.cafe/util/temple/lib/sanctum"
)

var Sanctum *sanctum.Sanctum

func init() {
	Sanctum = sanctum.New("./")
}

var curPreset preset.Preset

func SelectPreset(p preset.Preset) {
	curPreset = p
}

func ReadObjectFile(item any, path ...string) {
	Sanctum.ReadObjectFile(item, path...)
}
func RegisterTemplateFile(name string) {
	Sanctum.RegisterTemplateFile(name)
}
func RegisterTemplateDir(path string) {
	Sanctum.RegisterTemplateDir(path)
}
func RegisterTemplate(name string, content string) {
	Sanctum.RegisterTemplate(name, content)
}

func RegisterFunc(s string, fn any) {
	Sanctum.RegisterFunc(s, fn)
}

func RegisterFuncVar(s string, val any) {
	Sanctum.RegisterFuncVar(s, val)
}

func Prepare(p prayer.Prayer) {
	Sanctum.Prepare(p)
}

func Pray() error {
	if curPreset != nil {
		curPreset.Bind(Sanctum)
		if !curPreset.Initialized() {
			err := curPreset.Initialize()
			if err != nil {
				return err
			}
		}
		return curPreset.Run()
	}
	return Sanctum.Pray()
}
