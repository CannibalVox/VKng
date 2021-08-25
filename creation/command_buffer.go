package creation

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoalloc"
	"github.com/palantir/stacktrace"
	"unsafe"
)

type CommandBufferLevel uint

const (
	Primary CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_PRIMARY
	Secondary CommandBufferLevel = C.VK_COMMAND_BUFFER_LEVEL_SECONDARY
	Unset CommandBufferLevel = 0xFFFFFFFF
)

type CommandBufferOptions struct {
	Level CommandBufferLevel
	BufferCount int
	CommandBufferHandle unsafe.Pointer

	Next Options
}

func (o *CommandBufferOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	if o.Level == Unset {
		return nil, stacktrace.NewError("attempted to create command buffers without setting Level")
	}
	if o.BufferCount == 0 {
		return nil, stacktrace.NewError("attempted to create 0 command buffers")
	}

	createInfo := (*C.VkCommandBufferAllocateInfo)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkCommandBufferAllocateInfo{}))))
	createInfo.sType = C.VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO
	createInfo.level = C.VkCommandBufferLevel(o.Level)
	createInfo.commandBufferCount = C.uint32_t(o.BufferCount)
	createInfo.commandPool = (C.VkCommandPool)(o.CommandBufferHandle)

	var next unsafe.Pointer
	var err error
	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}
	if err != nil {
		return nil, err
	}
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}
