package objects

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
 */
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/creation"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type Device struct {
	handle C.VkDevice
}

func (d *Device) Handle() uintptr {
	return uintptr(unsafe.Pointer(d.handle))
}

func (d *Device) Destroy() {
	C.vkDestroyDevice(d.handle, nil)
}

func (d *Device) CreateCommandPool(allocator cgoalloc.Allocator, o *creation.CommandPoolOptions) (*CommandPool, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var cmdPoolHandle C.VkCommandPool
	res := C.vkCreateCommandPool(d.handle, (*C.VkCommandPoolCreateInfo)(createInfo), nil, &cmdPoolHandle)
	err = VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &CommandPool{handle: cmdPoolHandle, device: d.handle}, nil
}

func (d *Device) CreateCommandBuffers(allocator cgoalloc.Allocator, o *creation.CommandBufferOptions) ([]*CommandBuffer, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	commandBufferPtr := (*C.VkCommandBuffer)(arena.Malloc(o.BufferCount*int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))

	res := C.vkAllocateCommandBuffers(d.handle, (*C.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = VKng.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	commandBufferArray := ([]C.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []*CommandBuffer
	for i := 0; i < o.BufferCount; i++ {
		result = append(result, &CommandBuffer{handle: commandBufferArray[i]})
	}

	return result, nil
}
