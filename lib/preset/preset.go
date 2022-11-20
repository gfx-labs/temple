package preset

import "gfx.cafe/util/temple/lib/sanctum"

type Preset interface {
	Bind(*sanctum.Sanctum)

	Initialized() bool

	Initialize() error
	Load() error
	Run() error
}
