package ext_descriptor_indexing

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type PhysicalDeviceDescriptorIndexingFeaturesOutData struct {
	ShaderInputAttachmentArrayDynamicIndexing          bool
	ShaderUniformTexelBufferArrayDynamicIndexing       bool
	ShaderStorageTexelBufferArrayDynamicIndexing       bool
	ShaderUniformBufferArrayNonUniformIndexing         bool
	ShaderSampledImageArrayNonUniformIndexing          bool
	ShaderStorageBufferArrayNonUniformIndexing         bool
	ShaderStorageImageArrayNonUniformIndexing          bool
	ShaderInputAttachmentArrayNonUniformIndexing       bool
	ShaderUniformTexelBufferArrayNonUniformIndexing    bool
	ShaderStorageTexelBufferArrayNonUniformIndexing    bool
	DescriptorBindingUniformBufferUpdateAfterBind      bool
	DescriptorBindingSampledImageUpdateAfterBind       bool
	DescriptorBindingStorageImageUpdateAfterBind       bool
	DescriptorBindingStorageBufferUpdateAfterBind      bool
	DescriptorBindingUniformTexelBufferUpdateAfterBind bool
	DescriptorBindingStorageTexelBufferUpdateAfterBind bool
	DescriptorBindingUpdateUnusedWhilePending          bool
	DescriptorBindingPartiallyBound                    bool
	DescriptorBindingVariableDescriptorCount           bool
	RuntimeDescriptorArray                             bool

	common.HaveNext
}

func (o *PhysicalDeviceDescriptorIndexingFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDescriptorIndexingFeaturesEXT{})))
	}

	info := (*C.VkPhysicalDeviceDescriptorIndexingFeaturesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDescriptorIndexingFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceDescriptorIndexingFeaturesEXT)(cDataPointer)

	o.ShaderInputAttachmentArrayDynamicIndexing = info.shaderInputAttachmentArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderUniformTexelBufferArrayDynamicIndexing = info.shaderUniformTexelBufferArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderStorageTexelBufferArrayDynamicIndexing = info.shaderStorageTexelBufferArrayDynamicIndexing != C.VkBool32(0)
	o.ShaderUniformBufferArrayNonUniformIndexing = info.shaderUniformBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderSampledImageArrayNonUniformIndexing = info.shaderSampledImageArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageBufferArrayNonUniformIndexing = info.shaderStorageBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageImageArrayNonUniformIndexing = info.shaderStorageImageArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderInputAttachmentArrayNonUniformIndexing = info.shaderInputAttachmentArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderUniformTexelBufferArrayNonUniformIndexing = info.shaderUniformTexelBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.ShaderStorageTexelBufferArrayNonUniformIndexing = info.shaderStorageTexelBufferArrayNonUniformIndexing != C.VkBool32(0)
	o.DescriptorBindingUniformBufferUpdateAfterBind = info.descriptorBindingUniformBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingSampledImageUpdateAfterBind = info.descriptorBindingSampledImageUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageImageUpdateAfterBind = info.descriptorBindingStorageImageUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageBufferUpdateAfterBind = info.descriptorBindingStorageBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingUniformTexelBufferUpdateAfterBind = info.descriptorBindingUniformTexelBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingStorageTexelBufferUpdateAfterBind = info.descriptorBindingStorageTexelBufferUpdateAfterBind != C.VkBool32(0)
	o.DescriptorBindingUpdateUnusedWhilePending = info.descriptorBindingUpdateUnusedWhilePending != C.VkBool32(0)
	o.DescriptorBindingPartiallyBound = info.descriptorBindingPartiallyBound != C.VkBool32(0)
	o.DescriptorBindingVariableDescriptorCount = info.descriptorBindingVariableDescriptorCount != C.VkBool32(0)
	o.RuntimeDescriptorArray = info.runtimeDescriptorArray != C.VkBool32(0)

	return info.pNext, nil
}
