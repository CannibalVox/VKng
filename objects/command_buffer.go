package objects

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "unsafe"

type CommandPool struct {
	handle C.VkCommandPool
	device C.VkDevice
}

func (p *CommandPool) Handle() uintptr {
	return uintptr(unsafe.Pointer(p.handle))
}

func (p *CommandPool) Destroy() {
	C.vkDestroyCommandPool(p.device, p.handle, nil)
}


type CommandBuffer struct {
	device C.VkDevice
	pool C.VkCommandPool
	handle C.VkCommandBuffer
}

func (c *CommandBuffer) Handle() uintptr {
	return uintptr(unsafe.Pointer(c.handle))
}

func (c *CommandBuffer) Destroy() {
	C.vkFreeCommandBuffers(c.device, c.pool, 1, &c.handle)
}

func DestroyBuffers(buffers []*CommandBuffer) {
	for _, buffer := range buffers {
		buffer.Destroy()
	}
}
