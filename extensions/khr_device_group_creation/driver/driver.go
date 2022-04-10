package khr_device_group_creation_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoEnumeratePhysicalDeviceGroupsKHR(PFN_vkEnumeratePhysicalDeviceGroupsKHR fn, VkInstance instance, uint32_t *pPhysicalDeviceGroupCount, VkPhysicalDeviceGroupPropertiesKHR *pPhysicalDeviceGroupProperties) {
	return fn(instance, pPhysicalDeviceGroupCount, pPhysicalDeviceGroupProperties);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_device_group_creation

type Driver interface {
	VkEnumeratePhysicalDeviceGroupsKHR(instance driver.VkInstance, pPhysicalDeviceGroupCount *driver.Uint32, pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error)
}

type VkPhysicalDeviceGroupPropertiesKHR C.VkPhysicalDeviceGroupPropertiesKHR
type VkDeviceGroupDeviceCreateInfoKHR C.VkDeviceGroupDeviceCreateInfoKHR

type CDriver struct {
	coreDriver driver.Driver

	enumeratePhysicalDeviceGroups C.PFN_vkEnumeratePhysicalDeviceGroupsKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		enumeratePhysicalDeviceGroups: (C.PFN_vkEnumeratePhysicalDeviceGroupsKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkEnumeratePhysicalDeviceGroupsKHR")))),
	}
}

func (d *CDriver) VkEnumeratePhysicalDeviceGroupsKHR(instance driver.VkInstance, pPhysicalDeviceGroupCount *driver.Uint32, pPhysicalDeviceGroupProperties *VkPhysicalDeviceGroupPropertiesKHR) (common.VkResult, error) {
	if d.enumeratePhysicalDeviceGroups == nil {
		panic("attempt to call extension method vkEnumeratePhysicalDeviceGroupsKHR when extension not present")
	}

	res := common.VkResult(C.cgoEnumeratePhysicalDeviceGroupsKHR(
		d.enumeratePhysicalDeviceGroups,
		(C.VkInstance)(unsafe.Pointer(instance)),
		(*C.uint32_t)(pPhysicalDeviceGroupCount),
		(*C.VkPhysicalDeviceGroupPropertiesKHR)(pPhysicalDeviceGroupProperties),
	))

	return res, res.ToError()
}
