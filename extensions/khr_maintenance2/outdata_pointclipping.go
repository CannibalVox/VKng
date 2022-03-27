package khr_maintenance2

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

type PointClippingOutData struct {
	PointClippingBehavior PointClippingBehavior

	common.HaveNext
}

func (o *PointClippingOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDevicePointClippingPropertiesKHR{})))
	}

	properties := (*C.VkPhysicalDevicePointClippingPropertiesKHR)(preallocatedPointer)
	properties.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES_KHR
	properties.pNext = next

	return preallocatedPointer, nil
}

func (o *PointClippingOutData) PopulateOutData(cPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	properties := (*C.VkPhysicalDevicePointClippingPropertiesKHR)(cPointer)
	o.PointClippingBehavior = PointClippingBehavior(properties.pointClippingBehavior)

	return properties.pNext, nil
}

var _ common.Options = &PointClippingOutData{}
