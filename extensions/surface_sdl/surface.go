package ext_surface_sdl2

import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resource"
	ext_surface2 "github.com/CannibalVox/VKng/extensions/surface"
	"github.com/CannibalVox/cgoalloc"
)

func CreateSurface(allocator cgoalloc.Allocator, instance *resource.Instance, options *CreationOptions) (*ext_surface2.Surface, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	return ext_surface2.CreateSurface(allocator, createInfo, instance)
}
