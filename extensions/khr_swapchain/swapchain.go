package khr_swapchain

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type vulkanSwapchain struct {
	handle VkSwapchainKHR
	device core.VkDevice
	driver Driver
}

type Swapchain interface {
	Handle() VkSwapchainKHR
	Destroy() error
	Images() ([]core.Image, core.VkResult, error)
	AcquireNextImage(timeout time.Duration, semaphore core.Semaphore, fence core.Fence) (int, core.VkResult, error)
	PresentToQueue(queue core.Queue, o *PresentOptions) (resultBySwapchain []core.VkResult, res core.VkResult, anyError error)
}

func (s *vulkanSwapchain) Handle() VkSwapchainKHR {
	return s.handle
}

func (s *vulkanSwapchain) Destroy() error {
	return s.driver.VkDestroySwapchainKHR(s.device, s.handle, nil)
}

func (s *vulkanSwapchain) Images() ([]core.Image, core.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	imageCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	imageCountRef := (*core.Uint32)(imageCountPtr)

	res, err := s.driver.VkGetSwapchainImagesKHR(s.device, s.handle, imageCountRef, nil)
	if err != nil {
		return nil, res, err
	}

	imageCount := int(*imageCountRef)
	if imageCount == 0 {
		return nil, res, nil
	}

	imagesPtr := (*core.VkImage)(allocator.Malloc(imageCount * int(unsafe.Sizeof([1]C.VkImage{}))))

	res, err = s.driver.VkGetSwapchainImagesKHR(s.device, s.handle, imageCountRef, imagesPtr)
	if err != nil {
		return nil, res, err
	}

	imagesSlice := ([]core.VkImage)(unsafe.Slice(imagesPtr, imageCount))
	var result []core.Image
	deviceHandle := (core.VkDevice)(unsafe.Pointer(s.device))
	for i := 0; i < imageCount; i++ {
		result = append(result, core.CreateImageFromHandles(imagesSlice[i], deviceHandle, s.driver.coreDriver()))
	}

	return result, res, nil
}

func (s *vulkanSwapchain) AcquireNextImage(timeout time.Duration, semaphore core.Semaphore, fence core.Fence) (int, core.VkResult, error) {
	var imageIndex core.Uint32

	var semaphoreHandle core.VkSemaphore
	var fenceHandle core.VkFence

	if semaphore != nil {
		semaphoreHandle = semaphore.Handle()
	}
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	res, err := s.driver.VkAcquireNextImageKHR(s.device, s.handle, core.Uint64(common.TimeoutNanoseconds(timeout)), semaphoreHandle, fenceHandle, &imageIndex)

	return int(imageIndex), res, err
}

func (s *vulkanSwapchain) PresentToQueue(queue core.Queue, o *PresentOptions) (resultBySwapchain []core.VkResult, res core.VkResult, anyError error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core.VKErrorUnknown, err
	}

	createInfoPtr := (*VkPresentInfoKHR)(createInfo)
	res, err = s.driver.VkQueuePresentKHR(queue.Handle(), createInfoPtr)

	resSlice := unsafe.Slice(createInfoPtr.pResults, len(o.Swapchains))
	for i := 0; i < len(o.Swapchains); i++ {
		singleRes := core.VkResult(resSlice[i])
		resultBySwapchain = append(resultBySwapchain, singleRes)
	}

	return resultBySwapchain, res, res.ToError()
}
