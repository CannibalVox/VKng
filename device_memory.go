package VKng

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type DeviceMemoryOptions struct {
	AllocationSize  uint64
	MemoryTypeIndex int

	Next core.Options
}

func (o *DeviceMemoryOptions) AllocForC(allocator *cgoalloc.ArenaAllocator) (unsafe.Pointer, error) {
	createInfo := (*C.VkMemoryAllocateInfo)(allocator.Malloc(C.sizeof_struct_VkMemoryAllocateInfo))

	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
	createInfo.allocationSize = C.VkDeviceSize(o.AllocationSize)
	createInfo.memoryTypeIndex = C.uint32_t(o.MemoryTypeIndex)

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

type DeviceMemoryHandle C.VkDeviceMemory
type DeviceMemory struct {
	device C.VkDevice
	handle C.VkDeviceMemory
}

func (m *DeviceMemory) Handle() DeviceMemoryHandle {
	return DeviceMemoryHandle(m.handle)
}

func (m *DeviceMemory) Free() {
	C.vkFreeMemory(m.device, m.handle, nil)
}

func (m *DeviceMemory) MapMemory(offset uint64, size uint64) (unsafe.Pointer, core.Result, error) {
	var data unsafe.Pointer
	res := core.Result(C.vkMapMemory(m.device, m.handle, C.VkDeviceSize(offset), C.VkDeviceSize(size), 0, &data))
	err := res.ToError()
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (m *DeviceMemory) UnmapMemory() {
	C.vkUnmapMemory(m.device, m.handle)
}
