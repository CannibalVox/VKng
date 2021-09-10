package ext_swapchain

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"

VkResult cgoAcquireNextImageKHR(PFN_vkAcquireNextImageKHR fn, VkDevice device, VkSwapchainKHR swapchain, uint64_t timeout, VkSemaphore semaphore, VkFence fence, uint32_t* pImageIndex) {
	return fn(device, swapchain, timeout, semaphore, fence, pImageIndex);
}

VkResult cgoCreateSwapchainKHR(PFN_vkCreateSwapchainKHR fn, VkDevice device, VkSwapchainCreateInfoKHR* pCreateInfo, VkAllocationCallbacks* pAllocator, VkSwapchainKHR* pSwapchain) {
	return fn(device, pCreateInfo, pAllocator, pSwapchain);
}

void cgoDestroySwapchainKHR(PFN_vkDestroySwapchainKHR fn, VkDevice device, VkSwapchainKHR swapchain, VkAllocationCallbacks* pAllocator) {
	fn(device, swapchain, pAllocator);
}

VkResult cgoGetSwapchainImagesKHR(PFN_vkGetSwapchainImagesKHR fn, VkDevice device, VkSwapchainKHR swapchain, uint32_t* pSwapchainImageCount, VkImage* pSwapchainImages) {
	return fn(device, swapchain, pSwapchainImageCount, pSwapchainImages);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resource"
	"github.com/CannibalVox/cgoalloc"
	"time"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SWAPCHAIN_EXTENSION_NAME

type SwapchainHandle C.VkSwapchainKHR
type Swapchain struct {
	handle C.VkSwapchainKHR
	device C.VkDevice

	destroyFunc      C.PFN_vkDestroySwapchainKHR
	getImagesFunc    C.PFN_vkGetSwapchainImagesKHR
	acquireNextFunc  C.PFN_vkAcquireNextImageKHR
	queuePresentFunc C.PFN_vkQueuePresentKHR
}

func CreateSwapchain(allocator cgoalloc.Allocator, device *resource.Device, options *CreationOptions) (*Swapchain, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := options.AllocForC(arena)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var swapchain C.VkSwapchainKHR
	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))
	createFuncPtr := device.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(arena, "vkCreateSwapchainKHR")))

	res := loader.VkResult(C.cgoCreateSwapchainKHR((C.PFN_vkCreateSwapchainKHR)(createFuncPtr), deviceHandle, (*C.VkSwapchainCreateInfoKHR)(createInfo), nil, &swapchain))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	destroyFunc := (C.PFN_vkDestroySwapchainKHR)(device.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(arena, "vkDestroySwapchainKHR"))))
	getImagesFunc := (C.PFN_vkGetSwapchainImagesKHR)(device.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(arena, "vkGetSwapchainImagesKHR"))))
	acquireNextFunc := (C.PFN_vkAcquireNextImageKHR)(device.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(arena, "vkAcquireNextImageKHR"))))
	queuePresentFunc := (C.PFN_vkQueuePresentKHR)(device.Loader().LoadProcAddr((*loader.Char)(cgoalloc.CString(arena, "vkQueuePresentKHR"))))

	return &Swapchain{
		handle: swapchain,
		device: deviceHandle,

		destroyFunc:      destroyFunc,
		getImagesFunc:    getImagesFunc,
		acquireNextFunc:  acquireNextFunc,
		queuePresentFunc: queuePresentFunc,
	}, res, nil
}

func (s *Swapchain) Handle() SwapchainHandle {
	return SwapchainHandle(s.handle)
}

func (s *Swapchain) Destroy() {
	C.cgoDestroySwapchainKHR(s.destroyFunc, s.device, s.handle, nil)
}

func (s *Swapchain) Images(allocator cgoalloc.Allocator) ([]*resource.Image, loader.VkResult, error) {
	imageCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	defer allocator.Free(imageCountPtr)

	imageCountRef := (*C.uint32_t)(imageCountPtr)

	res := loader.VkResult(C.cgoGetSwapchainImagesKHR(s.getImagesFunc, s.device, s.handle, imageCountRef, nil))
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

	res = loader.VkResult(C.cgoGetSwapchainImagesKHR(s.getImagesFunc, s.device, s.handle, imageCountRef, (*C.VkImage)(imagesPtr)))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	imagesSlice := ([]loader.VkImage)(unsafe.Slice((*loader.VkImage)(imagesPtr), imageCount))
	var result []*resource.Image
	deviceHandle := (loader.VkDevice)(unsafe.Pointer(s.device))
	for i := 0; i < imageCount; i++ {
		result = append(result, resource.CreateFromHandles(imagesSlice[i], deviceHandle))
	}

	return result, res, nil
}

func (s *Swapchain) AcquireNextImage(timeout time.Duration, semaphore *resource.Semaphore, fence *resource.Fence) (int, loader.VkResult, error) {
	var imageIndex C.uint32_t

	var semaphoreHandle C.VkSemaphore
	var fenceHandle C.VkFence

	if semaphore != nil {
		semaphoreHandle = (C.VkSemaphore)(unsafe.Pointer(semaphore.Handle()))
	}

	if fence != nil {
		fenceHandle = (C.VkFence)(unsafe.Pointer(fence.Handle()))
	}

	res := C.cgoAcquireNextImageKHR(s.acquireNextFunc, s.device, s.handle, C.uint64_t(core.TimeoutNanoseconds(timeout)), semaphoreHandle, fenceHandle, &imageIndex)
	result := loader.VkResult(res)
	err := result.ToError()
	return int(imageIndex), result, err
}
