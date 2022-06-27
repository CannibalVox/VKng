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

type VulkanPipeline struct {
	deviceDriver   driver.Driver
	device         driver.VkDevice
	pipelineHandle driver.VkPipeline

	maximumAPIVersion common.APIVersion
}

func (p *VulkanPipeline) Handle() driver.VkPipeline {
	return p.pipelineHandle
}

func (p *VulkanPipeline) Driver() driver.Driver {
	return p.deviceDriver
}

func (p *VulkanPipeline) DeviceHandle() driver.VkDevice {
	return p.device
}

func (p *VulkanPipeline) APIVersion() common.APIVersion {
	return p.maximumAPIVersion
}

func (p *VulkanPipeline) Destroy(callbacks *driver.AllocationCallbacks) {
	p.deviceDriver.VkDestroyPipeline(p.device, p.pipelineHandle, callbacks.Handle())
	p.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(p.pipelineHandle))
}
