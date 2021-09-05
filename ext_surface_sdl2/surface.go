package ext_surface_sdl2

import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/ext_surface"
	"github.com/CannibalVox/cgoalloc"
)

func CreateSurface(allocator cgoalloc.Allocator, instance *VKng.Instance, options *CreationOptions) (*ext_surface.Surface, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	return ext_surface.CreateSurface(createInfo, instance)
}
