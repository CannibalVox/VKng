package khr_descriptor_update_template

import (
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_descriptor_update_template_driver "github.com/CannibalVox/VKng/extensions/khr_descriptor_update_template/driver"
)

//go:generate mockgen -source iface.go -destination ./mocks/extension.go -package mock_descriptor_update_template

type DescriptorUpdateTemplate interface {
	Handle() khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR
	Destroy(allocator *driver.AllocationCallbacks)

	UpdateDescriptorSetFromImage(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorImageInfo)
	UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo)
	UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle)
}

type Extension interface {
	CreateDescriptorUpdateTemplate(device core1_0.Device, o DescriptorUpdateTemplateCreateInfo, allocator *driver.AllocationCallbacks) (DescriptorUpdateTemplate, driver.VkResult, error)
}
