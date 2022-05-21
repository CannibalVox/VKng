package khr_portability_subset

/*
#define VK_ENABLE_BETA_EXTENSIONS 1
#include <stdlib.h>
#include "../vulkan/vulkan.h"
#include "../vulkan/vulkan_beta.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type PhysicalDevicePortabilitySubsetFeaturesOutData struct {
	ConstantAlphaColorBlendFactors          bool
	Events                                  bool
	ImageViewFormatReinterpretation         bool
	ImageViewFormatSwizzle                  bool
	ImageView2DOn3DImage                    bool
	MultisampleArrayImage                   bool
	MutableComparisonSamplers               bool
	PointPolygons                           bool
	SamplerMipLodBias                       bool
	SeparateStencilMaskRef                  bool
	ShaderSamplerRateInterpolationFunctions bool
	TessellationIsolines                    bool
	TessellationPointMode                   bool
	TriangleFans                            bool
	VertexAttributeAccessBeyondStride       bool

	common.HaveNext
}

func (o *PhysicalDevicePortabilitySubsetFeaturesOutData) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(C.sizeof_struct_VkPhysicalDevicePortabilitySubsetFeaturesKHR)
	}

	outData := (*C.VkPhysicalDevicePortabilitySubsetFeaturesKHR)(preallocatedPointer)
	outData.sType = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR
	outData.pNext = next

	return preallocatedPointer, nil
}

func (o *PhysicalDevicePortabilitySubsetFeaturesOutData) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	outData := (*C.VkPhysicalDevicePortabilitySubsetFeaturesKHR)(cDataPointer)
	o.ConstantAlphaColorBlendFactors = outData.constantAlphaColorBlendFactors != C.VkBool32(0)
	o.Events = outData.events != C.VkBool32(0)
	o.ImageViewFormatReinterpretation = outData.imageViewFormatReinterpretation != C.VkBool32(0)
	o.ImageViewFormatSwizzle = outData.imageViewFormatSwizzle != C.VkBool32(0)
	o.ImageView2DOn3DImage = outData.imageView2DOn3DImage != C.VkBool32(0)
	o.MultisampleArrayImage = outData.multisampleArrayImage != C.VkBool32(0)
	o.MutableComparisonSamplers = outData.mutableComparisonSamplers != C.VkBool32(0)
	o.PointPolygons = outData.pointPolygons != C.VkBool32(0)
	o.SamplerMipLodBias = outData.samplerMipLodBias != C.VkBool32(0)
	o.SeparateStencilMaskRef = outData.separateStencilMaskRef != C.VkBool32(0)
	o.ShaderSamplerRateInterpolationFunctions = outData.shaderSampleRateInterpolationFunctions != C.VkBool32(0)
	o.TessellationIsolines = outData.tessellationIsolines != C.VkBool32(0)
	o.TessellationPointMode = outData.tessellationPointMode != C.VkBool32(0)
	o.TriangleFans = outData.triangleFans != C.VkBool32(0)
	o.VertexAttributeAccessBeyondStride = outData.vertexAttributeAccessBeyondStride != C.VkBool32(0)

	return outData.pNext, nil
}
