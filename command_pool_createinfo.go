package VKng

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"github.com/palantir/stacktrace"
	"unsafe"
)

type CommandPoolOptions struct {
	GraphicsQueueFamily *uint32

	Next core.Options
}

func (o *CommandPoolOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	if o.GraphicsQueueFamily == nil {
		return nil, stacktrace.NewError("attempted to create a command pool without setting GraphicsQueueFamilyIndex")
	}

	familyIndex := *o.GraphicsQueueFamily

	cmdPoolCreate := (*C.VkCommandPoolCreateInfo)(allocator.Malloc(int(unsafe.Sizeof([1]C.VkCommandPoolCreateInfo{}))))
	cmdPoolCreate.sType = C.VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO
	cmdPoolCreate.flags = 0
	cmdPoolCreate.queueFamilyIndex = C.uint32_t(familyIndex)

	var next unsafe.Pointer
	var err error

	if o.Next != nil {
		next, err = o.Next.AllocForC(allocator)
	}
	if err != nil {
		return nil, err
	}

	cmdPoolCreate.pNext = next

	return unsafe.Pointer(cmdPoolCreate), nil
}
