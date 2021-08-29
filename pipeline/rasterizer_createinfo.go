package pipeline

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core"

type PolygonMode int32

const (
	ModeFill            PolygonMode = C.VK_POLYGON_MODE_FILL
	ModeLine            PolygonMode = C.VK_POLYGON_MODE_LINE
	ModePoint           PolygonMode = C.VK_POLYGON_MODE_POINT
	ModeFillRectangleNV PolygonMode = C.VK_POLYGON_MODE_FILL_RECTANGLE_NV
)

var polygonModeToString = map[PolygonMode]string{
	ModeFill:            "Fill",
	ModeLine:            "Line",
	ModePoint:           "Point",
	ModeFillRectangleNV: "Fill Rectangle (Nvidia)",
}

func (m PolygonMode) String() string {
	return polygonModeToString[m]
}

type RasterizationOptions struct {
	DepthClamp        bool
	RasterizerDiscard bool

	PolygonMode PolygonMode
	CullMode    core.CullMode
	FrontFace   core.FrontFace

	DepthBias              bool
	DepthBiasClamp         float32
	DeptBiasConstantFactor float32
	DepthBiasSlopeFactor   float32

	LineWidth float32

	Next core.Options
}
