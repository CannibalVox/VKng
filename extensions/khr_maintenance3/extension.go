package khr_maintenance3

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_maintenance3_driver "github.com/CannibalVox/VKng/extensions/khr_maintenance3/driver"
	"github.com/CannibalVox/cgoparam"
)

type VulkanExtension struct {
	driver khr_maintenance3_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	return &VulkanExtension{
		driver: khr_maintenance3_driver.CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver khr_maintenance3_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) DescriptorSetLayoutSupport(device core1_0.Device, setLayoutOptions core1_0.DescriptorSetLayoutCreateOptions, support *DescriptorSetLayoutSupportOutData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionsPtr, err := common.AllocOptions(arena, setLayoutOptions)
	if err != nil {
		return err
	}

	outDataPtr, err := common.AllocOptions(arena, support)
	if err != nil {
		return err
	}

	e.driver.VkGetDescriptorSetLayoutSupportKHR(device.Handle(), (*driver.VkDescriptorSetLayoutCreateInfo)(optionsPtr), (*khr_maintenance3_driver.VkDescriptorSetLayoutSupportKHR)(outDataPtr))

	return common.PopulateOutData(support, outDataPtr)
}

var _ Extension = &VulkanExtension{}
