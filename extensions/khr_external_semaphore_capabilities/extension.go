package khr_external_semaphore_capabilities

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	khr_external_semaphore_capabilities_driver "github.com/CannibalVox/VKng/extensions/khr_external_semaphore_capabilities/driver"
	"github.com/CannibalVox/cgoparam"
)

type VulkanExtension struct {
	driver khr_external_semaphore_capabilities_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	if !device.IsDeviceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionFromDriver(khr_external_semaphore_capabilities_driver.CreateDriverFromCore(device.Driver()))
}

func CreateExtensionFromDriver(driver khr_external_semaphore_capabilities_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) PhysicalDeviceExternalSemaphoreProperties(physicalDevice core1_0.PhysicalDevice, o PhysicalDeviceExternalSemaphoreInfo, outData *ExternalSemaphoreProperties) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOutDataHeader(arena, outData)
	if err != nil {
		return err
	}

	e.driver.VkGetPhysicalDeviceExternalSemaphorePropertiesKHR(
		physicalDevice.Handle(),
		(*khr_external_semaphore_capabilities_driver.VkPhysicalDeviceExternalSemaphoreInfoKHR)(infoPtr),
		(*khr_external_semaphore_capabilities_driver.VkExternalSemaphorePropertiesKHR)(outDataPtr),
	)

	return common.PopulateOutData(outData, outDataPtr)
}
