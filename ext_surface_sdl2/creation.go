package ext_surface_sdl2

import (
	"github.com/CannibalVox/VKng"
	"github.com/veandco/go-sdl2/sdl"
)

type CreationOptions struct {
	Window *sdl.Window

	Next VKng.Options
}
