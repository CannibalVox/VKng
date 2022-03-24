package khr_get_physical_device_properties2

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DeviceFeatureOptions struct {
	Features core1_0.PhysicalDeviceFeatures

	common.HaveNext
}

func (o DeviceFeatureOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceFeatures2KHR{})))
	}

	data := (*C.VkPhysicalDeviceFeatures2KHR)(preallocatedPointer)
	data.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR
	data.pNext = next
	_, err := o.Features.PopulateCPointer(allocator, unsafe.Pointer(&data.features))

	return preallocatedPointer, err
}

func (o DeviceFeatureOptions) PopulateOutData(cDataPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	data := (*C.VkPhysicalDeviceFeatures2KHR)(cDataPointer)

	return data.pNext, nil
}
