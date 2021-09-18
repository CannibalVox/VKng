package ext_surface_sdl2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	ext_surface "github.com/CannibalVox/VKng/extensions/surface"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

func CreateSurface(instance core.Instance, window *sdl.Window) (ext_surface.Surface, core.VkResult, error) {
	surfacePtrUnsafe, err := window.VulkanCreateSurface(instance.Handle())
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	surfacePtr := (*C.VkSurfaceKHR)(surfacePtrUnsafe)

	return ext_surface.CreateSurface(unsafe.Pointer(*surfacePtr), instance)
}
