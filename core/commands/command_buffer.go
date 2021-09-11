package commands

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/VKng/core/pipeline"
	"github.com/CannibalVox/VKng/core/resource"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type CommandBuffer struct {
	loader *loader.Loader
	device loader.VkDevice
	pool   loader.VkCommandPool
	handle loader.VkCommandBuffer
}

func CreateCommandBuffers(allocator cgoalloc.Allocator, device *resource.Device, o *CommandBufferOptions) ([]*CommandBuffer, loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, loader.VKErrorUnknown, err
	}

	commandBufferPtr := (*loader.VkCommandBuffer)(arena.Malloc(o.BufferCount * int(unsafe.Sizeof([1]loader.VkCommandBuffer{}))))

	res, err := device.Loader().VkAllocateCommandBuffers(device.Handle(), (*loader.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]loader.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []*CommandBuffer
	for i := 0; i < o.BufferCount; i++ {
		result = append(result, &CommandBuffer{loader: device.Loader(), pool: o.CommandPool.handle, device: o.CommandPool.device, handle: commandBufferArray[i]})
	}

	return result, res, nil
}

func (c *CommandBuffer) Handle() loader.VkCommandBuffer {
	return c.handle
}

func (c *CommandBuffer) Destroy() error {
	// cgocheckpointer considers &(c.handle) to be a go pointer containing a go pointer, probably
	// because loader is a go pointer?  Weird but passing a pointer just to the handle works
	handle := c.handle
	return c.loader.VkFreeCommandBuffers(c.device, c.pool, 1, &handle)
}

func DestroyBuffers(allocator cgoalloc.Allocator, pool *CommandPool, buffers []*CommandBuffer) error {
	bufferCount := len(buffers)
	if bufferCount == 0 {
		return nil
	}

	destroyPtr := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{})))
	defer allocator.Free(destroyPtr)

	destroySlice := ([]loader.VkCommandBuffer)(unsafe.Slice((*loader.VkCommandBuffer)(destroyPtr), bufferCount))
	for i := 0; i < bufferCount; i++ {
		destroySlice[i] = buffers[i].handle
	}

	return pool.loader.VkFreeCommandBuffers(pool.device, pool.handle, loader.Uint32(bufferCount), (*loader.VkCommandBuffer)(destroyPtr))
}

func (c *CommandBuffer) Begin(allocator cgoalloc.Allocator, o *BeginOptions) (loader.VkResult, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return loader.VKErrorUnknown, err
	}

	return c.loader.VkBeginCommandBuffer(c.handle, (*loader.VkCommandBufferBeginInfo)(createInfo))
}

func (c *CommandBuffer) End() (loader.VkResult, error) {
	return c.loader.VkEndCommandBuffer(c.handle)
}

func (c *CommandBuffer) CmdBeginRenderPass(allocator cgoalloc.Allocator, contents SubpassContents, o *RenderPassBeginOptions) error {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return err
	}

	return c.loader.VkCmdBeginRenderPass(c.handle, (*loader.VkRenderPassBeginInfo)(createInfo), loader.VkSubpassContents(contents))
}

func (c *CommandBuffer) CmdEndRenderPass() error {
	return c.loader.VkCmdEndRenderPass(c.handle)
}

func (c *CommandBuffer) CmdBindPipeline(bindPoint core.PipelineBindPoint, pipeline *pipeline.Pipeline) error {
	return c.loader.VkCmdBindPipeline(c.handle, loader.VkPipelineBindPoint(bindPoint), pipeline.Handle())
}

func (c *CommandBuffer) CmdDraw(vertexCount, instanceCount int, firstVertex, firstInstance uint32) error {
	return c.loader.VkCmdDraw(c.handle, loader.Uint32(vertexCount), loader.Uint32(instanceCount), loader.Uint32(firstVertex), loader.Uint32(firstInstance))
}

func (c *CommandBuffer) CmdDrawIndexed(indexCount, instanceCount int, firstIndex uint32, vertexOffset int, firstInstance uint32) error {
	return c.loader.VkCmdDrawIndexed(c.handle, loader.Uint32(indexCount), loader.Uint32(instanceCount), loader.Uint32(firstIndex), loader.Int32(vertexOffset), loader.Uint32(firstInstance))
}

func (c *CommandBuffer) CmdBindVertexBuffers(allocator cgoalloc.Allocator, firstBinding uint32, buffers []*resource.Buffer, bufferOffsets []int) error {
	bufferCount := len(buffers)

	bufferArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof([1]C.VkBuffer{})))
	defer allocator.Free(bufferArrayUnsafe)

	offsetArrayUnsafe := allocator.Malloc(bufferCount * int(unsafe.Sizeof(C.VkDeviceSize(0))))
	defer allocator.Free(offsetArrayUnsafe)

	bufferArrayPtr := (*loader.VkBuffer)(bufferArrayUnsafe)
	offsetArrayPtr := (*loader.VkDeviceSize)(offsetArrayUnsafe)

	bufferArraySlice := ([]loader.VkBuffer)(unsafe.Slice(bufferArrayPtr, bufferCount))
	offsetArraySlice := ([]loader.VkDeviceSize)(unsafe.Slice(offsetArrayPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		bufferArraySlice[i] = buffers[i].Handle()
		offsetArraySlice[i] = loader.VkDeviceSize(bufferOffsets[i])
	}

	return c.loader.VkCmdBindVertexBuffers(c.handle, loader.Uint32(firstBinding), loader.Uint32(bufferCount), bufferArrayPtr, offsetArrayPtr)
}

func (c *CommandBuffer) CmdBindIndexBuffer(buffer *resource.Buffer, offset int, indexType core.IndexType) error {
	return c.loader.VkCmdBindIndexBuffer(c.handle, buffer.Handle(), loader.VkDeviceSize(offset), loader.VkIndexType(indexType))
}