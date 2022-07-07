package khr_sampler_ycbcr_conversion

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_sampler_ycbcr_conversion_driver "github.com/CannibalVox/VKng/extensions/khr_sampler_ycbcr_conversion/driver"
)

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_sampler_ycbcr_conversion

type SamplerYcbcrConversion interface {
	Handle() khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR
	Destroy(allocator *driver.AllocationCallbacks)
}

type Extension interface {
	CreateSamplerYcbcrConversion(device core1_0.Device, o SamplerYcbcrConversionCreateInfo, allocator *driver.AllocationCallbacks) (SamplerYcbcrConversion, common.VkResult, error)
}
