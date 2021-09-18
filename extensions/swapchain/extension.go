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
	"github.com/CannibalVox/VKng/core/common"
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
	Images() ([]core.Image, core.VkResult, error)
	AcquireNextImage(timeout time.Duration, semaphore core.Semaphore, fence core.Fence) (int, core.VkResult, error)
	PresentToQueue(queue core.Queue, o *PresentOptions) (resultBySwapchain []core.VkResult, res core.VkResult, anyError error)
}

func CreateSwapchain(device core.Device, options *CreationOptions) (Swapchain, core.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	var swapchain C.VkSwapchainKHR
	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))
	createFuncPtr := device.Driver().LoadProcAddr((*core.Char)(arena.CString("vkCreateSwapchainKHR")))

	res := core.VkResult(C.cgoCreateSwapchainKHR((C.PFN_vkCreateSwapchainKHR)(createFuncPtr), deviceHandle, (*C.VkSwapchainCreateInfoKHR)(createInfo), nil, &swapchain))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	destroyFunc := (C.PFN_vkDestroySwapchainKHR)(device.Driver().LoadProcAddr((*core.Char)(arena.CString("vkDestroySwapchainKHR"))))
	getImagesFunc := (C.PFN_vkGetSwapchainImagesKHR)(device.Driver().LoadProcAddr((*core.Char)(arena.CString("vkGetSwapchainImagesKHR"))))
	acquireNextFunc := (C.PFN_vkAcquireNextImageKHR)(device.Driver().LoadProcAddr((*core.Char)(arena.CString("vkAcquireNextImageKHR"))))
	queuePresentFunc := (C.PFN_vkQueuePresentKHR)(device.Driver().LoadProcAddr((*core.Char)(arena.CString("vkQueuePresentKHR"))))

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

func (s *vulkanSwapchain) Images() ([]core.Image, core.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	imageCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	imageCountRef := (*C.uint32_t)(imageCountPtr)

	res := core.VkResult(C.cgoGetSwapchainImagesKHR(s.getImagesFunc, s.device, s.handle, imageCountRef, nil))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	imageCount := int(*imageCountRef)
	if imageCount == 0 {
		return nil, res, nil
	}

	imagesPtr := allocator.Malloc(imageCount * int(unsafe.Sizeof([1]C.VkImage{})))

	res = core.VkResult(C.cgoGetSwapchainImagesKHR(s.getImagesFunc, s.device, s.handle, imageCountRef, (*C.VkImage)(imagesPtr)))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	imagesSlice := ([]core.VkImage)(unsafe.Slice((*core.VkImage)(imagesPtr), imageCount))
	var result []core.Image
	deviceHandle := (core.VkDevice)(unsafe.Pointer(s.device))
	for i := 0; i < imageCount; i++ {
		result = append(result, core.CreateFromHandles(imagesSlice[i], deviceHandle))
	}

	return result, res, nil
}

func (s *vulkanSwapchain) AcquireNextImage(timeout time.Duration, semaphore core.Semaphore, fence core.Fence) (int, core.VkResult, error) {
	var imageIndex C.uint32_t

	var semaphoreHandle C.VkSemaphore
	var fenceHandle C.VkFence

	if semaphore != nil {
		semaphoreHandle = (C.VkSemaphore)(unsafe.Pointer(semaphore.Handle()))
	}

	if fence != nil {
		fenceHandle = (C.VkFence)(unsafe.Pointer(fence.Handle()))
	}

	res := C.cgoAcquireNextImageKHR(s.acquireNextFunc, s.device, s.handle, C.uint64_t(common.TimeoutNanoseconds(timeout)), semaphoreHandle, fenceHandle, &imageIndex)
	result := core.VkResult(res)
	err := result.ToError()
	return int(imageIndex), result, err
}
