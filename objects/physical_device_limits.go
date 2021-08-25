package objects

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type PhysicalDeviceLimits struct {
	MaxImageDimension1D uint32
	MaxImageDimension2D uint32
	MaxImageDimension3D uint32
	MaxImageDimensionCube uint32
	MaxImageArrayLayers uint32
	MaxTexelBufferElements uint32

	MaxUniformBufferRange uint32
	MaxStorageBufferRange uint32
	MaxPushConstantsSize uint32

	MaxMemoryAllocationCount uint32
	MaxSamplerAllocationCount uint32

	BufferImageGranularity uint64
	SparseAddressSpaceSize uint64

	MaxBoundDescriptorSets uint32
	MaxPerStageDescriptorSamplers uint32
	MaxPerStageDescriptorUniformBuffers uint32
	MaxPerStageDescriptorStorageBuffers uint32
	MaxPerStageDescriptorSampledImages uint32
	MaxPerStageDescriptorStorageImages uint32
	MaxPerStageDescriptorInputAttachments uint32
	MaxPerStageResources uint32

	MaxDescriptorSetSamplers uint32
	MaxDescriptorSetUniformBuffers uint32
	MaxDescriptorSetUniformBuffersDynamic uint32
	MaxDescriptorSetStorageBuffers uint32
	MaxDescriptorSetStorageBuffersDynamic uint32
	MaxDescriptorSetSampledImages uint32
	MaxDescriptorSetStorageImages uint32
	MaxDescriptorSetInputAttachments uint32

	MaxVertexInputAttributes uint32
	MaxVertexInputBindings uint32
	MaxVertexInputAttributeOffset uint32
	MaxVertexInputBindingStride uint32
	MaxVertexOutputComponents uint32

	MaxTessellationGenerationLevel uint32
	MaxTessellationPatchSize uint32
	MaxTessellationControlPerVertexInputComponents uint32
	MaxTessellationControlPerVertexOutputComponents uint32
	MaxTessellationControlPerPatchOutputComponents uint32
	MaxTessellationControlTotalOutputComponents uint32
	MaxTessellationEvaluationInputComponents uint32
	MaxTessellationEvaluationOutputComponents uint32

	MaxGeometryShaderInvocations uint32
	MaxGeometryInputComponents uint32
	MaxGeometryOutputComponents uint32
	MaxGeometryOutputVertices uint32
	MaxGeometryTotalOutputComponents uint32

	MaxFragmentInputComponents uint32
	MaxFragmentOutputAttachments uint32
	MaxFragmentDualSrcAttachments uint32
	MaxFragmentCombinedOutputResources uint32

	MaxComputeSharedMemorySize uint32
	MaxComputeWorkGroupCount [3]uint32
	MaxComputeWorkGroupInvocations uint32
	MaxComputeWorkGroupSize [3]uint32

	SubPixelPrecisionBits uint32
	SubTexelPrecisionBits uint32
	MipmapPrecisionBits uint32

	MaxDrawIndexedIndexValue uint32
	MaxDrawIndirectCount uint32

	MaxSamplerLodBias float32
	MaxSamplerAnisotropy float32

	MaxViewports uint32
	MaxViewportDimensions [2]uint32
	ViewportBoundsRange [2]float32
	ViewportSubPixelBits uint32

	MinMemoryMapAlignment uint
	MinTexelBufferOffsetAlignment uint64
	MinUniformBufferOffsetAlignment uint64
	MinStorageBufferOffsetAlignment uint64

	MinTexelOffset int32
	MaxTexelOffset uint32
	MinTexelGatherOffset int32
	MaxTexelGatherOffset uint32
	MinInterpolationOffset float32
	MaxInterpolationOffset float32
	SubPixelInterpolationOffsetBits uint32

	MaxFramebufferWidth uint32
	MaxFramebufferHeight uint32
	MaxFramebufferLayers uint32

	FramebufferColorSampleCounts SampleCount
	FramebufferDepthSampleCounts SampleCount
	FramebufferStencilSampleCounts SampleCount
	FramebufferNoAttachmentsSampleCounts SampleCount

	MaxColorAttachments uint32
	SampledImageColorSampleCounts SampleCount
	SampledImageIntegerSampleCounts SampleCount
	SampledImageDepthSampleCounts SampleCount
	SampledImageStencilSampleCounts SampleCount
	StorageImageSampleCounts SampleCount
	MaxSampleMaskWords uint32

	TimestampComputeAndGraphics bool
	TimestampPeriod float32

	MaxClipDistances uint32
	MaxCullDistances uint32
	MaxCombinedClipAndCullDistances uint32

	DiscreteQueuePriorities uint32

	PointSizeRange [2]float32
	LineWidthRange [2]float32
	PointSizeGranularity float32
	LineWidthGranularity float32

	StrictLines bool
	StandardSampleLocations bool

	OptimalBufferCopyOffsetAlignment uint64
	OptimalBufferCopyRowPitchAlignment uint64
	NonCoherentAtomSize uint64
}

func createPhysicalDeviceLimits(l *C.VkPhysicalDeviceLimits) *PhysicalDeviceLimits {
	return &PhysicalDeviceLimits{
		MaxImageDimension1D: uint32(l.maxImageDimension1D),
		MaxImageDimension2D: uint32(l.maxImageDimension2D),
		MaxImageDimension3D: uint32(l.maxImageDimension3D),
		MaxImageDimensionCube: uint32(l.maxImageDimensionCube),
		MaxImageArrayLayers: uint32(l.maxImageArrayLayers),
		MaxTexelBufferElements: uint32(l.maxTexelBufferElements),
		MaxUniformBufferRange: uint32(l.maxUniformBufferRange),
		MaxStorageBufferRange: uint32(l.maxStorageBufferRange),
		MaxPushConstantsSize: uint32(l.maxPushConstantsSize),
		MaxMemoryAllocationCount: uint32(l.maxMemoryAllocationCount),
		MaxSamplerAllocationCount: uint32(l.maxSamplerAllocationCount),
		BufferImageGranularity: uint64(l.bufferImageGranularity),
		SparseAddressSpaceSize: uint64(l.sparseAddressSpaceSize),
		MaxBoundDescriptorSets: uint32(l.maxBoundDescriptorSets),
		MaxPerStageDescriptorSamplers: uint32(l.maxPerStageDescriptorSamplers),
		MaxPerStageDescriptorUniformBuffers: uint32(l.maxPerStageDescriptorUniformBuffers),
		MaxPerStageDescriptorStorageBuffers: uint32(l.maxPerStageDescriptorStorageBuffers),
		MaxPerStageDescriptorSampledImages: uint32(l.maxPerStageDescriptorSampledImages),
		MaxPerStageDescriptorStorageImages: uint32(l.maxPerStageDescriptorStorageImages),
		MaxPerStageDescriptorInputAttachments: uint32(l.maxPerStageDescriptorInputAttachments),
		MaxPerStageResources: uint32(l.maxPerStageResources),
		MaxDescriptorSetSamplers: uint32(l.maxDescriptorSetSamplers),
		MaxDescriptorSetUniformBuffers: uint32(l.maxDescriptorSetUniformBuffers),
		MaxDescriptorSetUniformBuffersDynamic: uint32(l.maxDescriptorSetUniformBuffersDynamic),
		MaxDescriptorSetStorageBuffers: uint32(l.maxDescriptorSetStorageBuffers),
		MaxDescriptorSetStorageBuffersDynamic: uint32(l.maxDescriptorSetStorageBuffersDynamic),
		MaxDescriptorSetSampledImages: uint32(l.maxDescriptorSetSampledImages),
		MaxDescriptorSetStorageImages: uint32(l.maxDescriptorSetStorageImages),
		MaxDescriptorSetInputAttachments: uint32(l.maxDescriptorSetInputAttachments),
		MaxVertexInputAttributes: uint32(l.maxVertexInputAttributes),
		MaxVertexInputBindings: uint32(l.maxVertexInputBindings),
		MaxVertexInputAttributeOffset: uint32(l.maxVertexInputAttributeOffset),
		MaxVertexInputBindingStride: uint32(l.maxVertexInputBindingStride),
		MaxVertexOutputComponents: uint32(l.maxVertexOutputComponents),
		MaxTessellationGenerationLevel: uint32(l.maxTessellationGenerationLevel),
		MaxTessellationPatchSize: uint32(l.maxTessellationPatchSize),
		MaxTessellationControlPerVertexInputComponents: uint32(l.maxTessellationControlPerVertexInputComponents),
		MaxTessellationControlPerVertexOutputComponents: uint32(l.maxTessellationControlPerVertexOutputComponents),
		MaxTessellationControlPerPatchOutputComponents: uint32(l.maxTessellationControlPerPatchOutputComponents),
		MaxTessellationControlTotalOutputComponents: uint32(l.maxTessellationControlTotalOutputComponents),
		MaxTessellationEvaluationInputComponents: uint32(l.maxTessellationEvaluationInputComponents),
		MaxTessellationEvaluationOutputComponents: uint32(l.maxTessellationEvaluationOutputComponents),
		MaxGeometryShaderInvocations: uint32(l.maxGeometryShaderInvocations),
		MaxGeometryInputComponents: uint32(l.maxGeometryInputComponents),
		MaxGeometryOutputComponents: uint32(l.maxGeometryOutputComponents),
		MaxGeometryOutputVertices: uint32(l.maxGeometryOutputVertices),
		MaxGeometryTotalOutputComponents: uint32(l.maxGeometryTotalOutputComponents),
		MaxFragmentInputComponents: uint32(l.maxFragmentInputComponents),
		MaxFragmentOutputAttachments: uint32(l.maxFragmentOutputAttachments),
		MaxFragmentDualSrcAttachments: uint32(l.maxFragmentDualSrcAttachments),
		MaxFragmentCombinedOutputResources: uint32(l.maxFragmentCombinedOutputResources),
		MaxComputeSharedMemorySize: uint32(l.maxComputeSharedMemorySize),
		MaxComputeWorkGroupInvocations: uint32(l.maxComputeWorkGroupInvocations),
		SubPixelPrecisionBits: uint32(l.subPixelPrecisionBits),
		SubTexelPrecisionBits: uint32(l.subTexelPrecisionBits),
		MipmapPrecisionBits: uint32(l.mipmapPrecisionBits),
		MaxDrawIndexedIndexValue: uint32(l.maxDrawIndexedIndexValue),
		MaxDrawIndirectCount: uint32(l.maxDrawIndirectCount),
		MaxSamplerLodBias: float32(l.maxSamplerLodBias),
		MaxSamplerAnisotropy: float32(l.maxSamplerAnisotropy),
		MaxViewports: uint32(l.maxViewports),
		ViewportSubPixelBits: uint32(l.viewportSubPixelBits),
		MinMemoryMapAlignment: uint(l.minMemoryMapAlignment),
		MinTexelBufferOffsetAlignment: uint64(l.minTexelBufferOffsetAlignment),
		MinUniformBufferOffsetAlignment: uint64(l.minUniformBufferOffsetAlignment),
		MinStorageBufferOffsetAlignment: uint64(l.minStorageBufferOffsetAlignment),
		MinTexelOffset: int32(l.minTexelOffset),
		MaxTexelOffset: uint32(l.maxTexelOffset),
		MinTexelGatherOffset: int32(l.minTexelGatherOffset),
		MaxTexelGatherOffset: uint32(l.maxTexelGatherOffset),
		MinInterpolationOffset: float32(l.minInterpolationOffset),
		MaxInterpolationOffset: float32(l.maxInterpolationOffset),
		SubPixelInterpolationOffsetBits: uint32(l.subPixelInterpolationOffsetBits),
		MaxFramebufferWidth: uint32(l.maxFramebufferWidth),
		MaxFramebufferHeight: uint32(l.maxFramebufferHeight),
		MaxFramebufferLayers: uint32(l.maxFramebufferLayers),
		FramebufferColorSampleCounts: SampleCount(l.framebufferColorSampleCounts),
		FramebufferDepthSampleCounts: SampleCount(l.framebufferDepthSampleCounts),
		FramebufferStencilSampleCounts: SampleCount(l.framebufferStencilSampleCounts),
		FramebufferNoAttachmentsSampleCounts: SampleCount(l.framebufferNoAttachmentsSampleCounts),
		MaxColorAttachments: uint32(l.maxColorAttachments),
		SampledImageColorSampleCounts: SampleCount(l.sampledImageColorSampleCounts),
		SampledImageIntegerSampleCounts: SampleCount(l.sampledImageIntegerSampleCounts),
		SampledImageDepthSampleCounts: SampleCount(l.sampledImageDepthSampleCounts),
		SampledImageStencilSampleCounts: SampleCount(l.sampledImageStencilSampleCounts),
		StorageImageSampleCounts: SampleCount(l.storageImageSampleCounts),
		MaxSampleMaskWords: uint32(l.maxSampleMaskWords),
		TimestampComputeAndGraphics: l.timestampComputeAndGraphics != C.VK_FALSE,
		TimestampPeriod: float32(l.timestampPeriod),
		MaxClipDistances: uint32(l.maxClipDistances),
		MaxCullDistances: uint32(l.maxCullDistances),
		MaxCombinedClipAndCullDistances: uint32(l.maxCombinedClipAndCullDistances),
		DiscreteQueuePriorities: uint32(l.discreteQueuePriorities),
		PointSizeGranularity: float32(l.pointSizeGranularity),
		LineWidthGranularity: float32(l.lineWidthGranularity),
		StrictLines: l.strictLines != C.VK_FALSE,
		StandardSampleLocations: l.standardSampleLocations != C.VK_FALSE,
		OptimalBufferCopyOffsetAlignment: uint64(l.optimalBufferCopyOffsetAlignment),
		OptimalBufferCopyRowPitchAlignment: uint64(l.optimalBufferCopyRowPitchAlignment),
		NonCoherentAtomSize: uint64(l.nonCoherentAtomSize),
		MaxComputeWorkGroupCount: [3]uint32 {
			uint32(l.maxComputeWorkGroupCount[0]),
			uint32(l.maxComputeWorkGroupCount[1]),
			uint32(l.maxComputeWorkGroupCount[2]),
		},
		MaxComputeWorkGroupSize: [3]uint32 {
			uint32(l.maxComputeWorkGroupSize[0]),
			uint32(l.maxComputeWorkGroupSize[1]),
			uint32(l.maxComputeWorkGroupSize[2]),
		},
		MaxViewportDimensions: [2]uint32 {
			uint32(l.maxViewportDimensions[0]),
			uint32(l.maxViewportDimensions[1]),
		},
		ViewportBoundsRange: [2]float32 {
			float32(l.viewportBoundsRange[0]),
			float32(l.viewportBoundsRange[1]),
		},
		PointSizeRange: [2]float32 {
			float32(l.pointSizeRange[0]),
			float32(l.pointSizeRange[1]),
		},
		LineWidthRange: [2]float32 {
			float32(l.lineWidthRange[0]),
			float32(l.lineWidthRange[1]),
		},
	}
}

