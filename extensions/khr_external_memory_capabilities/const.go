package khr_external_memory_capabilities

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import "github.com/CannibalVox/VKng/core/common"

type ExternalMemoryFeatures int32

var externalMemoryFeaturesMapping = common.NewFlagStringMapping[ExternalMemoryFeatures]()

func (f ExternalMemoryFeatures) Register(str string) {
	externalMemoryFeaturesMapping.Register(f, str)
}

func (f ExternalMemoryFeatures) String() string {
	return externalMemoryFeaturesMapping.FlagsToString(f)
}

////

type ExternalMemoryHandleTypes int32

var externalMemoryHandleTypesMapping = common.NewFlagStringMapping[ExternalMemoryHandleTypes]()

func (f ExternalMemoryHandleTypes) Register(str string) {
	externalMemoryHandleTypesMapping.Register(f, str)
}

func (f ExternalMemoryHandleTypes) String() string {
	return externalMemoryHandleTypesMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME

	LUIDSize int = C.VK_LUID_SIZE_KHR

	ExternalMemoryFeatureDedicatedOnly ExternalMemoryFeatures = C.VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_KHR
	ExternalMemoryFeatureExportable    ExternalMemoryFeatures = C.VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_KHR
	ExternalMemoryFeatureImportable    ExternalMemoryFeatures = C.VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_KHR

	ExternalMemoryHandleTypeD3D11Texture    ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT_KHR
	ExternalMemoryHandleTypeD3D11TextureKMT ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT_KHR
	ExternalMemoryHandleTypeD3D12Heap       ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT_KHR
	ExternalMemoryHandleTypeD3D12Resource   ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT_KHR
	ExternalMemoryHandleTypeOpaqueFD        ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT_KHR
	ExternalMemoryHandleTypeOpaqueWin32     ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR
	ExternalMemoryHandleTypeOpaqueWin32KMT  ExternalMemoryHandleTypes = C.VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT_KHR
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
