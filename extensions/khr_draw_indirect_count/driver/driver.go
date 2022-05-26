package khr_draw_indirect_count_driver

/*
#include <stdlib.h>
#include "../../vulkan/vulkan.h"

void cgoCmdDrawIndexedIndirectCountKHR(PFN_vkCmdDrawIndexedIndirectCountKHR fn, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	fn(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}

void cgoCmdDrawIndirectCountKHR(PFN_vkCmdDrawIndirectCountKHR fn, VkCommandBuffer commandBuffer, VkBuffer buffer, VkDeviceSize offset, VkBuffer countBuffer, VkDeviceSize countBufferOffset, uint32_t maxDrawCount, uint32_t stride) {
	fn(commandBuffer, buffer, offset, countBuffer, countBufferOffset, maxDrawCount, stride);
}
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

//go:generate mockgen -source driver.go -destination ../mocks/driver.go -package mock_draw_indirect_count

type Driver interface {
	VkCmdDrawIndexedIndirectCountKHR(commandBuffer driver.VkCommandBuffer, buffer driver.VkBuffer, offset driver.VkDeviceSize, countBuffer driver.VkBuffer, countBufferOffset driver.VkDeviceSize, maxDrawCount driver.Uint32, stride driver.Uint32)
	VkCmdDrawIndirectCountKHR(commandBuffer driver.VkCommandBuffer, buffer driver.VkBuffer, offset driver.VkDeviceSize, countBuffer driver.VkBuffer, countBufferOffset driver.VkDeviceSize, maxDrawCount driver.Uint32, stride driver.Uint32)
}

type CDriver struct {
	coreDriver driver.Driver

	drawIndexedIndirectCount C.PFN_vkCmdDrawIndexedIndirectCountKHR
	drawIndirectCount        C.PFN_vkCmdDrawIndirectCountKHR
}

func CreateDriverFromCore(coreDriver driver.Driver) *CDriver {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	return &CDriver{
		coreDriver: coreDriver,

		drawIndexedIndirectCount: (C.PFN_vkCmdDrawIndexedIndirectCountKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCmdDrawIndexedIndirectCountKHR")))),
		drawIndirectCount:        (C.PFN_vkCmdDrawIndirectCountKHR)(coreDriver.LoadProcAddr((*driver.Char)(arena.CString("vkCmdDrawIndirectCountKHR")))),
	}
}

func (d *CDriver) VkCmdDrawIndexedIndirectCountKHR(commandBuffer driver.VkCommandBuffer, buffer driver.VkBuffer, offset driver.VkDeviceSize, countBuffer driver.VkBuffer, countBufferOffset driver.VkDeviceSize, maxDrawCount driver.Uint32, stride driver.Uint32) {
	if d.drawIndexedIndirectCount == nil {
		panic("attempt to call extension method vkCmdDrawIndexedIndirectCountKHR when extension not present")
	}

	C.cgoCmdDrawIndexedIndirectCountKHR(
		d.drawIndexedIndirectCount,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.VkBuffer(unsafe.Pointer(buffer)),
		C.VkDeviceSize(offset),
		C.VkBuffer(unsafe.Pointer(countBuffer)),
		C.VkDeviceSize(countBufferOffset),
		C.uint32_t(maxDrawCount),
		C.uint32_t(stride),
	)
}

func (d *CDriver) VkCmdDrawIndirectCountKHR(commandBuffer driver.VkCommandBuffer, buffer driver.VkBuffer, offset driver.VkDeviceSize, countBuffer driver.VkBuffer, countBufferOffset driver.VkDeviceSize, maxDrawCount driver.Uint32, stride driver.Uint32) {
	if d.drawIndirectCount == nil {
		panic("attempt to call extension method vkCmdDrawIndirectCountKHR when extension not present")
	}

	C.cgoCmdDrawIndirectCountKHR(
		d.drawIndirectCount,
		C.VkCommandBuffer(unsafe.Pointer(commandBuffer)),
		C.VkBuffer(unsafe.Pointer(buffer)),
		C.VkDeviceSize(offset),
		C.VkBuffer(unsafe.Pointer(countBuffer)),
		C.VkDeviceSize(countBufferOffset),
		C.uint32_t(maxDrawCount),
		C.uint32_t(stride),
	)
}
