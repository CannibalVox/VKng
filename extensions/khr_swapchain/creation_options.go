package khr_swapchain

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	ext_surface "github.com/CannibalVox/VKng/extensions/khr_surface"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const (
	ObjectTypeSwapchain common.ObjectType = C.VK_OBJECT_TYPE_SWAPCHAIN_KHR

	ImageLayoutPresentSrc common.ImageLayout = C.VK_IMAGE_LAYOUT_PRESENT_SRC_KHR

	VKErrorOutOfDate common.VkResult = C.VK_ERROR_OUT_OF_DATE_KHR
	VKSuboptimal     common.VkResult = C.VK_SUBOPTIMAL_KHR
)

func init() {
	ObjectTypeSwapchain.Register("Swapchain")

	ImageLayoutPresentSrc.Register("Present Src")

	VKErrorOutOfDate.Register("out of date")
	VKSuboptimal.Register("Suboptimal")
}

type CreationOptions struct {
	Surface ext_surface.Surface

	MinImageCount uint32

	ImageFormat      common.DataFormat
	ImageColorSpace  ext_surface.ColorSpace
	ImageExtent      common.Extent2D
	ImageArrayLayers uint32
	ImageUsage       common.ImageUsages

	SharingMode        common.SharingMode
	QueueFamilyIndices []int

	PreTransform   ext_surface.SurfaceTransforms
	CompositeAlpha ext_surface.CompositeAlphaModes
	PresentMode    ext_surface.PresentMode

	Clipped      bool
	OldSwapchain Swapchain

	common.HaveNext
}

func (o *CreationOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkSwapchainCreateInfoKHR{})))
	}
	createInfo := (*C.VkSwapchainCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
	createInfo.flags = 0
	createInfo.pNext = next

	createInfo.surface = C.VkSurfaceKHR(unsafe.Pointer(o.Surface.Handle()))
	createInfo.minImageCount = C.uint32_t(o.MinImageCount)

	createInfo.imageFormat = C.VkFormat(o.ImageFormat)
	createInfo.imageColorSpace = C.VkColorSpaceKHR(o.ImageColorSpace)
	createInfo.imageExtent.width = C.uint32_t(o.ImageExtent.Width)
	createInfo.imageExtent.height = C.uint32_t(o.ImageExtent.Height)
	createInfo.imageArrayLayers = C.uint32_t(o.ImageArrayLayers)
	createInfo.imageUsage = C.VkImageUsageFlags(o.ImageUsage)

	createInfo.imageSharingMode = C.VkSharingMode(o.SharingMode)
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
		createInfo.oldSwapchain = (C.VkSwapchainKHR)(o.OldSwapchain.Handle())
	}

	return preallocatedPointer, nil
}
