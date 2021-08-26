package ext_surface

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/objects"
	"unsafe"
)

type Surface struct {
	instance C.VkInstance
	handle C.VkSurfaceKHR
}

func (s *Surface) Handle() uintptr {
	return uintptr(unsafe.Pointer(s.handle))
}

func (s *Surface) Destroy() {
	C.vkDestroySurfaceKHR(s.instance, s.handle, nil)
}

func (s *Surface) CanBePresentedBy(physicalDevice *objects.PhysicalDevice, queueFamilyIndex int) bool {
	deviceHandle := (C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice.Handle()))
	var canPresent C.VkBool32
	C.vkGetPhysicalDeviceSurfaceSupportKHR(deviceHandle, C.uint(queueFamilyIndex), s.handle, &canPresent)

	return canPresent != C.VK_FALSE
}
