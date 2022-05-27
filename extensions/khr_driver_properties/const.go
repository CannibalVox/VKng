package khr_driver_properties

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type DriverID int32

var driverIDMapping = make(map[DriverID]string)

func (e DriverID) Register(str string) {
	driverIDMapping[e] = str
}

func (e DriverID) String() string {
	return driverIDMapping[e]
}

////

const (
	ExtensionName     string = C.VK_KHR_DRIVER_PROPERTIES_EXTENSION_NAME
	MaxDriverInfoSize int    = C.VK_MAX_DRIVER_INFO_SIZE_KHR
	MaxDriverNameSize int    = C.VK_MAX_DRIVER_NAME_SIZE_KHR

	DriverIDAmdOpenSource           DriverID = C.VK_DRIVER_ID_AMD_OPEN_SOURCE_KHR
	DriverIDAmdProprietary          DriverID = C.VK_DRIVER_ID_AMD_PROPRIETARY_KHR
	DriverIDArmProprietary          DriverID = C.VK_DRIVER_ID_ARM_PROPRIETARY_KHR
	DriverIDBroadcomProprietary     DriverID = C.VK_DRIVER_ID_BROADCOM_PROPRIETARY_KHR
	DriverIDGgpProprietary          DriverID = C.VK_DRIVER_ID_GGP_PROPRIETARY_KHR
	DriverIDGoogleSwiftshader       DriverID = C.VK_DRIVER_ID_GOOGLE_SWIFTSHADER_KHR
	DriverIDImaginationProprietary  DriverID = C.VK_DRIVER_ID_IMAGINATION_PROPRIETARY_KHR
	DriverIDIntelOpenSourceMesa     DriverID = C.VK_DRIVER_ID_INTEL_OPEN_SOURCE_MESA_KHR
	DriverIDIntelProprietaryWindows DriverID = C.VK_DRIVER_ID_INTEL_PROPRIETARY_WINDOWS_KHR
	DriverIDMesaRadV                DriverID = C.VK_DRIVER_ID_MESA_RADV_KHR
	DriverIDNvidiaProprietary       DriverID = C.VK_DRIVER_ID_NVIDIA_PROPRIETARY_KHR
	DriverIDQualcommProprietary     DriverID = C.VK_DRIVER_ID_QUALCOMM_PROPRIETARY_KHR
)
