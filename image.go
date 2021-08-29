package VKng

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
)

type ImageHandle C.VkImage
type Image struct {
	handle C.VkImage
	device C.VkDevice
}

func CreateFromHandles(handle ImageHandle, device DeviceHandle) *Image {
	return &Image{handle: C.VkImage(handle), device: C.VkDevice(device)}
}

func (i *Image) Handle() ImageHandle {
	return ImageHandle(i.handle)
}

func (i *Image) CreateImageView(allocator cgoalloc.Allocator, o *ImageViewOptions) (*ImageView, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfoPtr, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	createInfo := (*C.VkImageViewCreateInfo)(createInfoPtr)
	createInfo.image = i.handle

	var imageViewHandle C.VkImageView

	res := C.vkCreateImageView(i.device, createInfo, nil, &imageViewHandle)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &ImageView{handle: imageViewHandle, device: i.device}, nil
}
