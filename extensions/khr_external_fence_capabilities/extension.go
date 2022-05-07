package khr_external_fence_capabilities

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	khr_external_fence_capabilities_driver "github.com/CannibalVox/VKng/extensions/khr_external_fence_capabilities/driver"
	"github.com/CannibalVox/cgoparam"
)

type VulkanExtension struct {
	driver khr_external_fence_capabilities_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	return CreateExtensionFromDriver(khr_external_fence_capabilities_driver.CreateDriverFromCore(device.Driver()))
}

func CreateExtensionFromDriver(driver khr_external_fence_capabilities_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) ExternalFenceProperties(physicalDevice core1_0.PhysicalDevice, o ExternalFenceOptions, outData *ExternalFenceOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOptions(arena, outData)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceExternalFencePropertiesKHR(
		physicalDevice.Handle(),
		(*khr_external_fence_capabilities_driver.VkPhysicalDeviceExternalFenceInfoKHR)(infoPtr),
		(*khr_external_fence_capabilities_driver.VkExternalFencePropertiesKHR)(outDataPtr),
	)

	return common.PopulateOutData(outData, outDataPtr)
}
