package khr_swapchain

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

VkResult cgoQueuePresentKHR(PFN_vkQueuePresentKHR fn, VkQueue queue, VkPresentInfoKHR* pPresentInfo) {
	return fn(queue, pPresentInfo);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SWAPCHAIN_EXTENSION_NAME

type CDriver struct {
	driver           driver.Driver
	createFunc       C.PFN_vkCreateSwapchainKHR
	destroyFunc      C.PFN_vkDestroySwapchainKHR
	getImagesFunc    C.PFN_vkGetSwapchainImagesKHR
	acquireNextFunc  C.PFN_vkAcquireNextImageKHR
	queuePresentFunc C.PFN_vkQueuePresentKHR
}

type VkSwapchainKHR C.VkSwapchainKHR
type VkSwapchainCreateInfoKHR C.VkSwapchainCreateInfoKHR
type VkPresentInfoKHR C.VkPresentInfoKHR
type Driver interface {
	coreDriver() driver.Driver
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

func (d *CDriver) coreDriver() driver.Driver {
	return d.driver
}

func (d *CDriver) VkCreateSwapchainKHR(device driver.VkDevice, pCreateInfo *VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *VkSwapchainKHR) (common.VkResult, error) {
	if d.createFunc == nil {
		panic("attempt to call extension method vkCreateSwapchainKHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateSwapchainKHR(d.createFunc,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSwapchainCreateInfoKHR)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkSwapchainKHR)(pSwapchain)))

	return res, res.ToError()
}

func (d *CDriver) VkDestroySwapchainKHR(device driver.VkDevice, swapchain VkSwapchainKHR, pAllocator *driver.VkAllocationCallbacks) {
	if d.destroyFunc == nil {
		panic("attempt to call extension method vkDestroySwapchainKHR when extension not present")
	}

	C.cgoDestroySwapchainKHR(d.destroyFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(swapchain),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (d *CDriver) VkGetSwapchainImagesKHR(device driver.VkDevice, swapchain VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error) {
	if d.getImagesFunc == nil {
		panic("attempt to call extension method vkGetSwapchainImagesKHR when extension not present")
	}

	res := common.VkResult(C.cgoGetSwapchainImagesKHR(d.getImagesFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(swapchain),
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
		C.VkSwapchainKHR(swapchain),
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

type VulkanExtension struct {
	driver Driver
}

type Extension interface {
	CreateSwapchain(device core1_0.Device, allocation *driver.AllocationCallbacks, options *CreationOptions) (Swapchain, common.VkResult, error)
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	return &VulkanExtension{
		driver: CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (l *VulkanExtension) CreateSwapchain(device core1_0.Device, allocation *driver.AllocationCallbacks, options *CreationOptions) (Swapchain, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var swapchain VkSwapchainKHR

	res, err := l.driver.VkCreateSwapchainKHR(device.Handle(), (*VkSwapchainCreateInfoKHR)(createInfo), allocation.Handle(), &swapchain)
	if err != nil {
		return nil, res, err
	}

	return &vulkanSwapchain{
		handle: swapchain,
		device: device.Handle(),
		driver: l.driver,
	}, res, nil
}
