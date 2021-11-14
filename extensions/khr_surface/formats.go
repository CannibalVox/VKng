package khr_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
)

type ColorSpace int32

const (
	SRGBNonlinear            ColorSpace = C.VK_COLOR_SPACE_SRGB_NONLINEAR_KHR
	DisplayP3NonlinearEXT    ColorSpace = C.VK_COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT
	ExtendedSRGBLinearEXT    ColorSpace = C.VK_COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT
	DisplayP3LinearEXT       ColorSpace = C.VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT
	DCIP3NonlinearEXT        ColorSpace = C.VK_COLOR_SPACE_DCI_P3_NONLINEAR_EXT
	BT709LinearEXT           ColorSpace = C.VK_COLOR_SPACE_BT709_LINEAR_EXT
	BT709NonlinearEXT        ColorSpace = C.VK_COLOR_SPACE_BT709_NONLINEAR_EXT
	BT2020LinearEXT          ColorSpace = C.VK_COLOR_SPACE_BT2020_LINEAR_EXT
	HDR10ST2084EXT           ColorSpace = C.VK_COLOR_SPACE_HDR10_ST2084_EXT
	DolbyVisionEXT           ColorSpace = C.VK_COLOR_SPACE_DOLBYVISION_EXT
	HDR10HLGEXT              ColorSpace = C.VK_COLOR_SPACE_HDR10_HLG_EXT
	AdobeRGBLinearEXT        ColorSpace = C.VK_COLOR_SPACE_ADOBERGB_LINEAR_EXT
	AdobeRGBNonlinearEXT     ColorSpace = C.VK_COLOR_SPACE_ADOBERGB_NONLINEAR_EXT
	PassThroughEXT           ColorSpace = C.VK_COLOR_SPACE_PASS_THROUGH_EXT
	ExtendedSRGBNonlinearEXT ColorSpace = C.VK_COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT
	DisplayNativeAMD         ColorSpace = C.VK_COLOR_SPACE_DISPLAY_NATIVE_AMD
)

var colorSpaceToString = map[ColorSpace]string{
	SRGBNonlinear:            "sRGB Non-Linear",
	DisplayP3NonlinearEXT:    "Display-P3 Non-Linear (Extension)",
	ExtendedSRGBNonlinearEXT: "Extended sRGB Non-Linear (Extension)",
	ExtendedSRGBLinearEXT:    "Extended sRGB Linear (Extension)",
	DisplayP3LinearEXT:       "Display-P3 Linear (Extension)",
	DCIP3NonlinearEXT:        "DCI-P3 Non-Linear (Extension)",
	BT709LinearEXT:           "BT709 Linear (Extension)",
	BT709NonlinearEXT:        "BT709 non-Linear (Extension)",
	BT2020LinearEXT:          "BT2020 Linear (Extension)",
	HDR10ST2084EXT:           "HDR10 (BT2020 Color) - SMPTE ST2084 (Extension)",
	DolbyVisionEXT:           "Dolby Vision (BT2020 Color) - SMPTE ST2084 (Extension)",
	HDR10HLGEXT:              "HDR10 (BT2020 Color) - Hybrid Log Gamma (Extension)",
	AdobeRGBNonlinearEXT:     "AdobeRGB Non-Linear (Extension)",
	AdobeRGBLinearEXT:        "AdobeRGB Linear (Extension)",
	PassThroughEXT:           "Pass-Through (Extension)",
	DisplayNativeAMD:         "Display's Native Color Space (AMD Extension)",
}

func (s ColorSpace) String() string {
	return colorSpaceToString[s]
}

type Format struct {
	Format     common.DataFormat
	ColorSpace ColorSpace
}
