package khr_sampler_ycbcr_conversion

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/core1_0"
)

type ChromaLocation int32

var chromaLocationMapping = make(map[ChromaLocation]string)

func (e ChromaLocation) Register(str string) {
	chromaLocationMapping[e] = str
}

func (e ChromaLocation) String() string {
	return chromaLocationMapping[e]
}

////

type SamplerYcbcrModelConversion int32

var samplerModelConversionMapping = make(map[SamplerYcbcrModelConversion]string)

func (e SamplerYcbcrModelConversion) Register(str string) {
	samplerModelConversionMapping[e] = str
}

func (e SamplerYcbcrModelConversion) String() string {
	return samplerModelConversionMapping[e]
}

////

type SamplerYcbcrRange int32

var samplerRangeMapping = make(map[SamplerYcbcrRange]string)

func (e SamplerYcbcrRange) Register(str string) {
	samplerRangeMapping[e] = str
}

func (e SamplerYcbcrRange) String() string {
	return samplerRangeMapping[e]
}

////

const (
	ExtensionName string = C.VK_KHR_SAMPLER_YCBCR_CONVERSION_EXTENSION_NAME

	ChromaLocationCositedEven ChromaLocation = C.VK_CHROMA_LOCATION_COSITED_EVEN_KHR
	ChromaLocationMidpoint    ChromaLocation = C.VK_CHROMA_LOCATION_MIDPOINT_KHR

	ObjectTypeSamplerYcbcrConversion core1_0.ObjectType = C.VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR

	FormatB10X6G10X6R10X6G10X6HorizontalChromaComponentPacked     core1_0.Format = C.VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16_KHR
	FormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked     core1_0.Format = C.VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR
	FormatB16G16R16G16HorizontalChroma                            core1_0.Format = C.VK_FORMAT_B16G16R16G16_422_UNORM_KHR
	FormatB8G8R8G8HorizontalChroma                                core1_0.Format = C.VK_FORMAT_B8G8R8G8_422_UNORM_KHR
	FormatG10X6B10X6G10X6R10X6HorizontalChromaComponentPacked     core1_0.Format = C.VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16_KHR
	FormatG10X6_B10X6R10X6_2PlaneDualChromaComponentPacked        core1_0.Format = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16_KHR
	FormatG10X6_B10X6R10X6_2PlaneHorizontalChromaComponentPacked  core1_0.Format = C.VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16_KHR
	FormatG10X6_B10X6_R10X6_3PlaneDualChromaComponentPacked       core1_0.Format = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16_KHR
	FormatG10X6_B10X6_R10X6_3PlaneHorizontalChromaComponentPacked core1_0.Format = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16_KHR
	FormatG10X6_B10X6_R10X6_3PlaneNoChromaComponentPacked         core1_0.Format = C.VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16_KHR
	FormatG12X4B12X4G12X4R12X4_HorizontalChromaComponentPacked    core1_0.Format = C.VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16_KHR
	FormatG12X4_B12X4R12X4_2PlaneDualChromaComponentPacked        core1_0.Format = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16_KHR
	FormatG12X4_B12X4R12X4_2PlaneHorizontalChromaComponentPacked  core1_0.Format = C.VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16_KHR
	FormatG12X4_B12X4_R12X4_3PlaneDualChromaComponentPacked       core1_0.Format = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16_KHR
	FormatG12X4_B12X4_R12X4_3PlaneHorizontalChromaComponentPacked core1_0.Format = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16_KHR
	FormatG12X4_B12X4_R12X4_3PlaneNoChromaComponentPacked         core1_0.Format = C.VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16_KHR
	FormatG16B16G16R16_HorizontalChroma                           core1_0.Format = C.VK_FORMAT_G16B16G16R16_422_UNORM_KHR
	FormatG16_B16R16_2PlaneDualChroma                             core1_0.Format = C.VK_FORMAT_G16_B16R16_2PLANE_420_UNORM_KHR
	FormatG16_B16R16_2PlaneHorizontalChroma                       core1_0.Format = C.VK_FORMAT_G16_B16R16_2PLANE_422_UNORM_KHR
	FormatG16_B16_R16_3PlaneDualChroma                            core1_0.Format = C.VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM_KHR
	FormatG16_B16_R16_3PlaneHorizontalChroma                      core1_0.Format = C.VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM_KHR
	FormatG16_B16_R16_3PlaneNoChroma                              core1_0.Format = C.VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM_KHR
	FormatG8B8G8R8_HorizontalChroma                               core1_0.Format = C.VK_FORMAT_G8B8G8R8_422_UNORM_KHR
	FormatG8_B8R8_2PlaneDualChroma                                core1_0.Format = C.VK_FORMAT_G8_B8R8_2PLANE_420_UNORM_KHR
	FormatG8_B8R8_2PlaneHorizontalChroma                          core1_0.Format = C.VK_FORMAT_G8_B8R8_2PLANE_422_UNORM_KHR
	FormatG8_B8_R8_3PlaneDualChroma                               core1_0.Format = C.VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM_KHR
	FormatG8_B8_R8_3PlaneHorizontalChroma                         core1_0.Format = C.VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM_KHR
	FormatG8_B8_R8_3PlaneNoChroma                                 core1_0.Format = C.VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM_KHR
	FormatR10X6G10X6B10X6A10X6UnsignedNormalizedComponentPacked   core1_0.Format = C.VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16_KHR
	FormatR10X6G10X6UnsignedNormalizedComponentPacked             core1_0.Format = C.VK_FORMAT_R10X6G10X6_UNORM_2PACK16_KHR
	FormatR10X6UnsignedNormalizedComponentPacked                  core1_0.Format = C.VK_FORMAT_R10X6_UNORM_PACK16_KHR
	FormatR12X4G12X4B12X4A12X4UnsignedNormalizedComponentPacked   core1_0.Format = C.VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16_KHR
	FormatR12X4G12X4UnsignedNormalizedComponentPacked             core1_0.Format = C.VK_FORMAT_R12X4G12X4_UNORM_2PACK16_KHR
	FormatR12X4UnsignedNormalizedComponentPacked                  core1_0.Format = C.VK_FORMAT_R12X4_UNORM_PACK16_KHR

	FormatFeatureCositedChromaSamples                                             core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT_KHR
	FormatFeatureDisjoint                                                         core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_DISJOINT_BIT_KHR
	FormatFeatureMidpointChromaSamples                                            core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT_KHR
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicit          core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT_KHR
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceable core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT_KHR
	FormatFeatureSampledImageYcbcrConversionLinearFilter                          core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT_KHR
	FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilter          core1_0.FormatFeatureFlags = C.VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT_KHR

	ImageAspectPlane0 core1_0.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_0_BIT_KHR
	ImageAspectPlane1 core1_0.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_1_BIT_KHR
	ImageAspectPlane2 core1_0.ImageAspectFlags = C.VK_IMAGE_ASPECT_PLANE_2_BIT_KHR

	ImageCreateDisjoint core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_DISJOINT_BIT_KHR

	SamplerYcbcrModelConversionRGBIdentity   SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY_KHR
	SamplerYcbcrModelConversionYcbcr2020     SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020_KHR
	SamplerYcbcrModelConversionYcbcr601      SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601_KHR
	SamplerYcbcrModelConversionYcbcr709      SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR
	SamplerYcbcrModelConversionYcbcrIdentity SamplerYcbcrModelConversion = C.VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY_KHR

	SamplerYcbcrRangeITUFull   SamplerYcbcrRange = C.VK_SAMPLER_YCBCR_RANGE_ITU_FULL_KHR
	SamplerYcbcrRangeITUNarrow SamplerYcbcrRange = C.VK_SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR
)

func init() {
	ChromaLocationCositedEven.Register("Cosited Even")
	ChromaLocationMidpoint.Register("Midpoint")

	ObjectTypeSamplerYcbcrConversion.Register("Sampler Ycbcr Conversion")

	FormatB10X6G10X6R10X6G10X6HorizontalChromaComponentPacked.Register("B10(X6)G10(X6)R10(X6)G10(X6) Horizontal Chroma (Component-Packed)")
	FormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked.Register("B12(X4)G12(X4)R12(X4)G12(X4) Horizontal Chroma (Component-Packed)")
	FormatB16G16R16G16HorizontalChroma.Register("B16G16R16G16 Horizontal Chroma")
	FormatB8G8R8G8HorizontalChroma.Register("B8G8R8G8 Horizontal Chroma")
	FormatG10X6B10X6G10X6R10X6HorizontalChromaComponentPacked.Register("G10(X6)B10(X6)G10(X6)R10(X6) Horizontal Chroma (Component-Packed)")
	FormatG10X6_B10X6R10X6_2PlaneDualChromaComponentPacked.Register("2-Plane G10(X6) B10(X6)R10(X6) Dual Chroma (Component-Packed)")
	FormatG10X6_B10X6R10X6_2PlaneHorizontalChromaComponentPacked.Register("2-Plane G10(X6) B10(X6)R10(X6) Horizontal Chroma (Component-Packed)")
	FormatG10X6_B10X6_R10X6_3PlaneDualChromaComponentPacked.Register("3-Plane G10(X6) B10(X6) R10(X6) Dual Chroma (Component-Packed)")
	FormatG10X6_B10X6_R10X6_3PlaneHorizontalChromaComponentPacked.Register("3-Plane G10(X6) B10(X6) R10(X6) Horizontal Chroma (Component-Packed)")
	FormatG10X6_B10X6_R10X6_3PlaneNoChromaComponentPacked.Register("3-Plane G10(X6) B10(X6) R10(X6) No Chroma (Component-Packed)")
	FormatG12X4B12X4G12X4R12X4_HorizontalChromaComponentPacked.Register("G12(X4)B12(X4)G12(X4)R12(X4) Horizontal Chroma (Component-Packed)")
	FormatG12X4_B12X4R12X4_2PlaneDualChromaComponentPacked.Register("2-Plane G12(X4) B12(X4)R12(X4) Dual Chroma (Component-Packed)")
	FormatG12X4_B12X4R12X4_2PlaneHorizontalChromaComponentPacked.Register("2-Plane G12(X4) B12(X4)R12(X4) Horizontal Chroma (Component-Packed)")
	FormatG12X4_B12X4_R12X4_3PlaneDualChromaComponentPacked.Register("3-Plane G12(X4) B12(X4) R12(X4) Dual Chroma (Component-Packed)")
	FormatG12X4_B12X4_R12X4_3PlaneHorizontalChromaComponentPacked.Register("3-Plane G12(X4) B12(X4) R12(X4) Horizontal Chroma (Component-Packed)")
	FormatG12X4_B12X4_R12X4_3PlaneNoChromaComponentPacked.Register("3-Plane G12(X4) B12(X4) R12(X4) No Chroma (Component-Packed)")
	FormatG16B16G16R16_HorizontalChroma.Register("G16B16G16R16 Horizontal Chroma")
	FormatG16_B16R16_2PlaneDualChroma.Register("2-Plane G16 B16R16 Dual Chroma")
	FormatG16_B16R16_2PlaneHorizontalChroma.Register("2-Plane G16 B16R16 Horizontal Chroma")
	FormatG16_B16_R16_3PlaneDualChroma.Register("3-Plane G16 B16 R16 Dual Chroma")
	FormatG16_B16_R16_3PlaneHorizontalChroma.Register("3-Plane G16 B16 R16 Horizontal Chroma")
	FormatG16_B16_R16_3PlaneNoChroma.Register("3-Plane G16 B16 R16 No Chroma")
	FormatG8B8G8R8_HorizontalChroma.Register("G8B8G8R8 Horizontal Chroma")
	FormatG8_B8R8_2PlaneDualChroma.Register("2-Plane G8 B8R8 Dual Chroma")
	FormatG8_B8R8_2PlaneHorizontalChroma.Register("2-Plane G8 B8R8 Horizontal Chroma")
	FormatG8_B8_R8_3PlaneDualChroma.Register("3-Plane G8 B8 R8 Dual Chroma")
	FormatG8_B8_R8_3PlaneHorizontalChroma.Register("3-Plane G8 B8 R8 Horizontal Chroma")
	FormatG8_B8_R8_3PlaneNoChroma.Register("3-Plane G8 B8 R8 No Chroma")
	FormatR10X6G10X6B10X6A10X6UnsignedNormalizedComponentPacked.Register("R10(X6)G10(X6)B10(X6)A10(X6) Unsigned Normalized (Component-Packed)")
	FormatR10X6G10X6UnsignedNormalizedComponentPacked.Register("R10(X6)G10(X6) Unsigned Normalized (Component-Packed)")
	FormatR10X6UnsignedNormalizedComponentPacked.Register("R10(X6) Unsigned Normalized (Component-Packed)")
	FormatR12X4G12X4B12X4A12X4UnsignedNormalizedComponentPacked.Register("R12(X4)G12(X4)B12(X4)A12(X4) Unsigned Normalized (Component-Packed)")
	FormatR12X4G12X4UnsignedNormalizedComponentPacked.Register("R12(X4)G12(X4) Unsigned Normalized (Component-Packed)")
	FormatR12X4UnsignedNormalizedComponentPacked.Register("R12(X4) Unsigned Normalized (Component-Packed)")

	FormatFeatureCositedChromaSamples.Register("Cosited Chroma Samples")
	FormatFeatureDisjoint.Register("Disjoint")
	FormatFeatureMidpointChromaSamples.Register("Midpoint Chroma Samples")
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicit.Register("Sampled Image Ycbcr Conversion - Chroma Reconstruction (Explicit)")
	FormatFeatureSampledImageYcbcrConversionChromaReconstructionExplicitForceable.Register("Sampled Image Ycbcr Conversion - Chroma Reconstruction (Explicit, Forceable)")
	FormatFeatureSampledImageYcbcrConversionLinearFilter.Register("Sampled Image Ycbcr Conversion - Linear Filter")
	FormatFeatureSampledImageYcbcrConversionSeparateReconstructionFilter.Register("Sampled Image Ycbcr Conversion - Separate Reconstruction Filter")

	ImageAspectPlane0.Register("Plane 0")
	ImageAspectPlane1.Register("Plane 1")
	ImageAspectPlane2.Register("Plane 2")

	ImageCreateDisjoint.Register("Disjoint")

	SamplerYcbcrModelConversionRGBIdentity.Register("RGB Identity")
	SamplerYcbcrModelConversionYcbcr2020.Register("Ycbcr 2020")
	SamplerYcbcrModelConversionYcbcr601.Register("Ycbcr 601")
	SamplerYcbcrModelConversionYcbcr709.Register("Ycbcr 709")
	SamplerYcbcrModelConversionYcbcrIdentity.Register("Ycbcr Identity")

	SamplerYcbcrRangeITUFull.Register("ITU Full")
	SamplerYcbcrRangeITUNarrow.Register("ITU Narrow")
}
