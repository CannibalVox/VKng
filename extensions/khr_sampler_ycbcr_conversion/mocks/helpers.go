package mock_sampler_ycbcr_conversion

import (
	khr_sampler_ycbcr_conversion_driver "github.com/CannibalVox/VKng/extensions/khr_sampler_ycbcr_conversion/driver"
	"github.com/golang/mock/gomock"
	"math/rand"
	"unsafe"
)

func NewFakeSamplerYcbcrConversion() khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR {
	return khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockSamplerYcbcrConversion(ctrl *gomock.Controller) *MockSamplerYcbcrConversion {
	sampler := NewMockSamplerYcbcrConversion(ctrl)
	sampler.EXPECT().Handle().Return(NewFakeSamplerYcbcrConversion()).AnyTimes()

	return sampler
}