package render_pass

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type FramebufferHandle C.VkFramebuffer
type Framebuffer struct {
	device C.VkDevice
	handle C.VkFramebuffer
}

func CreateFrameBuffer(allocator cgoalloc.Allocator, device *VKng.Device, o *FramebufferOptions) (*Framebuffer, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	deviceHandle := C.VkDevice(unsafe.Pointer(device.Handle()))
	var framebuffer C.VkFramebuffer

	res := C.vkCreateFramebuffer(deviceHandle, (*C.VkFramebufferCreateInfo)(createInfo), nil, &framebuffer)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &Framebuffer{device: deviceHandle, handle: framebuffer}, nil
}

func (b *Framebuffer) Handle() FramebufferHandle {
	return FramebufferHandle(b.handle)
}

func (b *Framebuffer) Destroy() {
	C.vkDestroyFramebuffer(b.device, b.handle, nil)
}
