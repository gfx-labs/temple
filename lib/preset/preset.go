package preset

import "gfx.cafe/util/temple/lib/sanctum"

type Preset interface {
	Initialize(sanctum *sanctum.Sanctum, config any) error
	Load() error
	Run() error
}
