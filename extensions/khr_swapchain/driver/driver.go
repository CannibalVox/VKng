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

VkResult cgoSwapchainAcquireNextImage2KHR(PFN_vkAcquireNextImage2KHR fn, VkDevice device, VkAcquireNextImageInfoKHR *pAcquireInfo, uint32_t *pImageIndex) {
	return fn(device, pAcquireInfo, pImageIndex);
}

VkResult cgoSwapchainGetDeviceGroupPresentCapabilitiesKHR(PFN_vkGetDeviceGroupPresentCapabilitiesKHR fn, VkDevice device, VkDeviceGroupPresentCapabilitiesKHR *pDeviceGroupPresentCapabilities) {
	return fn(device, pDeviceGroupPresentCapabilities);
}

VkResult cgoSwapchainGetDeviceGroupSurfacePresentModesKHR(PFN_vkGetDeviceGroupSurfacePresentModesKHR fn, VkDevice device, VkSurfaceKHR surface, VkDeviceGroupPresentModeFlagsKHR *pModes) {
	return fn(device, surface, pModes);
}

VkResult cgoSwapchainGetPhysicalDevicePresentRectanglesKHR(PFN_vkGetPhysicalDevicePresentRectanglesKHR fn, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t *pRectCount, VkRect2D *pRects) {
	return fn(physicalDevice, surface, pRectCount, pRects);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	khr_surface_driver "github.com/CannibalVox/VKng/extensions/khr_surface/driver"
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

	acquireNextImageFunc                  C.PFN_vkAcquireNextImage2KHR
	getDeviceGroupPresentCapsFunc         C.PFN_vkGetDeviceGroupPresentCapabilitiesKHR
	getDeviceGroupSurfacePresentModesFunc C.PFN_vkGetDeviceGroupSurfacePresentModesKHR
	getPhysicalDevicePresentRectsFunc     C.PFN_vkGetPhysicalDevicePresentRectanglesKHR
}

type VkSwapchainKHR driver.VulkanHandle
type VkSwapchainCreateInfoKHR C.VkSwapchainCreateInfoKHR
type VkPresentInfoKHR C.VkPresentInfoKHR
type VkAcquireNextImageInfoKHR C.VkAcquireNextImageInfoKHR
type VkDeviceGroupPresentCapabilitiesKHR C.VkDeviceGroupPresentCapabilitiesKHR
type VkBindImageMemorySwapchainInfoKHR C.VkBindImageMemorySwapchainInfoKHR
type VkImageSwapchainCreateInfoKHR C.VkImageSwapchainCreateInfoKHR
type VkDeviceGroupPresentInfoKHR C.VkDeviceGroupPresentInfoKHR
type VkDeviceGroupSwapchainCreateInfoKHR C.VkDeviceGroupSwapchainCreateInfoKHR
type VkDeviceGroupPresentModeFlagsKHR C.VkDeviceGroupPresentModeFlagsKHR

type Driver interface {
	VkCreateSwapchainKHR(device driver.VkDevice, pCreateInfo *VkSwapchainCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pSwapchain *VkSwapchainKHR) (common.VkResult, error)
	VkDestroySwapchainKHR(device driver.VkDevice, swapchain VkSwapchainKHR, pAllocator *driver.VkAllocationCallbacks)
	VkGetSwapchainImagesKHR(device driver.VkDevice, swapchain VkSwapchainKHR, pSwapchainImageCount *driver.Uint32, pSwapchainImages *driver.VkImage) (common.VkResult, error)
	VkAcquireNextImageKHR(device driver.VkDevice, swapchain VkSwapchainKHR, timeout driver.Uint64, semaphore driver.VkSemaphore, fence driver.VkFence, pImageIndex *driver.Uint32) (common.VkResult, error)
	VkQueuePresentKHR(queue driver.VkQueue, pPresentInfo *VkPresentInfoKHR) (common.VkResult, error)
	VkAcquireNextImage2KHR(device driver.VkDevice, pAcquireInfo *VkAcquireNextImageInfoKHR, pImageIndex *driver.Uint32) (common.VkResult, error)
	VkGetDeviceGroupPresentCapabilitiesKHR(device driver.VkDevice, pDeviceGroupPresentCapabilities *VkDeviceGroupPresentCapabilitiesKHR) (common.VkResult, error)
	VkGetDeviceGroupSurfacePresentModesKHR(device driver.VkDevice, surface khr_surface_driver.VkSurfaceKHR, pModes *VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error)
	VkGetPhysicalDevicePresentRectanglesKHR(physicalDevice driver.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pRectCount *driver.Uint32, pRects *driver.VkRect2D) (common.VkResult, error)
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

		acquireNextImageFunc:                  (C.PFN_vkAcquireNextImage2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkAcquireNextImage2KHR")))),
		getDeviceGroupPresentCapsFunc:         (C.PFN_vkGetDeviceGroupPresentCapabilitiesKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetDeviceGroupPresentCapabilitiesKHR")))),
		getDeviceGroupSurfacePresentModesFunc: (C.PFN_vkGetDeviceGroupSurfacePresentModesKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetDeviceGroupSurfacePresentModesKHR")))),
		getPhysicalDevicePresentRectsFunc:     (C.PFN_vkGetPhysicalDevicePresentRectanglesKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDevicePresentRectanglesKHR")))),
	}
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

func (d *CDriver) VkAcquireNextImage2KHR(device driver.VkDevice, pAcquireInfo *VkAcquireNextImageInfoKHR, pImageIndex *driver.Uint32) (common.VkResult, error) {
	if d.acquireNextImageFunc == nil {
		panic("attempt to call extension method vkAcquireNextImage2KHR when extension (or core1.1) not present")
	}

	res := common.VkResult(C.cgoSwapchainAcquireNextImage2KHR(d.acquireNextImageFunc,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkAcquireNextImageInfoKHR)(pAcquireInfo),
		(*C.uint32_t)(pImageIndex),
	))

	return res, res.ToError()
}

func (d *CDriver) VkGetDeviceGroupPresentCapabilitiesKHR(device driver.VkDevice, pDeviceGroupPresentCapabilities *VkDeviceGroupPresentCapabilitiesKHR) (common.VkResult, error) {
	if d.getDeviceGroupPresentCapsFunc == nil {
		panic("attempt to call extension method vkGetDeviceGroupPresentCapabilitiesKHR when extension (or core1.1) not present")
	}

	res := common.VkResult(C.cgoSwapchainGetDeviceGroupPresentCapabilitiesKHR(d.getDeviceGroupPresentCapsFunc,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkDeviceGroupPresentCapabilitiesKHR)(pDeviceGroupPresentCapabilities),
	))

	return res, res.ToError()
}

func (d *CDriver) VkGetDeviceGroupSurfacePresentModesKHR(device driver.VkDevice, surface khr_surface_driver.VkSurfaceKHR, pModes *VkDeviceGroupPresentModeFlagsKHR) (common.VkResult, error) {
	if d.getDeviceGroupSurfacePresentModesFunc == nil {
		panic("attempt to call extension method vkGetDeviceGroupSurfacePresentModesKHR when extension (or core1.1) not present")
	}

	res := common.VkResult(C.cgoSwapchainGetDeviceGroupSurfacePresentModesKHR(d.getDeviceGroupSurfacePresentModesFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.VkDeviceGroupPresentModeFlagsKHR)(pModes),
	))

	return res, res.ToError()
}

func (d *CDriver) VkGetPhysicalDevicePresentRectanglesKHR(physicalDevice driver.VkPhysicalDevice, surface khr_surface_driver.VkSurfaceKHR, pRectCount *driver.Uint32, pRects *driver.VkRect2D) (common.VkResult, error) {
	if d.getPhysicalDevicePresentRectsFunc == nil {
		panic("attempt to call extension method vkGetDeviceGroupSurfacePresentModesKHR when extension (or core1.1) not present")
	}

	res := common.VkResult(C.cgoSwapchainGetPhysicalDevicePresentRectanglesKHR(d.getPhysicalDevicePresentRectsFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(unsafe.Pointer(surface)),
		(*C.uint32_t)(pRectCount),
		(*C.VkRect2D)(unsafe.Pointer(pRects)),
	))

	return res, res.ToError()
}
