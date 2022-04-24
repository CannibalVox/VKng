package khr_external_memory_capabilities_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetPhysicalDeviceExternalBufferPropertiesKHR(PFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalBufferInfoKHR *pExternalBufferInfo, VkExternalBufferPropertiesKHR *pExternalBufferProperties) {
	fn(physicalDevice, pExternalBufferInfo, pExternalBufferProperties);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_external_memory_capabilities

type Driver interface {
	VkGetPhysicalDeviceExternalBufferPropertiesKHR(physicalDevice driver.VkPhysicalDevice, pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfoKHR, pExternalBufferProperties *VkExternalBufferPropertiesKHR)
}

type VkPhysicalDeviceExternalBufferInfoKHR C.VkPhysicalDeviceExternalBufferInfoKHR
type VkExternalBufferPropertiesKHR C.VkExternalBufferPropertiesKHR
type VkExternalMemoryPropertiesKHR C.VkExternalMemoryPropertiesKHR
type VkExternalImageFormatPropertiesKHR C.VkExternalImageFormatPropertiesKHR
type VkPhysicalDeviceExternalImageFormatInfoKHR C.VkPhysicalDeviceExternalImageFormatInfoKHR
type VkPhysicalDeviceIDPropertiesKHR C.VkPhysicalDeviceIDPropertiesKHR

type CDriver struct {
	coreDriver driver.Driver

	getPhysicalDeviceExternalBufferProperties C.PFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		getPhysicalDeviceExternalBufferProperties: (C.PFN_vkGetPhysicalDeviceExternalBufferPropertiesKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceExternalBufferPropertiesKHR")))),
	}
}

func (d *CDriver) VkGetPhysicalDeviceExternalBufferPropertiesKHR(physicalDevice driver.VkPhysicalDevice, pExternalBufferInfo *VkPhysicalDeviceExternalBufferInfoKHR, pExternalBufferProperties *VkExternalBufferPropertiesKHR) {
	if d.getPhysicalDeviceExternalBufferProperties == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceExternalBufferPropertiesKHR when extension not present")
	}

	C.cgoGetPhysicalDeviceExternalBufferPropertiesKHR(
		d.getPhysicalDeviceExternalBufferProperties,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceExternalBufferInfoKHR)(pExternalBufferInfo),
		(*C.VkExternalBufferPropertiesKHR)(pExternalBufferProperties),
	)
}
