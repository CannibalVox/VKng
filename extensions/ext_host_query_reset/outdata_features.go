package ext_host_query_reset

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

type PhysicalDeviceHostQueryResetFeaturesOutData struct {
	HostQueryReset bool

	common.HaveNext
}

func (o *PhysicalDeviceHostQueryResetFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkPhysicalDeviceHostQueryResetFeaturesEXT{})))
	}

	info := (*C.VkPhysicalDeviceHostQueryResetFeaturesEXT)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_HOST_QUERY_RESET_FEATURES_EXT
	info.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDeviceHostQueryResetFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkPhysicalDeviceHostQueryResetFeaturesEXT)(cDataPointer)
	o.HostQueryReset = info.hostQueryReset != C.VkBool32(0)

	return info.pNext, nil
}
