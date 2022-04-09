package mock_descriptor_update_template

import (
	khr_descriptor_update_template_driver "github.com/CannibalVox/VKng/extensions/khr_descriptor_update_template/driver"
	"github.com/golang/mock/gomock"
	"math/rand"
	"unsafe"
)

func NewFakeDescriptorTemplate() khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR {
	return khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR(unsafe.Pointer(uintptr(rand.Int())))
}

func EasyMockDescriptorTemplate(ctrl *gomock.Controller) *MockDescriptorTemplate {
	swapchain := NewMockDescriptorTemplate(ctrl)
	swapchain.EXPECT().Handle().Return(NewFakeDescriptorTemplate()).AnyTimes()

	return swapchain
}
