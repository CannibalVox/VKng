package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/core1_1"
	"github.com/CannibalVox/VKng/core/driver"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type VulkanPipelineCache struct {
	Driver              driver.Driver
	Device              driver.VkDevice
	PipelineCacheHandle driver.VkPipelineCache

	MaximumAPIVersion common.APIVersion

	PipelineCache1_1 core1_1.PipelineCache
}

func (c *VulkanPipelineCache) Handle() driver.VkPipelineCache {
	return c.PipelineCacheHandle
}

func (c *VulkanPipelineCache) Core1_1() core1_1.PipelineCache {
	return c.PipelineCache1_1
}

func (c *VulkanPipelineCache) Destroy(callbacks *driver.AllocationCallbacks) {
	c.Driver.VkDestroyPipelineCache(c.Device, c.PipelineCacheHandle, callbacks.Handle())
	c.Driver.ObjectStore().Delete(driver.VulkanHandle(c.PipelineCacheHandle), c)
}

func (c *VulkanPipelineCache) CacheData() ([]byte, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	cacheSizePtr := arena.Malloc(int(unsafe.Sizeof(C.size_t(0))))
	cacheSize := (*driver.Size)(cacheSizePtr)

	res, err := c.Driver.VkGetPipelineCacheData(c.Device, c.PipelineCacheHandle, cacheSize, nil)
	if err != nil {
		return nil, res, err
	}

	cacheDataPtr := arena.Malloc(int(*cacheSize))

	res, err = c.Driver.VkGetPipelineCacheData(c.Device, c.PipelineCacheHandle, cacheSize, cacheDataPtr)
	if err != nil {
		return nil, res, err
	}

	outData := make([]byte, *cacheSize)
	inData := ([]byte)(unsafe.Slice((*byte)(cacheDataPtr), int(*cacheSize)))
	copy(outData, inData)

	return outData, res, nil
}

func (c *VulkanPipelineCache) MergePipelineCaches(srcCaches []core1_0.PipelineCache) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	srcCount := len(srcCaches)
	srcPtr := (*driver.VkPipelineCache)(arena.Malloc(srcCount * int(unsafe.Sizeof([1]driver.VkPipelineCache{}))))
	srcSlice := ([]driver.VkPipelineCache)(unsafe.Slice(srcPtr, srcCount))

	for i := 0; i < srcCount; i++ {
		srcSlice[i] = srcCaches[i].Handle()
	}

	return c.Driver.VkMergePipelineCaches(c.Device, c.PipelineCacheHandle, driver.Uint32(srcCount), srcPtr)
}