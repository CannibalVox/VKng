package vk_khr_maintenance1

//go:generate mockgen -source driver.go -destination ./mocks/driver.go -package mock_maintenance1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"

void cgoTrimCommandPoolKHR(PFN_vkTrimCommandPoolKHR fn, VkDevice device, VkCommandPool commandPool, VkCommandPoolTrimFlagsKHR flags) {
	fn(device, commandPool, flags);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VkCommandPoolTrimFlagsKHR C.VkCommandPoolTrimFlagsKHR
type Driver interface {
	VkTrimCommandPoolKHR(device driver.VkDevice, commandPool driver.VkCommandPool, flags VkCommandPoolTrimFlagsKHR)
}

type CDriver struct {
	driver driver.Driver

	trimCommandPool C.PFN_vkTrimCommandPoolKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		driver: coreDriver,

		trimCommandPool: (C.PFN_vkTrimCommandPoolKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkTrimCommandPoolKHR")))),
	}
}

func (d *CDriver) VkTrimCommandPoolKHR(device driver.VkDevice, commandPool driver.VkCommandPool, flags VkCommandPoolTrimFlagsKHR) {
	if d.trimCommandPool == nil {
		panic("attempt to call extension method vkTrimCommandPoolKHR when extension not present")
	}

	C.cgoTrimCommandPoolKHR(d.trimCommandPool,
		C.VkDevice(unsafe.Pointer(device)),
		C.VkCommandPool(unsafe.Pointer(commandPool)),
		C.VkCommandPoolTrimFlags(flags))
}
