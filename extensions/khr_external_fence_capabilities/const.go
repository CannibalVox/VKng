package khr_external_fence_capabilities

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type ExternalFenceFeatures int32

var externalFenceFeaturesMapping = common.NewFlagStringMapping[ExternalFenceFeatures]()

func (f ExternalFenceFeatures) Register(str string) {
	externalFenceFeaturesMapping.Register(f, str)
}

func (f ExternalFenceFeatures) String() string {
	return externalFenceFeaturesMapping.FlagsToString(f)
}

////

type ExternalFenceHandleTypes int32

var externalFenceHandleTypesMapping = common.NewFlagStringMapping[ExternalFenceHandleTypes]()

func (f ExternalFenceHandleTypes) Register(str string) {
	externalFenceHandleTypesMapping.Register(f, str)
}

func (f ExternalFenceHandleTypes) String() string {
	return externalFenceHandleTypesMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_EXTERNAL_FENCE_CAPABILITIES_EXTENSION_NAME

	LUIDSize int = C.VK_LUID_SIZE_KHR

	ExternalFenceFeatureExportable ExternalFenceFeatures = C.VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT_KHR
	ExternalFenceFeatureImportable ExternalFenceFeatures = C.VK_EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT_KHR

	ExternalFenceHandleTypeOpaqueFD       ExternalFenceHandleTypes = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR
	ExternalFenceHandleTypeOpaqueWin32    ExternalFenceHandleTypes = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR
	ExternalFenceHandleTypeOpaqueWin32KMT ExternalFenceHandleTypes = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
	ExternalFenceHandleTypeSyncFD         ExternalFenceHandleTypes = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT_KHR
)

func init() {
	ExternalFenceFeatureExportable.Register("Exportable")
	ExternalFenceFeatureImportable.Register("Importable")

	ExternalFenceHandleTypeOpaqueFD.Register("Opaque File Descriptor")
	ExternalFenceHandleTypeOpaqueWin32.Register("Opaque Win32")
	ExternalFenceHandleTypeOpaqueWin32KMT.Register("Opaque Win32 Kernel-Mode")
	ExternalFenceHandleTypeSyncFD.Register("Sync File Descriptor")
}
