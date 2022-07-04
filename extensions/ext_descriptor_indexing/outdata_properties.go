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

type PhysicalDeviceDescriptorIndexingOutData struct {
	MaxUpdateAfterBindDescriptorsInAllPools            int
	ShaderUniformBufferArrayNonUniformIndexingNative   bool
	ShaderSampledImageArrayNonUniformIndexingNative    bool
	ShaderStorageBufferArrayNonUniformIndexingNative   bool
	ShaderStorageImageArrayNonUniformIndexingNative    bool
	ShaderInputAttachmentArrayNonUniformIndexingNative bool
	RobustBufferAccessUpdateAfterBind                  bool
	QuadDivergentImplicitLod                           bool

	MaxPerStageDescriptorUpdateAfterBindSamplers         int
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers   int
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers   int
	MaxPerStageDescriptorUpdateAfterBindSampledImages    int
	MaxPerStageDescriptorUpdateAfterBindStorageImages    int
	MaxPerStageDescriptorUpdateAfterBindInputAttachments int
	MaxPerStageUpdateAfterBindResources                  int

	MaxDescriptorSetUpdateAfterBindSamplers              int
	MaxDescriptorSetUpdateAfterBindUniformBuffers        int
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic int
	MaxDescriptorSetUpdateAfterBindStorageBuffers        int
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic int
	MaxDescriptorSetUpdateAfterBindSampledImages         int
	MaxDescriptorSetUpdateAfterBindStorageImages         int
	MaxDescriptorSetUpdateAfterBindInputAttachments      int

	common.NextOutData
}

func (o *PhysicalDeviceDescriptorIndexingOutData) PopulateHeader(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceDescriptorIndexingPropertiesEXT{})))
	}

	info := (*C.VkPhysicalDeviceDescriptorIndexingPropertiesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceDescriptorIndexingOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceDescriptorIndexingPropertiesEXT)(cDataPointer)

	o.MaxUpdateAfterBindDescriptorsInAllPools = int(info.maxUpdateAfterBindDescriptorsInAllPools)
	o.ShaderUniformBufferArrayNonUniformIndexingNative = info.shaderUniformBufferArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderSampledImageArrayNonUniformIndexingNative = info.shaderSampledImageArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderStorageBufferArrayNonUniformIndexingNative = info.shaderStorageBufferArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderStorageImageArrayNonUniformIndexingNative = info.shaderStorageImageArrayNonUniformIndexingNative != C.VkBool32(0)
	o.ShaderInputAttachmentArrayNonUniformIndexingNative = info.shaderInputAttachmentArrayNonUniformIndexingNative != C.VkBool32(0)
	o.RobustBufferAccessUpdateAfterBind = info.robustBufferAccessUpdateAfterBind != C.VkBool32(0)
	o.QuadDivergentImplicitLod = info.quadDivergentImplicitLod != C.VkBool32(0)

	o.MaxPerStageDescriptorUpdateAfterBindSamplers = int(info.maxPerStageDescriptorUpdateAfterBindSamplers)
	o.MaxPerStageDescriptorUpdateAfterBindUniformBuffers = int(info.maxPerStageDescriptorUpdateAfterBindUniformBuffers)
	o.MaxPerStageDescriptorUpdateAfterBindStorageBuffers = int(info.maxPerStageDescriptorUpdateAfterBindStorageBuffers)
	o.MaxPerStageDescriptorUpdateAfterBindSampledImages = int(info.maxPerStageDescriptorUpdateAfterBindSampledImages)
	o.MaxPerStageDescriptorUpdateAfterBindStorageImages = int(info.maxPerStageDescriptorUpdateAfterBindStorageImages)
	o.MaxPerStageDescriptorUpdateAfterBindInputAttachments = int(info.maxPerStageDescriptorUpdateAfterBindInputAttachments)
	o.MaxPerStageUpdateAfterBindResources = int(info.maxPerStageUpdateAfterBindResources)

	o.MaxDescriptorSetUpdateAfterBindSamplers = int(info.maxDescriptorSetUpdateAfterBindSamplers)
	o.MaxDescriptorSetUpdateAfterBindUniformBuffers = int(info.maxDescriptorSetUpdateAfterBindUniformBuffers)
	o.MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic = int(info.maxDescriptorSetUpdateAfterBindUniformBuffersDynamic)
	o.MaxDescriptorSetUpdateAfterBindStorageBuffers = int(info.maxDescriptorSetUpdateAfterBindStorageBuffers)
	o.MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic = int(info.maxDescriptorSetUpdateAfterBindStorageBuffersDynamic)
	o.MaxDescriptorSetUpdateAfterBindSampledImages = int(info.maxDescriptorSetUpdateAfterBindSampledImages)
	o.MaxDescriptorSetUpdateAfterBindStorageImages = int(info.maxDescriptorSetUpdateAfterBindStorageImages)
	o.MaxDescriptorSetUpdateAfterBindInputAttachments = int(info.maxDescriptorSetUpdateAfterBindInputAttachments)

	return info.pNext, nil
}
