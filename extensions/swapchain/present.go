package ext_swapchain

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/resource"
	"github.com/CannibalVox/cgoalloc"
	"github.com/palantir/stacktrace"
	"unsafe"
)

type PresentOptions struct {
	WaitSemaphores []*resource.Semaphore
	Swapchains     []*Swapchain
	ImageIndices   []int

	Next core.Options
}

func (o *PresentOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	if len(o.Swapchains) != len(o.ImageIndices) {
		return nil, stacktrace.NewError("present: specified %d swapchains and %d image indices, but they should match")
	}

	createInfo := (*C.VkPresentInfoKHR)(allocator.Malloc(C.sizeof_struct_VkPresentInfoKHR))
	createInfo.sType = C.VK_STRUCTURE_TYPE_PRESENT_INFO_KHR

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

	var err error
	var next unsafe.Pointer
	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}

	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}

func PresentToQueue(allocator cgoalloc.Allocator, queue *resource.Queue, o *PresentOptions) (resultBySwapchain []core.Result, res core.Result, anyError error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	createInfoPtr := (*C.VkPresentInfoKHR)(createInfo)
	queueHandle := (C.VkQueue)(unsafe.Pointer(queue.Handle()))

	res = core.Result(C.vkQueuePresentKHR(queueHandle, createInfoPtr))

	resSlice := unsafe.Slice(createInfoPtr.pResults, len(o.Swapchains))
	for i := 0; i < len(o.Swapchains); i++ {
		singleRes := core.Result(resSlice[i])
		resultBySwapchain = append(resultBySwapchain, singleRes)
	}

	return resultBySwapchain, res, res.ToError()
}
