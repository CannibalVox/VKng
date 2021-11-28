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
	ColorSpaceSRGBNonlinear            ColorSpace = C.VK_COLOR_SPACE_SRGB_NONLINEAR_KHR
	ColorSpaceDisplayP3NonlinearEXT    ColorSpace = C.VK_COLOR_SPACE_DISPLAY_P3_NONLINEAR_EXT
	ColorSpaceExtendedSRGBLinearEXT    ColorSpace = C.VK_COLOR_SPACE_EXTENDED_SRGB_LINEAR_EXT
	ColorSpaceDisplayP3LinearEXT       ColorSpace = C.VK_COLOR_SPACE_DISPLAY_P3_LINEAR_EXT
	ColorSpaceDCIP3NonlinearEXT        ColorSpace = C.VK_COLOR_SPACE_DCI_P3_NONLINEAR_EXT
	ColorSpaceBT709LinearEXT           ColorSpace = C.VK_COLOR_SPACE_BT709_LINEAR_EXT
	ColorSpaceBT709NonlinearEXT        ColorSpace = C.VK_COLOR_SPACE_BT709_NONLINEAR_EXT
	ColorSpaceBT2020LinearEXT          ColorSpace = C.VK_COLOR_SPACE_BT2020_LINEAR_EXT
	ColorSpaceHDR10ST2084EXT           ColorSpace = C.VK_COLOR_SPACE_HDR10_ST2084_EXT
	ColorSpaceDolbyVisionEXT           ColorSpace = C.VK_COLOR_SPACE_DOLBYVISION_EXT
	ColorSpaceHDR10HLGEXT              ColorSpace = C.VK_COLOR_SPACE_HDR10_HLG_EXT
	ColorSpaceAdobeRGBLinearEXT        ColorSpace = C.VK_COLOR_SPACE_ADOBERGB_LINEAR_EXT
	ColorSpaceAdobeRGBNonlinearEXT     ColorSpace = C.VK_COLOR_SPACE_ADOBERGB_NONLINEAR_EXT
	ColorSpacePassThroughEXT           ColorSpace = C.VK_COLOR_SPACE_PASS_THROUGH_EXT
	ColorSpaceExtendedSRGBNonlinearEXT ColorSpace = C.VK_COLOR_SPACE_EXTENDED_SRGB_NONLINEAR_EXT
	ColorSpaceDisplayNativeAMD         ColorSpace = C.VK_COLOR_SPACE_DISPLAY_NATIVE_AMD
)

var colorSpaceToString = map[ColorSpace]string{
	ColorSpaceSRGBNonlinear:            "sRGB Non-Linear",
	ColorSpaceDisplayP3NonlinearEXT:    "Display-P3 Non-Linear (Extension)",
	ColorSpaceExtendedSRGBNonlinearEXT: "Extended sRGB Non-Linear (Extension)",
	ColorSpaceExtendedSRGBLinearEXT:    "Extended sRGB Linear (Extension)",
	ColorSpaceDisplayP3LinearEXT:       "Display-P3 Linear (Extension)",
	ColorSpaceDCIP3NonlinearEXT:        "DCI-P3 Non-Linear (Extension)",
	ColorSpaceBT709LinearEXT:           "BT709 Linear (Extension)",
	ColorSpaceBT709NonlinearEXT:        "BT709 non-Linear (Extension)",
	ColorSpaceBT2020LinearEXT:          "BT2020 Linear (Extension)",
	ColorSpaceHDR10ST2084EXT:           "HDR10 (BT2020 Color) - SMPTE ST2084 (Extension)",
	ColorSpaceDolbyVisionEXT:           "Dolby Vision (BT2020 Color) - SMPTE ST2084 (Extension)",
	ColorSpaceHDR10HLGEXT:              "HDR10 (BT2020 Color) - Hybrid Log Gamma (Extension)",
	ColorSpaceAdobeRGBNonlinearEXT:     "AdobeRGB Non-Linear (Extension)",
	ColorSpaceAdobeRGBLinearEXT:        "AdobeRGB Linear (Extension)",
	ColorSpacePassThroughEXT:           "Pass-Through (Extension)",
	ColorSpaceDisplayNativeAMD:         "Display's Native Color Space (AMD Extension)",
}

func (s ColorSpace) String() string {
	return colorSpaceToString[s]
}

type Format struct {
	Format     common.DataFormat
	ColorSpace ColorSpace
}
