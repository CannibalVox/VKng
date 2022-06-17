package ext_host_query_reset

import (
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	ext_host_query_reset_driver "github.com/CannibalVox/VKng/extensions/ext_host_query_reset/driver"
)

type VulkanExtension struct {
	driver ext_host_query_reset_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: ext_host_query_reset_driver.CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver ext_host_query_reset_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) ResetQueryPool(queryPool core1_0.QueryPool, firstQuery, queryCount int) {
	e.driver.VkResetQueryPool(queryPool.DeviceHandle(), queryPool.Handle(), driver.Uint32(firstQuery), driver.Uint32(queryCount))
}
