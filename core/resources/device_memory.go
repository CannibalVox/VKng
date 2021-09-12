package resources

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/loader"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DeviceMemoryOptions struct {
	AllocationSize  int
	MemoryTypeIndex int

	core.HaveNext
}

func (o *DeviceMemoryOptions) AllocForC(allocator *cgoparam.Allocator, next unsafe.Pointer) (unsafe.Pointer, error) {
	createInfo := (*C.VkMemoryAllocateInfo)(allocator.Malloc(C.sizeof_struct_VkMemoryAllocateInfo))

	createInfo.sType = C.VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO
	createInfo.allocationSize = C.VkDeviceSize(o.AllocationSize)
	createInfo.memoryTypeIndex = C.uint32_t(o.MemoryTypeIndex)
	createInfo.pNext = next

	return unsafe.Pointer(createInfo), nil
}

type vulkanDeviceMemory struct {
	loader loader.Loader
	device loader.VkDevice
	handle loader.VkDeviceMemory
}

func (m *vulkanDeviceMemory) Handle() loader.VkDeviceMemory {
	return m.handle
}

func (m *vulkanDeviceMemory) Free() error {
	return m.loader.VkFreeMemory(m.device, m.handle, nil)
}

func (m *vulkanDeviceMemory) MapMemory(offset int, size int) (unsafe.Pointer, loader.VkResult, error) {
	var data unsafe.Pointer
	res, err := m.loader.VkMapMemory(m.device, m.handle, loader.VkDeviceSize(offset), loader.VkDeviceSize(size), 0, &data)
	if err != nil {
		return nil, res, err
	}

	return data, res, nil
}

func (m *vulkanDeviceMemory) UnmapMemory() error {
	return m.loader.VkUnmapMemory(m.device, m.handle)
}

func (m *vulkanDeviceMemory) WriteData(offset int, data interface{}) (loader.VkResult, error) {
	bufferSize := binary.Size(data)

	memoryPtr, res, err := m.MapMemory(offset, bufferSize)
	if err != nil {
		return res, err
	}
	defer m.UnmapMemory()

	dataBuffer := unsafe.Slice((*byte)(memoryPtr), bufferSize)

	buf := &bytes.Buffer{}
	err = binary.Write(buf, core.ByteOrder, data)
	if err != nil {
		return loader.VKErrorUnknown, err
	}

	copy(dataBuffer, buf.Bytes())
	return res, nil
}