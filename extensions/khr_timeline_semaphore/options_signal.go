package khr_timeline_semaphore

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type SemaphoreSignalInfo struct {
	Semaphore core1_0.Semaphore
	Value     uint64

	common.NextOptions
}

func (o SemaphoreSignalInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreSignalInfoKHR{})))
	}

	if o.Semaphore == nil {
		return nil, errors.New("the 'Semaphore' field of SemaphoreSignalInfo must be non-nil")
	}

	info := (*C.VkSemaphoreSignalInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_SIGNAL_INFO_KHR
	info.pNext = next
	info.semaphore = C.VkSemaphore(unsafe.Pointer(o.Semaphore.Handle()))
	info.value = C.uint64_t(o.Value)

	return preallocatedPointer, nil
}
