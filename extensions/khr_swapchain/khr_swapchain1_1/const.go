package khr_swapchain1_1

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

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
	DeviceGroupPresentModeLocal            DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR
	DeviceGroupPresentModeRemote           DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR
	DeviceGroupPresentModeSum              DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR
	DeviceGroupPresentModeLocalMultiDevice DeviceGroupPresentModeFlags = C.VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR
)

func init() {
	DeviceGroupPresentModeLocal.Register("Local")
	DeviceGroupPresentModeRemote.Register("Remote")
	DeviceGroupPresentModeSum.Register("Sum")
	DeviceGroupPresentModeLocalMultiDevice.Register("Local Multi-Device")
}
