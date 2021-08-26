package objects

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type CommandPoolHandle C.VkCommandPool
type CommandPool struct {
	handle C.VkCommandPool
	device C.VkDevice
}

func (p *CommandPool) Handle() CommandPoolHandle {
	return CommandPoolHandle(p.handle)
}

func (p *CommandPool) Destroy() {
	C.vkDestroyCommandPool(p.device, p.handle, nil)
}

type CommandBufferHandle C.VkCommandBuffer
type CommandBuffer struct {
	device C.VkDevice
	pool C.VkCommandPool
	handle C.VkCommandBuffer
}

func (c *CommandBuffer) Handle() CommandBufferHandle {
	return CommandBufferHandle(c.handle)
}

func (c *CommandBuffer) Destroy() {
	C.vkFreeCommandBuffers(c.device, c.pool, 1, &c.handle)
}

func DestroyBuffers(buffers []*CommandBuffer) {
	for _, buffer := range buffers {
		buffer.Destroy()
	}
}
