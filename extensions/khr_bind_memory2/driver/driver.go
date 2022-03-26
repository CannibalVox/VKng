package khr_bind_memory2_driver

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_bind_memory2

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

VkResult cgoBindBufferMemory2KHR(PFN_vkBindBufferMemory2KHR fn, VkDevice device, uint32_t bindInfoCount, VkBindBufferMemoryInfoKHR *pBindInfos) {
	return fn(device, bindInfoCount, pBindInfos);
}

VkResult cgoBindImageMemory2KHR(PFN_vkBindImageMemory2KHR fn, VkDevice device, uint32_t bindInfoCount, VkBindImageMemoryInfoKHR *pBindInfos) {
	return fn(device, bindInfoCount, pBindInfos);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type CDriver struct {
	driver           driver.Driver
	bindBufferMemory C.PFN_vkBindBufferMemory2KHR
	bindImageMemory  C.PFN_vkBindImageMemory2KHR
}

type VkBindBufferMemoryInfoKHR C.VkBindBufferMemoryInfoKHR
type VkBindImageMemoryInfoKHR C.VkBindImageMemoryInfoKHR
type Driver interface {
	VkBindBufferMemory2KHR(device driver.VkDevice, bindInfoCount driver.Uint32, pBindInfos *VkBindBufferMemoryInfoKHR) (common.VkResult, error)
	VkBindImageMemory2KHR(device driver.VkDevice, bindInfoCount driver.Uint32, pBindInfos *VkBindImageMemoryInfoKHR) (common.VkResult, error)
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		driver:           coreDriver,
		bindBufferMemory: (C.PFN_vkBindBufferMemory2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkBindBufferMemory2KHR")))),
		bindImageMemory:  (C.PFN_vkBindImageMemory2KHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkBindImageMemory2KHR")))),
	}
}

func (d *CDriver) VkBindBufferMemory2KHR(device driver.VkDevice, bindInfoCount driver.Uint32, pBindInfos *VkBindBufferMemoryInfoKHR) (common.VkResult, error) {
	if d.bindBufferMemory == nil {
		panic("attempt to call extension method vkBindBufferMemory2KHR when extension not present")
	}

	res := common.VkResult(C.cgoBindBufferMemory2KHR(d.bindBufferMemory,
		C.VkDevice(unsafe.Pointer(device)),
		C.uint32_t(bindInfoCount),
		(*C.VkBindBufferMemoryInfo)(pBindInfos)))

	return res, res.ToError()
}

func (d *CDriver) VkBindImageMemory2KHR(device driver.VkDevice, bindInfoCount driver.Uint32, pBindInfos *VkBindImageMemoryInfoKHR) (common.VkResult, error) {
	if d.bindImageMemory == nil {
		panic("attempt to call extension method vkBindImageMemory2KHR when extension not present")
	}

	res := common.VkResult(C.cgoBindImageMemory2KHR(d.bindImageMemory,
		C.VkDevice(unsafe.Pointer(device)),
		C.uint32_t(bindInfoCount),
		(*C.VkBindImageMemoryInfo)(pBindInfos)))

	return res, res.ToError()
}
