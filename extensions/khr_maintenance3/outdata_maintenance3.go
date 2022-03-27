package khr_maintenance3

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

type Maintenance3OutData struct {
	MaxPerSetDescriptors    int
	MaxMemoryAllocationSize uint64

	common.HaveNext
}

func (o *Maintenance3OutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceMaintenance3PropertiesKHR{})))
	}

	outData := (*C.VkPhysicalDeviceMaintenance3PropertiesKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MAINTENANCE_3_PROPERTIES_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *Maintenance3OutData) PopulateOutData(cPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDeviceMaintenance3PropertiesKHR)(cPointer)

	o.MaxMemoryAllocationSize = uint64(outData.maxMemoryAllocationSize)
	o.MaxPerSetDescriptors = int(outData.maxPerSetDescriptors)

	return outData.pNext, nil
}
