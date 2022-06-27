package ext_host_query_reset_driver

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_host_query_reset

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoResetQueryPoolEXT(PFN_vkResetQueryPoolEXT fn, VkDevice device, VkQueryPool queryPool, uint32_t firstQuery, uint32_t queryCount) {
	fn(device, queryPool, firstQuery, queryCount);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type Driver interface {
	VkResetQueryPoolEXT(device driver.VkDevice, queryPool driver.VkQueryPool, firstQuery driver.Uint32, queryCount driver.Uint32)
}

type VkPhysicalDeviceHostQueryResetFeaturesEXT C.VkPhysicalDeviceHostQueryResetFeaturesEXT

type CDriver struct {
	coreDriver driver.Driver

	resetQueryPool C.PFN_vkResetQueryPoolEXT
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver:     coreDriver,
		resetQueryPool: (C.PFN_vkResetQueryPoolEXT)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkResetQueryPoolEXT")))),
	}
}

func (d *CDriver) VkResetQueryPoolEXT(device driver.VkDevice, queryPool driver.VkQueryPool, firstQuery driver.Uint32, queryCount driver.Uint32) {
	C.cgoResetQueryPoolEXT(
		d.resetQueryPool,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkQueryPool(unsafe.Pointer(queryPool)),
		C.uint32_t(firstQuery),
		C.uint32_t(queryCount),
	)
}
