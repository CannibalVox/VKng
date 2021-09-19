package khr_surface_sdl2

//go:generate mockgen -source extension.go -destination ./mocks/mocks.go -package mock_surface_sdl2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/extensions/khr_surface"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

type khrSurfaceSDl2Loader struct {
	instance core.Instance
	driver   khr_surface.Driver
}

type Loader interface {
	CreateSurface(window *sdl.Window) (khr_surface.Surface, core.VkResult, error)
}

func CreateLoaderFromInstance(instance core.Instance) Loader {
	driver := khr_surface.CreateDriverFromInstance(instance)
	return &khrSurfaceSDl2Loader{
		instance: instance,
		driver:   driver,
	}
}

func CreateLoaderFromDriver(instance core.Instance, driver khr_surface.Driver) Loader {
	return &khrSurfaceSDl2Loader{
		instance: instance,
		driver:   driver,
	}
}

func (l *khrSurfaceSDl2Loader) CreateSurface(window *sdl.Window) (khr_surface.Surface, core.VkResult, error) {
	surfacePtrUnsafe, err := window.VulkanCreateSurface(l.instance.Handle())
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	surfacePtr := (*C.VkSurfaceKHR)(surfacePtrUnsafe)

	return khr_surface.CreateSurface(unsafe.Pointer(*surfacePtr), l.instance, l.driver)
}
