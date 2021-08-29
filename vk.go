package VKng

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

func AvailableExtensions(alloc cgoalloc.Allocator) (map[string]*core.ExtensionProperties, error) {
	extensionCount := (*C.uint32_t)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer alloc.Free(unsafe.Pointer(extensionCount))

	res := C.vkEnumerateInstanceExtensionProperties(nil, extensionCount, nil)
	err := core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	if *extensionCount == 0 {
		return nil, nil
	}

	extensions := (*C.VkExtensionProperties)(alloc.Malloc(int(*extensionCount) * int(unsafe.Sizeof(C.VkExtensionProperties{}))))
	defer alloc.Free(unsafe.Pointer(extensions))

	res = C.vkEnumerateInstanceExtensionProperties(nil, extensionCount, extensions)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	intExtensionCount := int(*extensionCount)
	extensionArray := ([]C.VkExtensionProperties)(unsafe.Slice(extensions, intExtensionCount))
	outExtensions := make(map[string]*core.ExtensionProperties)
	for i := 0; i < intExtensionCount; i++ {
		extension := extensionArray[i]

		outExtension := &core.ExtensionProperties{
			ExtensionName: C.GoString((*C.char)(&extension.extensionName[0])),
			SpecVersion:   core.Version(extension.specVersion),
		}

		existingExtension, ok := outExtensions[outExtension.ExtensionName]
		if ok && existingExtension.SpecVersion >= outExtension.SpecVersion {
			continue
		}
		outExtensions[outExtension.ExtensionName] = outExtension
	}

	return outExtensions, nil
}

func AvailableLayers(alloc cgoalloc.Allocator) (map[string]*core.LayerProperties, error) {
	layerCount := (*C.uint32_t)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer alloc.Free(unsafe.Pointer(layerCount))

	res := C.vkEnumerateInstanceLayerProperties(layerCount, nil)
	err := core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	if *layerCount == 0 {
		return nil, nil
	}

	layers := (*C.VkLayerProperties)(alloc.Malloc(int(*layerCount) * int(unsafe.Sizeof(C.VkLayerProperties{}))))
	defer alloc.Free(unsafe.Pointer(layers))

	res = C.vkEnumerateInstanceLayerProperties(layerCount, layers)
	err = core.Result(res).ToError()
	if err != nil {
		return nil, err
	}

	intLayerCount := int(*layerCount)
	layerArray := ([]C.VkLayerProperties)(unsafe.Slice(layers, intLayerCount))
	outLayers := make(map[string]*core.LayerProperties)
	for i := 0; i < intLayerCount; i++ {
		layer := layerArray[i]

		outLayer := &core.LayerProperties{
			LayerName:             C.GoString((*C.char)(&layer.layerName[0])),
			SpecVersion:           core.Version(layer.specVersion),
			ImplementationVersion: core.Version(layer.implementationVersion),
			Description:           C.GoString((*C.char)(&layer.description[0])),
		}

		existingLayer, ok := outLayers[outLayer.LayerName]
		if ok && existingLayer.SpecVersion >= outLayer.SpecVersion {
			continue
		}
		outLayers[outLayer.LayerName] = outLayer
	}

	return outLayers, nil
}
