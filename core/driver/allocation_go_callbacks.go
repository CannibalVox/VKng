package driver

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"runtime/cgo"
	"unsafe"
)

//export goAllocationCallback
func goAllocationCallback(userData unsafe.Pointer, size C.size_t, alignment C.size_t, allocationScope C.VkSystemAllocationScope) unsafe.Pointer {
	callbacks := cgo.Handle(userData).Value().(*AllocationCallbacks)
	return callbacks.allocation(callbacks.userData, int(size), int(alignment), common.SystemAllocationScope(allocationScope))
}

//export goReallocationCallback
func goReallocationCallback(userData unsafe.Pointer, original unsafe.Pointer, size, alignment C.size_t, allocationScope C.VkSystemAllocationScope) unsafe.Pointer {
	callbacks := cgo.Handle(userData).Value().(*AllocationCallbacks)
	return callbacks.reallocation(callbacks.userData, original, int(size), int(alignment), common.SystemAllocationScope(allocationScope))
}

//export goFreeCallback
func goFreeCallback(userData unsafe.Pointer, memory unsafe.Pointer) {
	callbacks := cgo.Handle(userData).Value().(*AllocationCallbacks)
	callbacks.free(callbacks.userData, memory)
}

//export goInternalAllocationCallback
func goInternalAllocationCallback(userData unsafe.Pointer, size C.size_t, allocationType C.VkInternalAllocationType, allocationScope C.VkSystemAllocationScope) {
	callbacks := cgo.Handle(userData).Value().(*AllocationCallbacks)
	callbacks.internalAllocation(callbacks.userData, int(size), common.InternalAllocationType(allocationType), common.SystemAllocationScope(allocationScope))
}

//export goInternalFreeCallback
func goInternalFreeCallback(userData unsafe.Pointer, size C.size_t, allocationType C.VkInternalAllocationType, allocationScope C.VkSystemAllocationScope) {
	callbacks := cgo.Handle(userData).Value().(*AllocationCallbacks)
	callbacks.internalFree(callbacks.userData, int(size), common.InternalAllocationType(allocationType), common.SystemAllocationScope(allocationScope))
}
