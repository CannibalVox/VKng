package commands

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/pipeline"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type CommandBufferHandle C.VkCommandBuffer
type CommandBuffer struct {
	device C.VkDevice
	pool   C.VkCommandPool
	handle C.VkCommandBuffer
}

func CreateCommandBuffers(allocator cgoalloc.Allocator, device *VKng.Device, o *CommandBufferOptions) ([]*CommandBuffer, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	deviceHandle := (C.VkDevice)(unsafe.Pointer(device.Handle()))

	commandBufferPtr := (*C.VkCommandBuffer)(arena.Malloc(o.BufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))

	res := C.vkAllocateCommandBuffers(deviceHandle, (*C.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	commandBufferArray := ([]C.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []*CommandBuffer
	for i := 0; i < o.BufferCount; i++ {
		result = append(result, &CommandBuffer{pool: o.CommandPool.handle, device: deviceHandle, handle: commandBufferArray[i]})
	}

	return result, nil
}

func (c *CommandBuffer) Handle() CommandBufferHandle {
	return CommandBufferHandle(c.handle)
}

func (c *CommandBuffer) Destroy() {
	C.vkFreeCommandBuffers(c.device, c.pool, 1, &c.handle)
}

func DestroyBuffers(allocator cgoalloc.Allocator, pool *CommandPool, buffers []*CommandBuffer) {
	bufferCount := len(buffers)
	if bufferCount == 0 {
		return
	}

	destroyPtr := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{})))
	defer allocator.Free(destroyPtr)

	destroySlice := ([]C.VkCommandBuffer)(unsafe.Slice((*C.VkCommandBuffer)(destroyPtr), bufferCount))
	for i := 0; i < bufferCount; i++ {
		destroySlice[i] = buffers[i].handle
	}

	C.vkFreeCommandBuffers(pool.device, pool.handle, C.uint32_t(bufferCount), (*C.VkCommandBuffer)(destroyPtr))
}

func (c *CommandBuffer) Begin(allocator cgoalloc.Allocator, o *BeginOptions) error {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return err
	}

	res := C.vkBeginCommandBuffer(c.handle, (*C.VkCommandBufferBeginInfo)(createInfo))
	return core.Result(res).ToError()
}

func (c *CommandBuffer) End() error {
	res := C.vkEndCommandBuffer(c.handle)
	return core.Result(res).ToError()
}

func (c *CommandBuffer) CmdBeginRenderPass(allocator cgoalloc.Allocator, contents SubpassContents, o *RenderPassBeginOptions) error {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return err
	}

	C.vkCmdBeginRenderPass(c.handle, (*C.VkRenderPassBeginInfo)(createInfo), C.VkSubpassContents(contents))

	return nil
}

func (c *CommandBuffer) CmdEndRenderPass() {
	C.vkCmdEndRenderPass(c.handle)
}

func (c *CommandBuffer) CmdBindPipeline(bindPoint core.PipelineBindPoint, pipeline *pipeline.Pipeline) {
	pipelineHandle := (C.VkPipeline)(unsafe.Pointer(pipeline.Handle()))
	C.vkCmdBindPipeline(c.handle, C.VkPipelineBindPoint(bindPoint), pipelineHandle)
}

func (c *CommandBuffer) CmdDraw(vertexCount, instanceCount, firstVertex, firstInstance uint32) {
	C.vkCmdDraw(c.handle, C.uint32_t(vertexCount), C.uint32_t(instanceCount), C.uint32_t(firstVertex), C.uint32_t(firstInstance))
}
