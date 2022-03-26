package khr_get_physical_device_properties2_driver

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_get_physical_device_properties2

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetPhysicalDeviceFeatures2KHR(PFN_vkGetPhysicalDeviceFeatures2KHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceFeatures2KHR *pFeatures) {
	fn(physicalDevice, pFeatures);
}

void cgoGetPhysicalDeviceFormatProperties2KHR(PFN_vkGetPhysicalDeviceFormatProperties2KHR fn, VkPhysicalDevice physicalDevice, VkFormat format, VkFormatProperties2KHR *pFormatProperties) {
	fn(physicalDevice, format, pFormatProperties);
}

VkResult cgoGetPhysicalDeviceImageFormatProperties2KHR(PFN_vkGetPhysicalDeviceImageFormatProperties2KHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceImageFormatInfo2KHR *pImageFormatInfo, VkImageFormatProperties2KHR *pImageFormatProperties) {
	return fn(physicalDevice, pImageFormatInfo, pImageFormatProperties);
}

void cgoGetPhysicalDeviceMemoryProperties2KHR(PFN_vkGetPhysicalDeviceMemoryProperties2KHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceMemoryProperties2KHR *pMemoryProperties) {
	fn(physicalDevice, pMemoryProperties);
}

void cgoGetPhysicalDeviceProperties2KHR(PFN_vkGetPhysicalDeviceProperties2KHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceProperties2KHR *pProperties) {
	fn(physicalDevice, pProperties);
}

void cgoGetPhysicalDeviceQueueFamilyProperties2KHR(PFN_vkGetPhysicalDeviceQueueFamilyProperties2KHR fn, VkPhysicalDevice physicalDevice, uint32_t *pQueueFamilyPropertyCount, VkQueueFamilyProperties2KHR *pQueueFamilyProperties) {
	fn(physicalDevice, pQueueFamilyPropertyCount, pQueueFamilyProperties);
}

void cgoGetPhysicalDeviceSparseImageFormatProperties2KHR(PFN_vkGetPhysicalDeviceSparseImageFormatProperties2KHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceSparseImageFormatInfo2KHR *pFormatInfo, uint32_t *pPropertyCount, VkSparseImageFormatProperties2KHR *pProperties) {
	fn(physicalDevice, pFormatInfo, pPropertyCount, pProperties);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type Driver interface {
	VkGetPhysicalDeviceFeatures2KHR(physicalDevice driver.VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures2KHR)
	VkGetPhysicalDeviceFormatProperties2KHR(physicalDevice driver.VkPhysicalDevice, format driver.VkFormat, pFormatProperties *VkFormatProperties2KHR)
	VkGetPhysicalDeviceImageFormatProperties2KHR(physicalDevice driver.VkPhysicalDevice, pImageFormatInfo *VkPhysicalDeviceImageFormatInfo2KHR, pImageFormatProperties *VkImageFormatProperties2KHR) (common.VkResult, error)
	VkGetPhysicalDeviceMemoryProperties2KHR(physicalDevice driver.VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties2KHR)
	VkGetPhysicalDeviceProperties2KHR(physicalDevice driver.VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties2KHR)
	VkGetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice driver.VkPhysicalDevice, pQueueFamilyPropertyCount *driver.Uint32, pQueueFamilyProperties *VkQueueFamilyProperties2KHR)
	VkGetPhysicalDeviceSparseImageFormatProperties2KHR(physicalDevice driver.VkPhysicalDevice, pFormatInfo *VkPhysicalDeviceSparseImageFormatInfo2KHR, pPropertyCount *driver.Uint32, pProperties *VkSparseImageFormatProperties2KHR)
}

type VkPhysicalDeviceFeatures2KHR C.VkPhysicalDeviceFeatures2KHR
type VkFormatProperties2KHR C.VkFormatProperties2KHR
type VkPhysicalDeviceImageFormatInfo2KHR C.VkPhysicalDeviceImageFormatInfo2KHR
type VkImageFormatProperties2KHR C.VkImageFormatProperties2KHR
type VkPhysicalDeviceMemoryProperties2KHR C.VkPhysicalDeviceMemoryProperties2KHR
type VkPhysicalDeviceProperties2KHR C.VkPhysicalDeviceProperties2KHR
type VkQueueFamilyProperties2KHR C.VkQueueFamilyProperties2KHR
type VkPhysicalDeviceSparseImageFormatInfo2KHR C.VkPhysicalDeviceSparseImageFormatInfo2KHR
type VkSparseImageFormatProperties2KHR C.VkSparseImageFormatProperties2KHR

type CDriver struct {
	coreDriver driver.Driver

	getPhysicalDeviceFeatures2                    C.PFN_vkGetPhysicalDeviceFeatures2KHR
	getPhysicalDeviceFormatProperties2            C.PFN_vkGetPhysicalDeviceFormatProperties2KHR
	getPhysicalDeviceImageFormatProperties2       C.PFN_vkGetPhysicalDeviceImageFormatProperties2KHR
	getPhysicalDeviceMemoryProperties2            C.PFN_vkGetPhysicalDeviceMemoryProperties2KHR
	getPhysicalDeviceProperties2                  C.PFN_vkGetPhysicalDeviceProperties2KHR
	getPhysicalDeviceQueueFamilyProperties2       C.PFN_vkGetPhysicalDeviceQueueFamilyProperties2KHR
	getPhysicalDeviceSparseImageFormatProperties2 C.PFN_vkGetPhysicalDeviceSparseImageFormatProperties2KHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		getPhysicalDeviceFeatures2:                    (C.PFN_vkGetPhysicalDeviceFeatures2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceFeatures2KHR")))),
		getPhysicalDeviceFormatProperties2:            (C.PFN_vkGetPhysicalDeviceFormatProperties2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceFormatProperties2KHR")))),
		getPhysicalDeviceImageFormatProperties2:       (C.PFN_vkGetPhysicalDeviceImageFormatProperties2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceImageFormatProperties2KHR")))),
		getPhysicalDeviceMemoryProperties2:            (C.PFN_vkGetPhysicalDeviceMemoryProperties2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceMemoryProperties2KHR")))),
		getPhysicalDeviceProperties2:                  (C.PFN_vkGetPhysicalDeviceProperties2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceProperties2KHR")))),
		getPhysicalDeviceQueueFamilyProperties2:       (C.PFN_vkGetPhysicalDeviceQueueFamilyProperties2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceQueueFamilyProperties2KHR")))),
		getPhysicalDeviceSparseImageFormatProperties2: (C.PFN_vkGetPhysicalDeviceSparseImageFormatProperties2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceSparseImageFormatProperties2KHR")))),
	}
}

func (d *CDriver) VkGetPhysicalDeviceFeatures2KHR(physicalDevice driver.VkPhysicalDevice, pFeatures *VkPhysicalDeviceFeatures2KHR) {
	C.cgoGetPhysicalDeviceFeatures2KHR(d.getPhysicalDeviceFeatures2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceFeatures2KHR)(pFeatures),
	)
}

func (d *CDriver) VkGetPhysicalDeviceFormatProperties2KHR(physicalDevice driver.VkPhysicalDevice, format driver.VkFormat, pFormatProperties *VkFormatProperties2KHR) {
	C.cgoGetPhysicalDeviceFormatProperties2KHR(d.getPhysicalDeviceFormatProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		C.VkFormat(format),
		(*C.VkFormatProperties2KHR)(pFormatProperties),
	)
}

func (d *CDriver) VkGetPhysicalDeviceImageFormatProperties2KHR(physicalDevice driver.VkPhysicalDevice, pImageFormatInfo *VkPhysicalDeviceImageFormatInfo2KHR, pImageFormatProperties *VkImageFormatProperties2KHR) (common.VkResult, error) {
	res := common.VkResult(C.cgoGetPhysicalDeviceImageFormatProperties2KHR(d.getPhysicalDeviceImageFormatProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceImageFormatInfo2KHR)(pImageFormatInfo),
		(*C.VkImageFormatProperties2KHR)(pImageFormatProperties)))
	return res, res.ToError()
}

func (d *CDriver) VkGetPhysicalDeviceMemoryProperties2KHR(physicalDevice driver.VkPhysicalDevice, pMemoryProperties *VkPhysicalDeviceMemoryProperties2KHR) {
	C.cgoGetPhysicalDeviceMemoryProperties2KHR(d.getPhysicalDeviceMemoryProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceMemoryProperties2KHR)(pMemoryProperties))
}

func (d *CDriver) VkGetPhysicalDeviceProperties2KHR(physicalDevice driver.VkPhysicalDevice, pProperties *VkPhysicalDeviceProperties2KHR) {
	C.cgoGetPhysicalDeviceProperties2KHR(d.getPhysicalDeviceProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceProperties2KHR)(pProperties))
}

func (d *CDriver) VkGetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice driver.VkPhysicalDevice, pQueueFamilyPropertyCount *driver.Uint32, pQueueFamilyProperties *VkQueueFamilyProperties2KHR) {
	C.cgoGetPhysicalDeviceQueueFamilyProperties2KHR(d.getPhysicalDeviceQueueFamilyProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.uint32_t)(pQueueFamilyPropertyCount),
		(*C.VkQueueFamilyProperties2KHR)(pQueueFamilyProperties))
}

func (d *CDriver) VkGetPhysicalDeviceSparseImageFormatProperties2KHR(physicalDevice driver.VkPhysicalDevice, pFormatInfo *VkPhysicalDeviceSparseImageFormatInfo2KHR, pPropertyCount *driver.Uint32, pProperties *VkSparseImageFormatProperties2KHR) {
	C.cgoGetPhysicalDeviceSparseImageFormatProperties2KHR(d.getPhysicalDeviceImageFormatProperties2,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceSparseImageFormatInfo2KHR)(pFormatInfo),
		(*C.uint32_t)(pPropertyCount),
		(*C.VkSparseImageFormatProperties2KHR)(pProperties))
}
