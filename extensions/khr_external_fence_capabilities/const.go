package khr_external_fence_capabilities

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type ExternalFenceFeatureFlags int32

var externalFenceFeaturesMapping = common.NewFlagStringMapping[ExternalFenceFeatureFlags]()

func (f ExternalFenceFeatureFlags) Register(str string) {
	externalFenceFeaturesMapping.Register(f, str)
}

func (f ExternalFenceFeatureFlags) String() string {
	return externalFenceFeaturesMapping.FlagsToString(f)
}

////

type ExternalFenceHandleTypeFlags int32

var externalFenceHandleTypesMapping = common.NewFlagStringMapping[ExternalFenceHandleTypeFlags]()

func (f ExternalFenceHandleTypeFlags) Register(str string) {
	externalFenceHandleTypesMapping.Register(f, str)
}

func (f ExternalFenceHandleTypeFlags) String() string {
	return externalFenceHandleTypesMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_EXTERNAL_FENCE_CAPABILITIES_EXTENSION_NAME

	LUIDSize int = C.VK_LUID_SIZE_KHR

	ExternalFenceFeatureExportable ExternalFenceFeatureFlags = C.VK_EXTERNAL_FENCE_FEATURE_EXPORTABLE_BIT_KHR
	ExternalFenceFeatureImportable ExternalFenceFeatureFlags = C.VK_EXTERNAL_FENCE_FEATURE_IMPORTABLE_BIT_KHR

	ExternalFenceHandleTypeOpaqueFD       ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR
	ExternalFenceHandleTypeOpaqueWin32    ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR
	ExternalFenceHandleTypeOpaqueWin32KMT ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
	ExternalFenceHandleTypeSyncFD         ExternalFenceHandleTypeFlags = C.VK_EXTERNAL_FENCE_HANDLE_TYPE_SYNC_FD_BIT_KHR
)

func init() {
	ExternalFenceFeatureExportable.Register("Exportable")
	ExternalFenceFeatureImportable.Register("Importable")

	ExternalFenceHandleTypeOpaqueFD.Register("Opaque File Descriptor")
	ExternalFenceHandleTypeOpaqueWin32.Register("Opaque Win32")
	ExternalFenceHandleTypeOpaqueWin32KMT.Register("Opaque Win32 Kernel-Mode")
	ExternalFenceHandleTypeSyncFD.Register("Sync File Descriptor")
}
