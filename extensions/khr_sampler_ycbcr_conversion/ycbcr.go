package khr_sampler_ycbcr_conversion

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	khr_sampler_ycbcr_conversion_driver "github.com/CannibalVox/VKng/extensions/khr_sampler_ycbcr_conversion/driver"
)

type vulkanSamplerYcbcrConversion struct {
	coreDriver  driver.Driver
	driver      khr_sampler_ycbcr_conversion_driver.Driver
	device      driver.VkDevice
	ycbcrHandle khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR

	maximumAPIVersion common.APIVersion
}

func (y *vulkanSamplerYcbcrConversion) Handle() khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR {
	return y.ycbcrHandle
}

func (y *vulkanSamplerYcbcrConversion) Destroy(allocator *driver.AllocationCallbacks) {
	y.driver.VkDestroySamplerYcbcrConversionKHR(y.device, y.ycbcrHandle, allocator.Handle())
}
