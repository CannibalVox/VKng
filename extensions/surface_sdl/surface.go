package ext_surface_sdl2

import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/resource"
	ext_surface2 "github.com/CannibalVox/VKng/extensions/surface"
	"github.com/CannibalVox/cgoalloc"
)

func CreateSurface(allocator cgoalloc.Allocator, instance *resource.Instance, options *CreationOptions) (*ext_surface2.Surface, core.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	return ext_surface2.CreateSurface(createInfo, instance)
}
