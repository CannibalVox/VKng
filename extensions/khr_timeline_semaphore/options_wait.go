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

type SemaphoreWaitOptions struct {
	Flags      SemaphoreWaitFlags
	Semaphores []core1_0.Semaphore
	Values     []uint64

	common.NextOptions
}

func (o SemaphoreWaitOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkSemaphoreWaitInfoKHR{})))
	}

	if len(o.Semaphores) != len(o.Values) {
		return nil, errors.Newf("the SemaphoreWaitOptions 'Semaphores' list has %d elements, but the 'Values' list has %d elements- these lists must be the same size", len(o.Semaphores), len(o.Values))
	}

	info := (*C.VkSemaphoreWaitInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_SEMAPHORE_WAIT_INFO_KHR
	info.pNext = next
	info.flags = C.VkSemaphoreWaitFlags(o.Flags)

	count := len(o.Semaphores)
	info.semaphoreCount = C.uint32_t(count)
	info.pSemaphores = nil
	info.pValues = nil

	if count > 0 {
		info.pSemaphores = (*C.VkSemaphore)(allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkSemaphore{}))))
		info.pValues = (*C.uint64_t)(allocator.Malloc(count * int(unsafe.Sizeof(C.uint64_t(0)))))

		semaphoreSlice := unsafe.Slice(info.pSemaphores, count)
		valueSlice := unsafe.Slice(info.pValues, count)

		for i := 0; i < count; i++ {
			if o.Semaphores[i] == nil {
				return nil, errors.Newf("the SemaphoreWaitOptions 'Semaphores' list has a nil semaphore at element %d- all elements must be non-nil", i)
			}

			semaphoreSlice[i] = C.VkSemaphore(unsafe.Pointer(o.Semaphores[i].Handle()))
			valueSlice[i] = C.uint64_t(o.Values[i])
		}
	}

	return preallocatedPointer, nil
}
