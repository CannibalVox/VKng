package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
)

type VulkanShaderModule struct {
	deviceDriver       driver.Driver
	device             driver.VkDevice
	shaderModuleHandle driver.VkShaderModule

	maximumAPIVersion common.APIVersion
}

func (m *VulkanShaderModule) Handle() driver.VkShaderModule {
	return m.shaderModuleHandle
}

func (m *VulkanShaderModule) Driver() driver.Driver {
	return m.deviceDriver
}

func (m *VulkanShaderModule) DeviceHandle() driver.VkDevice {
	return m.device
}

func (m *VulkanShaderModule) APIVersion() common.APIVersion {
	return m.maximumAPIVersion
}

func (m *VulkanShaderModule) Destroy(callbacks *driver.AllocationCallbacks) {
	m.deviceDriver.VkDestroyShaderModule(m.device, m.shaderModuleHandle, callbacks.Handle())
	m.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(m.shaderModuleHandle))
}
