package ext_surface_sdl2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resources"
	ext_surface "github.com/CannibalVox/VKng/extensions/surface"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

func CreateSurface(instance resources.Instance, window *sdl.Window) (ext_surface.Surface, loader.VkResult, error) {
	surfacePtrUnsafe, err := window.VulkanCreateSurface(instance.Handle())
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	surfacePtr := (*C.VkSurfaceKHR)(surfacePtrUnsafe)

	return ext_surface.CreateSurface(unsafe.Pointer(*surfacePtr), instance)
}
