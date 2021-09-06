package commands

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
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

type CommandPoolHandle C.VkCommandPool
type CommandPool struct {
	handle C.VkCommandPool
	device C.VkDevice
}

func CreateCommandPool(allocator cgoalloc.Allocator, device *core.Device, o *CommandPoolOptions) (*CommandPool, VKng.Result, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, VKng.VKErrorUnknown, err
	}

	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))

	var cmdPoolHandle C.VkCommandPool
	res := VKng.Result(C.vkCreateCommandPool(deviceHandle, (*C.VkCommandPoolCreateInfo)(createInfo), nil, &cmdPoolHandle))
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	return &CommandPool{handle: cmdPoolHandle, device: deviceHandle}, res, nil
}

func (p *CommandPool) Handle() CommandPoolHandle {
	return CommandPoolHandle(p.handle)
}

func (p *CommandPool) Destroy() {
	C.vkDestroyCommandPool(p.device, p.handle, nil)
}
