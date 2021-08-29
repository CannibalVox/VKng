package ext_surface_sdl2

import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/ext_surface"
	"github.com/CannibalVox/cgoalloc"
)

func CreateSurface(allocator cgoalloc.Allocator, instance *VKng.Instance, options *CreationOptions) (*ext_surface.Surface, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	return ext_surface.CreateSurface(createInfo, instance)
}
