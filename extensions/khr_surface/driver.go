package khr_surface

//go:generate mockgen -source driver.go -destination ./mocks/driver.go -package mock_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"

VkResult cgoGetPhysicalDeviceSurfaceCapabilitiesKHR(PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR fn, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, VkSurfaceCapabilitiesKHR* pSurfaceCapabilities) {
	return fn(physicalDevice, surface, pSurfaceCapabilities);
}

VkResult cgoGetPhysicalDeviceSurfaceSupportKHR(PFN_vkGetPhysicalDeviceSurfaceSupportKHR fn, VkPhysicalDevice physicalDevice, uint32_t queueFamilyIndex, VkSurfaceKHR surface, VkBool32* pSupported) {
	return fn(physicalDevice, queueFamilyIndex, surface, pSupported);
}

void cgoDestroySurfaceKHR(PFN_vkDestroySurfaceKHR fn, VkInstance instance, VkSurfaceKHR surface, VkAllocationCallbacks* pAllocator) {
	fn(instance, surface, pAllocator);
}

VkResult cgoGetPhysicalDeviceSurfaceFormatsKHR(PFN_vkGetPhysicalDeviceSurfaceFormatsKHR fn,VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pSurfaceFormatCount, VkSurfaceFormatKHR* pSurfaceFormats) {
	return fn(physicalDevice, surface, pSurfaceFormatCount, pSurfaceFormats);
}

VkResult cgoGetPhysicalDeviceSurfacePresentModesKHR(PFN_vkGetPhysicalDeviceSurfacePresentModesKHR fn, VkPhysicalDevice physicalDevice, VkSurfaceKHR surface, uint32_t* pPresentModeCount, VkPresentModeKHR* pPresentModes) {
	return fn(physicalDevice, surface, pPresentModeCount, pPresentModes);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type khrSurfaceDriver struct {
	physicalSurfaceCapabilitiesFunc C.PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR
	physicalSurfaceSupportFunc      C.PFN_vkGetPhysicalDeviceSurfaceSupportKHR
	surfaceFormatsFunc              C.PFN_vkGetPhysicalDeviceSurfaceFormatsKHR
	presentModesFunc                C.PFN_vkGetPhysicalDeviceSurfacePresentModesKHR
	destroyFunc                     C.PFN_vkDestroySurfaceKHR
}

type VkSurfaceKHR C.VkSurfaceKHR
type VkSurfaceCapabilitiesKHR C.VkSurfaceCapabilitiesKHR
type VkSurfaceFormatKHR C.VkSurfaceFormatKHR
type VkPresentModeKHR C.VkPresentModeKHR

type Driver interface {
	VkDestroySurfaceKHR(instance driver.VkInstance, surface VkSurfaceKHR, pAllocator *driver.VkAllocationCallbacks)
	VkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice driver.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceCapabilities *VkSurfaceCapabilitiesKHR) (common.VkResult, error)
	VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice driver.VkPhysicalDevice, queueFamilyIndex driver.Uint32, surface VkSurfaceKHR, pSupported *driver.VkBool32) (common.VkResult, error)
	VkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice driver.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceFormatCount *driver.Uint32, pSurfaceFormats *VkSurfaceFormatKHR) (common.VkResult, error)
	VkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice driver.VkPhysicalDevice, surface VkSurfaceKHR, pPresentModeCount *driver.Uint32, pPresentModes *VkPresentModeKHR) (common.VkResult, error)
}

func CreateDriverFromCore(coreDriver driver.Driver) Driver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	physicalSurfaceCapabilitiesFunc := (C.PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceSurfaceCapabilitiesKHR"))))
	physicalSurfaceSupportFunc := (C.PFN_vkGetPhysicalDeviceSurfaceSupportKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceSurfaceSupportKHR"))))
	surfaceFormatsFunc := (C.PFN_vkGetPhysicalDeviceSurfaceFormatsKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceSurfaceFormatsKHR"))))
	presentModesFunc := (C.PFN_vkGetPhysicalDeviceSurfacePresentModesKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceSurfacePresentModesKHR"))))
	destroyFunc := (C.PFN_vkDestroySurfaceKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkDestroySurfaceKHR"))))

	return &khrSurfaceDriver{
		physicalSurfaceSupportFunc:      physicalSurfaceSupportFunc,
		physicalSurfaceCapabilitiesFunc: physicalSurfaceCapabilitiesFunc,
		surfaceFormatsFunc:              surfaceFormatsFunc,
		presentModesFunc:                presentModesFunc,
		destroyFunc:                     destroyFunc,
	}
}

func (d *khrSurfaceDriver) VkDestroySurfaceKHR(instance driver.VkInstance, surface VkSurfaceKHR, pAllocator *driver.VkAllocationCallbacks) {
	if d.destroyFunc == nil {
		panic("attempt to call extension method vkDestroySurfaceKHR when extension not present")
	}

	C.cgoDestroySurfaceKHR(d.destroyFunc,
		C.VkInstance(unsafe.Pointer(instance)),
		C.VkSurfaceKHR(surface),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (d *khrSurfaceDriver) VkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice driver.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceCapabilities *VkSurfaceCapabilitiesKHR) (common.VkResult, error) {
	if d.physicalSurfaceCapabilitiesFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceCapabilitiesKHR")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfaceCapabilitiesKHR(d.physicalSurfaceCapabilitiesFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(surface),
		(*C.VkSurfaceCapabilitiesKHR)(pSurfaceCapabilities)))

	return res, res.ToError()
}

func (d *khrSurfaceDriver) VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice driver.VkPhysicalDevice, queueFamilyIndex driver.Uint32, surface VkSurfaceKHR, pSupported *driver.VkBool32) (common.VkResult, error) {
	if d.physicalSurfaceSupportFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceSupportKHR")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfaceSupportKHR(d.physicalSurfaceSupportFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.uint32_t(queueFamilyIndex),
		C.VkSurfaceKHR(surface),
		(*C.VkBool32)(unsafe.Pointer(pSupported))))

	return res, res.ToError()
}

func (d *khrSurfaceDriver) VkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice driver.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceFormatCount *driver.Uint32, pSurfaceFormats *VkSurfaceFormatKHR) (common.VkResult, error) {
	if d.surfaceFormatsFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceFormatsKHR")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfaceFormatsKHR(d.surfaceFormatsFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(surface),
		(*C.uint32_t)(unsafe.Pointer(pSurfaceFormatCount)),
		(*C.VkSurfaceFormatKHR)(pSurfaceFormats)))
	return res, res.ToError()
}

func (d *khrSurfaceDriver) VkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice driver.VkPhysicalDevice, surface VkSurfaceKHR, pPresentModeCount *driver.Uint32, pPresentModes *VkPresentModeKHR) (common.VkResult, error) {
	if d.presentModesFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfacePresentModesKHR")
	}

	res := common.VkResult(C.cgoGetPhysicalDeviceSurfacePresentModesKHR(d.presentModesFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(surface),
		(*C.uint32_t)(unsafe.Pointer(pPresentModeCount)),
		(*C.VkPresentModeKHR)(pPresentModes)))

	return res, res.ToError()
}
