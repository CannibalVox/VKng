package khr_external_memory_capabilities

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type ExternalMemoryFeatureFlags int32

var externalMemoryFeaturesMapping = common.NewFlagStringMapping[ExternalMemoryFeatureFlags]()

func (f ExternalMemoryFeatureFlags) Register(str string) {
	externalMemoryFeaturesMapping.Register(f, str)
}

func (f ExternalMemoryFeatureFlags) String() string {
	return externalMemoryFeaturesMapping.FlagsToString(f)
}

////

type ExternalMemoryHandleTypeFlags int32

var externalMemoryHandleTypesMapping = common.NewFlagStringMapping[ExternalMemoryHandleTypeFlags]()

func (f ExternalMemoryHandleTypeFlags) Register(str string) {
	externalMemoryHandleTypesMapping.Register(f, str)
}

func (f ExternalMemoryHandleTypeFlags) String() string {
	return externalMemoryHandleTypesMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME

	LUIDSize int = C.VK_LUID_SIZE_KHR

	ExternalMemoryFeatureDedicatedOnly ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_KHR
	ExternalMemoryFeatureExportable    ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_KHR
	ExternalMemoryFeatureImportable    ExternalMemoryFeatureFlags = C.VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_KHR

	ExternalMemoryHandleTypeD3D11Texture    ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR
	ExternalMemoryHandleTypeD3D11TextureKMT ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR
	ExternalMemoryHandleTypeD3D12Heap       ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR
	ExternalMemoryHandleTypeD3D12Resource   ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT_KHR
	ExternalMemoryHandleTypeOpaqueFD        ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT_KHR
	ExternalMemoryHandleTypeOpaqueWin32     ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR
	ExternalMemoryHandleTypeOpaqueWin32KMT  ExternalMemoryHandleTypeFlags = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
)

func init() {
	ExternalMemoryFeatureDedicatedOnly.Register("Dedicated Only")
	ExternalMemoryFeatureExportable.Register("Exportable")
	ExternalMemoryFeatureImportable.Register("Importable")

	ExternalMemoryHandleTypeD3D11Texture.Register("D3D11 Texture")
	ExternalMemoryHandleTypeD3D11TextureKMT.Register("D3D11 Texture (Kernel-Mode)")
	ExternalMemoryHandleTypeD3D12Heap.Register("D3D12 Heap")
	ExternalMemoryHandleTypeD3D12Resource.Register("D3D12 Resource")
	ExternalMemoryHandleTypeOpaqueFD.Register("Opaque File Descriptor")
	ExternalMemoryHandleTypeOpaqueWin32.Register("Opaque Win32")
	ExternalMemoryHandleTypeOpaqueWin32KMT.Register("Opaque Win32 (Kernel-Mode)")
}
