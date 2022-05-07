package khr_swapchain

//go:generate mockgen -source swapchain.go -destination ./mocks/swapchain.go -package mock_swapchain

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
	khr_swapchain_driver "github.com/CannibalVox/VKng/extensions/khr_swapchain/driver"
	"github.com/CannibalVox/cgoparam"
	"time"
	"unsafe"
)

type vulkanSwapchain struct {
	handle     khr_swapchain_driver.VkSwapchainKHR
	device     driver.VkDevice
	driver     khr_swapchain_driver.Driver
	coreDriver driver.Driver

	minimumAPIVersion common.APIVersion
}

type Swapchain interface {
	Handle() khr_swapchain_driver.VkSwapchainKHR

	Destroy(callbacks *driver.AllocationCallbacks)
	Images() ([]core.Image, common.VkResult, error)
	AcquireNextImage(timeout time.Duration, semaphore core1_0.Semaphore, fence core1_0.Fence) (int, common.VkResult, error)
}

func (s *vulkanSwapchain) Handle() khr_swapchain_driver.VkSwapchainKHR {
	return s.handle
}

func (s *vulkanSwapchain) Destroy(callbacks *driver.AllocationCallbacks) {
	s.driver.VkDestroySwapchainKHR(s.device, s.handle, callbacks.Handle())
	s.coreDriver.ObjectStore().Delete(driver.VulkanHandle(s.handle), s)
}

func (s *vulkanSwapchain) attemptImages() ([]core.Image, common.VkResult, error) {
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
	var result []core.Image
	for i := 0; i < imageCount; i++ {
		image := core.CreateImage(s.coreDriver, s.device, imagesSlice[i], s.minimumAPIVersion)
		s.coreDriver.ObjectStore().SetParent(driver.VulkanHandle(s.handle), driver.VulkanHandle(imagesSlice[i]))
		result = append(result, image)
	}

	return result, res, nil
}

func (s *vulkanSwapchain) Images() ([]core.Image, common.VkResult, error) {
	var result []core.Image
	var res common.VkResult
	var err error

	for doWhile := true; doWhile; doWhile = (res == core1_0.VKIncomplete) {
		result, res, err = s.attemptImages()
	}

	return result, res, err
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
