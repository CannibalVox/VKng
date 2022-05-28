package internal1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"github.com/cockroachdb/errors"
	"time"
	"unsafe"
)

type VulkanDevice struct {
	DeviceDriver driver.Driver
	DeviceHandle driver.VkDevice

	MaximumAPIVersion      common.APIVersion
	ActiveDeviceExtensions map[string]struct{}
}

func (d *VulkanDevice) Driver() driver.Driver {
	return d.DeviceDriver
}

func (d *VulkanDevice) Handle() driver.VkDevice {
	return d.DeviceHandle
}

func (d *VulkanDevice) APIVersion() common.APIVersion {
	return d.MaximumAPIVersion
}

func (d *VulkanDevice) IsDeviceExtensionActive(extensionName string) bool {
	_, active := d.ActiveDeviceExtensions[extensionName]
	return active
}

func (d *VulkanDevice) Destroy(callbacks *driver.AllocationCallbacks) {
	d.DeviceDriver.VkDestroyDevice(d.DeviceHandle, callbacks.Handle())
	d.DeviceDriver.ObjectStore().Delete(driver.VulkanHandle(d.DeviceHandle))
}

func (d *VulkanDevice) WaitForIdle() (common.VkResult, error) {
	return d.DeviceDriver.VkDeviceWaitIdle(d.DeviceHandle)
}

func (d *VulkanDevice) WaitForFences(waitForAll bool, timeout time.Duration, fences []core1_0.Fence) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)

	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	waitAllConst := C.VK_FALSE
	if waitForAll {
		waitAllConst = C.VK_TRUE
	}

	return d.DeviceDriver.VkWaitForFences(d.DeviceHandle, driver.Uint32(fenceCount), fencePtr, driver.VkBool32(waitAllConst), driver.Uint64(common.TimeoutNanoseconds(timeout)))
}

func (d *VulkanDevice) ResetFences(fences []core1_0.Fence) (common.VkResult, error) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	fenceCount := len(fences)
	fenceUnsafePtr := allocator.Malloc(fenceCount * int(unsafe.Sizeof([1]C.VkFence{})))

	fencePtr := (*driver.VkFence)(fenceUnsafePtr)
	fenceSlice := ([]driver.VkFence)(unsafe.Slice(fencePtr, fenceCount))
	for i := 0; i < fenceCount; i++ {
		fenceSlice[i] = fences[i].Handle()
	}

	return d.DeviceDriver.VkResetFences(d.DeviceHandle, driver.Uint32(fenceCount), fencePtr)
}

func (d *VulkanDevice) UpdateDescriptorSets(writes []core1_0.WriteDescriptorSetOptions, copies []core1_0.CopyDescriptorSetOptions) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	writeCount := len(writes)
	copyCount := len(copies)

	var err error
	var writePtr *C.VkWriteDescriptorSet
	var copyPtr *C.VkCopyDescriptorSet

	if writeCount > 0 {
		writePtr, err = common.AllocOptionSlice[C.VkWriteDescriptorSet, core1_0.WriteDescriptorSetOptions](arena, writes)
		if err != nil {
			return err
		}
	}

	if copyCount > 0 {
		copyPtr, err = common.AllocOptionSlice[C.VkCopyDescriptorSet, core1_0.CopyDescriptorSetOptions](arena, copies)
		if err != nil {
			return err
		}
	}

	d.DeviceDriver.VkUpdateDescriptorSets(d.DeviceHandle, driver.Uint32(writeCount), (*driver.VkWriteDescriptorSet)(unsafe.Pointer(writePtr)), driver.Uint32(copyCount), (*driver.VkCopyDescriptorSet)(unsafe.Pointer(copyPtr)))
	return nil
}

func (d *VulkanDevice) FlushMappedMemoryRanges(ranges []core1_0.MappedMemoryRangeOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos, err := common.AllocOptionSlice[C.VkMappedMemoryRange, core1_0.MappedMemoryRangeOptions](arena, ranges)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.DeviceDriver.VkFlushMappedMemoryRanges(d.DeviceHandle, driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (d *VulkanDevice) InvalidateMappedMemoryRanges(ranges []core1_0.MappedMemoryRangeOptions) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	rangeCount := len(ranges)
	createInfos, err := common.AllocOptionSlice[C.VkMappedMemoryRange, core1_0.MappedMemoryRangeOptions](arena, ranges)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return d.DeviceDriver.VkInvalidateMappedMemoryRanges(d.DeviceHandle, driver.Uint32(rangeCount), (*driver.VkMappedMemoryRange)(unsafe.Pointer(createInfos)))
}

func (d *VulkanDevice) CreateBufferView(allocationCallbacks *driver.AllocationCallbacks, options core1_0.BufferViewCreateOptions) (core1_0.BufferView, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, options)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var bufferViewHandle driver.VkBufferView

	res, err := d.DeviceDriver.VkCreateBufferView(d.DeviceHandle, (*driver.VkBufferViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferViewHandle)
	if err != nil {
		return nil, res, err
	}

	bufferView := CreateBufferViewObject(d.DeviceDriver, d.DeviceHandle, bufferViewHandle, d.MaximumAPIVersion)

	return bufferView, res, nil
}

func (d *VulkanDevice) CreateShaderModule(allocationCallbacks *driver.AllocationCallbacks, o core1_0.ShaderModuleCreateOptions) (core1_0.ShaderModule, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var shaderModuleHandle driver.VkShaderModule
	res, err := d.DeviceDriver.VkCreateShaderModule(d.DeviceHandle, (*driver.VkShaderModuleCreateInfo)(createInfo), allocationCallbacks.Handle(), &shaderModuleHandle)
	if err != nil {
		return nil, res, err
	}

	shaderModule := CreateShaderModuleObject(d.DeviceDriver, d.DeviceHandle, shaderModuleHandle, d.MaximumAPIVersion)

	return shaderModule, res, nil
}

func (d *VulkanDevice) CreateImageView(allocationCallbacks *driver.AllocationCallbacks, o core1_0.ImageViewCreateOptions) (core1_0.ImageView, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var imageViewHandle driver.VkImageView

	res, err := d.DeviceDriver.VkCreateImageView(d.DeviceHandle, (*driver.VkImageViewCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageViewHandle)
	if err != nil {
		return nil, res, err
	}

	imageView := CreateImageViewObject(d.DeviceDriver, d.DeviceHandle, imageViewHandle, d.MaximumAPIVersion)

	return imageView, res, nil
}

func (d *VulkanDevice) CreateSemaphore(allocationCallbacks *driver.AllocationCallbacks, o core1_0.SemaphoreCreateOptions) (core1_0.Semaphore, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var semaphoreHandle driver.VkSemaphore

	res, err := d.DeviceDriver.VkCreateSemaphore(d.DeviceHandle, (*driver.VkSemaphoreCreateInfo)(createInfo), allocationCallbacks.Handle(), &semaphoreHandle)
	if err != nil {
		return nil, res, err
	}

	semaphore := CreateSemaphoreObject(d.DeviceDriver, d.DeviceHandle, semaphoreHandle, d.MaximumAPIVersion)

	return semaphore, res, nil
}

func (d *VulkanDevice) CreateFence(allocationCallbacks *driver.AllocationCallbacks, o core1_0.FenceCreateOptions) (core1_0.Fence, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var fenceHandle driver.VkFence

	res, err := d.DeviceDriver.VkCreateFence(d.DeviceHandle, (*driver.VkFenceCreateInfo)(createInfo), allocationCallbacks.Handle(), &fenceHandle)
	if err != nil {
		return nil, res, err
	}

	fence := CreateFenceObject(d.DeviceDriver, d.DeviceHandle, fenceHandle, d.MaximumAPIVersion)

	return fence, res, nil
}

func (d *VulkanDevice) CreateBuffer(allocationCallbacks *driver.AllocationCallbacks, o core1_0.BufferCreateOptions) (core1_0.Buffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var bufferHandle driver.VkBuffer

	res, err := d.DeviceDriver.VkCreateBuffer(d.DeviceHandle, (*driver.VkBufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &bufferHandle)
	if err != nil {
		return nil, res, err
	}

	buffer := CreateBufferObject(d.DeviceDriver, d.DeviceHandle, bufferHandle, d.MaximumAPIVersion)

	return buffer, res, nil
}

func (d *VulkanDevice) CreateDescriptorSetLayout(allocationCallbacks *driver.AllocationCallbacks, o core1_0.DescriptorSetLayoutCreateOptions) (core1_0.DescriptorSetLayout, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var descriptorSetLayoutHandle driver.VkDescriptorSetLayout

	res, err := d.DeviceDriver.VkCreateDescriptorSetLayout(d.DeviceHandle, (*driver.VkDescriptorSetLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorSetLayoutHandle)
	if err != nil {
		return nil, res, err
	}

	descriptorSetLayout := CreateDescriptorSetLayoutObject(d.DeviceDriver, d.DeviceHandle, descriptorSetLayoutHandle, d.MaximumAPIVersion)

	return descriptorSetLayout, res, nil
}

func (d *VulkanDevice) CreateDescriptorPool(allocationCallbacks *driver.AllocationCallbacks, o core1_0.DescriptorPoolCreateOptions) (core1_0.DescriptorPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var descriptorPoolHandle driver.VkDescriptorPool

	res, err := d.DeviceDriver.VkCreateDescriptorPool(d.DeviceHandle, (*driver.VkDescriptorPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &descriptorPoolHandle)
	if err != nil {
		return nil, res, err
	}

	descriptorPool := CreateDescriptorPoolObject(d.DeviceDriver, d.DeviceHandle, descriptorPoolHandle, d.MaximumAPIVersion)

	return descriptorPool, res, nil
}

func (d *VulkanDevice) CreateCommandPool(allocationCallbacks *driver.AllocationCallbacks, o core1_0.CommandPoolCreateOptions) (core1_0.CommandPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var cmdPoolHandle driver.VkCommandPool
	res, err := d.DeviceDriver.VkCreateCommandPool(d.DeviceHandle, (*driver.VkCommandPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &cmdPoolHandle)
	if err != nil {
		return nil, res, err
	}

	commandPool := CreateCommandPoolObject(d.DeviceDriver, d.DeviceHandle, cmdPoolHandle, d.MaximumAPIVersion)

	return commandPool, res, nil
}

func (d *VulkanDevice) CreateEvent(allocationCallbacks *driver.AllocationCallbacks, o core1_0.EventCreateOptions) (core1_0.Event, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var eventHandle driver.VkEvent
	res, err := d.DeviceDriver.VkCreateEvent(d.DeviceHandle, (*driver.VkEventCreateInfo)(createInfo), allocationCallbacks.Handle(), &eventHandle)
	if err != nil {
		return nil, res, err
	}

	event := CreateEventObject(d.DeviceDriver, d.DeviceHandle, eventHandle, d.MaximumAPIVersion)

	return event, res, nil
}

func (d *VulkanDevice) CreateFramebuffer(allocationCallbacks *driver.AllocationCallbacks, o core1_0.FramebufferCreateOptions) (core1_0.Framebuffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var framebufferHandle driver.VkFramebuffer

	res, err := d.DeviceDriver.VkCreateFramebuffer(d.DeviceHandle, (*driver.VkFramebufferCreateInfo)(createInfo), allocationCallbacks.Handle(), &framebufferHandle)
	if err != nil {
		return nil, res, err
	}

	framebuffer := CreateFramebufferObject(d.DeviceDriver, d.DeviceHandle, framebufferHandle, d.MaximumAPIVersion)

	return framebuffer, res, nil
}

func (d *VulkanDevice) CreateGraphicsPipelines(pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.GraphicsPipelineCreateOptions) ([]core1_0.Pipeline, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtr, err := common.AllocOptionSlice[C.VkGraphicsPipelineCreateInfo, core1_0.GraphicsPipelineCreateOptions](arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	pipelinePtr := (*driver.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]driver.VkPipeline{}))))

	var pipelineCacheHandle driver.VkPipelineCache
	if pipelineCache != nil {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := d.DeviceDriver.VkCreateGraphicsPipelines(d.DeviceHandle, pipelineCacheHandle, driver.Uint32(pipelineCount), (*driver.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(pipelineCreateInfosPtr)), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []core1_0.Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))

	for i := 0; i < pipelineCount; i++ {
		pipeline := CreatePipelineObject(d.DeviceDriver, d.DeviceHandle, pipelineSlice[i], d.MaximumAPIVersion)
		output = append(output, pipeline)
	}

	return output, res, nil
}

func (d *VulkanDevice) CreateComputePipelines(pipelineCache core1_0.PipelineCache, allocationCallbacks *driver.AllocationCallbacks, o []core1_0.ComputePipelineCreateOptions) ([]core1_0.Pipeline, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	pipelineCount := len(o)

	pipelineCreateInfosPtr, err := common.AllocOptionSlice[C.VkComputePipelineCreateInfo, core1_0.ComputePipelineCreateOptions](arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	pipelinePtr := (*driver.VkPipeline)(arena.Malloc(pipelineCount * int(unsafe.Sizeof([1]driver.VkPipeline{}))))

	var pipelineCacheHandle driver.VkPipelineCache
	if pipelineCache != nil {
		pipelineCacheHandle = pipelineCache.Handle()
	}

	res, err := d.DeviceDriver.VkCreateComputePipelines(d.DeviceHandle, pipelineCacheHandle, driver.Uint32(pipelineCount), (*driver.VkComputePipelineCreateInfo)(unsafe.Pointer(pipelineCreateInfosPtr)), allocationCallbacks.Handle(), pipelinePtr)
	if err != nil {
		return nil, res, err
	}

	var output []core1_0.Pipeline
	pipelineSlice := ([]driver.VkPipeline)(unsafe.Slice(pipelinePtr, pipelineCount))

	for i := 0; i < pipelineCount; i++ {
		pipeline := CreatePipelineObject(d.DeviceDriver, d.DeviceHandle, pipelineSlice[i], d.MaximumAPIVersion)

		output = append(output, pipeline)
	}

	return output, res, nil
}

func (d *VulkanDevice) CreateImage(allocationCallbacks *driver.AllocationCallbacks, o core1_0.ImageCreateOptions) (core1_0.Image, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var imageHandle driver.VkImage
	res, err := d.DeviceDriver.VkCreateImage(d.DeviceHandle, (*driver.VkImageCreateInfo)(createInfo), allocationCallbacks.Handle(), &imageHandle)
	if err != nil {
		return nil, res, err
	}

	image := CreateImageObject(d.DeviceDriver, d.DeviceHandle, imageHandle, d.MaximumAPIVersion)

	return image, res, nil
}

func (d *VulkanDevice) CreatePipelineCache(allocationCallbacks *driver.AllocationCallbacks, o core1_0.PipelineCacheCreateOptions) (core1_0.PipelineCache, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var pipelineCacheHandle driver.VkPipelineCache
	res, err := d.DeviceDriver.VkCreatePipelineCache(d.DeviceHandle, (*driver.VkPipelineCacheCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineCacheHandle)
	if err != nil {
		return nil, res, err
	}

	pipelineCache := CreatePipelineCacheObject(d.DeviceDriver, d.DeviceHandle, pipelineCacheHandle, d.MaximumAPIVersion)

	return pipelineCache, res, nil
}

func (d *VulkanDevice) CreatePipelineLayout(allocationCallbacks *driver.AllocationCallbacks, o core1_0.PipelineLayoutCreateOptions) (core1_0.PipelineLayout, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var pipelineLayoutHandle driver.VkPipelineLayout
	res, err := d.DeviceDriver.VkCreatePipelineLayout(d.DeviceHandle, (*driver.VkPipelineLayoutCreateInfo)(createInfo), allocationCallbacks.Handle(), &pipelineLayoutHandle)
	if err != nil {
		return nil, res, err
	}

	pipelineLayout := CreatePipelineLayoutObject(d.DeviceDriver, d.DeviceHandle, pipelineLayoutHandle, d.MaximumAPIVersion)

	return pipelineLayout, res, nil
}

func (d *VulkanDevice) CreateQueryPool(allocationCallbacks *driver.AllocationCallbacks, o core1_0.QueryPoolCreateOptions) (core1_0.QueryPool, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var queryPoolHandle driver.VkQueryPool

	res, err := d.DeviceDriver.VkCreateQueryPool(d.DeviceHandle, (*driver.VkQueryPoolCreateInfo)(createInfo), allocationCallbacks.Handle(), &queryPoolHandle)
	if err != nil {
		return nil, res, err
	}

	queryPool := CreateQueryPoolObject(d.DeviceDriver, d.DeviceHandle, queryPoolHandle, d.MaximumAPIVersion)
	return queryPool, res, nil
}

func (d *VulkanDevice) CreateRenderPass(allocationCallbacks *driver.AllocationCallbacks, o core1_0.RenderPassCreateOptions) (core1_0.RenderPass, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var renderPassHandle driver.VkRenderPass

	res, err := d.DeviceDriver.VkCreateRenderPass(d.DeviceHandle, (*driver.VkRenderPassCreateInfo)(createInfo), allocationCallbacks.Handle(), &renderPassHandle)
	if err != nil {
		return nil, res, err
	}

	renderPass := CreateRenderPassObject(d.DeviceDriver, d.DeviceHandle, renderPassHandle, d.MaximumAPIVersion)

	return renderPass, res, nil
}

func (d *VulkanDevice) CreateSampler(allocationCallbacks *driver.AllocationCallbacks, o core1_0.SamplerCreateOptions) (core1_0.Sampler, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var samplerHandle driver.VkSampler

	res, err := d.DeviceDriver.VkCreateSampler(d.DeviceHandle, (*driver.VkSamplerCreateInfo)(createInfo), allocationCallbacks.Handle(), &samplerHandle)
	if err != nil {
		return nil, res, err
	}

	sampler := CreateSamplerObject(d.DeviceDriver, d.DeviceHandle, samplerHandle, d.MaximumAPIVersion)

	return sampler, res, nil
}

func (d *VulkanDevice) GetQueue(queueFamilyIndex int, queueIndex int) core1_0.Queue {

	var queueHandle driver.VkQueue

	d.DeviceDriver.VkGetDeviceQueue(d.DeviceHandle, driver.Uint32(queueFamilyIndex), driver.Uint32(queueIndex), &queueHandle)

	queue := CreateQueueObject(d.DeviceDriver, d.DeviceHandle, queueHandle, d.MaximumAPIVersion)

	return queue
}

func (d *VulkanDevice) AllocateMemory(allocationCallbacks *driver.AllocationCallbacks, o core1_0.MemoryAllocateOptions) (core1_0.DeviceMemory, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var deviceMemoryHandle driver.VkDeviceMemory

	deviceDriver := d.DeviceDriver
	deviceHandle := d.DeviceHandle

	res, err := deviceDriver.VkAllocateMemory(deviceHandle, (*driver.VkMemoryAllocateInfo)(createInfo), allocationCallbacks.Handle(), &deviceMemoryHandle)
	if err != nil {
		return nil, res, err
	}

	deviceMemory := CreateDeviceMemoryObject(deviceDriver, deviceHandle, deviceMemoryHandle, d.MaximumAPIVersion, o.AllocationSize)

	return deviceMemory, res, nil
}

func (d *VulkanDevice) FreeMemory(deviceMemory core1_0.DeviceMemory, allocationCallbacks *driver.AllocationCallbacks) {
	// This is really only here for a kind of API symmetry
	deviceMemory.Free(allocationCallbacks)
}

// Free a slice of command buffers which should all have the same device/driver/pool
// guaranteed to have at least one element
func (d *VulkanDevice) freeCommandBufferSlice(buffers []core1_0.CommandBuffer) {
	allocator := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(allocator)

	bufferCount := len(buffers)
	bufferDriver := buffers[0].Driver()
	bufferDevice := buffers[0].DeviceHandle()
	bufferPool := buffers[0].CommandPoolHandle()

	size := bufferCount * int(unsafe.Sizeof([1]C.VkCommandBuffer{}))
	bufferArrayPtr := (*driver.VkCommandBuffer)(allocator.Malloc(size))
	bufferArraySlice := ([]driver.VkCommandBuffer)(unsafe.Slice(bufferArrayPtr, bufferCount))

	for i := 0; i < bufferCount; i++ {
		bufferArraySlice[i] = buffers[i].Handle()
	}

	bufferDriver.VkFreeCommandBuffers(bufferDevice, bufferPool, driver.Uint32(bufferCount), bufferArrayPtr)

	objStore := d.DeviceDriver.ObjectStore()
	for i := 0; i < bufferCount; i++ {
		objStore.Delete(driver.VulkanHandle(buffers[i].Handle()))
	}
}

func (d *VulkanDevice) FreeCommandBuffers(buffers []core1_0.CommandBuffer) {
	bufferCount := len(buffers)
	if bufferCount == 0 {
		return
	}

	multimap := make(map[driver.VkCommandPool][]core1_0.CommandBuffer)
	for _, buffer := range buffers {
		poolHandle := buffer.CommandPoolHandle()
		existingSet := multimap[poolHandle]
		multimap[poolHandle] = append(existingSet, buffer)
	}

	for _, setBuffers := range multimap {
		d.freeCommandBufferSlice(setBuffers)
	}
}

func (d *VulkanDevice) AllocateCommandBuffers(o core1_0.CommandBufferAllocateOptions) ([]core1_0.CommandBuffer, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	if o.CommandPool == nil {
		return nil, core1_0.VKErrorUnknown, errors.New("no command pool provided to allocate from")
	}

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	device := o.CommandPool.DeviceHandle()

	commandBufferPtr := (*driver.VkCommandBuffer)(arena.Malloc(o.BufferCount * int(unsafe.Sizeof([1]driver.VkCommandBuffer{}))))

	res, err := o.CommandPool.Driver().VkAllocateCommandBuffers(device, (*driver.VkCommandBufferAllocateInfo)(createInfo), commandBufferPtr)
	err = res.ToError()
	if err != nil {
		return nil, res, err
	}

	commandBufferArray := ([]driver.VkCommandBuffer)(unsafe.Slice(commandBufferPtr, o.BufferCount))
	var result []core1_0.CommandBuffer

	for i := 0; i < o.BufferCount; i++ {
		commandBuffer := CreateCommandBufferObject(o.CommandPool.Driver(), o.CommandPool.Handle(), device, commandBufferArray[i], o.CommandPool.APIVersion())

		result = append(result, commandBuffer)
	}

	return result, res, nil
}

func (d *VulkanDevice) AllocateDescriptorSets(o core1_0.DescriptorSetAllocateOptions) ([]core1_0.DescriptorSet, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	if o.DescriptorPool == nil {
		return nil, core1_0.VKErrorUnknown, errors.New("no descriptor pool provided to allocate from")
	}

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	device := o.DescriptorPool.DeviceHandle()
	poolDriver := o.DescriptorPool.Driver()

	setCount := len(o.AllocationLayouts)
	descriptorSets := (*driver.VkDescriptorSet)(arena.Malloc(setCount * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))))

	res, err := poolDriver.VkAllocateDescriptorSets(device, (*driver.VkDescriptorSetAllocateInfo)(createInfo), descriptorSets)
	if err != nil {
		return nil, res, err
	}

	var sets []core1_0.DescriptorSet
	descriptorSetSlice := ([]driver.VkDescriptorSet)(unsafe.Slice(descriptorSets, setCount))

	for i := 0; i < setCount; i++ {
		descriptorSet := CreateDescriptorSetObject(poolDriver, device, o.DescriptorPool.Handle(), descriptorSetSlice[i], o.DescriptorPool.APIVersion())

		sets = append(sets, descriptorSet)
	}

	return sets, res, nil
}

// Free a slice of descriptor sets which should all have the same device/driver/pool
// guaranteed to have at least one element
func (d *VulkanDevice) freeDescriptorSetSlice(sets []core1_0.DescriptorSet) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	setSize := len(sets)
	arraySize := setSize * int(unsafe.Sizeof([1]C.VkDescriptorSet{}))

	arrayPtr := (*driver.VkDescriptorSet)(arena.Malloc(arraySize))
	arraySlice := ([]driver.VkDescriptorSet)(unsafe.Slice(arrayPtr, setSize))

	for i := 0; i < setSize; i++ {
		arraySlice[i] = sets[i].Handle()
	}

	setDriver := sets[0].Driver()
	pool := sets[0].DescriptorPoolHandle()
	device := sets[0].DeviceHandle()

	res, err := setDriver.VkFreeDescriptorSets(device, pool, driver.Uint32(setSize), arrayPtr)
	if err != nil {
		return res, err
	}

	objStore := setDriver.ObjectStore()
	for i := 0; i < setSize; i++ {
		objStore.Delete(driver.VulkanHandle(sets[i].Handle()))
	}

	return res, nil
}

func (d *VulkanDevice) FreeDescriptorSets(sets []core1_0.DescriptorSet) (common.VkResult, error) {
	poolMultimap := make(map[driver.VkDescriptorPool][]core1_0.DescriptorSet)

	for _, set := range sets {
		poolHandle := set.DescriptorPoolHandle()
		existingSet := poolMultimap[poolHandle]
		poolMultimap[poolHandle] = append(existingSet, set)
	}

	var res common.VkResult
	var err error
	for _, set := range poolMultimap {
		res, err = d.freeDescriptorSetSlice(set)
		if err != nil {
			return res, err
		}
	}

	return res, err
}
