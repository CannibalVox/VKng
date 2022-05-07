package khr_descriptor_update_template

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_descriptor_update_template_driver "github.com/CannibalVox/VKng/extensions/khr_descriptor_update_template/driver"
	"github.com/CannibalVox/cgoparam"
)

type VulkanExtension struct {
	driver khr_descriptor_update_template_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	return &VulkanExtension{
		driver: khr_descriptor_update_template_driver.CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver khr_descriptor_update_template_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) CreateDescriptorUpdateTemplate(device core1_0.Device, o DescriptorUpdateTemplateCreateOptions, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfoPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var templateHandle khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR
	res, err := e.driver.VkCreateDescriptorUpdateTemplateKHR(device.Handle(),
		(*khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateCreateInfoKHR)(createInfoPtr),
		allocator.Handle(),
		&templateHandle,
	)
	if err != nil {
		return nil, res, err
	}

	descriptorTemplate := device.Driver().ObjectStore().GetOrCreate(driver.VulkanHandle(templateHandle),
		func() interface{} {
			template := &vulkanDescriptorUpdateTemplate{
				driver:                   e.driver,
				coreDriver:               device.Driver(),
				device:                   device.Handle(),
				descriptorTemplateHandle: templateHandle,
				maximumAPIVersion:        device.APIVersion(),
			}

			return template
		}).(*vulkanDescriptorUpdateTemplate)
	device.Driver().ObjectStore().SetParent(driver.VulkanHandle(device.Handle()), driver.VulkanHandle(templateHandle))

	return descriptorTemplate, res, nil
}
