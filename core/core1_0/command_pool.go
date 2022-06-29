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

type VulkanCommandPool struct {
	deviceDriver      driver.Driver
	commandPoolHandle driver.VkCommandPool
	device            driver.VkDevice
	maximumAPIVersion common.APIVersion
}

func (p *VulkanCommandPool) Handle() driver.VkCommandPool {
	return p.commandPoolHandle
}

func (p *VulkanCommandPool) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p *VulkanCommandPool) Driver() driver.Driver {
	return p.deviceDriver
}

func (p *VulkanCommandPool) APIVersion() common.APIVersion {
	return p.maximumAPIVersion
}

func (p *VulkanCommandPool) Destroy(callbacks *driver.AllocationCallbacks) {
	p.deviceDriver.VkDestroyCommandPool(p.device, p.commandPoolHandle, callbacks.Handle())
	p.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.commandPoolHandle))
}

func (p *VulkanCommandPool) Reset(flags CommandPoolResetFlags) (common.VkResult, error) {
	return p.deviceDriver.VkResetCommandPool(p.device, p.commandPoolHandle, driver.VkCommandPoolResetFlags(flags))
}
