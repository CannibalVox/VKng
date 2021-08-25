package objects

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng"

func createPhysicalDeviceFeatures(f *C.VkPhysicalDeviceFeatures) *VKng.PhysicalDeviceFeatures {
	return &VKng.PhysicalDeviceFeatures{
		RobustBufferAccess: f.robustBufferAccess != C.VK_FALSE,
		FullDrawIndexUint32: f.fullDrawIndexUint32 != C.VK_FALSE,
		ImageCubeArray: f.imageCubeArray != C.VK_FALSE,
		IndependentBlend: f.independentBlend != C.VK_FALSE,
		GeometryShader: f.geometryShader != C.VK_FALSE,
		TessellationShader: f.tessellationShader != C.VK_FALSE,
		SampleRateShading: f.sampleRateShading != C.VK_FALSE,
		DualSrcBlend: f.dualSrcBlend != C.VK_FALSE,
		LogicOp: f.logicOp != C.VK_FALSE,
		MultiDrawIndirect: f.multiDrawIndirect != C.VK_FALSE,
		DrawIndirectFirstInstance:              f.drawIndirectFirstInstance != C.VK_FALSE,
		DepthClamp:                             f.depthClamp != C.VK_FALSE,
		DepthBiasClamp:                         f.depthBiasClamp != C.VK_FALSE,
		FillModeNonSolid:                       f.fillModeNonSolid != C.VK_FALSE,
		DepthBounds:                            f.depthBounds != C.VK_FALSE,
		WideLines:                              f.wideLines != C.VK_FALSE,
		LargePoints:                            f.largePoints != C.VK_FALSE,
		AlphaToOne:                             f.alphaToOne != C.VK_FALSE,
		MultiViewport:                          f.multiViewport != C.VK_FALSE,
		SamplerAnisotropy:                      f.samplerAnisotropy != C.VK_FALSE,
		TextureCompressionEtc2:                 f.textureCompressionETC2 != C.VK_FALSE,
		TextureCompressionAstcLdc:              f.textureCompressionASTC_LDR != C.VK_FALSE,
		TextureCompressionBc:                   f.textureCompressionBC != C.VK_FALSE,
		OcclusionQueryPrecise:                  f.occlusionQueryPrecise != C.VK_FALSE,
		PipelineStatisticsQuery:                f.pipelineStatisticsQuery != C.VK_FALSE,
		VertexPipelineStoresAndAtomics:         f.vertexPipelineStoresAndAtomics != C.VK_FALSE,
		FragmentStoresAndAtomics:               f.fragmentStoresAndAtomics != C.VK_FALSE,
		ShaderTessellationAndGeometryPointSize: f.shaderTessellationAndGeometryPointSize != C.VK_FALSE,
		ShaderImageGatherExtended:              f.shaderImageGatherExtended != C.VK_FALSE,
		ShaderStorageImageExtendedFormats:      f.shaderStorageImageExtendedFormats != C.VK_FALSE,
		ShaderStorageImageMultisample:          f.shaderStorageImageMultisample != C.VK_FALSE,
		ShaderStorageImageReadWithoutFormat:    f.shaderStorageImageReadWithoutFormat != C.VK_FALSE,
		ShaderStorageImageWriteWithoutFormat:   f.shaderStorageImageWriteWithoutFormat != C.VK_FALSE,
		ShaderUniformBufferArrayDynamicIndexing: f.shaderUniformBufferArrayDynamicIndexing != C.VK_FALSE,
		ShaderSampledImageArrayDynamicIndexing: f.shaderSampledImageArrayDynamicIndexing != C.VK_FALSE,
		ShaderStorageBufferArrayDynamicIndexing: f.shaderStorageBufferArrayDynamicIndexing != C.VK_FALSE,
		ShaderStorageImageArrayDynamicIndexing: f.shaderStorageImageArrayDynamicIndexing != C.VK_FALSE,
		ShaderClipDistance: f.shaderClipDistance != C.VK_FALSE,
		ShaderCullDistance: f.shaderCullDistance != C.VK_FALSE,
		ShaderFloat64: f.shaderFloat64 != C.VK_FALSE,
		ShaderInt64: f.shaderInt64 != C.VK_FALSE,
		ShaderInt16: f.shaderInt16 != C.VK_FALSE,
		ShaderResourceResidency: f.shaderResourceResidency != C.VK_FALSE,
		ShaderResourceMinLod: f.shaderResourceMinLod != C.VK_FALSE,
		SparseBinding: f.sparseBinding != C.VK_FALSE,
		SparseResidencyBuffer: f.sparseResidencyBuffer != C.VK_FALSE,
		SparseResidencyImage2D: f.sparseResidencyImage2D != C.VK_FALSE,
		SparseResidencyImage3D: f.sparseResidencyImage3D != C.VK_FALSE,
		SparseResidency2Samples: f.sparseResidency2Samples != C.VK_FALSE,
		SparseResidency4Samples: f.sparseResidency4Samples != C.VK_FALSE,
		SparseResidency8Samples: f.sparseResidency8Samples != C.VK_FALSE,
		SparseResidency16Samples: f.sparseResidency16Samples != C.VK_FALSE,
		SparseResidencyAliased: f.sparseResidencyAliased != C.VK_FALSE,
		VariableMultisampleRate: f.variableMultisampleRate != C.VK_FALSE,
		InheritedQueries: f.inheritedQueries != C.VK_FALSE,
	}
}

