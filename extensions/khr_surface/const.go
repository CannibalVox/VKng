package khr_surface

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

type SurfaceTransforms int32

var surfaceTransformsMapping = common.NewFlagStringMapping[SurfaceTransforms]()

func (f SurfaceTransforms) Register(str string) {
	surfaceTransformsMapping.Register(f, str)
}
func (f SurfaceTransforms) String() string {
	return surfaceTransformsMapping.FlagsToString(f)
}

////

type CompositeAlphaModes int32

var compositeAlphaModeMapping = make(map[CompositeAlphaModes]string)

func (e CompositeAlphaModes) Register(str string) {
	compositeAlphaModeMapping[e] = str
}

func (e CompositeAlphaModes) String() string {
	return compositeAlphaModeMapping[e]
}

////

type PresentMode int32

var presentModeMapping = make(map[PresentMode]string)

func (e PresentMode) Register(str string) {
	presentModeMapping[e] = str
}

func (e PresentMode) String() string {
	return presentModeMapping[e]
}

////

type ColorSpace int32

var colorSpaceMapping = make(map[ColorSpace]string)

func (e ColorSpace) Register(str string) {
	colorSpaceMapping[e] = str
}

func (e ColorSpace) String() string {
	return colorSpaceMapping[e]
}

////

const (
	ExtensionName string = C.VK_KHR_SURFACE_EXTENSION_NAME

	ObjectTypeSurface core1_0.ObjectType = C.VK_OBJECT_TYPE_SURFACE_KHR

	TransformIdentity                  SurfaceTransforms = C.VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR
	TransformRotate90                  SurfaceTransforms = C.VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR
	TransformRotate180                 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_ROTATE_180_BIT_KHR
	TransformRotate270                 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR
	TransformHorizontalMirror          SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_BIT_KHR
	TransformHorizontalMirrorRotate90  SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_90_BIT_KHR
	TransformHorizontalMirrorRotate180 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_180_BIT_KHR
	TransformHorizontalMirrorRotate270 SurfaceTransforms = C.VK_SURFACE_TRANSFORM_HORIZONTAL_MIRROR_ROTATE_270_BIT_KHR
	TransformInherit                   SurfaceTransforms = C.VK_SURFACE_TRANSFORM_INHERIT_BIT_KHR

	CompositeAlphaModeOpaque         CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR
	CompositeAlphaModePreMultiplied  CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR
	CompositeAlphaModePostMultiplied CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR
	CompositeAlphaModeInherit        CompositeAlphaModes = C.VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR

	PresentImmediate   PresentMode = C.VK_PRESENT_MODE_IMMEDIATE_KHR
	PresentMailbox     PresentMode = C.VK_PRESENT_MODE_MAILBOX_KHR
	PresentFIFO        PresentMode = C.VK_PRESENT_MODE_FIFO_KHR
	PresentFIFORelaxed PresentMode = C.VK_PRESENT_MODE_FIFO_RELAXED_KHR

	ColorSpaceSRGBNonlinear ColorSpace = C.VK_COLOR_SPACE_SRGB_NONLINEAR_KHR

	VKErrorSurfaceLost       common.VkResult = C.VK_ERROR_SURFACE_LOST_KHR
	VKErrorNativeWindowInUse common.VkResult = C.VK_ERROR_NATIVE_WINDOW_IN_USE_KHR
)

func init() {
	ObjectTypeSurface.Register("Surface")

	TransformIdentity.Register("Identity")
	TransformRotate90.Register("Rotate 90")
	TransformRotate180.Register("Rotate 180")
	TransformRotate270.Register("Rotate 270")
	TransformHorizontalMirror.Register("Horizontal Mirror")
	TransformHorizontalMirrorRotate90.Register("Horizontal Mirror & Rotate 90")
	TransformHorizontalMirrorRotate180.Register("Horizontal Mirror & Rotate 180")
	TransformHorizontalMirrorRotate270.Register("Horizontal Mirror & Rotate 270")
	TransformInherit.Register("Inherit")

	CompositeAlphaModeOpaque.Register("Opaque")
	CompositeAlphaModePreMultiplied.Register("Pre-Multiplied")
	CompositeAlphaModePostMultiplied.Register("Post-Multiplied")
	CompositeAlphaModeInherit.Register("Inherited")

	PresentImmediate.Register("Immediate")
	PresentMailbox.Register("Mailbox")
	PresentFIFO.Register("FIFO")
	PresentFIFORelaxed.Register("FIFO Relaxed")

	ColorSpaceSRGBNonlinear.Register("sRGB Non-Linear")

	VKErrorSurfaceLost.Register("surface lost")
	VKErrorNativeWindowInUse.Register("native window in use")
}
