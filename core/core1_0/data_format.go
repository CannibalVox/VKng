package core1_0

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

const (
	FormatUndefined                    Format = C.VK_FORMAT_UNDEFINED
	FormatR4G4UnsignedNormalizedPacked Format = C.VK_FORMAT_R4G4_UNORM_PACK8

	FormatR4G4B4A4UnsignedNormalizedPacked Format = C.VK_FORMAT_R4G4B4A4_UNORM_PACK16
	FormatB4G4R4A4UnsignedNormalizedPacked Format = C.VK_FORMAT_B4G4R4A4_UNORM_PACK16
	FormatR5G6B5UnsignedNormalizedPacked   Format = C.VK_FORMAT_R5G6B5_UNORM_PACK16
	FormatB5G6R5UnsignedNormalizedPacked   Format = C.VK_FORMAT_B5G6R5_UNORM_PACK16
	FormatR5G5B5A1UnsignedNormalizedPacked Format = C.VK_FORMAT_R5G5B5A1_UNORM_PACK16
	FormatB5G5R5A1UnsignedNormalizedPacked Format = C.VK_FORMAT_B5G5R5A1_UNORM_PACK16
	FormatA1R5G5B5UnsignedNormalizedPacked Format = C.VK_FORMAT_A1R5G5B5_UNORM_PACK16

	FormatR8UnsignedNormalized Format = C.VK_FORMAT_R8_UNORM
	FormatR8SignedNormalized   Format = C.VK_FORMAT_R8_SNORM
	FormatR8UnsignedScaled     Format = C.VK_FORMAT_R8_USCALED
	FormatR8SignedScaled       Format = C.VK_FORMAT_R8_SSCALED
	FormatR8UnsignedInt        Format = C.VK_FORMAT_R8_UINT
	FormatR8SignedInt          Format = C.VK_FORMAT_R8_SINT
	FormatR8SRGB               Format = C.VK_FORMAT_R8_SRGB

	FormatR8G8UnsignedNormalized Format = C.VK_FORMAT_R8G8_UNORM
	FormatR8G8SignedNormalized   Format = C.VK_FORMAT_R8G8_SNORM
	FormatR8G8UnsignedScaled     Format = C.VK_FORMAT_R8G8_USCALED
	FormatR8G8SignedScaled       Format = C.VK_FORMAT_R8G8_SSCALED
	FormatR8G8UnsignedInt        Format = C.VK_FORMAT_R8G8_UINT
	FormatR8G8SignedInt          Format = C.VK_FORMAT_R8G8_SINT
	FormatR8G8SRGB               Format = C.VK_FORMAT_R8G8_SRGB

	FormatR8G8B8UnsignedNormalized Format = C.VK_FORMAT_R8G8B8_UNORM
	FormatR8G8B8SignedNormalized   Format = C.VK_FORMAT_R8G8B8_SNORM
	FormatR8G8B8UnsignedScaled     Format = C.VK_FORMAT_R8G8B8_USCALED
	FormatR8G8B8SignedScaled       Format = C.VK_FORMAT_R8G8B8_SSCALED
	FormatR8G8B8UnsignedInt        Format = C.VK_FORMAT_R8G8B8_UINT
	FormatR8G8B8SignedInt          Format = C.VK_FORMAT_R8G8B8_SINT
	FormatR8G8B8SRGB               Format = C.VK_FORMAT_R8G8B8_SRGB

	FormatB8G8R8UnsignedNormalized Format = C.VK_FORMAT_B8G8R8_UNORM
	FormatB8G8R8SignedNormalized   Format = C.VK_FORMAT_B8G8R8_SNORM
	FormatB8G8R8UnsignedScaled     Format = C.VK_FORMAT_B8G8R8_USCALED
	FormatB8G8R8SignedScaled       Format = C.VK_FORMAT_B8G8R8_SSCALED
	FormatB8G8R8UnsignedInt        Format = C.VK_FORMAT_B8G8R8_UINT
	FormatB8G8R8SignedInt          Format = C.VK_FORMAT_B8G8R8_SINT
	FormatB8G8R8SRGB               Format = C.VK_FORMAT_B8G8R8_SRGB

	FormatR8G8B8A8UnsignedNormalized Format = C.VK_FORMAT_R8G8B8A8_UNORM
	FormatR8G8B8A8SignedNormalized   Format = C.VK_FORMAT_R8G8B8A8_SNORM
	FormatR8G8B8A8UnsignedScaled     Format = C.VK_FORMAT_R8G8B8A8_USCALED
	FormatR8G8B8A8SignedScaled       Format = C.VK_FORMAT_R8G8B8A8_SSCALED
	FormatR8G8B8A8UnsignedInt        Format = C.VK_FORMAT_R8G8B8A8_UINT
	FormatR8G8B8A8SignedInt          Format = C.VK_FORMAT_R8G8B8A8_SINT
	FormatR8G8B8A8SRGB               Format = C.VK_FORMAT_R8G8B8A8_SRGB

	FormatB8G8R8A8UnsignedNormalized Format = C.VK_FORMAT_B8G8R8A8_UNORM
	FormatB8G8R8A8SignedNormalized   Format = C.VK_FORMAT_B8G8R8A8_SNORM
	FormatB8G8R8A8UnsignedScaled     Format = C.VK_FORMAT_B8G8R8A8_USCALED
	FormatB8G8R8A8SignedScaled       Format = C.VK_FORMAT_B8G8R8A8_SSCALED
	FormatB8G8R8A8UnsignedInt        Format = C.VK_FORMAT_B8G8R8A8_UINT
	FormatB8G8R8A8SignedInt          Format = C.VK_FORMAT_B8G8R8A8_SINT
	FormatB8G8R8A8SRGB               Format = C.VK_FORMAT_B8G8R8A8_SRGB

	FormatA8B8G8R8UnsignedNormalizedPacked Format = C.VK_FORMAT_A8B8G8R8_UNORM_PACK32
	FormatA8B8G8R8SignedNormalizedPacked   Format = C.VK_FORMAT_A8B8G8R8_SNORM_PACK32
	FormatA8B8G8R8UnsignedScaledPacked     Format = C.VK_FORMAT_A8B8G8R8_USCALED_PACK32
	FormatA8B8G8R8SignedScaledPacked       Format = C.VK_FORMAT_A8B8G8R8_SSCALED_PACK32
	FormatA8B8G8R8UnsignedIntPacked        Format = C.VK_FORMAT_A8B8G8R8_UINT_PACK32
	FormatA8B8G8R8SignedIntPacked          Format = C.VK_FORMAT_A8B8G8R8_SINT_PACK32
	FormatA8B8G8R8SRGBPacked               Format = C.VK_FORMAT_A8B8G8R8_SRGB_PACK32

	FormatA2R10G10B10UnsignedNormalizedPacked Format = C.VK_FORMAT_A2R10G10B10_UNORM_PACK32
	FormatA2R10G10B10SignedNormalizedPacked   Format = C.VK_FORMAT_A2R10G10B10_SNORM_PACK32
	FormatA2R10G10B10UnsignedScaledPacked     Format = C.VK_FORMAT_A2R10G10B10_USCALED_PACK32
	FormatA2R10G10B10SignedScaledPacked       Format = C.VK_FORMAT_A2R10G10B10_SSCALED_PACK32
	FormatA2R10G10B10UnsignedIntPacked        Format = C.VK_FORMAT_A2R10G10B10_UINT_PACK32
	FormatA2R10G10B10SignedIntPacked          Format = C.VK_FORMAT_A2R10G10B10_SINT_PACK32

	FormatA2B10G10R10UnsignedNormalizedPacked Format = C.VK_FORMAT_A2B10G10R10_UNORM_PACK32
	FormatA2B10G10R10SignedNormalizedPacked   Format = C.VK_FORMAT_A2B10G10R10_SNORM_PACK32
	FormatA2B10G10R10UnsignedScaledPacked     Format = C.VK_FORMAT_A2B10G10R10_USCALED_PACK32
	FormatA2B10G10R10SignedScaledPacked       Format = C.VK_FORMAT_A2B10G10R10_SSCALED_PACK32
	FormatA2B10G10R10UnsignedIntPacked        Format = C.VK_FORMAT_A2B10G10R10_UINT_PACK32
	FormatA2B10G10R10SignedIntPacked          Format = C.VK_FORMAT_A2B10G10R10_SINT_PACK32

	FormatR16UnsignedNormalized Format = C.VK_FORMAT_R16_UNORM
	FormatR16SignedNormalized   Format = C.VK_FORMAT_R16_SNORM
	FormatR16UnsignedScaled     Format = C.VK_FORMAT_R16_USCALED
	FormatR16SignedScaled       Format = C.VK_FORMAT_R16_SSCALED
	FormatR16UnsignedInt        Format = C.VK_FORMAT_R16_UINT
	FormatR16SignedInt          Format = C.VK_FORMAT_R16_SINT
	FormatR16SignedFloat        Format = C.VK_FORMAT_R16_SFLOAT

	FormatR16G16UnsignedNormalized Format = C.VK_FORMAT_R16G16_UNORM
	FormatR16G16SignedNormalized   Format = C.VK_FORMAT_R16G16_SNORM
	FormatR16G16UnsignedScaled     Format = C.VK_FORMAT_R16G16_USCALED
	FormatR16G16SignedScaled       Format = C.VK_FORMAT_R16G16_SSCALED
	FormatR16G16UnsignedInt        Format = C.VK_FORMAT_R16G16_UINT
	FormatR16G16SignedInt          Format = C.VK_FORMAT_R16G16_SINT
	FormatR16G16SignedFloat        Format = C.VK_FORMAT_R16G16_SFLOAT

	FormatR16G16B16UnsignedNormalized Format = C.VK_FORMAT_R16G16B16_UNORM
	FormatR16G16B16SignedNormalized   Format = C.VK_FORMAT_R16G16B16_SNORM
	FormatR16G16B16UnsignedScaled     Format = C.VK_FORMAT_R16G16B16_USCALED
	FormatR16G16B16SignedScaled       Format = C.VK_FORMAT_R16G16B16_SSCALED
	FormatR16G16B16UnsignedInt        Format = C.VK_FORMAT_R16G16B16_UINT
	FormatR16G16B16SignedInt          Format = C.VK_FORMAT_R16G16B16_SINT
	FormatR16G16B16SignedFloat        Format = C.VK_FORMAT_R16G16B16_SFLOAT

	FormatR16G16B16A16UnsignedNormalized Format = C.VK_FORMAT_R16G16B16A16_UNORM
	FormatR16G16B16A16SignedNormalized   Format = C.VK_FORMAT_R16G16B16A16_SNORM
	FormatR16G16B16A16UnsignedScaled     Format = C.VK_FORMAT_R16G16B16A16_USCALED
	FormatR16G16B16A16SignedScaled       Format = C.VK_FORMAT_R16G16B16A16_SSCALED
	FormatR16G16B16A16UnsignedInt        Format = C.VK_FORMAT_R16G16B16A16_UINT
	FormatR16G16B16A16SignedInt          Format = C.VK_FORMAT_R16G16B16A16_SINT
	FormatR16G16B16A16SignedFloat        Format = C.VK_FORMAT_R16G16B16A16_SFLOAT

	FormatR32UnsignedInt          Format = C.VK_FORMAT_R32_UINT
	FormatR32SignedInt            Format = C.VK_FORMAT_R32_SINT
	FormatR32SignedFloat          Format = C.VK_FORMAT_R32_SFLOAT
	FormatR32G32UnsignedInt       Format = C.VK_FORMAT_R32G32_UINT
	FormatR32G32SignedInt         Format = C.VK_FORMAT_R32G32_SINT
	FormatR32G32SignedFloat       Format = C.VK_FORMAT_R32G32_SFLOAT
	FormatR32G32B32UnsignedInt    Format = C.VK_FORMAT_R32G32B32_UINT
	FormatR32G32B32SignedInt      Format = C.VK_FORMAT_R32G32B32_SINT
	FormatR32G32B32SignedFloat    Format = C.VK_FORMAT_R32G32B32_SFLOAT
	FormatR32G32B32A32UnsignedInt Format = C.VK_FORMAT_R32G32B32A32_UINT
	FormatR32G32B32A32SignedInt   Format = C.VK_FORMAT_R32G32B32A32_SINT
	FormatR32G32B32A32SignedFloat Format = C.VK_FORMAT_R32G32B32A32_SFLOAT

	FormatR64UnsignedInt          Format = C.VK_FORMAT_R64_UINT
	FormatR64SignedInt            Format = C.VK_FORMAT_R64_SINT
	FormatR64SignedFloat          Format = C.VK_FORMAT_R64_SFLOAT
	FormatR64G64UnsignedInt       Format = C.VK_FORMAT_R64G64_UINT
	FormatR64G64SignedInt         Format = C.VK_FORMAT_R64G64_SINT
	FormatR64G64SignedFloat       Format = C.VK_FORMAT_R64G64_SFLOAT
	FormatR64G64B64UnsignedInt    Format = C.VK_FORMAT_R64G64B64_UINT
	FormatR64G64B64SignedInt      Format = C.VK_FORMAT_R64G64B64_SINT
	FormatR64G64B64SignedFloat    Format = C.VK_FORMAT_R64G64B64_SFLOAT
	FormatR64G64B64A64UnsignedInt Format = C.VK_FORMAT_R64G64B64A64_UINT
	FormatR64G64B64A64SignedInt   Format = C.VK_FORMAT_R64G64B64A64_SINT
	FormatR64G64B64A64SignedFloat Format = C.VK_FORMAT_R64G64B64A64_SFLOAT

	FormatB10G11R11UnsignedFloatPacked  Format = C.VK_FORMAT_B10G11R11_UFLOAT_PACK32
	FormatE5B9G9R9UnsignedFloatPacked   Format = C.VK_FORMAT_E5B9G9R9_UFLOAT_PACK32
	FormatD16UnsignedNormalized         Format = C.VK_FORMAT_D16_UNORM
	FormatD24X8UnsignedNormalizedPacked Format = C.VK_FORMAT_X8_D24_UNORM_PACK32
	FormatD32SignedFloat                Format = C.VK_FORMAT_D32_SFLOAT
	FormatS8UnsignedInt                 Format = C.VK_FORMAT_S8_UINT

	FormatD16UnsignedNormalizedS8UnsignedInt Format = C.VK_FORMAT_D16_UNORM_S8_UINT
	FormatD24UnsignedNormalizedS8UnsignedInt Format = C.VK_FORMAT_D24_UNORM_S8_UINT
	FormatD32SignedFloatS8UnsignedInt        Format = C.VK_FORMAT_D32_SFLOAT_S8_UINT

	FormatBC1_RGBUnsignedNormalized  Format = C.VK_FORMAT_BC1_RGB_UNORM_BLOCK
	FormatBC1_RGBsRGB                Format = C.VK_FORMAT_BC1_RGB_SRGB_BLOCK
	FormatBC1_RGBAUnsignedNormalized Format = C.VK_FORMAT_BC1_RGBA_UNORM_BLOCK
	FormatBC1_RGBAsRGB               Format = C.VK_FORMAT_BC1_RGBA_SRGB_BLOCK

	FormatBC2_UnsignedNormalized Format = C.VK_FORMAT_BC2_UNORM_BLOCK
	FormatBC2_sRGB               Format = C.VK_FORMAT_BC2_SRGB_BLOCK

	FormatBC3_UnsignedNormalized Format = C.VK_FORMAT_BC3_UNORM_BLOCK
	FormatBC3_sRGB               Format = C.VK_FORMAT_BC3_SRGB_BLOCK

	FormatBC4_UnsignedNormalized Format = C.VK_FORMAT_BC4_UNORM_BLOCK
	FormatBC4_SignedNormalized   Format = C.VK_FORMAT_BC4_SNORM_BLOCK

	FormatBC5_UnsignedNormalized Format = C.VK_FORMAT_BC5_UNORM_BLOCK
	FormatBC5_SignedNormalized   Format = C.VK_FORMAT_BC5_SNORM_BLOCK

	FormatBC6_UnsignedFloat Format = C.VK_FORMAT_BC6H_UFLOAT_BLOCK
	FormatBC6_SignedFloat   Format = C.VK_FORMAT_BC6H_SFLOAT_BLOCK

	FormatBC7_UnsignedNormalized Format = C.VK_FORMAT_BC7_UNORM_BLOCK
	FormatBC7_sRGB               Format = C.VK_FORMAT_BC7_SRGB_BLOCK

	FormatETC2_R8G8B8UnsignedNormalized   Format = C.VK_FORMAT_ETC2_R8G8B8_UNORM_BLOCK
	FormatETC2_R8G8B8sRGB                 Format = C.VK_FORMAT_ETC2_R8G8B8_SRGB_BLOCK
	FormatETC2_R8G8B8A1UnsignedNormalized Format = C.VK_FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK
	FormatETC2_R8G8B8A1sRGB               Format = C.VK_FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK
	FormatETC2_R8G8B8A8UnsignedNormalized Format = C.VK_FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK
	FormatETC2_R8G8B8A8sRGB               Format = C.VK_FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK

	FormatEAC_R11UnsignedNormalized    Format = C.VK_FORMAT_EAC_R11_UNORM_BLOCK
	FormatEAC_R11SignedNormalized      Format = C.VK_FORMAT_EAC_R11_SNORM_BLOCK
	FormatEAC_R11G11UnsignedNormalized Format = C.VK_FORMAT_EAC_R11G11_UNORM_BLOCK
	FormatEAC_R11G11SignedNormalized   Format = C.VK_FORMAT_EAC_R11G11_SNORM_BLOCK

	FormatASTC4x4_UnsignedNormalized   Format = C.VK_FORMAT_ASTC_4x4_UNORM_BLOCK
	FormatASTC4x4_sRGB                 Format = C.VK_FORMAT_ASTC_4x4_SRGB_BLOCK
	FormatASTC5x4_UnsignedNormalized   Format = C.VK_FORMAT_ASTC_5x4_UNORM_BLOCK
	FormatASTC5x4_sRGB                 Format = C.VK_FORMAT_ASTC_5x4_SRGB_BLOCK
	FormatASTC5x5_UnsignedNormalized   Format = C.VK_FORMAT_ASTC_5x5_UNORM_BLOCK
	FormatASTC5x5_sRGB                 Format = C.VK_FORMAT_ASTC_5x5_SRGB_BLOCK
	FormatASTC6x5_UnsignedNormalized   Format = C.VK_FORMAT_ASTC_6x5_UNORM_BLOCK
	FormatASTC6x5_sRGB                 Format = C.VK_FORMAT_ASTC_6x5_SRGB_BLOCK
	FormatASTC6x6_UnsignedNormalized   Format = C.VK_FORMAT_ASTC_6x6_UNORM_BLOCK
	FormatASTC6x6_sRGB                 Format = C.VK_FORMAT_ASTC_6x6_SRGB_BLOCK
	FormatASTC8x5_UnsignedNormalized   Format = C.VK_FORMAT_ASTC_8x5_UNORM_BLOCK
	FormatASTC8x5_sRGB                 Format = C.VK_FORMAT_ASTC_8x5_SRGB_BLOCK
	FormatASTC8x6_UnsignedNormalized   Format = C.VK_FORMAT_ASTC_8x6_UNORM_BLOCK
	FormatASTC8x6_sRGB                 Format = C.VK_FORMAT_ASTC_8x6_SRGB_BLOCK
	FormatASTC8x8_UnsignedNormalized   Format = C.VK_FORMAT_ASTC_8x8_UNORM_BLOCK
	FormatASTC8x8_sRGB                 Format = C.VK_FORMAT_ASTC_8x8_SRGB_BLOCK
	FormatASTC10x5_UnsignedNormalized  Format = C.VK_FORMAT_ASTC_10x5_UNORM_BLOCK
	FormatASTC10x5_sRGB                Format = C.VK_FORMAT_ASTC_10x5_SRGB_BLOCK
	FormatASTC10x6_UnsignedNormalized  Format = C.VK_FORMAT_ASTC_10x6_UNORM_BLOCK
	FormatASTC10x6_sRGB                Format = C.VK_FORMAT_ASTC_10x6_SRGB_BLOCK
	FormatASTC10x8_UnsignedNormalized  Format = C.VK_FORMAT_ASTC_10x8_UNORM_BLOCK
	FormatASTC10x8_sRGB                Format = C.VK_FORMAT_ASTC_10x8_SRGB_BLOCK
	FormatASTC10x10_UnsignedNormalized Format = C.VK_FORMAT_ASTC_10x10_UNORM_BLOCK
	FormatASTC10x10_sRGB               Format = C.VK_FORMAT_ASTC_10x10_SRGB_BLOCK
	FormatASTC12x10_UnsignedNormalized Format = C.VK_FORMAT_ASTC_12x10_UNORM_BLOCK
	FormatASTC12x10_sRGB               Format = C.VK_FORMAT_ASTC_12x10_SRGB_BLOCK
	FormatASTC12x12_UnsignedNormalized Format = C.VK_FORMAT_ASTC_12x12_UNORM_BLOCK
	FormatASTC12x12_sRGB               Format = C.VK_FORMAT_ASTC_12x12_SRGB_BLOCK
)

func init() {
	FormatUndefined.Register("Undefined")
	FormatR4G4UnsignedNormalizedPacked.Register("R4G4 Unsigned Normalized (Packed)")

	FormatR4G4B4A4UnsignedNormalizedPacked.Register("R4G4B4A4 Unsigned Normalized (Packed)")
	FormatB4G4R4A4UnsignedNormalizedPacked.Register("B4G4R4A4 Unsigned Normalized (Packed)")
	FormatR5G6B5UnsignedNormalizedPacked.Register("R5G6B5 Unsigned Normalized (Packed)")
	FormatB5G6R5UnsignedNormalizedPacked.Register("B5G6R5 Unsigned Normalized (Packed)")
	FormatR5G5B5A1UnsignedNormalizedPacked.Register("R5G5B5A1 Unsigned Normalized (Packed)")
	FormatB5G5R5A1UnsignedNormalizedPacked.Register("B5G5R5A1 Unsigned Normalized (Packed)")
	FormatA1R5G5B5UnsignedNormalizedPacked.Register("A1R5G5B5 Unsigned Normalized (Packed)")

	FormatR8UnsignedNormalized.Register("R8 Unsigned Normalized")
	FormatR8SignedNormalized.Register("R8 Signed Normalized")
	FormatR8UnsignedScaled.Register("R8 Unsigned Scaled")
	FormatR8SignedScaled.Register("R8 Signed Scaled")
	FormatR8UnsignedInt.Register("R8 Unsigned Int")
	FormatR8SignedInt.Register("R8 Signed Int")
	FormatR8SRGB.Register("R8 sRGB")

	FormatR8G8UnsignedNormalized.Register("R8G8 Unsigned Normalized")
	FormatR8G8SignedNormalized.Register("R8G8 Signed Normalized")
	FormatR8G8UnsignedScaled.Register("R8G8 Unsigned Scaled")
	FormatR8G8SignedScaled.Register("R8G8 Signed Scaled")
	FormatR8G8UnsignedInt.Register("R8G8 Unsigned Int")
	FormatR8G8SignedInt.Register("R8G8 Signed Int")
	FormatR8G8SRGB.Register("R8G8 sRGB")

	FormatR8G8B8UnsignedNormalized.Register("R8G8B8 Unsigned Normalized")
	FormatR8G8B8SignedNormalized.Register("R8G8B8 Signed Normalized")
	FormatR8G8B8UnsignedScaled.Register("R8G8B8 Unsigned Scaled")
	FormatR8G8B8SignedScaled.Register("R8G8B8 Signed Scaled")
	FormatR8G8B8UnsignedInt.Register("R8G8B8 Unsigned Int")
	FormatR8G8B8SignedInt.Register("R8G8B8 Signed Int")
	FormatR8G8B8SRGB.Register("R8G8B8 sRGB")

	FormatB8G8R8UnsignedNormalized.Register("B8G8R8 Unsigned Normalized")
	FormatB8G8R8SignedNormalized.Register("B8G8R8 Signed Normalized")
	FormatB8G8R8UnsignedScaled.Register("B8G8R8 Unsigned Scaled")
	FormatB8G8R8SignedScaled.Register("B8G8R8 Signed Scaled")
	FormatB8G8R8UnsignedInt.Register("B8G8R8 Unsigned Int")
	FormatB8G8R8SignedInt.Register("B8G8R8 Signed Int")
	FormatB8G8R8SRGB.Register("B8G8R8 sRGB")

	FormatR8G8B8A8UnsignedNormalized.Register("R8G8B8A8 Unsigned Normalized")
	FormatR8G8B8A8SignedNormalized.Register("R8G8B8A8 Signed Normalized")
	FormatR8G8B8A8UnsignedScaled.Register("R8G8B8A8 Unsigned Scaled")
	FormatR8G8B8A8SignedScaled.Register("R8G8B8A8 Signed Scaled")
	FormatR8G8B8A8UnsignedInt.Register("R8G8B8A8 Unsigned Int")
	FormatR8G8B8A8SignedInt.Register("R8G8B8A8 Signed Int")
	FormatR8G8B8A8SRGB.Register("R8G8B8A8 sRGB")

	FormatB8G8R8A8UnsignedNormalized.Register("B8G8R8A8 Unsigned Normalized")
	FormatB8G8R8A8SignedNormalized.Register("B8G8R8A8 Signed Normalized")
	FormatB8G8R8A8UnsignedScaled.Register("B8G8R8A8 Unsigned Scaled")
	FormatB8G8R8A8SignedScaled.Register("B8G8R8A8 Signed Scaled")
	FormatB8G8R8A8UnsignedInt.Register("B8G8R8A8 Unsigned Int")
	FormatB8G8R8A8SignedInt.Register("B8G8R8A8 Signed Int")
	FormatB8G8R8A8SRGB.Register("B8G8R8A8 sRGB")

	FormatA8B8G8R8UnsignedNormalizedPacked.Register("A8B8G8R8 Unsigned Normalized (Packed)")
	FormatA8B8G8R8SignedNormalizedPacked.Register("A8B8G8R8 Signed Normalized (Packed)")
	FormatA8B8G8R8UnsignedScaledPacked.Register("A8B8G8R8 Unsigned Scaled (Packed)")
	FormatA8B8G8R8SignedScaledPacked.Register("A8B8G8R8 Signed Scaled (Packed)")
	FormatA8B8G8R8UnsignedIntPacked.Register("A8B8G8R8 Unsigned Int (Packed)")
	FormatA8B8G8R8SignedIntPacked.Register("A8B8G8R8 Signed Int (Packed)")
	FormatA8B8G8R8SRGBPacked.Register("A8B8G8R8 sRGB (Packed)")

	FormatA2R10G10B10UnsignedNormalizedPacked.Register("A2R10G10B10 Unsigned Normalized (Packed)")
	FormatA2R10G10B10SignedNormalizedPacked.Register("A2R10G10B10 Signed Normalized (Packed)")
	FormatA2R10G10B10UnsignedScaledPacked.Register("A2R10G10B10 Unsigned Scaled (Packed)")
	FormatA2R10G10B10SignedScaledPacked.Register("A2R10G10B10 Signed Scaled (Packed)")
	FormatA2R10G10B10UnsignedIntPacked.Register("A2R10G10B10 Unsigned Int (Packed)")
	FormatA2R10G10B10SignedIntPacked.Register("A2R10G10B10 Signed Int (Packed)")

	FormatA2B10G10R10UnsignedNormalizedPacked.Register("A2B10G10R10 Unsigned Normalized (Packed)")
	FormatA2B10G10R10SignedNormalizedPacked.Register("A2B10G10R10 Signed Normalized (Packed)")
	FormatA2B10G10R10UnsignedScaledPacked.Register("A2B10G10R10 Unsigned Scaled (Packed)")
	FormatA2B10G10R10SignedScaledPacked.Register("A2B10G10R10 Signed Scaled (Packed)")
	FormatA2B10G10R10UnsignedIntPacked.Register("A2B10G10R10 Unsigned Int (Packed)")
	FormatA2B10G10R10SignedIntPacked.Register("A2B10G10R10 Signed Int (Packed)")

	FormatR16UnsignedNormalized.Register("R16 Unsigned Normalized")
	FormatR16SignedNormalized.Register("R16 Signed Normalized")
	FormatR16UnsignedScaled.Register("R16 Unsigned Scaled")
	FormatR16SignedScaled.Register("R16 Signed Scaled")
	FormatR16UnsignedInt.Register("R16 Unsigned Int")
	FormatR16SignedInt.Register("R16 Signed Int")
	FormatR16SignedFloat.Register("R16 Signed Float")

	FormatR16G16UnsignedNormalized.Register("R16G16 Unsigned Normalized")
	FormatR16G16SignedNormalized.Register("R16G16 Signed Normalized")
	FormatR16G16UnsignedScaled.Register("R16G16 Unsigned Scaled")
	FormatR16G16SignedScaled.Register("R16G16 Signed Scaled")
	FormatR16G16UnsignedInt.Register("R16G16 Unsigned Int")
	FormatR16G16SignedInt.Register("R16G16 Signed Int")
	FormatR16G16SignedFloat.Register("R16G16 Signed Float")

	FormatR16G16B16UnsignedNormalized.Register("R16G16B16 Unsigned Normalized")
	FormatR16G16B16SignedNormalized.Register("R16G16B16 Signed Normalized")
	FormatR16G16B16UnsignedScaled.Register("R16G16B16 Unsigned Scaled")
	FormatR16G16B16SignedScaled.Register("R16G16B16 Signed Scaled")
	FormatR16G16B16UnsignedInt.Register("R16G16B16 Unsigned Int")
	FormatR16G16B16SignedInt.Register("R16G16B16 Signed Int")
	FormatR16G16B16SignedFloat.Register("R16G16B16 Signed Float")

	FormatR16G16B16A16UnsignedNormalized.Register("R16G16B16A16 Unsigned Normalized")
	FormatR16G16B16A16SignedNormalized.Register("R16G16B16A16 Signed Normalized")
	FormatR16G16B16A16UnsignedScaled.Register("R16G16B16A16 Unsigned Scaled")
	FormatR16G16B16A16SignedScaled.Register("R16G16B16A16 Signed Scaled")
	FormatR16G16B16A16UnsignedInt.Register("R16G16B16A16 Unsigned Int")
	FormatR16G16B16A16SignedInt.Register("R16G16B16A16 Signed Int")
	FormatR16G16B16A16SignedFloat.Register("R16G16B16A16 Signed Float")

	FormatR32UnsignedInt.Register("R32 Unsigned Int")
	FormatR32SignedInt.Register("R32 Signed Int")
	FormatR32SignedFloat.Register("R32 Signed Float")
	FormatR32G32UnsignedInt.Register("R32G32 Unsigned Int")
	FormatR32G32SignedInt.Register("R32G32 Signed Int")
	FormatR32G32SignedFloat.Register("R32G32 Signed Float")
	FormatR32G32B32UnsignedInt.Register("R32G32B32 Unsigned Int")
	FormatR32G32B32SignedInt.Register("R32G32B32 Signed Int")
	FormatR32G32B32SignedFloat.Register("R32G32B32 Signed Float")
	FormatR32G32B32A32UnsignedInt.Register("R32G32B32A32 Unsigned Int")
	FormatR32G32B32A32SignedInt.Register("R32G32B32A32 Signed Int")
	FormatR32G32B32A32SignedFloat.Register("R32G32B32A32 Signed Float")

	FormatR64UnsignedInt.Register("R64 Unsigned Int")
	FormatR64SignedInt.Register("R64 Signed Int")
	FormatR64SignedFloat.Register("R64 Signed Float")
	FormatR64G64UnsignedInt.Register("R64G64 Unsigned Int")
	FormatR64G64SignedInt.Register("R64G64 Signed Int")
	FormatR64G64SignedFloat.Register("R64G64 Signed Float")
	FormatR64G64B64UnsignedInt.Register("R64G64B64 Unsigned Int")
	FormatR64G64B64SignedInt.Register("R64G64B64 Signed Int")
	FormatR64G64B64SignedFloat.Register("R64G64B64 Signed Float")
	FormatR64G64B64A64UnsignedInt.Register("R64G64B64A64 Unsigned Int")
	FormatR64G64B64A64SignedInt.Register("R64G64B64A64 Signed Int")
	FormatR64G64B64A64SignedFloat.Register("R64G64B64A64 Signed Float")

	FormatB10G11R11UnsignedFloatPacked.Register("B10G11R11 Unsigned Float (Packed)")
	FormatE5B9G9R9UnsignedFloatPacked.Register("E5B9G9R9 Unsigned Float (Packed)")
	FormatD16UnsignedNormalized.Register("D16 Unsigned Normalized")
	FormatD24X8UnsignedNormalizedPacked.Register("D24X8 Unsigned Normalized (Packed)")
	FormatD32SignedFloat.Register("D32 Signed Float")
	FormatS8UnsignedInt.Register("S8 Unsigned Int")

	FormatD16UnsignedNormalizedS8UnsignedInt.Register("D16 Unsigned Normalized S8 Unsigned Int")
	FormatD24UnsignedNormalizedS8UnsignedInt.Register("D24 Unsigned Normalized S8 Unsigned Int")
	FormatD32SignedFloatS8UnsignedInt.Register("D32 Signed Float S8 Unsigned Int")

	FormatBC1_RGBUnsignedNormalized.Register("BC1-Compressed -Compressed RGB Unsigned Normalized")
	FormatBC1_RGBsRGB.Register("BC1-Compressed -Compressed RGB sRGB")
	FormatBC1_RGBAUnsignedNormalized.Register("BC1-Compressed -Compressed RGBA Unsigned Normalized")
	FormatBC1_RGBAsRGB.Register("BC1-Compressed RGBA sRGB")

	FormatBC2_UnsignedNormalized.Register("BC2-Compressed Unsigned Normalized")
	FormatBC2_sRGB.Register("BC2-Compressed sRGB")

	FormatBC3_UnsignedNormalized.Register("BC3-Compressed Unsigned Normalized")
	FormatBC3_sRGB.Register("BC3-Compressed sRGB")

	FormatBC4_UnsignedNormalized.Register("BC4-Compressed Unsigned Normalized")
	FormatBC4_SignedNormalized.Register("BC4-Compressed Signed Normalized")

	FormatBC5_UnsignedNormalized.Register("BC5-Compressed Unsigned Normalized")
	FormatBC5_SignedNormalized.Register("BC5-Compressed Signed Normalized")

	FormatBC6_UnsignedFloat.Register("BC6-Compressed Unsigned Float")
	FormatBC6_SignedFloat.Register("BC6-Compressed Signed Float")

	FormatBC7_UnsignedNormalized.Register("BC7-Compressed Unsigned Normalized")
	FormatBC7_sRGB.Register("BC7-Compressed sRGB")

	FormatETC2_R8G8B8UnsignedNormalized.Register("ETC2-Compressed R8G8B8 Unsigned Normalized")
	FormatETC2_R8G8B8sRGB.Register("ETC2-Compressed R8G8B8 sRGB")
	FormatETC2_R8G8B8A1UnsignedNormalized.Register("ETC2-Compressed R8G8B8A1 Unsigned Normalized")
	FormatETC2_R8G8B8A1sRGB.Register("ETC2-Compressed R8G8B8A1 sRGB")
	FormatETC2_R8G8B8A8UnsignedNormalized.Register("ETC2-Compressed R8G8B8A8 Unsigned Normalized")
	FormatETC2_R8G8B8A8sRGB.Register("ETC2-Compressed R8G8B8A8 sRGB")

	FormatEAC_R11UnsignedNormalized.Register("EAC-Compressed R11 Unsigned Normalized")
	FormatEAC_R11SignedNormalized.Register("EAC-Compressed R11 Signed Normalized")
	FormatEAC_R11G11UnsignedNormalized.Register("EAC-Compressed R11G11 Unsigned Normalized")
	FormatEAC_R11G11SignedNormalized.Register("EAC-Compressed R11G11 Signed Normalized")

	FormatASTC4x4_UnsignedNormalized.Register("ASTC-Compressed (4x4) Unsigned Normalized")
	FormatASTC4x4_sRGB.Register("ASTC-Compressed (4x4) sRGB")
	FormatASTC5x4_UnsignedNormalized.Register("ASTC-Compressed (5x4) Unsigned Normalized")
	FormatASTC5x4_sRGB.Register("ASTC-Compressed (5x4) sRGB")
	FormatASTC5x5_UnsignedNormalized.Register("ASTC-Compressed (5x5) Unsigned Normalized")
	FormatASTC5x5_sRGB.Register("ASTC-Compressed (5x5) sRGB")
	FormatASTC6x5_UnsignedNormalized.Register("ASTC-Compressed (6x5) Unsigned Normalized")
	FormatASTC6x5_sRGB.Register("ASTC-Compressed (6x5) sRGB")
	FormatASTC6x6_UnsignedNormalized.Register("ASTC-Compressed (6x6) Unsigned Normalized")
	FormatASTC6x6_sRGB.Register("ASTC-Compressed (6x6) sRGB")
	FormatASTC8x5_UnsignedNormalized.Register("ASTC-Compressed (8x5) Unsigned Normalized")
	FormatASTC8x5_sRGB.Register("ASTC-Compressed (8x5) sRGB")
	FormatASTC8x6_UnsignedNormalized.Register("ASTC-Compressed (8x6) Unsigned Normalized")
	FormatASTC8x6_sRGB.Register("ASTC-Compressed (8x6) sRGB")
	FormatASTC8x8_UnsignedNormalized.Register("ASTC-Compressed (8x8) Unsigned Normalized")
	FormatASTC8x8_sRGB.Register("ASTC-Compressed (8x8) sRGB")
	FormatASTC10x5_UnsignedNormalized.Register("ASTC-Compressed (10x5) Unsigned Normalized")
	FormatASTC10x5_sRGB.Register("ASTC-Compressed (10x5) sRGB")
	FormatASTC10x6_UnsignedNormalized.Register("ASTC-Compressed (10x6) Unsigned Normalized")
	FormatASTC10x6_sRGB.Register("ASTC-Compressed (10x6) sRGB")
	FormatASTC10x8_UnsignedNormalized.Register("ASTC-Compressed (10x8) Unsigned Normalized")
	FormatASTC10x8_sRGB.Register("ASTC-Compressed (10x8) sRGB")
	FormatASTC10x10_UnsignedNormalized.Register("ASTC-Compressed (10x10) Unsigned Normalized")
	FormatASTC10x10_sRGB.Register("ASTC-Compressed (10x10) sRGB")
	FormatASTC12x10_UnsignedNormalized.Register("ASTC-Compressed (12x10) Unsigned Normalized")
	FormatASTC12x10_sRGB.Register("ASTC-Compressed (12x10) sRGB")
	FormatASTC12x12_UnsignedNormalized.Register("ASTC-Compressed (12x12) Unsigned Normalized")
	FormatASTC12x12_sRGB.Register("ASTC-Compressed (12x12) sRGB")
}
