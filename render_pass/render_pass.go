package render_pass

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

type RenderPassHandle C.VkRenderPass
type RenderPass struct {
	device C.VkDevice
	handle C.VkRenderPass
}

func CreateRenderPass(allocator cgoalloc.Allocator, device *VKng.Device, o *RenderPassOptions) (*RenderPass, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var renderPass C.VkRenderPass
	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))

	res := C.vkCreateRenderPass(deviceHandle, (*C.VkRenderPassCreateInfo)(createInfo), nil, &renderPass)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &RenderPass{device: deviceHandle, handle: renderPass}, nil
}

func (p *RenderPass) Handle() RenderPassHandle {
	return RenderPassHandle(p.handle)
}

func (p *RenderPass) Destroy() {
	C.vkDestroyRenderPass(p.device, p.handle, nil)
}
