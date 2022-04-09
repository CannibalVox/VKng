package khr_descriptor_update_template

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	khr_descriptor_update_template_driver "github.com/CannibalVox/VKng/extensions/khr_descriptor_update_template/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type vulkanDescriptorTemplate struct {
	coreDriver               driver.Driver
	driver                   khr_descriptor_update_template_driver.Driver
	device                   driver.VkDevice
	descriptorTemplateHandle khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR

	maximumAPIVersion common.APIVersion
}

func (t *vulkanDescriptorTemplate) Handle() khr_descriptor_update_template_driver.VkDescriptorUpdateTemplateKHR {
	return t.descriptorTemplateHandle
}

func (t *vulkanDescriptorTemplate) Destroy(allocator *driver.AllocationCallbacks) {
	t.driver.VkDestroyDescriptorUpdateTemplateKHR(t.device, t.descriptorTemplateHandle, allocator.Handle())
}

func (t *vulkanDescriptorTemplate) UpdateDescriptorSetFromImage(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorImageInfo) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorImageInfo)
	info := (*C.VkDescriptorImageInfo)(infoUnsafe)
	info.sampler = C.VkSampler(unsafe.Pointer(data.Sampler.Handle()))
	info.imageView = C.VkImageView(unsafe.Pointer(data.ImageView.Handle()))
	info.imageLayout = C.VkImageLayout(data.ImageLayout)

	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		t.device,
		descriptorSet.Handle(),
		t.descriptorTemplateHandle,
		infoUnsafe,
	)
}

func (t *vulkanDescriptorTemplate) UpdateDescriptorSetFromBuffer(descriptorSet core1_0.DescriptorSet, data core1_0.DescriptorBufferInfo) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	infoUnsafe := arena.Malloc(C.sizeof_struct_VkDescriptorBufferInfo)
	info := (*C.VkDescriptorBufferInfo)(infoUnsafe)
	info.buffer = C.VkBuffer(unsafe.Pointer(data.Buffer.Handle()))
	info.offset = C.VkDeviceSize(data.Offset)
	info._range = C.VkDeviceSize(data.Range)

	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		t.device,
		descriptorSet.Handle(),
		t.descriptorTemplateHandle,
		infoUnsafe,
	)
}

func (t *vulkanDescriptorTemplate) UpdateDescriptorSetFromObjectHandle(descriptorSet core1_0.DescriptorSet, data driver.VulkanHandle) {
	t.driver.VkUpdateDescriptorSetWithTemplateKHR(
		t.device,
		descriptorSet.Handle(),
		t.descriptorTemplateHandle,
		unsafe.Pointer(data),
	)
}
