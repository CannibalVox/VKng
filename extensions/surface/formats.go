package ext_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
)

type ColorSpace int32

const (
	SRGBNonlinear         ColorSpace = C.VK_COLOR_SPACE_SRGB_NONLINEAR_KHR
	DisplayP3Nonlinear    ColorSpace = C.VK_COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT
	ExtendedSRGBLinear    ColorSpace = C.VK_COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT
	DisplayP3Linear       ColorSpace = C.VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT
	DCIP3Nonlinear        ColorSpace = C.VK_COLOR_SPACE_DCI_P3_NONLINEAR_EXT
	BT709Linear           ColorSpace = C.VK_COLOR_SPACE_BT709_LINEAR_EXT
	BT709Nonlinear        ColorSpace = C.VK_COLOR_SPACE_BT709_NONLINEAR_EXT
	BT2020Linear          ColorSpace = C.VK_COLOR_SPACE_BT2020_LINEAR_EXT
	HDR10ST2084           ColorSpace = C.VK_COLOR_SPACE_HDR10_ST2084_EXT
	DolbyVision           ColorSpace = C.VK_COLOR_SPACE_DOLBYVISION_EXT
	HDR10HLG              ColorSpace = C.VK_COLOR_SPACE_HDR10_HLG_EXT
	AdobeRGBLinear        ColorSpace = C.VK_COLOR_SPACE_ADOBERGB_LINEAR_EXT
	AdobeRGBNonlinear     ColorSpace = C.VK_COLOR_SPACE_ADOBERGB_NONLINEAR_EXT
	PassThrough           ColorSpace = C.VK_COLOR_SPACE_PASS_THROUGH_EXT
	ExtendedSRGBNonlinear ColorSpace = C.VK_COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT
	DisplayNativeAMD      ColorSpace = C.VK_COLOR_SPACE_DISPLAY_NATIVE_AMD
)

var colorSpaceToString = map[ColorSpace]string{
	SRGBNonlinear:         "sRGB Non-Linear",
	DisplayP3Nonlinear:    "Display-P3 Non-Linear",
	ExtendedSRGBNonlinear: "Extended sRGB Non-Linear",
	ExtendedSRGBLinear:    "Extended sRGB Linear",
	DisplayP3Linear:       "Display-P3 Linear",
	DCIP3Nonlinear:        "DCI-P3 Non-Linear",
	BT709Linear:           "BT709 Linear",
	BT709Nonlinear:        "BT709 non-Linear",
	BT2020Linear:          "BT2020 Linear",
	HDR10ST2084:           "HDR10 (BT2020 Color) - SMPTE ST2084",
	DolbyVision:           "Dolby Vision (BT2020 Color) - SMPTE ST2084",
	HDR10HLG:              "HDR10 (BT2020 Color) - Hybrid Log Gamma",
	AdobeRGBNonlinear:     "AdobeRGB Non-Linear",
	AdobeRGBLinear:        "AdobeRGB Linear",
	PassThrough:           "Pass-Through",
	DisplayNativeAMD:      "Display's Native Color Space (AMD)",
}

func (s ColorSpace) String() string {
	return colorSpaceToString[s]
}

type Format struct {
	Format     core.DataFormat
	ColorSpace ColorSpace
}
