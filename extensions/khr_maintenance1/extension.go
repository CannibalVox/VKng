package khr_maintenance1

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_maintenance1

import "github.com/CannibalVox/VKng/core/core1_0"

type Extension interface {
	TrimCommandPool(commandPool core1_0.CommandPool, flags CommandPoolTrimFlags)
}

type VulkanExtension struct {
	driver Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	return &VulkanExtension{
		driver: CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) TrimCommandPool(commandPool core1_0.CommandPool, flags CommandPoolTrimFlags) {
	e.driver.VkTrimCommandPoolKHR(commandPool.Device(), commandPool.Handle(), VkCommandPoolTrimFlagsKHR(flags))
}
