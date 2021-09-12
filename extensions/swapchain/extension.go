package ext_swapchain

//go:generate mockgen -source extension.go -destination ./mocks/mocks.go -package mock_swapchain

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
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SWAPCHAIN_EXTENSION_NAME

type SwapchainHandle C.VkSwapchainKHR
type vulkanSwapchain struct {
	handle C.VkSwapchainKHR
	device C.VkDevice

	destroyFunc      C.PFN_vkDestroySwapchainKHR
	getImagesFunc    C.PFN_vkGetSwapchainImagesKHR
	acquireNextFunc  C.PFN_vkAcquireNextImageKHR
	queuePresentFunc C.PFN_vkQueuePresentKHR
}

type Swapchain interface {
	Handle() SwapchainHandle
	Destroy()
	Images() ([]resources.Image, loader.VkResult, error)
	AcquireNextImage(timeout time.Duration, semaphore resources.Semaphore, fence resources.Fence) (int, loader.VkResult, error)
	PresentToQueue(queue resources.Queue, o *PresentOptions) (resultBySwapchain []loader.VkResult, res loader.VkResult, anyError error)
}

func CreateSwapchain(device resources.Device, options *CreationOptions) (Swapchain, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, options)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var swapchain C.VkSwapchainKHR
	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))
	createFuncPtr := device.Loader().LoadProcAddr((*loader.Char)(arena.CString("vkCreateSwapchainKHR")))

	res := loader.VkResult(C.cgoCreateSwapchainKHR((C.PFN_vkCreateSwapchainKHR)(createFuncPtr), deviceHandle, (*C.VkSwapchainCreateInfoKHR)(createInfo), nil, &swapchain))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	destroyFunc := (C.PFN_vkDestroySwapchainKHR)(device.Loader().LoadProcAddr((*loader.Char)(arena.CString("vkDestroySwapchainKHR"))))
	getImagesFunc := (C.PFN_vkGetSwapchainImagesKHR)(device.Loader().LoadProcAddr((*loader.Char)(arena.CString("vkGetSwapchainImagesKHR"))))
	acquireNextFunc := (C.PFN_vkAcquireNextImageKHR)(device.Loader().LoadProcAddr((*loader.Char)(arena.CString("vkAcquireNextImageKHR"))))
	queuePresentFunc := (C.PFN_vkQueuePresentKHR)(device.Loader().LoadProcAddr((*loader.Char)(arena.CString("vkQueuePresentKHR"))))

	return &vulkanSwapchain{
		handle: swapchain,
		device: deviceHandle,

		destroyFunc:      destroyFunc,
		getImagesFunc:    getImagesFunc,
		acquireNextFunc:  acquireNextFunc,
		queuePresentFunc: queuePresentFunc,
	}, res, nil
}

func (s *vulkanSwapchain) Handle() SwapchainHandle {
	return SwapchainHandle(s.handle)
}

func (s *vulkanSwapchain) Destroy() {
	C.cgoDestroySwapchainKHR(s.destroyFunc, s.device, s.handle, nil)
}

func (s *vulkanSwapchain) Images() ([]resources.Image, loader.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	imageCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
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

	res = loader.VkResult(C.cgoGetSwapchainImagesKHR(s.getImagesFunc, s.device, s.handle, imageCountRef, (*C.VkImage)(imagesPtr)))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	imagesSlice := ([]loader.VkImage)(unsafe.Slice((*loader.VkImage)(imagesPtr), imageCount))
	var result []resources.Image
	deviceHandle := (loader.VkDevice)(unsafe.Pointer(s.device))
	for i := 0; i < imageCount; i++ {
		result = append(result, resources.CreateFromHandles(imagesSlice[i], deviceHandle))
	}

	return result, res, nil
}

func (s *vulkanSwapchain) AcquireNextImage(timeout time.Duration, semaphore resources.Semaphore, fence resources.Fence) (int, loader.VkResult, error) {
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
