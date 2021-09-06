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

func CreateSwapchain(allocator cgoalloc.Allocator, device *core.Device, options *CreationOptions) (*Swapchain, VKng.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, VKng.VKErrorUnknown, err
	}

	var swapchain C.VkSwapchainKHR
	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))

	res := VKng.Result(C.vkCreateSwapchainKHR(deviceHandle, (*C.VkSwapchainCreateInfoKHR)(createInfo), nil, &swapchain))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &Swapchain{handle: swapchain, device: deviceHandle}, res, nil
}

func (s *Swapchain) Handle() SwapchainHandle {
	return SwapchainHandle(s.handle)
}

func (s *Swapchain) Destroy() {
	C.vkDestroySwapchainKHR(s.device, s.handle, nil)
}

func (s *Swapchain) Images(allocator cgoalloc.Allocator) ([]*core.Image, VKng.Result, error) {
	imageCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(imageCountPtr)

	imageCountRef := (*C.uint32_t)(imageCountPtr)

	res := VKng.Result(C.vkGetSwapchainImagesKHR(s.device, s.handle, imageCountRef, nil))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	imageCount := int(*imageCountRef)
	if imageCount == 0 {
		return nil, res, nil
	}

	imagesPtr := allocator.Malloc(imageCount * int(unsafe.Sizeof([1]C.VkImage{})))
	defer allocator.Free(imagesPtr)

	imagesRef := (*C.VkImage)(imagesPtr)

	res = VKng.Result(C.vkGetSwapchainImagesKHR(s.device, s.handle, imageCountRef, imagesRef))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	imagesSlice := ([]C.VkImage)(unsafe.Slice(imagesRef, imageCount))
	var result []*core.Image
	for i := 0; i < imageCount; i++ {
		result = append(result, core.CreateFromHandles(core.ImageHandle(unsafe.Pointer(imagesSlice[i])), core.DeviceHandle(unsafe.Pointer(s.device))))
	}

	return result, res, nil
}

func (s *Swapchain) AcquireNextImage(timeout time.Duration, semaphore *core.Semaphore, fence *core.Fence) (int, VKng.Result, error) {
	var imageIndex C.uint32_t

	var semaphoreHandle C.VkSemaphore
	var fenceHandle C.VkFence

	if semaphore != nil {
		semaphoreHandle = (C.VkSemaphore)(unsafe.Pointer(semaphore.Handle()))
	}

	if fence != nil {
		fenceHandle = (C.VkFence)(unsafe.Pointer(fence.Handle()))
	}

	res := C.vkAcquireNextImageKHR(s.device, s.handle, C.uint64_t(VKng.TimeoutNanoseconds(timeout)), semaphoreHandle, fenceHandle, &imageIndex)
	result := VKng.Result(res)
	err := result.ToError()
	return int(imageIndex), result, err
}
