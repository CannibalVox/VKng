package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type PipelineLayoutHandle C.VkPipelineLayout
type PipelineLayout struct {
	device C.VkDevice
	handle C.VkPipelineLayout
}

func (l *PipelineLayout) Handle() PipelineLayoutHandle {
	return PipelineLayoutHandle(l.handle)
}

func (l *PipelineLayout) Destroy() {
	C.vkDestroyPipelineLayout(l.device, l.handle, nil)
}

func CreatePipelineLayout(allocator cgoalloc.Allocator, device *core.Device, o *PipelineLayoutOptions) (*PipelineLayout, VKng.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, VKng.VKErrorUnknown, err
	}

	deviceHandle := C.VkDevice(unsafe.Pointer(device.Handle()))

	var pipelineLayout C.VkPipelineLayout
	res := VKng.Result(C.vkCreatePipelineLayout(deviceHandle, (*C.VkPipelineLayoutCreateInfo)(createInfo), nil, &pipelineLayout))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &PipelineLayout{handle: pipelineLayout, device: deviceHandle}, res, nil
}
