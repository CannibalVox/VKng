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

type VulkanEvent struct {
	eventHandle  driver.VkEvent
	device       driver.VkDevice
	deviceDriver driver.Driver

	maximumAPIVersion common.APIVersion
}

func (e *VulkanEvent) Handle() driver.VkEvent {
	return e.eventHandle
}

func (e *VulkanEvent) DeviceHandle() driver.VkDevice {
	return e.device
}

func (e *VulkanEvent) Driver() driver.Driver {
	return e.deviceDriver
}

func (e *VulkanEvent) APIVersion() common.APIVersion {
	return e.maximumAPIVersion
}

func (e *VulkanEvent) Destroy(callbacks *driver.AllocationCallbacks) {
	e.deviceDriver.VkDestroyEvent(e.device, e.eventHandle, callbacks.Handle())
	e.deviceDriver.ObjectStore().Delete(driver.VulkanHandle(e.eventHandle))
}

func (e *VulkanEvent) Set() (common.VkResult, error) {
	return e.deviceDriver.VkSetEvent(e.device, e.eventHandle)
}

func (e *VulkanEvent) Reset() (common.VkResult, error) {
	return e.deviceDriver.VkResetEvent(e.device, e.eventHandle)
}

func (e *VulkanEvent) Status() (common.VkResult, error) {
	return e.deviceDriver.VkGetEventStatus(e.device, e.eventHandle)
}
