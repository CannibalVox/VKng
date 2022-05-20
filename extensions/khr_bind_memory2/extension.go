package khr_bind_memory2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/VKng/extensions/khr_bind_memory2/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanExtension struct {
	driver khr_bind_memory2_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}
	return CreateExtensionFromDriver(khr_bind_memory2_driver.CreateDriverFromCore(device.Driver()))
}

func CreateExtensionFromDriver(driver khr_bind_memory2_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) BindBufferMemory(device core1_0.Device, options []BindBufferMemoryOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindBufferMemoryInfoKHR, BindBufferMemoryOptions](arena, options)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkBindBufferMemory2KHR(device.Handle(), driver.Uint32(len(options)), (*khr_bind_memory2_driver.VkBindBufferMemoryInfoKHR)(unsafe.Pointer(optionPtr)))
}

func (e *VulkanExtension) BindImageMemory(device core1_0.Device, options []BindImageMemoryOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptionSlice[C.VkBindImageMemoryInfoKHR, BindImageMemoryOptions](arena, options)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return e.driver.VkBindImageMemory2KHR(device.Handle(), driver.Uint32(len(options)), (*khr_bind_memory2_driver.VkBindImageMemoryInfoKHR)(unsafe.Pointer(optionPtr)))
}

var _ Extension = &VulkanExtension{}
