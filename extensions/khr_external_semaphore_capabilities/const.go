package khr_external_semaphore_capabilities

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type ExternalSemaphoreFeatureFlags int32

var externalSemaphoreFeaturesMapping = common.NewFlagStringMapping[ExternalSemaphoreFeatureFlags]()

func (f ExternalSemaphoreFeatureFlags) Register(str string) {
	externalSemaphoreFeaturesMapping.Register(f, str)
}

func (f ExternalSemaphoreFeatureFlags) String() string {
	return externalSemaphoreFeaturesMapping.FlagsToString(f)
}

////

type ExternalSemaphoreHandleTypeFlags int32

var externalSemaphoreHandleTypesMapping = common.NewFlagStringMapping[ExternalSemaphoreHandleTypeFlags]()

func (f ExternalSemaphoreHandleTypeFlags) Register(str string) {
	externalSemaphoreHandleTypesMapping.Register(f, str)
}

func (f ExternalSemaphoreHandleTypeFlags) String() string {
	return externalSemaphoreHandleTypesMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_EXTENSION_NAME

	LUIDSize int = C.VK_LUID_SIZE_KHR

	ExternalSemaphoreFeatureExportable ExternalSemaphoreFeatureFlags = C.VK_EXTERNAL_SEMAPHORE_FEATURE_EXPORTABLE_BIT_KHR
	ExternalSemaphoreFeatureImportable ExternalSemaphoreFeatureFlags = C.VK_EXTERNAL_SEMAPHORE_FEATURE_IMPORTABLE_BIT_KHR

	ExternalSemaphoreHandleTypeOpaqueFD       ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_FD_BIT_KHR
	ExternalSemaphoreHandleTypeOpaqueWin32    ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR
	ExternalSemaphoreHandleTypeOpaqueWin32KMT ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
	ExternalSemaphoreHandleTypeD3D12Fence     ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT_KHR
	ExternalSemaphoreHandleTypeSyncFD         ExternalSemaphoreHandleTypeFlags = C.VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_SYNC_FD_BIT_KHR
)

func init() {
	ExternalSemaphoreFeatureExportable.Register("Exportable")
	ExternalSemaphoreFeatureImportable.Register("Importable")

	ExternalSemaphoreHandleTypeOpaqueFD.Register("Opaque File Descriptor")
	ExternalSemaphoreHandleTypeOpaqueWin32.Register("Opaque Win32 Handle")
	ExternalSemaphoreHandleTypeOpaqueWin32KMT.Register("Opaque Win32 Handle (Kernel Mode)")
	ExternalSemaphoreHandleTypeD3D12Fence.Register("D3D Fence")
	ExternalSemaphoreHandleTypeSyncFD.Register("Sync File Descriptor")
}
