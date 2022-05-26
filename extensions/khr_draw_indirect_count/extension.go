package khr_draw_indirect_count

import (
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_draw_indirect_count_driver "github.com/CannibalVox/VKng/extensions/khr_draw_indirect_count/driver"
)

type VulkanExtension struct {
	driver khr_draw_indirect_count_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device, instance core1_0.Instance) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionFromDriver(khr_draw_indirect_count_driver.CreateDriverFromCore(device.Driver()))
}

func CreateExtensionFromDriver(driver khr_draw_indirect_count_driver.Driver) *VulkanExtension {
	ext := &VulkanExtension{
		driver: driver,
	}

	return ext
}

func (e *VulkanExtension) CmdDrawIndexedIndirectCount(commandBuffer core1_0.CommandBuffer, buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	e.driver.VkCmdDrawIndexedIndirectCountKHR(
		commandBuffer.Handle(),
		buffer.Handle(),
		driver.VkDeviceSize(offset),
		countBuffer.Handle(),
		driver.VkDeviceSize(countBufferOffset),
		driver.Uint32(maxDrawCount),
		driver.Uint32(stride),
	)
}

func (e *VulkanExtension) CmdDrawIndirectCount(commandBuffer core1_0.CommandBuffer, buffer core1_0.Buffer, offset uint64, countBuffer core1_0.Buffer, countBufferOffset uint64, maxDrawCount, stride int) {
	e.driver.VkCmdDrawIndirectCountKHR(
		commandBuffer.Handle(),
		buffer.Handle(),
		driver.VkDeviceSize(offset),
		countBuffer.Handle(),
		driver.VkDeviceSize(countBufferOffset),
		driver.Uint32(maxDrawCount),
		driver.Uint32(stride),
	)
}
