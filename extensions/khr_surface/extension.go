package khr_surface

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_surface

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
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const ExtensionName = C.VK_KHR_SURFACE_EXTENSION_NAME

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
	VkDestroySurfaceKHR(instance core.VkInstance, surface VkSurfaceKHR, pAllocator *core.VkAllocationCallbacks)
	VkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice core.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceCapabilities *VkSurfaceCapabilitiesKHR) (core.VkResult, error)
	VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice core.VkPhysicalDevice, queueFamilyIndex core.Uint32, surface VkSurfaceKHR, pSupported *core.VkBool32) (core.VkResult, error)
	VkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice core.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceFormatCount *core.Uint32, pSurfaceFormats *VkSurfaceFormatKHR) (core.VkResult, error)
	VkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice core.VkPhysicalDevice, surface VkSurfaceKHR, pPresentModeCount *core.Uint32, pPresentModes *VkPresentModeKHR) (core.VkResult, error)
}

func CreateDriverFromCore(driver core.Driver) Driver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	physicalSurfaceCapabilitiesFunc := (C.PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR)(driver.LoadProcAddr((*core.Char)(arena.CString("vkGetPhysicalDeviceSurfaceCapabilitiesKHR"))))
	physicalSurfaceSupportFunc := (C.PFN_vkGetPhysicalDeviceSurfaceSupportKHR)(driver.LoadProcAddr((*core.Char)(arena.CString("vkGetPhysicalDeviceSurfaceSupportKHR"))))
	surfaceFormatsFunc := (C.PFN_vkGetPhysicalDeviceSurfaceFormatsKHR)(driver.LoadProcAddr((*core.Char)(arena.CString("vkGetPhysicalDeviceSurfaceFormatsKHR"))))
	presentModesFunc := (C.PFN_vkGetPhysicalDeviceSurfacePresentModesKHR)(driver.LoadProcAddr((*core.Char)(arena.CString("vkGetPhysicalDeviceSurfacePresentModesKHR"))))
	destroyFunc := (C.PFN_vkDestroySurfaceKHR)(driver.LoadProcAddr((*core.Char)(arena.CString("vkDestroySurfaceKHR"))))

	return &khrSurfaceDriver{
		physicalSurfaceSupportFunc:      physicalSurfaceSupportFunc,
		physicalSurfaceCapabilitiesFunc: physicalSurfaceCapabilitiesFunc,
		surfaceFormatsFunc:              surfaceFormatsFunc,
		presentModesFunc:                presentModesFunc,
		destroyFunc:                     destroyFunc,
	}
}

func (d *khrSurfaceDriver) VkDestroySurfaceKHR(instance core.VkInstance, surface VkSurfaceKHR, pAllocator *core.VkAllocationCallbacks) {
	if d.destroyFunc == nil {
		panic("attempt to call extension method vkDestroySurfaceKHR when extension not present")
	}

	C.cgoDestroySurfaceKHR(d.destroyFunc,
		C.VkInstance(unsafe.Pointer(instance)),
		C.VkSurfaceKHR(surface),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)))
}

func (d *khrSurfaceDriver) VkGetPhysicalDeviceSurfaceCapabilitiesKHR(physicalDevice core.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceCapabilities *VkSurfaceCapabilitiesKHR) (core.VkResult, error) {
	if d.physicalSurfaceCapabilitiesFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceCapabilitiesKHR")
	}

	res := core.VkResult(C.cgoGetPhysicalDeviceSurfaceCapabilitiesKHR(d.physicalSurfaceCapabilitiesFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(surface),
		(*C.VkSurfaceCapabilitiesKHR)(pSurfaceCapabilities)))

	return res, res.ToError()
}

func (d *khrSurfaceDriver) VkGetPhysicalDeviceSurfaceSupportKHR(physicalDevice core.VkPhysicalDevice, queueFamilyIndex core.Uint32, surface VkSurfaceKHR, pSupported *core.VkBool32) (core.VkResult, error) {
	if d.physicalSurfaceSupportFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceSupportKHR")
	}

	res := core.VkResult(C.cgoGetPhysicalDeviceSurfaceSupportKHR(d.physicalSurfaceSupportFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.uint32_t(queueFamilyIndex),
		C.VkSurfaceKHR(surface),
		(*C.VkBool32)(unsafe.Pointer(pSupported))))

	return res, res.ToError()
}

func (d *khrSurfaceDriver) VkGetPhysicalDeviceSurfaceFormatsKHR(physicalDevice core.VkPhysicalDevice, surface VkSurfaceKHR, pSurfaceFormatCount *core.Uint32, pSurfaceFormats *VkSurfaceFormatKHR) (core.VkResult, error) {
	if d.surfaceFormatsFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfaceFormatsKHR")
	}

	res := core.VkResult(C.cgoGetPhysicalDeviceSurfaceFormatsKHR(d.surfaceFormatsFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(surface),
		(*C.uint32_t)(unsafe.Pointer(pSurfaceFormatCount)),
		(*C.VkSurfaceFormatKHR)(pSurfaceFormats)))
	return res, res.ToError()
}

func (d *khrSurfaceDriver) VkGetPhysicalDeviceSurfacePresentModesKHR(physicalDevice core.VkPhysicalDevice, surface VkSurfaceKHR, pPresentModeCount *core.Uint32, pPresentModes *VkPresentModeKHR) (core.VkResult, error) {
	if d.presentModesFunc == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceSurfacePresentModesKHR")
	}

	res := core.VkResult(C.cgoGetPhysicalDeviceSurfacePresentModesKHR(d.presentModesFunc,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkSurfaceKHR(surface),
		(*C.uint32_t)(unsafe.Pointer(pPresentModeCount)),
		(*C.VkPresentModeKHR)(pPresentModes)))

	return res, res.ToError()
}
