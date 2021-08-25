package VKng

/*
#cgo windows LDFLAGS: -lvulkan
#cgo linux freebsd darwin openbsd pkg-config: vulkan
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/cgoalloc"
	"unsafe"
)

type VKExtensionProperties struct {
	ExtensionName string
	SpecVersion   Version
}

func AvailableExtensions(alloc cgoalloc.Allocator) (map[string]*VKExtensionProperties, error) {
	extensionCount := (*C.uint32_t)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer alloc.Free(unsafe.Pointer(extensionCount))

	res := C.vkEnumerateInstanceExtensionProperties(nil, extensionCount, nil)
	err := Result(res).ToError()
	if err != nil {
		return nil, err
	}

	if *extensionCount == 0 {
		return nil, nil
	}

	extensions := (*C.VkExtensionProperties)(alloc.Malloc(int(*extensionCount) * int(unsafe.Sizeof(C.VkExtensionProperties{}))))
	defer alloc.Free(unsafe.Pointer(extensions))

	res = C.vkEnumerateInstanceExtensionProperties(nil, extensionCount, extensions)
	err = Result(res).ToError()
	if err != nil {
		return nil, err
	}

	intExtensionCount := int(*extensionCount)
	extensionArray := ([]C.VkExtensionProperties)(unsafe.Slice(extensions, intExtensionCount))
	outExtensions := make(map[string]*VKExtensionProperties)
	for i := 0; i < intExtensionCount; i++ {
		extension := extensionArray[i]

		outExtension := &VKExtensionProperties{
			ExtensionName: C.GoString((*C.char)(&extension.extensionName[0])),
			SpecVersion:   Version(extension.specVersion),
		}

		existingExtension, ok := outExtensions[outExtension.ExtensionName]
		if ok && existingExtension.SpecVersion >= outExtension.SpecVersion {
			continue
		}
		outExtensions[outExtension.ExtensionName] = outExtension
	}

	return outExtensions, nil
}

type VKLayerProperties struct {
	LayerName             string
	SpecVersion           Version
	ImplementationVersion Version
	Description           string
}

func AvailableLayers(alloc cgoalloc.Allocator) (map[string]*VKLayerProperties, error) {
	layerCount := (*C.uint32_t)(alloc.Malloc(int(unsafe.Sizeof(C.uint32_t(0)))))
	defer alloc.Free(unsafe.Pointer(layerCount))

	res := C.vkEnumerateInstanceLayerProperties(layerCount, nil)
	err := Result(res).ToError()
	if err != nil {
		return nil, err
	}

	if *layerCount == 0 {
		return nil, nil
	}

	layers := (*C.VkLayerProperties)(alloc.Malloc(int(*layerCount) * int(unsafe.Sizeof(C.VkLayerProperties{}))))
	defer alloc.Free(unsafe.Pointer(layers))

	res = C.vkEnumerateInstanceLayerProperties(layerCount, layers)
	err = Result(res).ToError()
	if err != nil {
		return nil, err
	}

	intLayerCount := int(*layerCount)
	layerArray := ([]C.VkLayerProperties)(unsafe.Slice(layers, intLayerCount))
	outLayers := make(map[string]*VKLayerProperties)
	for i := 0; i < intLayerCount; i++ {
		layer := layerArray[i]

		outLayer := &VKLayerProperties{
			LayerName:             C.GoString((*C.char)(&layer.layerName[0])),
			SpecVersion:           Version(layer.specVersion),
			ImplementationVersion: Version(layer.implementationVersion),
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