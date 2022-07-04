package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

const (
	BufferCreateSparseBinding   BufferCreateFlags = C.VK_BUFFER_CREATE_SPARSE_BINDING_BIT
	BufferCreateSparseResidency BufferCreateFlags = C.VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT
	BufferCreateSparseAliased   BufferCreateFlags = C.VK_BUFFER_CREATE_SPARSE_ALIASED_BIT

	BufferUsageTransferSrc        BufferUsages = C.VK_BUFFER_USAGE_TRANSFER_SRC_BIT
	BufferUsageTransferDst        BufferUsages = C.VK_BUFFER_USAGE_TRANSFER_DST_BIT
	BufferUsageUniformTexelBuffer BufferUsages = C.VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT
	BufferUsageStorageTexelBuffer BufferUsages = C.VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT
	BufferUsageUniformBuffer      BufferUsages = C.VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT
	BufferUsageStorageBuffer      BufferUsages = C.VK_BUFFER_USAGE_STORAGE_BUFFER_BIT
	BufferUsageIndexBuffer        BufferUsages = C.VK_BUFFER_USAGE_INDEX_BUFFER_BIT
	BufferUsageVertexBuffer       BufferUsages = C.VK_BUFFER_USAGE_VERTEX_BUFFER_BIT
	BufferUsageIndirectBuffer     BufferUsages = C.VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT

	SharingExclusive  SharingMode = C.VK_SHARING_MODE_EXCLUSIVE
	SharingConcurrent SharingMode = C.VK_SHARING_MODE_CONCURRENT
)

func init() {
	BufferCreateSparseBinding.Register("Sparse Binding")
	BufferCreateSparseResidency.Register("Sparse Residency")
	BufferCreateSparseAliased.Register("Sparse Aliased")

	BufferUsageTransferSrc.Register("Transfer Source")
	BufferUsageTransferDst.Register("Transfer Destination")
	BufferUsageUniformTexelBuffer.Register("Uniform Texel Buffer")
	BufferUsageStorageTexelBuffer.Register("Storage Texel Buffer")
	BufferUsageUniformBuffer.Register("Uniform Buffer")
	BufferUsageStorageBuffer.Register("Storage Buffer")
	BufferUsageIndexBuffer.Register("Index Buffer")
	BufferUsageVertexBuffer.Register("Vertex Buffer")
	BufferUsageIndirectBuffer.Register("Indirect Buffer")

	SharingExclusive.Register("Exclusive")
	SharingConcurrent.Register("Concurrent")
}

type BufferCreateOptions struct {
	Flags              BufferCreateFlags
	BufferSize         int
	Usage              BufferUsages
	SharingMode        SharingMode
	QueueFamilyIndices []int

	common.NextOptions
}

func (o BufferCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == unsafe.Pointer(nil) {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkBufferCreateInfo)
	}
	createInfo := (*C.VkBufferCreateInfo)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO
	createInfo.flags = C.VkBufferCreateFlags(o.Flags)
	createInfo.pNext = next
	createInfo.size = C.VkDeviceSize(o.BufferSize)
	createInfo.usage = C.VkBufferUsageFlags(o.Usage)
	createInfo.sharingMode = C.VkSharingMode(o.SharingMode)

	queueFamilyCount := len(o.QueueFamilyIndices)
	createInfo.queueFamilyIndexCount = C.uint32_t(queueFamilyCount)
	createInfo.pQueueFamilyIndices = nil

	if queueFamilyCount > 0 {
		indicesPtr := (*C.uint32_t)(allocator.Malloc(queueFamilyCount * int(unsafe.Sizeof(C.uint32_t(0)))))
		indicesSlice := ([]C.uint32_t)(unsafe.Slice(indicesPtr, queueFamilyCount))

		for i := 0; i < queueFamilyCount; i++ {
			indicesSlice[i] = C.uint32_t(o.QueueFamilyIndices[i])
		}

		createInfo.pQueueFamilyIndices = indicesPtr
	}

	return preallocatedPointer, nil
}
