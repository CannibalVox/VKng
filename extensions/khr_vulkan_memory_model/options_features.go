package khr_vulkan_memory_model

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type PhysicalDeviceVulkanMemoryModelFeaturesOptions struct {
	VulkanMemoryModel                             bool
	VulkanMemoryModelDeviceScope                  bool
	VulkanMemoryModelAvailabilityVisibilityChains bool

	common.HaveNext
}

func (o PhysicalDeviceVulkanMemoryModelFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR
	info.pNext = next
	info.vulkanMemoryModel = C.VkBool32(0)
	info.vulkanMemoryModelDeviceScope = C.VkBool32(0)
	info.vulkanMemoryModelAvailabilityVisibilityChains = C.VkBool32(0)

	if o.VulkanMemoryModel {
		info.vulkanMemoryModel = C.VkBool32(1)
	}

	if o.VulkanMemoryModelDeviceScope {
		info.vulkanMemoryModelDeviceScope = C.VkBool32(1)
	}

	if o.VulkanMemoryModelAvailabilityVisibilityChains {
		info.vulkanMemoryModelAvailabilityVisibilityChains = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o PhysicalDeviceVulkanMemoryModelFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceVulkanMemoryModelFeaturesKHR)(cDataPointer)
	return info.pNext, nil
}
