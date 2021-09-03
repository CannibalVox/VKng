package VKng

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"time"
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

func (d *Device) CreateImageView(allocator cgoalloc.Allocator, o *ImageViewOptions) (*ImageView, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var imageViewHandle C.VkImageView

	res := C.vkCreateImageView(d.handle, (*C.VkImageViewCreateInfo)(createInfo), nil, &imageViewHandle)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &ImageView{handle: imageViewHandle, device: d.handle}, nil
}

func (d *Device) CreateSemaphore(allocator cgoalloc.Allocator, o *SemaphoreOptions) (*Semaphore, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var semaphoreHandle C.VkSemaphore

	res := C.vkCreateSemaphore(d.handle, (*C.VkSemaphoreCreateInfo)(createInfo), nil, &semaphoreHandle)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &Semaphore{device: d.handle, handle: semaphoreHandle}, nil
}

func (d *Device) CreateFence(allocator cgoalloc.Allocator, o *FenceOptions) (*Fence, error) {
	arena := cgoalloc.CreateArenaAllocator(allocator)
	defer arena.FreeAll()

	createInfo, err := o.AllocForC(arena)
	if err != nil {
		return nil, err
	}

	var fenceHandle C.VkFence

	res := C.vkCreateFence(d.handle, (*C.VkFenceCreateInfo)(createInfo), nil, &fenceHandle)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	return &Fence{device: d.handle, handle: fenceHandle}, nil
}

func (d *Device) WaitForIdle() error {
	res := C.vkDeviceWaitIdle(d.handle)
	return core.Result(res).ToError()
}

func (d *Device) WaitForFences(allocator cgoalloc.Allocator, waitForAll bool, timeout time.Duration, fences []*Fence) error {
	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))
	defer allocator.Free(fenceUnsafePtr)

	fencePtr := (*C.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]C.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].handle
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	res := C.vkWaitForFences(d.handle, C.uint32_t(fenceCount), fencePtr, C.uint(waitAllConst), C.uint64_t(core.TimeoutNanoseconds(timeout)))
	return core.Result(res).ToError()
}

func (d *Device) ResetFences(allocator cgoalloc.Allocator, fences []*Fence) error {
	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))
	defer allocator.Free(fenceUnsafePtr)

	fencePtr := (*C.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]C.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].handle
	}

	res := C.vkResetFences(d.handle, C.uint32_t(fenceCount), fencePtr)
	return core.Result(res).ToError()
}
