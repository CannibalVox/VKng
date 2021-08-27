package ext_swapchain

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/objects"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SWAPCHAIN_EXTENSION_NAME

type SwapchainHandle C.VkSwapchainKHR
type Swapchain struct {
	handle C.VkSwapchainKHR
	device C.VkDevice
}

func CreateSwapchain(allocator cgoalloc.Allocator, device *objects.Device, options *CreationOptions) (*Swapchain, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var swapchain C.VkSwapchainKHR
	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))

	res := C.vkCreateSwapchainKHR(deviceHandle, (*C.VkSwapchainCreateInfoKHR)(createInfo), nil, &swapchain)
	err = VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &Swapchain{handle: swapchain, device: deviceHandle}, nil
}

func (s *Swapchain) Handle() SwapchainHandle {
	return SwapchainHandle(s.handle)
}

func (s *Swapchain) Destroy() {
	C.vkDestroySwapchainKHR(s.device, s.handle, nil)
}
