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
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SWAPCHAIN_EXTENSION_NAME

type khrSwapchainDriver struct {
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
	VkCreateSwapchainKHR(device core.VkDevice, pCreateInfo *VkSwapchainCreateInfoKHR, pAllocator *core.VkAllocationCallbacks, pSwapchain *VkSwapchainKHR) (core.VkResult, error)
	VkDestroySwapchainKHR(device core.VkDevice, swapchain VkSwapchainKHR, pAllocator *core.VkAllocationCallbacks) error
	VkGetSwapchainImagesKHR(device core.VkDevice, swapchain VkSwapchainKHR, pSwapchainImageCount *core.Uint32, pSwapchainImages *core.VkImage) (core.VkResult, error)
	VkAcquireNextImageKHR(device core.VkDevice, swapchain VkSwapchainKHR, timeout core.Uint64, semaphore core.VkSemaphore, fence core.VkFence, pImageIndex *core.Uint32) (core.VkResult, error)
	VkQueuePresentKHR(queue core.VkQueue, pPresentInfo *VkPresentInfoKHR) (core.VkResult, error)
}

func createDriverFromCore(driver core.Driver) *khrSwapchainDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &khrSwapchainDriver{
		createFunc:       (C.PFN_vkCreateSwapchainKHR)(driver.LoadProcAddr((*core.Char)(arena.CString("vkCreateSwapchainKHR")))),
		destroyFunc:      (C.PFN_vkDestroySwapchainKHR)(driver.LoadProcAddr((*core.Char)(arena.CString("vkDestroySwapchainKHR")))),
		getImagesFunc:    (C.PFN_vkGetSwapchainImagesKHR)(driver.LoadProcAddr((*core.Char)(arena.CString("vkGetSwapchainImagesKHR")))),
		acquireNextFunc:  (C.PFN_vkAcquireNextImageKHR)(driver.LoadProcAddr((*core.Char)(arena.CString("vkAcquireNextImageKHR")))),
		queuePresentFunc: (C.PFN_vkQueuePresentKHR)(driver.LoadProcAddr((*core.Char)(arena.CString("vkQueuePresentKHR")))),
	}
}

func (d *khrSwapchainDriver) VkCreateSwapchainKHR(device core.VkDevice, pCreateInfo *VkSwapchainCreateInfoKHR, pAllocator *core.VkAllocationCallbacks, pSwapchain *VkSwapchainKHR) (core.VkResult, error) {
	if d.createFunc == nil {
		return core.VKErrorUnknown, errors.New("attempt to call extension method vkCreateSwapchainKHR when extension not present")
	}

	res := core.VkResult(C.cgoCreateSwapchainKHR(d.createFunc,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSwapchainCreateInfoKHR)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkSwapchainKHR)(pSwapchain)))

	return res, res.ToError()
}

func (d *khrSwapchainDriver) VkDestroySwapchainKHR(device core.VkDevice, swapchain VkSwapchainKHR, pAllocator *core.VkAllocationCallbacks) error {
	if d.destroyFunc == nil {
		return errors.New("attempt to call extension method vkDestroySwapchainKHR when extension not present")
	}

	C.cgoDestroySwapchainKHR(d.destroyFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(swapchain),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))

	return nil
}

func (d *khrSwapchainDriver) VkGetSwapchainImagesKHR(device core.VkDevice, swapchain VkSwapchainKHR, pSwapchainImageCount *core.Uint32, pSwapchainImages *core.VkImage) (core.VkResult, error) {
	if d.getImagesFunc == nil {
		return core.VKErrorUnknown, errors.New("attempt to call extension method vkGetSwapchainImagesKHR when extension not present")
	}

	res := core.VkResult(C.cgoGetSwapchainImagesKHR(d.getImagesFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(swapchain),
		(*C.uint32_t)(unsafe.Pointer(pSwapchainImageCount)),
		(*C.VkImage)(unsafe.Pointer(pSwapchainImages))))

	return res, res.ToError()
}

func (d *khrSwapchainDriver) VkAcquireNextImageKHR(device core.VkDevice, swapchain VkSwapchainKHR, timeout core.Uint64, semaphore core.VkSemaphore, fence core.VkFence, pImageIndex *core.Uint32) (core.VkResult, error) {
	if d.acquireNextFunc == nil {
		return core.VKErrorUnknown, errors.New("attempt to call extension method vkAcquireNextImageKHR when extension not present")
	}

	res := core.VkResult(C.cgoAcquireNextImageKHR(d.acquireNextFunc,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSwapchainKHR(swapchain),
		C.uint64_t(timeout),
		C.VkSemaphore(unsafe.Pointer(semaphore)),
		C.VkFence(unsafe.Pointer(fence)),
		(*C.uint32_t)(unsafe.Pointer(pImageIndex)),
	))

	return res, res.ToError()
}

func (d *khrSwapchainDriver) VkQueuePresentKHR(queue core.VkQueue, pPresentInfo *VkPresentInfoKHR) (core.VkResult, error) {
	if d.queuePresentFunc == nil {
		return core.VKErrorUnknown, errors.New("attempt to call extension method vkQueuePresentKHR when extension not present")
	}

	res := core.VkResult(C.cgoQueuePresentKHR(d.queuePresentFunc,
		C.VkQueue(unsafe.Pointer(queue)),
		(*C.VkPresentInfoKHR)(pPresentInfo)))

	return res, res.ToError()
}

type khrSwapchainLoader struct {
	driver Driver
}

type Loader interface {
	CreateSwapchain(device core.Device, options *CreationOptions) (Swapchain, core.VkResult, error)
}

func CreateLoaderFromDevice(device core.Device) Loader {
	return &khrSwapchainLoader{
		driver: createDriverFromCore(device.Driver()),
	}
}

func CreateLoaderFromDriver(driver Driver) Loader {
	return &khrSwapchainLoader{
		driver: driver,
	}
}

func (l *khrSwapchainLoader) CreateSwapchain(device core.Device, options *CreationOptions) (Swapchain, core.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	var swapchain VkSwapchainKHR

	res, err := l.driver.VkCreateSwapchainKHR(device.Handle(), (*VkSwapchainCreateInfoKHR)(createInfo), nil, &swapchain)
	if err != nil {
		return nil, res, err
	}

	return &vulkanSwapchain{
		handle: swapchain,
		device: device.Handle(),
		driver: l.driver,
	}, res, nil
}
