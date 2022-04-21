package khr_external_fence_capabilities_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetPhysicalDeviceExternalFencePropertiesKHR(PFN_vkGetPhysicalDeviceExternalFencePropertiesKHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalFenceInfoKHR *pExternalFenceInfo, VkExternalFencePropertiesKHR *pExternalFenceProperties) {
	fn(physicalDevice, pExternalFenceInfo, pExternalFenceProperties);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_external_fence_capabilities

type Driver interface {
	VkGetPhysicalDeviceExternalFencePropertiesKHR(physicalDevice driver.VkPhysicalDevice, pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfoKHR, pExternalFenceProperties *VkExternalFencePropertiesKHR)
}

type VkPhysicalDeviceExternalFenceInfoKHR C.VkPhysicalDeviceExternalFenceInfoKHR
type VkExternalFencePropertiesKHR C.VkExternalFencePropertiesKHR
type VkPhysicalDeviceIDPropertiesKHR C.VkPhysicalDeviceIDPropertiesKHR

type CDriver struct {
	coreDriver driver.Driver

	getPhysicalDeviceExternalFenceProperties C.PFN_vkGetPhysicalDeviceExternalFencePropertiesKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		getPhysicalDeviceExternalFenceProperties: (C.PFN_vkGetPhysicalDeviceExternalFencePropertiesKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceExternalFencePropertiesKHR")))),
	}
}

func (d *CDriver) VkGetPhysicalDeviceExternalFencePropertiesKHR(physicalDevice driver.VkPhysicalDevice, pExternalFenceInfo *VkPhysicalDeviceExternalFenceInfoKHR, pExternalFenceProperties *VkExternalFencePropertiesKHR) {
	if d.getPhysicalDeviceExternalFenceProperties == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceExternalFencePropertiesKHR when extension not present")
	}

	C.cgoGetPhysicalDeviceExternalFencePropertiesKHR(
		d.getPhysicalDeviceExternalFenceProperties,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceExternalFenceInfoKHR)(pExternalFenceInfo),
		(*C.VkExternalFencePropertiesKHR)(pExternalFenceProperties),
	)
}
