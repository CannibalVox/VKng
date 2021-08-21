package VKng

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoalloc"
	"github.com/palantir/stacktrace"
	"unsafe"
)

type CommandPoolBuilder struct {
	deviceHandle C.VkDevice
	graphicsQueueFamily uint32
}

func (b *CommandPoolBuilder) GraphicsQueueFamilyIndex(index uint32) *CommandPoolBuilder {
	b.graphicsQueueFamily = index
	return b
}

func (b *CommandPoolBuilder) Build(allocator cgoalloc.Allocator) (*CommandPool, error) {
	if b.graphicsQueueFamily == 0xFFFFFFFF {
		return nil, stacktrace.NewError("attempted to create a command pool without setting GraphicsQueueFamilyIndex")
	}

	cmdPoolCreate := &C.VkCommandPoolCreateInfo  {
		sType: C.VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO,
		queueFamilyIndex: C.uint32_t(b.graphicsQueueFamily),
	}

	var cmdPoolHandle C.VkCommandPool
	res := C.vkCreateCommandPool(b.deviceHandle, cmdPoolCreate, nil, &cmdPoolHandle)
	err := VKResult(res).ToError()
	if err != nil {
		return nil, err
	}

	return &CommandPool{handle: cmdPoolHandle, device: b.deviceHandle}, nil
}

type CommandPool struct {
	handle C.VkCommandPool
	device C.VkDevice
}

func (p *CommandPool) Destroy() {
	C.vkDestroyCommandPool(p.device, p.handle, nil)
}

func (p *CommandPool) CommandBufferBuilder() *CommandBufferBuilder {
	return &CommandBufferBuilder{
		deviceHandle: p.device,
		commandPoolHandle: p.handle,
		level: Unset,
	}
}

type CommandBufferLevel uint

const (
	Primary CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	Secondary CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_SECONDARY
	Unset CommandBufferLevel = 0xFFFFFFFF
)

type CommandBufferBuilder struct {
	deviceHandle C.VkDevice
	commandPoolHandle C.VkCommandPool

	level CommandBufferLevel
}

func (b *CommandBufferBuilder) Level(l CommandBufferLevel) *CommandBufferBuilder {
	b.level = l
	return b
}

func (b *CommandBufferBuilder) Build(allocator cgoalloc.Allocator, count int) ([]*CommandBuffer, error) {
	if b.level == Unset {
		return nil, stacktrace.NewError("attempted to create command buffers without setting Level")
	}
	if count == 0 {
		return nil, stacktrace.NewError("attempted to create 0 command buffers")
	}

	commandBufferPtr := allocator.Malloc(count * int(unsafe.Sizeof([1]C.VkCommandBuffer{})))
	defer allocator.Free(commandBufferPtr)

	cmdBufferCreate := &C.VkCommandBufferAllocateInfo {
		sType: C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO,
		commandPool: b.commandPoolHandle,
		level: C.VkCommandBufferLevel(b.level),
		commandBufferCount: C.uint32_t(count),
	}

	res := C.vkAllocateCommandBuffers(b.deviceHandle, cmdBufferCreate, (*C.VkCommandBuffer)(commandBufferPtr))
	err := VKResult(res).ToError()
	if err != nil {
		return nil, err
	}

	commandBufferArray := (*[1<<30]C.VkCommandBuffer)(commandBufferPtr)
	var result []*CommandBuffer
	for i := 0; i < count; i++ {
		result = append(result, &CommandBuffer{handle: commandBufferArray[i]})
	}

	return result, nil
}

type CommandBuffer struct {
	device C.VkDevice
	pool C.VkCommandPool
	handle C.VkCommandBuffer
}

func (c *CommandBuffer) Destroy() {
	C.vkFreeCommandBuffers(c.device, c.pool, 1, &c.handle)
}

func DestroyBuffers(buffers []*CommandBuffer) {
	for _, buffer := range buffers {
		buffer.Destroy()
	}
}
