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

type DeviceHandle C.VkDevice
type Device struct {
	handle C.VkDevice
}

func (d *Device) Handle() DeviceHandle {
	return DeviceHandle(d.handle)
}

func (d *Device) Destroy() {
	C.vkDestroyDevice(d.handle, nil)
}

func (d *Device) GetQueue(queueFamilyIndex int, queueIndex int) (*Queue, error) {
	var queueHandle C.VkQueue

	C.vkGetDeviceQueue(d.handle, C.uint32_t(queueFamilyIndex), C.uint32_t(queueIndex), &queueHandle)

	return &Queue{handle: QueueHandle(queueHandle)}, nil
}

func (d *Device) CreateCommandPool(allocator cgoalloc.Allocator, o *CommandPoolOptions) (*CommandPool, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var cmdPoolHandle C.VkCommandPool
	res := C.vkCreateCommandPool(d.handle, (*C.VkCommandPoolCreateInfo)(createInfo), nil, &cmdPoolHandle)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &CommandPool{handle: cmdPoolHandle, device: d.handle}, nil
}

func (d *Device) CreateCommandBuffers(allocator cgoalloc.Allocator, o *CommandBufferOptions) ([]*CommandBuffer, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	commandBufferPtr := (*C.VkCommandBuffer)(arena.Malloc(o.BufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))))

	res := C.vkAllocateCommandBuffers(d.handle, (*C.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	commandBufferArray := ([]C.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []*CommandBuffer
	for i := 0; i < o.BufferCount; i++ {
		result = append(result, &CommandBuffer{handle: commandBufferArray[i]})
	}

	return result, nil
}

func (d *Device) CreateShaderModule(allocator cgoalloc.Allocator, o *ShaderModuleOptions) (*ShaderModule, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var shaderModule C.VkShaderModule
	res := C.vkCreateShaderModule(d.handle, (*C.VkShaderModuleCreateInfo)(createInfo), nil, &shaderModule)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &ShaderModule{handle: shaderModule, device: d.handle}, nil
}
