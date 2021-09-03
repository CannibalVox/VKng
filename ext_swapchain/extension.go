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
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"time"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SWAPCHAIN_EXTENSION_NAME

type SwapchainHandle C.VkSwapchainKHR
type Swapchain struct {
	handle C.VkSwapchainKHR
	device C.VkDevice
}

func CreateSwapchain(allocator cgoalloc.Allocator, device *VKng.Device, options *CreationOptions) (*Swapchain, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var swapchain C.VkSwapchainKHR
	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))

	res := C.vkCreateSwapchainKHR(deviceHandle, (*C.VkSwapchainCreateInfoKHR)(createInfo), nil, &swapchain)
	err = core.Result(res).ToError()
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

func (s *Swapchain) Images(allocator cgoalloc.Allocator) ([]*VKng.Image, error) {
	imageCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(imageCountPtr)

	imageCountRef := (*C.uint32_t)(imageCountPtr)

	res := C.vkGetSwapchainImagesKHR(s.device, s.handle, imageCountRef, nil)
	err := core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	imageCount := int(*imageCountRef)
	if imageCount == 0 {
		return nil, nil
	}

	imagesPtr := allocator.Malloc(imageCount * int(unsafe.Sizeof([1]C.VkImage{})))
	defer allocator.Free(imagesPtr)

	imagesRef := (*C.VkImage)(imagesPtr)

	res = C.vkGetSwapchainImagesKHR(s.device, s.handle, imageCountRef, imagesRef)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	imagesSlice := ([]C.VkImage)(unsafe.Slice(imagesRef, imageCount))
	var result []*VKng.Image
	for i := 0; i < imageCount; i++ {
		result = append(result, VKng.CreateFromHandles(VKng.ImageHandle(unsafe.Pointer(imagesSlice[i])), VKng.DeviceHandle(unsafe.Pointer(s.device))))
	}

	return result, nil
}

const NoTimeout = time.Duration(^int64(0))

func (s *Swapchain) AcquireNextImage(timeout time.Duration, semaphore *VKng.Semaphore, fence *VKng.Fence) (int, error) {
	var imageIndex C.uint32_t

	var semaphoreHandle C.VkSemaphore
	var fenceHandle C.VkFence

	if semaphore != nil {
		semaphoreHandle = (C.VkSemaphore)(unsafe.Pointer(semaphore.Handle()))
	}

	if fence != nil {
		fenceHandle = (C.VkFence)(unsafe.Pointer(fence.Handle()))
	}

	var cTimeout C.uint64_t
	if timeout == NoTimeout {
		cTimeout = C.uint64_t(timeout)
	} else {
		cTimeout = C.uint64_t(timeout.Nanoseconds())
	}

	res := C.vkAcquireNextImageKHR(s.device, s.handle, cTimeout, semaphoreHandle, fenceHandle, &imageIndex)
	err := core.Result(res).ToError()
	return int(imageIndex), err
}
