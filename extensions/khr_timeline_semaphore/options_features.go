package khr_timeline_semaphore

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

type PhysicalDeviceTimelineSemaphoreFeaturesOptions struct {
	TimelineSemaphore bool

	common.HaveNext
}

func (o PhysicalDeviceTimelineSemaphoreFeaturesOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR{})))
	}

	info := (*C.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TIMELINE_SEMAPHORE_FEATURES_KHR
	info.pNext = next
	info.timelineSemaphore = C.VkBool32(0)

	if o.TimelineSemaphore {
		info.timelineSemaphore = C.VkBool32(1)
	}

	return preallocatedPointer, nil
}

func (o PhysicalDeviceTimelineSemaphoreFeaturesOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceTimelineSemaphoreFeaturesKHR)(cDataPointer)
	return info.pNext, nil
}
