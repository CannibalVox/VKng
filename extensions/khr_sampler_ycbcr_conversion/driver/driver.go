package khr_sampler_ycbcr_conversion_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoCreateSamplerYcbcrConversionKHR(PFN_vkCreateSamplerYcbcrConversionKHR fn, VkDevice device, VkSamplerYcbcrConversionCreateInfoKHR *pCreateInfo, VkAllocationCallbacks *pAllocator, VkSamplerYcbcrConversionKHR *pYcbcrConversion) {
	return fn(device, pCreateInfo, pAllocator, pYcbcrConversion);
}

void cgoDestroySamplerYcbcrConversionKHR(PFN_vkDestroySamplerYcbcrConversionKHR fn, VkDevice device, VkSamplerYcbcrConversionKHR ycbcrConversion, VkAllocationCallbacks *pAllocator) {
	fn(device, ycbcrConversion, pAllocator);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_sampler_ycbcr_conversion

type Driver interface {
	VkCreateSamplerYcbcrConversionKHR(device driver.VkDevice, pCreateInfo *VkSamplerYcbcrConversionCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pYcbcrConversion *VkSamplerYcbcrConversionKHR) (common.VkResult, error)
	VkDestroySamplerYcbcrConversionKHR(device driver.VkDevice, ycbcrConversion VkSamplerYcbcrConversionKHR, pAllocator *driver.VkAllocationCallbacks)
}

type VkSamplerYcbcrConversionKHR driver.VulkanHandle
type VkSamplerYcbcrConversionCreateInfoKHR C.VkSamplerYcbcrConversionCreateInfoKHR
type VkBindImagePlaneMemoryInfoKHR C.VkBindImagePlaneMemoryInfoKHR
type VkSamplerYcbcrConversionImageFormatPropertiesKHR C.VkSamplerYcbcrConversionImageFormatPropertiesKHR
type VkImagePlaneMemoryRequirementsInfoKHR C.VkImagePlaneMemoryRequirementsInfoKHR
type VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR C.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR
type VkSamplerYcbcrConversionInfoKHR C.VkSamplerYcbcrConversionInfoKHR

type CDriver struct {
	coreDriver driver.Driver

	createSamplerYcbcrConversion  C.PFN_vkCreateSamplerYcbcrConversionKHR
	destroySamplerYcbcrConversion C.PFN_vkDestroySamplerYcbcrConversionKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		createSamplerYcbcrConversion:  (C.PFN_vkCreateSamplerYcbcrConversionKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCreateSamplerYcbcrConversionKHR")))),
		destroySamplerYcbcrConversion: (C.PFN_vkDestroySamplerYcbcrConversionKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkDestroySamplerYcbcrConversionKHR")))),
	}
}

func (d *CDriver) VkCreateSamplerYcbcrConversionKHR(device driver.VkDevice, pCreateInfo *VkSamplerYcbcrConversionCreateInfoKHR, pAllocator *driver.VkAllocationCallbacks, pYcbcrConversion *VkSamplerYcbcrConversionKHR) (common.VkResult, error) {
	if d.createSamplerYcbcrConversion == nil {
		panic("attempt to call extension method vkCreateSamplerYcbcrConversionKHR when extension not present")
	}

	res := common.VkResult(C.cgoCreateSamplerYcbcrConversionKHR(
		d.createSamplerYcbcrConversion,
		C.VkDevice(unsafe.Pointer(device)),
		(*C.VkSamplerYcbcrConversionCreateInfoKHR)(pCreateInfo),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
		(*C.VkSamplerYcbcrConversionKHR)(unsafe.Pointer(pYcbcrConversion)),
	))
	return res, res.ToError()
}

func (d *CDriver) VkDestroySamplerYcbcrConversionKHR(device driver.VkDevice, ycbcrConversion VkSamplerYcbcrConversionKHR, pAllocator *driver.VkAllocationCallbacks) {
	if d.destroySamplerYcbcrConversion == nil {
		panic("attempt to call extension method vkDestroySamplerYcbcrConversionKHR when extension not present")
	}

	C.cgoDestroySamplerYcbcrConversionKHR(
		d.destroySamplerYcbcrConversion,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkSamplerYcbcrConversionKHR(unsafe.Pointer(ycbcrConversion)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)),
	)
}
