package khr_sampler_ycbcr_conversion

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_sampler_ycbcr_conversion_driver "github.com/CannibalVox/VKng/extensions/khr_sampler_ycbcr_conversion/driver"
	"github.com/CannibalVox/cgoparam"
)

type VulkanExtension struct {
	driver khr_sampler_ycbcr_conversion_driver.Driver
}

func CreateExtensionFromDevice(device core1_0.Device) *VulkanExtension {
	return &VulkanExtension{
		driver: khr_sampler_ycbcr_conversion_driver.CreateDriverFromCore(device.Driver()),
	}
}

func CreateExtensionFromDriver(driver khr_sampler_ycbcr_conversion_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (e *VulkanExtension) CreateSamplerYcbcrConversion(device core1_0.Device, o SamplerYcbcrConversionCreateOptions, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	optionPtr, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var ycbcrHandle khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR
	res, err := e.driver.VkCreateSamplerYcbcrConversionKHR(
		device.Handle(),
		(*khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionCreateInfoKHR)(optionPtr),
		allocator.Handle(),
		&ycbcrHandle,
	)
	if err != nil {
		return nil, res, err
	}

	ycbcr := device.Driver().ObjectStore().GetOrCreate(driver.VulkanHandle(ycbcrHandle), driver.Core1_1,
		func() any {
			return &vulkanSamplerYcbcrConversion{
				coreDriver:        device.Driver(),
				driver:            e.driver,
				device:            device.Handle(),
				ycbcrHandle:       ycbcrHandle,
				maximumAPIVersion: device.APIVersion(),
			}
		}).(*vulkanSamplerYcbcrConversion)

	return ycbcr, res, nil
}
