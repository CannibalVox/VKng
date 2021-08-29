package pipeline

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "strings"

type BlendFactor int32

const (
	BlendZero                  BlendFactor = C.VK_BLEND_FACTOR_ZERO
	BlendOne                   BlendFactor = C.VK_BLEND_FACTOR_ONE
	BlendSrcColor              BlendFactor = C.VK_BLEND_FACTOR_SRC_COLOR
	BlendOneMinusSrcColor      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_COLOR
	BlendDstColor              BlendFactor = C.VK_BLEND_FACTOR_DST_COLOR
	BlendOneMinusDstColor      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_DST_COLOR
	BlendSrcAlpha              BlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA
	BlendOneMinusSrcAlpha      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC_ALPHA
	BlendDstAlpha              BlendFactor = C.VK_BLEND_FACTOR_DST_ALPHA
	BlendOneMinusDstAlpha      BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_DST_ALPHA
	BlendConstantColor         BlendFactor = C.VK_BLEND_FACTOR_CONSTANT_COLOR
	BlendOneMinusConstantColor BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_COLOR
	BlendConstantAlpha         BlendFactor = C.VK_BLEND_FACTOR_CONSTANT_ALPHA
	BlendOneMinusConstantAlpha BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_CONSTANT_ALPHA
	BlendSrcAlphaSaturate      BlendFactor = C.VK_BLEND_FACTOR_SRC_ALPHA_SATURATE
	BlendSrc1Color             BlendFactor = C.VK_BLEND_FACTOR_SRC1_COLOR
	BlendOneMinusSrc1Color     BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR
	BlendSrc1Alpha             BlendFactor = C.VK_BLEND_FACTOR_SRC1_ALPHA
	BlendOneMinusSrc1Alpha     BlendFactor = C.VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA
)

var blendFactorToString = map[BlendFactor]string{
	BlendZero:                  "0",
	BlendOne:                   "1",
	BlendSrcColor:              "Src Color",
	BlendOneMinusSrcColor:      "1 - Src Color",
	BlendDstColor:              "Dst Color",
	BlendOneMinusDstColor:      "1 - Dst Color",
	BlendSrcAlpha:              "Src Alpha",
	BlendOneMinusSrcAlpha:      "1 - Src Alpha",
	BlendDstAlpha:              "Dst Alpha",
	BlendOneMinusDstAlpha:      "1 - Dst Alpha",
	BlendConstantColor:         "Constant Color",
	BlendOneMinusConstantColor: "1 - Constant Color",
	BlendConstantAlpha:         "Constant Alpha",
	BlendOneMinusConstantAlpha: "1 - Constant Alpha",
	BlendSrcAlphaSaturate:      "Alpha Saturate",
	BlendSrc1Color:             "Src1 Color",
	BlendOneMinusSrc1Color:     "1 - Src1 Color",
	BlendSrc1Alpha:             "Src1 Alpha",
	BlendOneMinusSrc1Alpha:     "1 - Src1 Alpha",
}

func (f BlendFactor) String() string {
	return blendFactorToString[f]
}

type BlendOp int32

const (
	OpAdd              BlendOp = C.VK_BLEND_OP_ADD
	OpSubtract         BlendOp = C.VK_BLEND_OP_SUBTRACT
	OpMin              BlendOp = C.VK_BLEND_OP_MIN
	OpMax              BlendOp = C.VK_BLEND_OP_MAX
	OpZero             BlendOp = C.VK_BLEND_OP_ZERO_EXT
	OpSrc              BlendOp = C.VK_BLEND_OP_SRC_EXT
	OpDst              BlendOp = C.VK_BLEND_OP_DST_EXT
	OpSrcOver          BlendOp = C.VK_BLEND_OP_SRC_OVER_EXT
	OpDstOver          BlendOp = C.VK_BLEND_OP_DST_OVER_EXT
	OpSrcIn            BlendOp = C.VK_BLEND_OP_SRC_IN_EXT
	OpDstIn            BlendOp = C.VK_BLEND_OP_DST_IN_EXT
	OpSrcOut           BlendOp = C.VK_BLEND_OP_SRC_OUT_EXT
	OpDstOut           BlendOp = C.VK_BLEND_OP_DST_OUT_EXT
	OpSrcAtop          BlendOp = C.VK_BLEND_OP_SRC_ATOP_EXT
	OpDstAtop          BlendOp = C.VK_BLEND_OP_DST_ATOP_EXT
	OpXor              BlendOp = C.VK_BLEND_OP_XOR_EXT
	OpMultiply         BlendOp = C.VK_BLEND_OP_MULTIPLY_EXT
	OpScreen           BlendOp = C.VK_BLEND_OP_SCREEN_EXT
	OpOverlay          BlendOp = C.VK_BLEND_OP_OVERLAY_EXT
	OpDarken           BlendOp = C.VK_BLEND_OP_DARKEN_EXT
	OpLighten          BlendOp = C.VK_BLEND_OP_LIGHTEN_EXT
	OpColorDodge       BlendOp = C.VK_BLEND_OP_COLORDODGE_EXT
	OpColorBurn        BlendOp = C.VK_BLEND_OP_COLORBURN_EXT
	OpHardLight        BlendOp = C.VK_BLEND_OP_HARDLIGHT_EXT
	OpSoftLight        BlendOp = C.VK_BLEND_OP_SOFTLIGHT_EXT
	OpDifference       BlendOp = C.VK_BLEND_OP_DIFFERENCE_EXT
	OpExclusion        BlendOp = C.VK_BLEND_OP_EXCLUSION_EXT
	OpInvert           BlendOp = C.VK_BLEND_OP_INVERT_EXT
	OpInvertRGB        BlendOp = C.VK_BLEND_OP_INVERT_RGB_EXT
	OpLinearDodge      BlendOp = C.VK_BLEND_OP_LINEARDODGE_EXT
	OpLinearBurn       BlendOp = C.VK_BLEND_OP_LINEARBURN_EXT
	OpVividLight       BlendOp = C.VK_BLEND_OP_VIVIDLIGHT_EXT
	OpLinearLight      BlendOp = C.VK_BLEND_OP_LINEARLIGHT_EXT
	OpPinLight         BlendOp = C.VK_BLEND_OP_PINLIGHT_EXT
	OpHardMix          BlendOp = C.VK_BLEND_OP_HARDMIX_EXT
	OpHSLHue           BlendOp = C.VK_BLEND_OP_HSL_HUE_EXT
	OpHSLSaturation    BlendOp = C.VK_BLEND_OP_HSL_SATURATION_EXT
	OpHSLColor         BlendOp = C.VK_BLEND_OP_HSL_COLOR_EXT
	OpHSLLuminosity    BlendOp = C.VK_BLEND_OP_HSL_LUMINOSITY_EXT
	OpPlus             BlendOp = C.VK_BLEND_OP_PLUS_EXT
	OpPlusClamped      BlendOp = C.VK_BLEND_OP_PLUS_CLAMPED_EXT
	OpPlusClampedAlpha BlendOp = C.VK_BLEND_OP_PLUS_CLAMPED_ALPHA_EXT
	OpPlusDarker       BlendOp = C.VK_BLEND_OP_PLUS_DARKER_EXT
	OpMinus            BlendOp = C.VK_BLEND_OP_MINUS_EXT
	OpMinusClamped     BlendOp = C.VK_BLEND_OP_MINUS_CLAMPED_EXT
	OpContrast         BlendOp = C.VK_BLEND_OP_CONTRAST_EXT
	OpInvertOVG        BlendOp = C.VK_BLEND_OP_INVERT_OVG_EXT
	OpRed              BlendOp = C.VK_BLEND_OP_RED_EXT
	OpGreen            BlendOp = C.VK_BLEND_OP_GREEN_EXT
	OpBlue             BlendOp = C.VK_BLEND_OP_BLUE_EXT
)

var blendOpToString = map[BlendOp]string{
	OpAdd:              "Add",
	OpSubtract:         "Subtract",
	OpMin:              "Min",
	OpMax:              "Max",
	OpZero:             "Zero",
	OpSrc:              "Src",
	OpDst:              "Dst",
	OpSrcOver:          "Src Over",
	OpDstOver:          "Dst Over",
	OpSrcIn:            "Src In",
	OpDstIn:            "Dst In",
	OpSrcOut:           "Src Out",
	OpDstOut:           "Dst Out",
	OpSrcAtop:          "Src Atop",
	OpDstAtop:          "Dst Atop",
	OpXor:              "Xor",
	OpMultiply:         "Multiply",
	OpScreen:           "Screen",
	OpOverlay:          "Overlay",
	OpDarken:           "Darken",
	OpLighten:          "Lighten",
	OpColorDodge:       "Color Dodge",
	OpColorBurn:        "Color Burn",
	OpHardLight:        "Hard Light",
	OpSoftLight:        "Soft Light",
	OpDifference:       "Difference",
	OpExclusion:        "Exclusion",
	OpInvert:           "Invert",
	OpInvertRGB:        "Invert RGB",
	OpLinearDodge:      "Linear Dodge",
	OpLinearBurn:       "Linear Burn",
	OpVividLight:       "Vivid Light",
	OpLinearLight:      "Linear Light",
	OpPinLight:         "Pin Light",
	OpHardMix:          "Hard Mix",
	OpHSLHue:           "Hue (HSL)",
	OpHSLSaturation:    "Saturation (HSL)",
	OpHSLColor:         "Color (HSL)",
	OpHSLLuminosity:    "Luminosity (HSL)",
	OpPlus:             "Plus",
	OpPlusClamped:      "Plus Clamped",
	OpPlusClampedAlpha: "Plus Clamped Alpha",
	OpPlusDarker:       "Plus Darker",
	OpMinus:            "Minus",
	OpMinusClamped:     "Minus Clamped",
	OpContrast:         "Contrast",
	OpInvertOVG:        "Invert OVG",
	OpRed:              "Red",
	OpGreen:            "Green",
	OpBlue:             "Blue",
}

func (o BlendOp) String() string {
	return blendOpToString[o]
}

type ColorComponents int32

const (
	ComponentRed   ColorComponents = C.VK_COLOR_COMPONENT_R_BIT
	ComponentGreen ColorComponents = C.VK_COLOR_COMPONENT_G_BIT
	ComponentBlue  ColorComponents = C.VK_COLOR_COMPONENT_B_BIT
	ComponentAlpha ColorComponents = C.VK_COLOR_COMPONENT_A_BIT
)

func (c ColorComponents) String() string {
	var sb strings.Builder
	if (c & ComponentRed) != 0 {
		sb.WriteRune('R')
	}
	if (c & ComponentGreen) != 0 {
		sb.WriteRune('G')
	}
	if (c & ComponentBlue) != 0 {
		sb.WriteRune('B')
	}
	if (c & ComponentAlpha) != 0 {
		sb.WriteRune('A')
	}

	return sb.String()
}
