package khr_device_group

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/extensions/khr_swapchain"
)

type MemoryAllocateFlags int32

var memoryAllocateFlagsMapping = common.NewFlagStringMapping[MemoryAllocateFlags]()

func (f MemoryAllocateFlags) Register(str string) {
	memoryAllocateFlagsMapping.Register(f, str)
}

func (f MemoryAllocateFlags) String() string {
	return memoryAllocateFlagsMapping.FlagsToString(f)
}

////

type PeerMemoryFeatures int32

var peerMemoryFeaturesMapping = common.NewFlagStringMapping[PeerMemoryFeatures]()

func (f PeerMemoryFeatures) Register(str string) {
	peerMemoryFeaturesMapping.Register(f, str)
}
func (f PeerMemoryFeatures) String() string {
	return peerMemoryFeaturesMapping.FlagsToString(f)
}

///

type DeviceGroupPresentModeFlags int32

var deviceGroupPresentModeFlagsMapping = common.NewFlagStringMapping[DeviceGroupPresentModeFlags]()

func (f DeviceGroupPresentModeFlags) Register(str string) {
	deviceGroupPresentModeFlagsMapping.Register(f, str)
}

func (f DeviceGroupPresentModeFlags) String() string {
	return deviceGroupPresentModeFlagsMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_DEVICE_GROUP_EXTENSION_NAME

	DependencyDeviceGroup common.DependencyFlags = C.VK_DEPENDENCY_DEVICE_GROUP_BIT_KHR

	MemoryAllocateDeviceMask MemoryAllocateFlags = C.VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT_KHR

	DeviceGroupPresentModeLocal            DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR
	DeviceGroupPresentModeRemote           DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR
	DeviceGroupPresentModeSum              DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR
	DeviceGroupPresentModeLocalMultiDevice DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR

	PeerMemoryFeatureCopyDst    PeerMemoryFeatures = C.VK_PEER_MEMORY_FEATURE_COPY_DST_BIT_KHR
	PeerMemoryFeatureCopySrc    PeerMemoryFeatures = C.VK_PEER_MEMORY_FEATURE_COPY_SRC_BIT_KHR
	PeerMemoryFeatureGenericDst PeerMemoryFeatures = C.VK_PEER_MEMORY_FEATURE_GENERIC_DST_BIT_KHR
	PeerMemoryFeatureGenericSrc PeerMemoryFeatures = C.VK_PEER_MEMORY_FEATURE_GENERIC_SRC_BIT_KHR

	PipelineCreateDispatchBase             common.PipelineCreateFlags = C.VK_PIPELINE_CREATE_DISPATCH_BASE_KHR
	PipelineCreateViewIndexFromDeviceIndex common.PipelineCreateFlags = C.VK_PIPELINE_CREATE_VIEW_INDEX_FROM_DEVICE_INDEX_BIT_KHR

	ImageCreateSplitInstanceBindRegions common.ImageCreateFlags = C.VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR

	SwapchainCreateSplitInstanceBindRegions khr_swapchain.SwapchainCreateFlags = C.VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR
)

func init() {
	DependencyDeviceGroup.Register("Device Group")

	MemoryAllocateDeviceMask.Register("Device Mask")

	DeviceGroupPresentModeLocal.Register("Local")
	DeviceGroupPresentModeRemote.Register("Remote")
	DeviceGroupPresentModeSum.Register("Sum")
	DeviceGroupPresentModeLocalMultiDevice.Register("Local Multi-Device")

	PeerMemoryFeatureCopyDst.Register("Copy Dst")
	PeerMemoryFeatureCopySrc.Register("Copy Src")
	PeerMemoryFeatureGenericDst.Register("Generic Dst")
	PeerMemoryFeatureGenericSrc.Register("Generic Src")

	PipelineCreateDispatchBase.Register("Dispatch Base")
	PipelineCreateViewIndexFromDeviceIndex.Register("View Index From Device Index")

	ImageCreateSplitInstanceBindRegions.Register("Split Instance Bind Regions")
	SwapchainCreateSplitInstanceBindRegions.Register("Split Instance Bind Regions")
}
