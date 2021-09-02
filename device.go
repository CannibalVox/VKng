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

type DeviceHandle C.VkDevice
type Device struct {
	handle C.VkDevice
}

func (d *Device) Handle() DeviceHandle {
	return DeviceHandle(d.handle)
}

func (d *Device) Destroy() {
	C.vkDestroyDevice(d.handle, nil)
}

func (d *Device) GetQueue(queueFamilyIndex int, queueIndex int) (*Queue, error) {
	var queueHandle C.VkQueue

	C.vkGetDeviceQueue(d.handle, C.uint32_t(queueFamilyIndex), C.uint32_t(queueIndex), &queueHandle)

	return &Queue{handle: QueueHandle(queueHandle)}, nil
}

func (d *Device) CreateShaderModule(allocator cgoalloc.Allocator, o *ShaderModuleOptions) (*ShaderModule, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var shaderModule C.VkShaderModule
	res := C.vkCreateShaderModule(d.handle, (*C.VkShaderModuleCreateInfo)(createInfo), nil, &shaderModule)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &ShaderModule{handle: shaderModule, device: d.handle}, nil
}

func (d *Device) CreateImageView(allocator cgoalloc.Allocator, o *ImageViewOptions) (*ImageView, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var imageViewHandle C.VkImageView

	res := C.vkCreateImageView(d.handle, (*C.VkImageViewCreateInfo)(createInfo), nil, &imageViewHandle)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &ImageView{handle: imageViewHandle, device: d.handle}, nil
}
