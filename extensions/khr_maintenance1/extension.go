package khr_maintenance1

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_maintenance1

import (
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/extensions/khr_maintenance1/driver"
)

type Extension interface {
	TrimCommandPool(commandPool core1_0.CommandPool, flags CommandPoolTrimFlags)
}

type VulkanExtension struct {
	driver khr_maintenance1_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return &VulkanExtension{
		driver: khr_maintenance1_driver.CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver khr_maintenance1_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) TrimCommandPool(commandPool core1_0.CommandPool, flags CommandPoolTrimFlags) {
	e.driver.VkTrimCommandPoolKHR(commandPool.DeviceHandle(), commandPool.Handle(), khr_maintenance1_driver.VkCommandPoolTrimFlagsKHR(flags))
}
