package khr_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
)

type SurfaceTransforms int32

const (
	TransformIdentity                  SurfaceTransforms = C.VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR
	TransformRotate90                  SurfaceTransforms = C.VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR
	TransformRotate180                 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR
	TransformRotate270                 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR
	TransformHorizontalMirror          SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR
	TransformHorizontalMirrorRotate90  SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR
	TransformHorizontalMirrorRotate180 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR
	TransformHorizontalMirrorRotate270 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR
	TransformInherit                   SurfaceTransforms = C.VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR
)

var surfaceTransformsToString = map[SurfaceTransforms]string{
	TransformIdentity:                  "Identity",
	TransformRotate90:                  "Rotate 90",
	TransformRotate180:                 "Rotate 180",
	TransformRotate270:                 "Rotate 270",
	TransformHorizontalMirror:          "Horizontal Mirror",
	TransformHorizontalMirrorRotate90:  "Horizontal Mirror & Rotate 90",
	TransformHorizontalMirrorRotate180: "Horizontal Mirror & Rotate 180",
	TransformHorizontalMirrorRotate270: "Horizontal Mirror & Rotate 270",
	TransformInherit:                   "Inherit",
}

func (t SurfaceTransforms) String() string {
	return common.FlagsToString(t, surfaceTransformsToString)
}

type CompositeAlphaModes int32

const (
	AlphaModeOpaque         CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR
	AlphaModePreMultiplied  CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR
	AlphaModePostMultiplied CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR
	AlphaModeInherit        CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR
)

var compositeAlphaModesToString = map[CompositeAlphaModes]string{
	AlphaModeOpaque:         "Opaque",
	AlphaModePreMultiplied:  "Pre-Multiplied",
	AlphaModePostMultiplied: "Post-Multiplied",
	AlphaModeInherit:        "Inherited",
}

func (m CompositeAlphaModes) String() string {
	return common.FlagsToString(m, compositeAlphaModesToString)
}

type PresentMode int32

const (
	PresentImmediate               PresentMode = C.VK_PRESENT_MODE_IMMEDIATE_KHR
	PresentMailbox                 PresentMode = C.VK_PRESENT_MODE_MAILBOX_KHR
	PresentFIFO                    PresentMode = C.VK_PRESENT_MODE_FIFO_KHR
	PresentFIFORelaxed             PresentMode = C.VK_PRESENT_MODE_FIFO_RELAXED_KHR
	PresentSharedDemandRefresh     PresentMode = C.VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR
	PresentSharedContinuousRefresh PresentMode = C.VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR
)

var presentModeToString = map[PresentMode]string{
	PresentImmediate:               "Immediate",
	PresentMailbox:                 "Mailbox",
	PresentFIFO:                    "FIFO",
	PresentFIFORelaxed:             "FIFO Relaxed",
	PresentSharedDemandRefresh:     "Shared: Demand Refresh",
	PresentSharedContinuousRefresh: "Shared: Continuous Refresh",
}

func (m PresentMode) String() string {
	return presentModeToString[m]
}

type Capabilities struct {
	MinImageCount uint32
	MaxImageCount uint32

	CurrentExtent  common.Extent2D
	MinImageExtent common.Extent2D
	MaxImageExtent common.Extent2D

	MaxImageArrayLayers uint32
	SupportedTransforms SurfaceTransforms
	CurrentTransform    SurfaceTransforms

	SupportedCompositeAlpha CompositeAlphaModes
	SupportedImageUsage     common.ImageUsages
}
