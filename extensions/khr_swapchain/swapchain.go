package khr_swapchain

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type vulkanSwapchain struct {
	handle VkSwapchainKHR
	device driver.VkDevice
	driver Driver
}

type CommonSwapchain interface {
	Handle() VkSwapchainKHR
}

type Swapchain interface {
	CommonSwapchain
	Destroy(callbacks *driver.AllocationCallbacks)
	Images() ([]core1_0.Image, common.VkResult, error)
	AcquireNextImage(timeout time.Duration, semaphore core1_0.Semaphore, fence core1_0.Fence) (int, common.VkResult, error)
	PresentToQueue(queue core1_0.Queue, o *PresentOptions) (resultBySwapchain []common.VkResult, res common.VkResult, anyError error)
}

func (s *vulkanSwapchain) Handle() VkSwapchainKHR {
	return s.handle
}

func (s *vulkanSwapchain) Destroy(callbacks *driver.AllocationCallbacks) {
	s.driver.VkDestroySwapchainKHR(s.device, s.handle, callbacks.Handle())
}

func (s *vulkanSwapchain) Images() ([]core1_0.Image, common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	imageCountPtr := allocator.Malloc(int(unsafe.Sizeof(C.uint32_t(0))))
	imageCountRef := (*driver.Uint32)(imageCountPtr)

	res, err := s.driver.VkGetSwapchainImagesKHR(s.device, s.handle, imageCountRef, nil)
	if err != nil {
		return nil, res, err
	}

	imageCount := int(*imageCountRef)
	if imageCount == 0 {
		return nil, res, nil
	}

	imagesPtr := (*driver.VkImage)(allocator.Malloc(imageCount * int(unsafe.Sizeof([1]C.VkImage{}))))

	res, err = s.driver.VkGetSwapchainImagesKHR(s.device, s.handle, imageCountRef, imagesPtr)
	if err != nil {
		return nil, res, err
	}

	imagesSlice := ([]driver.VkImage)(unsafe.Slice(imagesPtr, imageCount))
	var result []core1_0.Image
	deviceHandle := (driver.VkDevice)(unsafe.Pointer(s.device))
	for i := 0; i < imageCount; i++ {
		image := core.CreateImageFromHandles(imagesSlice[i], deviceHandle, s.driver.coreDriver())
		result = append(result, image)
	}

	return result, res, nil
}

func (s *vulkanSwapchain) AcquireNextImage(timeout time.Duration, semaphore core1_0.Semaphore, fence core1_0.Fence) (int, common.VkResult, error) {
	var imageIndex driver.Uint32

	var semaphoreHandle driver.VkSemaphore
	var fenceHandle driver.VkFence

	if semaphore != nil {
		semaphoreHandle = semaphore.Handle()
	}
	if fence != nil {
		fenceHandle = fence.Handle()
	}

	res, err := s.driver.VkAcquireNextImageKHR(s.device, s.handle, driver.Uint64(common.TimeoutNanoseconds(timeout)), semaphoreHandle, fenceHandle, &imageIndex)

	return int(imageIndex), res, err
}

func (s *vulkanSwapchain) PresentToQueue(queue core1_0.Queue, o *PresentOptions) (resultBySwapchain []common.VkResult, res common.VkResult, anyError error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	createInfoPtr := (*VkPresentInfoKHR)(createInfo)
	res, err = s.driver.VkQueuePresentKHR(queue.Handle(), createInfoPtr)

	resSlice := unsafe.Slice(createInfoPtr.pResults, len(o.Swapchains))
	for i := 0; i < len(o.Swapchains); i++ {
		singleRes := common.VkResult(resSlice[i])
		resultBySwapchain = append(resultBySwapchain, singleRes)
	}

	return resultBySwapchain, res, res.ToError()
}
