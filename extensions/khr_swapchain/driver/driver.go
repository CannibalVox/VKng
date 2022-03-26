package khr_swapchain_driver

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_swapchain

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

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

VkResult cgoQueuePresentKHR(PFN_vkQueuePresentKHR fn, VkQueue queue, VkPresentInfoKHR* pPresentInfo) {
	return fn(queue, pPresentInfo);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type CDriver struct {
	driver           driver.Driver
	createFunc       C.PFN_vkCreateSwapchainKHR
	destroyFunc      C.PFN_vkDestroySwapchainKHR
	getImagesFunc    C.PFN_vkGetSwapchainImagesKHR
	acquireNextFunc  C.PFN_vkAcquireNextImageKHR
	queuePresentFunc C.PFN_vkQueuePresentKHR
}

type VkSwapchainKHR driver.VulkanHandle
type VkSwapchainCreateInfoKHR C.VkSwapchainCreateInfoKHR
type VkPresentInfoKHR C.VkPresentInfoKHR
type Driver interface {
	CreateImage(imageHandle driver.VkImage, deviceHandle driver.VkDevice) core1_0.Image
	VkCreateSwapchainKHR(device driver.VkDevice, pCreateInfo *VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *VkSwapchainKHR) (common.VkResult, error)
	VkDestroySwapchainKHR(device driver.VkDevice, swapchain VkSwapchainKHR, pAllocator *driver.VkAllocationCallbacks)
	VkGetSwapchainImagesKHR(device driver.VkDevice, swapchain VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error)
	VkAcquireNextImageKHR(device driver.VkDevice, swapchain VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error)
	VkQueuePresentKHR(queue driver.VkQueue, pPresentInfo *VkPresentInfoKHR) (common.VkResult, error)
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		driver:           coreDriver,
		createFunc:       (C.PFN_vkCreateSwapchainKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCreateSwapchainKHR")))),
		destroyFunc:      (C.PFN_vkDestroySwapchainKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkDestroySwapchainKHR")))),
		getImagesFunc:    (C.PFN_vkGetSwapchainImagesKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetSwapchainImagesKHR")))),
		acquireNextFunc:  (C.PFN_vkAcquireNextImageKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkAcquireNextImageKHR")))),
		queuePresentFunc: (C.PFN_vkQueuePresentKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkQueuePresentKHR")))),
	}
}

func (d *CDriver) CreateImage(imageHandle driver.VkImage, deviceHandle driver.VkDevice) core1_0.Image {
	return core.CreateImageFromHandles(imageHandle, deviceHandle, d.driver)
}

func (d *CDriver) VkCreateSwapchainKHR(device driver.VkDevice, pCreateInfo *VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *VkSwapchainKHR) (common.VkResult, error) {
	if d.createFunc == nil {
		panic("attempt to call extension method vkCreateSwapchainKHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateSwapchainKHR(d.createFunc,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSwapchainCreateInfoKHR)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkSwapchainKHR)(unsafe.Pointer(pSwapchain))))

	return res, res.ToError()
}

func (d *CDriver) VkDestroySwapchainKHR(device driver.VkDevice, swapchain VkSwapchainKHR, pAllocator *driver.VkAllocationCallbacks) {
	if d.destroyFunc == nil {
		panic("attempt to call extension method vkDestroySwapchainKHR when extension not present")
	}

	C.cgoDestroySwapchainKHR(d.destroyFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(unsafe.Pointer(swapchain)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (d *CDriver) VkGetSwapchainImagesKHR(device driver.VkDevice, swapchain VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
	if d.getImagesFunc == nil {
		panic("attempt to call extension method vkGetSwapchainImagesKHR when extension not present")
	}

	res := common.VkResult(C.cgoGetSwapchainImagesKHR(d.getImagesFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(unsafe.Pointer(swapchain)),
		(*C.uint32_t)(unsafe.Pointer(pSwapchainImageCount)),
		(*C.VkImage)(unsafe.Pointer(pSwapchainImages))))

	return res, res.ToError()
}

func (d *CDriver) VkAcquireNextImageKHR(device driver.VkDevice, swapchain VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error) {
	if d.acquireNextFunc == nil {
		panic("attempt to call extension method vkAcquireNextImageKHR when extension not present")
	}

	res := common.VkResult(C.cgoAcquireNextImageKHR(d.acquireNextFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(unsafe.Pointer(swapchain)),
		C.uint64_t(timeout),
		C.VkSemaphore(unsafe.Pointer(semaphore)),
		C.VkFence(unsafe.Pointer(fence)),
		(*C.uint32_t)(unsafe.Pointer(pImageIndex)),
	))

	return res, res.ToError()
}

func (d *CDriver) VkQueuePresentKHR(queue driver.VkQueue, pPresentInfo *VkPresentInfoKHR) (common.VkResult, error) {
	if d.queuePresentFunc == nil {
		panic("attempt to call extension method vkQueuePresentKHR when extension not present")
	}

	res := common.VkResult(C.cgoQueuePresentKHR(d.queuePresentFunc,
		C.VkQueue(unsafe.Pointer(queue)),
		(*C.VkPresentInfoKHR)(pPresentInfo)))

	return res, res.ToError()
}
