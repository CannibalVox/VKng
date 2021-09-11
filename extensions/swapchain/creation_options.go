package ext_swapchain

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	ext_surface2 "github.com/CannibalVox/VKng/extensions/surface"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type CreationOptions struct {
	Surface ext_surface2.Surface

	MinImageCount uint32

	ImageFormat      core.DataFormat
	ImageColorSpace  ext_surface2.ColorSpace
	ImageExtent      core.Extent2D
	ImageArrayLayers uint32
	ImageUsage       core.ImageUsages

	SharingMode        core.SharingMode
	QueueFamilyIndices []int

	PreTransform   ext_surface2.SurfaceTransforms
	CompositeAlpha ext_surface2.CompositeAlphaModes
	PresentMode    ext_surface2.PresentMode

	Clipped      bool
	OldSwapchain Swapchain

	Next core.Options
}

func (o *CreationOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkSwapchainCreateInfoKHR)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkSwapchainCreateInfoKHR{}))))
	createInfo.sType = C.VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR
	createInfo.flags = 0

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
