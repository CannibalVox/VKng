package ext_surface_sdl2

import (
	"github.com/CannibalVox/VKng/creation"
	"github.com/veandco/go-sdl2/sdl"
)

type CreationOptions struct {
	Window *sdl.Window

	Next creation.Options
}
