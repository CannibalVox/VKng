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

type SemaphoreTypeCreateOptions struct {
	SemaphoreType SemaphoreType
	InitialValue  uint64

	common.NextOptions
}

func (o SemaphoreTypeCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreTypeCreateInfoKHR{})))
	}

	info := (*C.VkSemaphoreTypeCreateInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO_KHR
	info.pNext = next
	info.semaphoreType = C.VkSemaphoreType(o.SemaphoreType)
	info.initialValue = C.uint64_t(o.InitialValue)

	return preallocatedPointer, nil
}
