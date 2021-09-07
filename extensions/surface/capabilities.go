package ext_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"strings"
)

type SurfaceTransforms int32

const (
	Identity                  SurfaceTransforms = C.VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR
	Rotate90                  SurfaceTransforms = C.VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR
	Rotate180                 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR
	Rotate270                 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR
	HorizontalMirror          SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR
	HorizontalMirrorRotate90  SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR
	HorizontalMirrorRotate180 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR
	HorizontalMirrorRotate270 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR
	InheritTransform          SurfaceTransforms = C.VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR
	AllTransforms             SurfaceTransforms = Identity | Rotate90 | Rotate180 | Rotate270 | HorizontalMirror |
		HorizontalMirrorRotate90 | HorizontalMirrorRotate180 |
		HorizontalMirrorRotate270 | InheritTransform
)

var surfaceTransformsToString = map[SurfaceTransforms]string{
	Identity:                  "Identity",
	Rotate90:                  "Rotate 90",
	Rotate180:                 "Rotate 180",
	Rotate270:                 "Rotate 270",
	HorizontalMirror:          "Horizontal Mirror",
	HorizontalMirrorRotate90:  "Horizontal Mirror & Rotate 90",
	HorizontalMirrorRotate180: "Horizontal Mirror & Rotate 180",
	HorizontalMirrorRotate270: "Horizontal Mirror & Rotate 270",
	InheritTransform:          "Inherit",
}

func (t SurfaceTransforms) String() string {
	hasOne := false
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		shiftedBit := SurfaceTransforms(1 << i)
		if t&shiftedBit != 0 {
			strVal, exists := surfaceTransformsToString[shiftedBit]
			if exists {
				if hasOne {
					sb.WriteString("|")
				}
				sb.WriteString(strVal)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type CompositeAlphaModes int32

const (
	Opaque         CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR
	PreMultiplied  CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR
	PostMultiplied CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR
	InheritAlpha   CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR
	AllAlphaModes  CompositeAlphaModes = Opaque | PreMultiplied | PostMultiplied | InheritAlpha
)

var compositeAlphaModesToString = map[CompositeAlphaModes]string{
	Opaque:         "Opaque",
	PreMultiplied:  "Pre-Multiplied",
	PostMultiplied: "Post-Multiplied",
	InheritAlpha:   "Inherited",
}

func (m CompositeAlphaModes) String() string {
	hasOne := false
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		shiftedBit := CompositeAlphaModes(1 << i)
		if m&shiftedBit != 0 {
			strVal, exists := compositeAlphaModesToString[shiftedBit]
			if exists {
				if hasOne {
					sb.WriteString("|")
				}
				sb.WriteString(strVal)
				hasOne = true
			}
		}
	}

	return sb.String()
}

type PresentMode int32

const (
	Immediate               PresentMode = C.VK_PRESENT_MODE_IMMEDIATE_KHR
	Mailbox                 PresentMode = C.VK_PRESENT_MODE_MAILBOX_KHR
	FIFO                    PresentMode = C.VK_PRESENT_MODE_FIFO_KHR
	FIFORelaxed             PresentMode = C.VK_PRESENT_MODE_FIFO_RELAXED_KHR
	SharedDemandRefresh     PresentMode = C.VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR
	SharedContinuousRefresh PresentMode = C.VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR
)

var presentModeToString = map[PresentMode]string{
	Immediate:               "Immediate",
	Mailbox:                 "Mailbox",
	FIFO:                    "FIFO",
	FIFORelaxed:             "FIFO Relaxed",
	SharedDemandRefresh:     "Shared: Demand Refresh",
	SharedContinuousRefresh: "Shared: Continuous Refresh",
}

func (m PresentMode) String() string {
	return presentModeToString[m]
}

type Capabilities struct {
	MinImageCount uint32
	MaxImageCount uint32

	CurrentExtent  core.Extent2D
	MinImageExtent core.Extent2D
	MaxImageExtent core.Extent2D

	MaxImageArrayLayers uint32
	SupportedTransforms SurfaceTransforms
	CurrentTransform    SurfaceTransforms

	SupportedCompositeAlpha CompositeAlphaModes
	SupportedImageUsage     core.ImageUsages
}
