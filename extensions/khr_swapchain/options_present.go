package khr_swapchain

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

type PresentOptionsOutData struct {
	Results []common.VkResult
}

type PresentOptions struct {
	WaitSemaphores []core1_0.Semaphore
	Swapchains     []Swapchain
	ImageIndices   []int

	common.HaveNext

	OutData *PresentOptionsOutData
}

func (o *PresentOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPresentInfoKHR)
	}
	if len(o.Swapchains) != len(o.ImageIndices) {
		return nil, errors.Newf("present: specified %d swapchains and %d image indices, but they should match")
	}

	createInfo := (*C.VkPresentInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_PRESENT_INFO_KHR
	createInfo.pNext = next

	waitSemaphoreCount := len(o.WaitSemaphores)
	createInfo.waitSemaphoreCount = C.uint32_t(waitSemaphoreCount)
	createInfo.pWaitSemaphores = nil
	if waitSemaphoreCount > 0 {
		semaphorePtr := (*C.VkSemaphore)(allocator.Malloc(waitSemaphoreCount * int(unsafe.Sizeof([1]C.VkSemaphore{}))))
		semaphoreSlice := ([]C.VkSemaphore)(unsafe.Slice(semaphorePtr, waitSemaphoreCount))

		for i := 0; i < waitSemaphoreCount; i++ {
			semaphoreHandle := (C.VkSemaphore)(unsafe.Pointer(o.WaitSemaphores[i].Handle()))
			semaphoreSlice[i] = semaphoreHandle
		}

		createInfo.pWaitSemaphores = semaphorePtr
	}

	swapchainCount := len(o.Swapchains)
	createInfo.swapchainCount = C.uint32_t(swapchainCount)
	createInfo.pSwapchains = nil
	createInfo.pImageIndices = nil
	createInfo.pResults = nil
	if swapchainCount > 0 {
		swapchainPtr := (*C.VkSwapchainKHR)(allocator.Malloc(swapchainCount * int(unsafe.Sizeof([1]C.VkSwapchainKHR{}))))
		swapchainSlice := ([]C.VkSwapchainKHR)(unsafe.Slice(swapchainPtr, swapchainCount))

		imageIndexPtr := (*C.uint32_t)(allocator.Malloc(swapchainCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		imageIndexSlice := ([]C.uint32_t)(unsafe.Slice(imageIndexPtr, swapchainCount))

		resultPtr := (*C.VkResult)(allocator.Malloc(swapchainCount * int(unsafe.Sizeof(C.VkResult(0)))))

		for i := 0; i < swapchainCount; i++ {
			swapchainSlice[i] = (C.VkSwapchainKHR)(unsafe.Pointer(o.Swapchains[i].Handle()))
			imageIndexSlice[i] = (C.uint32_t)(o.ImageIndices[i])
		}

		createInfo.pSwapchains = swapchainPtr
		createInfo.pImageIndices = imageIndexPtr
		createInfo.pResults = resultPtr
	}

	return preallocatedPointer, nil
}

func (o *PresentOptions) PopulateOutData(cDataPointer unsafe.Pointer) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkPresentInfoKHR)(cDataPointer)

	if o.OutData == nil {
		return createInfo.pNext, nil
	}

	resultCount := len(o.Swapchains)
	o.OutData.Results = make([]common.VkResult, resultCount)

	resultSlice := ([]C.VkResult)(unsafe.Slice(createInfo.pResults, resultCount))
	for i := 0; i < resultCount; i++ {
		o.OutData.Results[i] = common.VkResult(resultSlice[i])
	}

	return createInfo.pNext, nil
}
