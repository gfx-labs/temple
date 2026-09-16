package preset

import "github.com/gfx-labs/temple/lib/sanctum"

type Preset interface {
	Bind(*sanctum.Sanctum)

	Initialized() bool

	Initialize() error
	Load() error
	Run() error
}
