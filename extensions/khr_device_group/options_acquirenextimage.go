package khr_device_group

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/extensions/khr_swapchain"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"time"
	"unsafe"
)

type AcquireNextImageOptions struct {
	Swapchain  khr_swapchain.Swapchain
	Timeout    time.Duration
	Semaphore  core1_0.Semaphore
	Fence      core1_0.Fence
	DeviceMask uint32

	common.HaveNext
}

func (o AcquireNextImageOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Swapchain == nil {
		return nil, errors.New("field Swapchain of AcquireNextImageOptions must contain a valid swapchain")
	}

	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkAcquireNextImageInfoKHR{})))
	}

	info := (*C.VkAcquireNextImageInfoKHR)(preallocatedPointer)
	info.sType = C.VK_STRUCTURE_TYPE_ACQUIRE_NEXT_IMAGE_INFO_KHR
	info.pNext = next
	info.swapchain = C.VkSwapchainKHR(unsafe.Pointer(o.Swapchain.Handle()))
	info.semaphore = nil
	info.fence = nil
	info.timeout = C.uint64_t(common.TimeoutNanoseconds(o.Timeout))
	info.deviceMask = C.uint32_t(o.DeviceMask)

	if o.Semaphore != nil {
		info.semaphore = C.VkSemaphore(unsafe.Pointer(o.Semaphore.Handle()))
	}
	if o.Fence != nil {
		info.fence = C.VkFence(unsafe.Pointer(o.Fence.Handle()))
	}

	return preallocatedPointer, nil
}

func (o AcquireNextImageOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	info := (*C.VkAcquireNextImageInfoKHR)(cDataPointer)
	return info.pNext, nil
}
