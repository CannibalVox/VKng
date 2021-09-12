package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/resources"
	"github.com/CannibalVox/cgoparam"
)

type vulkanPipelineLayout struct {
	loader loader.Loader
	device loader.VkDevice
	handle loader.VkPipelineLayout
}

func (l *vulkanPipelineLayout) Handle() loader.VkPipelineLayout {
	return l.handle
}

func (l *vulkanPipelineLayout) Destroy() error {
	return l.loader.VkDestroyPipelineLayout(l.device, l.handle, nil)
}

func CreatePipelineLayout(device resources.Device, o *PipelineLayoutOptions) (PipelineLayout, loader.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := core.AllocOptions(arena, o)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	var pipelineLayout loader.VkPipelineLayout
	res, err := device.Loader().VkCreatePipelineLayout(device.Handle(), (*loader.VkPipelineLayoutCreateInfo)(createInfo), nil, &pipelineLayout)
	if err != nil {
		return nil, res, err
	}

	return &vulkanPipelineLayout{loader: device.Loader(), handle: pipelineLayout, device: device.Handle()}, res, nil
}
