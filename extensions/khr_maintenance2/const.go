package khr_maintenance2

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/core1_0"
)

type PointClippingBehavior int32

var pointClippingBehaviorMapping = make(map[PointClippingBehavior]string)

func (e PointClippingBehavior) Register(str string) {
	pointClippingBehaviorMapping[e] = str
}

func (e PointClippingBehavior) String() string {
	return pointClippingBehaviorMapping[e]
}

////

type TessellationDomainOrigin int32

var tessellationDomainOriginMapping = make(map[TessellationDomainOrigin]string)

func (e TessellationDomainOrigin) Register(str string) {
	tessellationDomainOriginMapping[e] = str
}

func (e TessellationDomainOrigin) String() string {
	return tessellationDomainOriginMapping[e]
}

////

const (
	ExtensionName string = C.VK_KHR_MAINTENANCE2_EXTENSION_NAME

	ImageCreateBlockTexelViewCompatible core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT_KHR
	ImageCreateExtendedUsage            core1_0.ImageCreateFlags = C.VK_IMAGE_CREATE_EXTENDED_USAGE_BIT_KHR

	ImageLayoutDepthAttachmentStencilReadOnlyOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL_KHR
	ImageLayoutDepthReadOnlyStencilAttachmentOptimal core1_0.ImageLayout = C.VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL_KHR

	PointClippingAllClipPlanes      PointClippingBehavior = C.VK_POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES_KHR
	PointClippingUserClipPlanesOnly PointClippingBehavior = C.VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY_KHR

	TessellationDomainOriginUpperLeft TessellationDomainOrigin = C.VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT_KHR
	TessellationDomainOriginLowerLeft TessellationDomainOrigin = C.VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT_KHR
)

func init() {
	ImageCreateBlockTexelViewCompatible.Register("Block Texel View Compatible")
	ImageCreateExtendedUsage.Register("Extended Usage")

	ImageLayoutDepthReadOnlyStencilAttachmentOptimal.Register("Depth Read-Only Stencil Attachment Optimal")
	ImageLayoutDepthAttachmentStencilReadOnlyOptimal.Register("Depth Attachment Stencil Read-Only Optimal")

	PointClippingAllClipPlanes.Register("All Clip Planes")
	PointClippingUserClipPlanesOnly.Register("User Clip Planes Only")

	TessellationDomainOriginUpperLeft.Register("Upper Left")
	TessellationDomainOriginLowerLeft.Register("Lower Left")
}
