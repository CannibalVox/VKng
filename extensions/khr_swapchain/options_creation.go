package khr_swapchain

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	ext_surface "github.com/CannibalVox/VKng/extensions/khr_surface"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type SwapchainCreateInfo struct {
	Surface ext_surface.Surface

	Flags SwapchainCreateFlags

	MinImageCount int

	ImageFormat      core1_0.Format
	ImageColorSpace  ext_surface.ColorSpace
	ImageExtent      core1_0.Extent2D
	ImageArrayLayers int
	ImageUsage       core1_0.ImageUsageFlags

	ImageSharingMode   core1_0.SharingMode
	QueueFamilyIndices []int

	PreTransform   ext_surface.SurfaceTransformFlags
	CompositeAlpha ext_surface.CompositeAlphaFlags
	PresentMode    ext_surface.PresentMode

	Clipped      bool
	OldSwapchain Swapchain

	common.NextOptions
}

func (o SwapchainCreateInfo) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Surface == nil {
		return nil, errors.New("khr_swapchain.SwapchainCreateInfo.Surface cannot be nil")
	}

	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkSwapchainCreateInfoKHR{})))
	}
	createInfo := (*C.VkSwapchainCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
	createInfo.flags = C.VkSwapchainCreateFlagsKHR(o.Flags)
	createInfo.pNext = next

	createInfo.surface = C.VkSurfaceKHR(unsafe.Pointer(o.Surface.Handle()))
	createInfo.minImageCount = C.uint32_t(o.MinImageCount)

	createInfo.imageFormat = C.VkFormat(o.ImageFormat)
	createInfo.imageColorSpace = C.VkColorSpaceKHR(o.ImageColorSpace)
	createInfo.imageExtent.width = C.uint32_t(o.ImageExtent.Width)
	createInfo.imageExtent.height = C.uint32_t(o.ImageExtent.Height)
	createInfo.imageArrayLayers = C.uint32_t(o.ImageArrayLayers)
	createInfo.imageUsage = C.VkImageUsageFlags(o.ImageUsage)

	createInfo.imageSharingMode = C.VkSharingMode(o.ImageSharingMode)
	createInfo.queueFamilyIndexCount = C.uint32_t(len(o.QueueFamilyIndices))

	if len(o.QueueFamilyIndices) == 0 {
		createInfo.pQueueFamilyIndices = nil
	} else {
		familyIndexPtr := (*C.uint32_t)(allocator.Malloc(len(o.QueueFamilyIndices) * int(unsafe.Sizeof(C.uint32_t(0)))))
		createInfo.pQueueFamilyIndices = familyIndexPtr

		familyIndexSlice := ([]C.uint32_t)(unsafe.Slice(familyIndexPtr, len(o.QueueFamilyIndices)))
		for i, index := range o.QueueFamilyIndices {
			familyIndexSlice[i] = C.uint32_t(index)
		}
	}

	createInfo.preTransform = C.VkSurfaceTransformFlagBitsKHR(o.PreTransform)
	createInfo.compositeAlpha = C.VkCompositeAlphaFlagBitsKHR(o.CompositeAlpha)
	createInfo.presentMode = C.VkPresentModeKHR(o.PresentMode)

	createInfo.clipped = C.VK_FALSE
	if o.Clipped {
		createInfo.clipped = C.VK_TRUE
	}

	createInfo.oldSwapchain = nil
	if o.OldSwapchain != nil {
		createInfo.oldSwapchain = (C.VkSwapchainKHR)(unsafe.Pointer(o.OldSwapchain.Handle()))
	}

	return preallocatedPointer, nil
}
