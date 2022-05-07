package core1_0

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

type SemaphoreCreateOptions struct {
	common.HaveNext
}

func (o SemaphoreCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkSemaphoreCreateInfo)
	}
	createInfo := (*C.VkSemaphoreCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO
	createInfo.flags = 0
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}

func (o SemaphoreCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkSemaphoreCreateInfo)(cDataPointer)
	return createInfo.pNext, nil
}
