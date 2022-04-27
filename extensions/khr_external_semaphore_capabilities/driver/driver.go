package khr_external_semaphore_capabilities_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoGetPhysicalDeviceExternalSemaphorePropertiesKHR(PFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR fn, VkPhysicalDevice physicalDevice, VkPhysicalDeviceExternalSemaphoreInfoKHR *pExternalSemaphoreInfo, VkExternalSemaphorePropertiesKHR *pExternalSemaphoreProperties) {
	fn(physicalDevice, pExternalSemaphoreInfo, pExternalSemaphoreProperties);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_external_semaphore_capabilities

type Driver interface {
	VkGetPhysicalDeviceExternalSemaphorePropertiesKHR(physicalDevice driver.VkPhysicalDevice, pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfoKHR, pExternalSemaphoreProperties *VkExternalSemaphorePropertiesKHR)
}

type VkPhysicalDeviceExternalSemaphoreInfoKHR C.VkPhysicalDeviceExternalSemaphoreInfoKHR
type VkExternalSemaphorePropertiesKHR C.VkExternalSemaphorePropertiesKHR
type VkPhysicalDeviceIDPropertiesKHR C.VkPhysicalDeviceIDPropertiesKHR

type CDriver struct {
	coreDriver driver.Driver

	getPhysicalDeviceExternalSemaphoreProperties C.PFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		getPhysicalDeviceExternalSemaphoreProperties: (C.PFN_vkGetPhysicalDeviceExternalSemaphorePropertiesKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkGetPhysicalDeviceExternalSemaphorePropertiesKHR")))),
	}
}

func (d *CDriver) VkGetPhysicalDeviceExternalSemaphorePropertiesKHR(physicalDevice driver.VkPhysicalDevice, pExternalSemaphoreInfo *VkPhysicalDeviceExternalSemaphoreInfoKHR, pExternalSemaphoreProperties *VkExternalSemaphorePropertiesKHR) {
	if d.getPhysicalDeviceExternalSemaphoreProperties == nil {
		panic("attempt to call extension method vkGetPhysicalDeviceExternalSemaphorePropertiesKHR when extension not present")
	}

	C.cgoGetPhysicalDeviceExternalSemaphorePropertiesKHR(
		d.getPhysicalDeviceExternalSemaphoreProperties,
		C.VkPhysicalDevice(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceExternalSemaphoreInfoKHR)(pExternalSemaphoreInfo),
		(*C.VkExternalSemaphorePropertiesKHR)(pExternalSemaphoreProperties),
	)
}
