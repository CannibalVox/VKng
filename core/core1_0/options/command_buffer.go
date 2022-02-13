package options

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/iface"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"unsafe"
)

type CommandBufferOptions struct {
	Level       common.CommandBufferLevel
	BufferCount int
	CommandPool iface.CommandPool

	core.HaveNext
}

func (o CommandBufferOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if o.Level == common.LevelUnset {
		return nil, errors.New("attempted to create command buffers without setting Level")
	}
	if o.BufferCount == 0 {
		return nil, errors.New("attempted to create 0 command buffers")
	}

	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof([1]C.VkCommandBufferAllocateInfo{})))
	}

	createInfo := (*C.VkCommandBufferAllocateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	createInfo.pNext = next

	createInfo.level = C.VkCommandBufferLevel(o.Level)
	createInfo.commandBufferCount = C.uint32_t(o.BufferCount)
	createInfo.commandPool = C.VkCommandPool(unsafe.Pointer(o.CommandPool.Handle()))

	return unsafe.Pointer(createInfo), nil
}
