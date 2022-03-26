package khr_surface_sdl2

//go:generate mockgen -source extension.go -destination ./mocks/mocks.go -package mock_surface_sdl2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/extensions/khr_surface"
	khr_surface_driver "github.com/CannibalVox/VKng/extensions/khr_surface/driver"
	"github.com/cockroachdb/errors"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

type VulkanExtension struct {
	driver khr_surface_driver.Driver
}

type Extension interface {
	CreateSurface(instance core1_0.Instance, window *sdl.Window) (khr_surface.Surface, common.VkResult, error)
}

func CreateExtensionFromInstance(instance core1_0.Instance) *VulkanExtension {
	driver := khr_surface_driver.CreateDriverFromCore(instance.Driver())
	return &VulkanExtension{
		driver: driver,
	}
}

func CreateExtensionFromDriver(driver khr_surface_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (l *VulkanExtension) CreateSurface(instance core1_0.Instance, window *sdl.Window) (khr_surface.Surface, common.VkResult, error) {
	surfacePtrUnsafe, err := window.VulkanCreateSurface((*C.VkInstance)(unsafe.Pointer(instance.Handle())))
	if err != nil {
		return nil, core1_0.VKErrorUnknown, errors.Wrap(err, "could not retrieve surface from SDL")
	}

	surfacePtr := (*C.VkSurfaceKHR)(surfacePtrUnsafe)

	return khr_surface.CreateSurface(unsafe.Pointer(*surfacePtr), instance, l.driver)
}
