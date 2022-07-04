package khr_get_physical_device_properties2_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/common/extensions"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanExtension_PhysicalDeviceFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {
			val := reflect.ValueOf(pFeatures).Elem()

			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			featureVal := val.FieldByName("features")
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("robustBufferAccess").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fullDrawIndexUint32").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("imageCubeArray").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("independentBlend").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("geometryShader").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("tessellationShader").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sampleRateShading").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("dualSrcBlend").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("logicOp").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("multiDrawIndirect").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("drawIndirectFirstInstance").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthClamp").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthBiasClamp").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fillModeNonSolid").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("depthBounds").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("wideLines").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("largePoints").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("alphaToOne").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("multiViewport").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("samplerAnisotropy").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionETC2").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionASTC_LDR").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("textureCompressionBC").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("occlusionQueryPrecise").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("pipelineStatisticsQuery").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("vertexPipelineStoresAndAtomics").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("fragmentStoresAndAtomics").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderTessellationAndGeometryPointSize").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderImageGatherExtended").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageExtendedFormats").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageMultisample").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageReadWithoutFormat").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageWriteWithoutFormat").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderUniformBufferArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderSampledImageArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageBufferArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderStorageImageArrayDynamicIndexing").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderClipDistance").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderCullDistance").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderFloat64").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderInt64").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderInt16").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderResourceResidency").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("shaderResourceMinLod").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseBinding").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyBuffer").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyImage2D").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyImage3D").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency2Samples").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency4Samples").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency8Samples").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidency16Samples").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("sparseResidencyAliased").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("variableMultisampleRate").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(featureVal.FieldByName("inheritedQueries").UnsafeAddr())) = driver.VkBool32(1)

		})

	outData := &khr_get_physical_device_properties2.DeviceFeatures{}
	err := extension.PhysicalDeviceFeatures2(physicalDevice, outData)
	require.NoError(t, err)

	features := outData.Features
	require.NotNil(t, features)
	require.True(t, features.RobustBufferAccess)
	require.False(t, features.FullDrawIndexUint32)
	require.True(t, features.ImageCubeArray)
	require.False(t, features.IndependentBlend)
	require.True(t, features.GeometryShader)
	require.False(t, features.TessellationShader)
	require.True(t, features.SampleRateShading)
	require.False(t, features.DualSrcBlend)
	require.True(t, features.LogicOp)
	require.False(t, features.MultiDrawIndirect)
	require.True(t, features.DrawIndirectFirstInstance)
	require.False(t, features.DepthClamp)
	require.True(t, features.DepthBiasClamp)
	require.False(t, features.FillModeNonSolid)
	require.True(t, features.DepthBounds)
	require.False(t, features.WideLines)
	require.True(t, features.LargePoints)
	require.False(t, features.AlphaToOne)
	require.True(t, features.MultiViewport)
	require.False(t, features.SamplerAnisotropy)
	require.True(t, features.TextureCompressionEtc2)
	require.False(t, features.TextureCompressionAstcLdc)
	require.True(t, features.TextureCompressionBc)
	require.False(t, features.OcclusionQueryPrecise)
	require.True(t, features.PipelineStatisticsQuery)
	require.False(t, features.VertexPipelineStoresAndAtomics)
	require.True(t, features.FragmentStoresAndAtomics)
	require.False(t, features.ShaderTessellationAndGeometryPointSize)
	require.True(t, features.ShaderImageGatherExtended)
	require.False(t, features.ShaderStorageImageExtendedFormats)
	require.True(t, features.ShaderStorageImageMultisample)
	require.False(t, features.ShaderStorageImageReadWithoutFormat)
	require.True(t, features.ShaderStorageImageWriteWithoutFormat)
	require.False(t, features.ShaderUniformBufferArrayDynamicIndexing)
	require.True(t, features.ShaderSampledImageArrayDynamicIndexing)
	require.False(t, features.ShaderStorageBufferArrayDynamicIndexing)
	require.True(t, features.ShaderStorageImageArrayDynamicIndexing)
	require.False(t, features.ShaderClipDistance)
	require.True(t, features.ShaderCullDistance)
	require.False(t, features.ShaderFloat64)
	require.True(t, features.ShaderInt64)
	require.False(t, features.ShaderInt16)
	require.True(t, features.ShaderResourceResidency)
	require.False(t, features.ShaderResourceMinLod)
	require.True(t, features.SparseBinding)
	require.False(t, features.SparseResidencyBuffer)
	require.True(t, features.SparseResidencyImage2D)
	require.False(t, features.SparseResidencyImage3D)
	require.True(t, features.SparseResidency2Samples)
	require.False(t, features.SparseResidency4Samples)
	require.True(t, features.SparseResidency8Samples)
	require.False(t, features.SparseResidency16Samples)
	require.True(t, features.SparseResidencyAliased)
	require.False(t, features.VariableMultisampleRate)
	require.True(t, features.InheritedQueries)
}

func TestVulkanDevice_CreateDeviceWithFeatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil).AnyTimes()
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := extensions.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_0, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(physicalDevice.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice, pCreateInfo *driver.VkDeviceCreateInfo, pAllocator *driver.VkAllocationCallbacks, pDevice *driver.VkDevice) (common.VkResult, error) {
			*pDevice = device.Handle()

			v := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(3), v.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO
			require.Equal(t, uint64(0), v.FieldByName("flags").Uint())
			require.Equal(t, uint64(2), v.FieldByName("queueCreateInfoCount").Uint())
			require.Equal(t, uint64(2), v.FieldByName("enabledExtensionCount").Uint())
			require.Equal(t, uint64(1), v.FieldByName("enabledLayerCount").Uint())
			require.True(t, v.FieldByName("pEnabledFeatures").IsNil())

			extensionNamePtr := (**driver.Char)(unsafe.Pointer(v.FieldByName("ppEnabledExtensionNames").Elem().UnsafeAddr()))
			extensionNameSlice := ([]*driver.Char)(unsafe.Slice(extensionNamePtr, 2))

			var extensionNames []string
			for _, extensionNameBytes := range extensionNameSlice {
				var extensionNameRunes []rune
				extensionNameByteSlice := ([]driver.Char)(unsafe.Slice(extensionNameBytes, 1<<30))
				for _, nameByte := range extensionNameByteSlice {
					if nameByte == 0 {
						break
					}

					extensionNameRunes = append(extensionNameRunes, rune(nameByte))
				}

				extensionNames = append(extensionNames, string(extensionNameRunes))
			}

			require.ElementsMatch(t, []string{"a", "b"}, extensionNames)

			layerNamePtr := (**driver.Char)(unsafe.Pointer(v.FieldByName("ppEnabledLayerNames").Elem().UnsafeAddr()))
			layerNameSlice := ([]*driver.Char)(unsafe.Slice(layerNamePtr, 1))

			var layerNames []string
			for _, layerNameBytes := range layerNameSlice {
				var layerNameRunes []rune
				layerNameByteSlice := ([]driver.Char)(unsafe.Slice(layerNameBytes, 1<<30))
				for _, nameByte := range layerNameByteSlice {
					if nameByte == 0 {
						break
					}

					layerNameRunes = append(layerNameRunes, rune(nameByte))
				}

				layerNames = append(layerNames, string(layerNameRunes))
			}

			require.ElementsMatch(t, []string{"c"}, layerNames)

			queueCreateInfoPtr := (*driver.VkDeviceQueueCreateInfo)(unsafe.Pointer(v.FieldByName("pQueueCreateInfos").Elem().UnsafeAddr()))
			queueCreateInfoSlice := ([]driver.VkDeviceQueueCreateInfo)(unsafe.Slice(queueCreateInfoPtr, 2))

			queueInfoV := reflect.ValueOf(queueCreateInfoSlice[0])
			require.Equal(t, uint64(2), queueInfoV.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
			require.True(t, queueInfoV.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), queueInfoV.FieldByName("flags").Uint())
			require.Equal(t, uint64(1), queueInfoV.FieldByName("queueFamilyIndex").Uint())
			require.Equal(t, uint64(3), queueInfoV.FieldByName("queueCount").Uint())

			priorityPtr := (*float32)(unsafe.Pointer(queueInfoV.FieldByName("pQueuePriorities").Elem().UnsafeAddr()))
			prioritySlice := ([]float32)(unsafe.Slice(priorityPtr, 3))
			require.Equal(t, float32(3), prioritySlice[0])
			require.Equal(t, float32(5), prioritySlice[1])
			require.Equal(t, float32(7), prioritySlice[2])

			queueInfoV = reflect.ValueOf(queueCreateInfoSlice[1])
			require.Equal(t, uint64(2), queueInfoV.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO
			require.True(t, queueInfoV.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), queueInfoV.FieldByName("flags").Uint())
			require.Equal(t, uint64(11), queueInfoV.FieldByName("queueFamilyIndex").Uint())
			require.Equal(t, uint64(1), queueInfoV.FieldByName("queueCount").Uint())

			priorityPtr = (*float32)(unsafe.Pointer(queueInfoV.FieldByName("pQueuePriorities").Elem().UnsafeAddr()))
			prioritySlice = ([]float32)(unsafe.Slice(priorityPtr, 1))
			require.Equal(t, float32(13), prioritySlice[0])

			nextPtr := (*khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR)(v.FieldByName("pNext").UnsafePointer())
			nextVal := reflect.ValueOf(nextPtr).Elem()
			require.Equal(t, uint64(1000059000), nextVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR
			require.True(t, nextVal.FieldByName("pNext").IsNil())

			features := nextVal.FieldByName("features")
			require.Equal(t, uint64(1), features.FieldByName("textureCompressionETC2").Uint())
			require.Equal(t, uint64(1), features.FieldByName("depthBounds").Uint())
			require.Equal(t, uint64(0), features.FieldByName("samplerAnisotropy").Uint())

			return core1_0.VKSuccess, nil
		})

	options := core1_0.DeviceCreateOptions{
		QueueFamilies: []core1_0.DeviceQueueCreateOptions{
			{
				QueueFamilyIndex:       1,
				CreatedQueuePriorities: []float32{3, 5, 7},
			},
			{
				QueueFamilyIndex:       11,
				CreatedQueuePriorities: []float32{13},
			},
		},
		ExtensionNames: []string{"a", "b"},
		LayerNames:     []string{"c"},
	}
	features := khr_get_physical_device_properties2.DeviceFeatures{
		Features: core1_0.PhysicalDeviceFeatures{
			TextureCompressionEtc2: true,
			DepthBounds:            true,
		},
	}
	options.Next = features

	actualDevice, _, err := physicalDevice.CreateDevice(nil, options)
	require.NoError(t, err)
	require.NotNil(t, actualDevice)
	require.Equal(t, device.Handle(), actualDevice.Handle())
}

func TestVulkanExtension_PhysicalDeviceFormatProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceFormatProperties2KHR(
		physicalDevice.Handle(),
		driver.VkFormat(64), // VK_FORMAT_A2B10G10R10_UNORM_PACK32
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		format driver.VkFormat,
		pFormatProperties *khr_get_physical_device_properties2_driver.VkFormatProperties2KHR) {

		val := reflect.ValueOf(pFormatProperties).Elem()
		require.Equal(t, uint64(1000059002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_FORMAT_PROPERTIES_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		properties := val.FieldByName("formatProperties")
		*(*uint32)(unsafe.Pointer(properties.FieldByName("optimalTilingFeatures").UnsafeAddr())) = uint32(0x00000100) // VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT
		*(*uint32)(unsafe.Pointer(properties.FieldByName("linearTilingFeatures").UnsafeAddr())) = uint32(0x00000400)  // VK_FORMAT_FEATURE_BLIT_SRC_BIT
		*(*uint32)(unsafe.Pointer(properties.FieldByName("bufferFeatures").UnsafeAddr())) = uint32(0x00000010)        // VK_FORMAT_FEATURE_STORAGE_TEXEL_BUFFER_BIT
	})

	outData := khr_get_physical_device_properties2.FormatPropertiesOutData{}
	err := extension.PhysicalDeviceFormatProperties2(physicalDevice,
		core1_0.DataFormatA2B10G10R10UnsignedNormalizedPacked,
		&outData)
	require.NoError(t, err)

	require.Equal(t, core1_0.FormatFeatureColorAttachmentBlend, outData.FormatProperties.OptimalTilingFeatures)
	require.Equal(t, core1_0.FormatFeatureBlitSource, outData.FormatProperties.LinearTilingFeatures)
	require.Equal(t, core1_0.FormatFeatureStorageTexelBuffer, outData.FormatProperties.BufferFeatures)

}

func TestVulkanExtension_PhysicalDeviceImageFormatProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceImageFormatProperties2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
			pImageFormatInfo *khr_get_physical_device_properties2_driver.VkPhysicalDeviceImageFormatInfo2KHR,
			pImageFormatProperties *khr_get_physical_device_properties2_driver.VkImageFormatProperties2KHR,
		) (common.VkResult, error) {
			optionVal := reflect.ValueOf(*pImageFormatInfo)

			require.Equal(t, uint64(1000059004), optionVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR
			require.True(t, optionVal.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(68), optionVal.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_UINT_PACK32
			require.Equal(t, uint64(1), optionVal.FieldByName("_type").Uint())          // VK_IMAGE_TYPE_2D
			require.Equal(t, uint64(0), optionVal.FieldByName("tiling").Uint())         // VK_IMAGE_TILING_OPTIMAL
			require.Equal(t, uint64(8), optionVal.FieldByName("usage").Uint())          // VK_IMAGE_USAGE_STORAGE_BIT
			require.Equal(t, uint64(0x00000010), optionVal.FieldByName("flags").Uint()) // VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT

			outDataVal := reflect.ValueOf(pImageFormatProperties).Elem()
			require.Equal(t, uint64(1000059003), outDataVal.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR
			require.True(t, outDataVal.FieldByName("pNext").IsNil())

			formatPropertiesVal := outDataVal.FieldByName("imageFormatProperties")

			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxMipLevels").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxArrayLayers").UnsafeAddr())) = uint32(3)
			*(*uint64)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxResourceSize").UnsafeAddr())) = uint64(5)
			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("sampleCounts").UnsafeAddr())) = uint32(core1_0.Samples8)
			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxExtent").FieldByName("width").UnsafeAddr())) = uint32(11)
			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxExtent").FieldByName("height").UnsafeAddr())) = uint32(13)
			*(*uint32)(unsafe.Pointer(formatPropertiesVal.FieldByName("maxExtent").FieldByName("depth").UnsafeAddr())) = uint32(17)

			return core1_0.VKSuccess, nil
		})

	outData := khr_get_physical_device_properties2.ImageFormatPropertiesOutData{}
	_, err := extension.PhysicalDeviceImageFormatProperties2(physicalDevice, khr_get_physical_device_properties2.ImageFormatOptions{
		Format: core1_0.DataFormatA2B10G10R10UnsignedIntPacked,
		Type:   core1_0.ImageType2D,
		Tiling: core1_0.ImageTilingOptimal,
		Usage:  core1_0.ImageUsageStorage,
		Flags:  core1_0.ImageCreateCubeCompatible,
	}, &outData)
	require.NoError(t, err)

	require.Equal(t, 1, outData.ImageFormatProperties.MaxMipLevels)
	require.Equal(t, 3, outData.ImageFormatProperties.MaxArrayLayers)
	require.Equal(t, 5, outData.ImageFormatProperties.MaxResourceSize)
	require.Equal(t, core1_0.Samples8, outData.ImageFormatProperties.SampleCounts)
	require.Equal(t, 11, outData.ImageFormatProperties.MaxExtent.Width)
	require.Equal(t, 13, outData.ImageFormatProperties.MaxExtent.Height)
	require.Equal(t, 17, outData.ImageFormatProperties.MaxExtent.Depth)
}

func TestVulkanExtension_PhysicalDeviceMemoryProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceMemoryProperties2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pMemoryProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceMemoryProperties2KHR) {
			val := reflect.ValueOf(pMemoryProperties).Elem()

			require.Equal(t, uint64(1000059006), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_PROPERTIES_2_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			memory := val.FieldByName("memoryProperties")
			*(*uint32)(unsafe.Pointer(memory.FieldByName("memoryTypeCount").UnsafeAddr())) = uint32(1)
			*(*uint32)(unsafe.Pointer(memory.FieldByName("memoryHeapCount").UnsafeAddr())) = uint32(1)

			memoryType := memory.FieldByName("memoryTypes").Index(0)
			*(*uint32)(unsafe.Pointer(memoryType.FieldByName("heapIndex").UnsafeAddr())) = uint32(3)
			*(*int32)(unsafe.Pointer(memoryType.FieldByName("propertyFlags").UnsafeAddr())) = int32(16) // VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT

			memoryHeap := memory.FieldByName("memoryHeaps").Index(0)
			*(*uint64)(unsafe.Pointer(memoryHeap.FieldByName("size").UnsafeAddr())) = uint64(99)
			*(*int32)(unsafe.Pointer(memoryHeap.FieldByName("flags").UnsafeAddr())) = int32(1) // VK_MEMORY_HEAP_DEVICE_LOCAL_BIT
		})

	outData := khr_get_physical_device_properties2.MemoryPropertiesOutData{}
	err := extension.PhysicalDeviceMemoryProperties2(physicalDevice, &outData)
	require.NoError(t, err)
	require.Equal(t, []core1_0.MemoryType{
		{
			Properties: core1_0.MemoryPropertyLazilyAllocated,
			HeapIndex:  3,
		},
	}, outData.MemoryProperties.MemoryTypes)
	require.Equal(t, []core1_0.MemoryHeap{
		{
			Flags: core1_0.MemoryHeapDeviceLocal,
			Size:  99,
		},
	}, outData.MemoryProperties.MemoryHeaps)
}

func TestVulkanExtension_PhysicalDeviceProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	deviceUUID, err := uuid.NewUUID()
	require.NoError(t, err)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {
			val := reflect.ValueOf(pProperties).Elem()

			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			properties := val.FieldByName("properties")

			*(*uint32)(unsafe.Pointer(properties.FieldByName("apiVersion").UnsafeAddr())) = uint32(common.Vulkan1_1)
			*(*uint32)(unsafe.Pointer(properties.FieldByName("driverVersion").UnsafeAddr())) = uint32(common.CreateVersion(3, 2, 1))
			*(*uint32)(unsafe.Pointer(properties.FieldByName("vendorID").UnsafeAddr())) = uint32(3)
			*(*uint32)(unsafe.Pointer(properties.FieldByName("deviceID").UnsafeAddr())) = uint32(5)
			*(*uint32)(unsafe.Pointer(properties.FieldByName("deviceType").UnsafeAddr())) = uint32(2) // VK_PHYSICAL_DEVICE_TYPE_DISCRETE_GPU
			deviceNamePtr := (*driver.Char)(unsafe.Pointer(properties.FieldByName("deviceName").UnsafeAddr()))
			deviceNameSlice := ([]driver.Char)(unsafe.Slice(deviceNamePtr, 256))
			deviceName := "Some Device"
			for i, r := range []byte(deviceName) {
				deviceNameSlice[i] = driver.Char(r)
			}
			deviceNameSlice[len(deviceName)] = 0

			uuidPtr := (*driver.Char)(unsafe.Pointer(properties.FieldByName("pipelineCacheUUID").UnsafeAddr()))
			uuidSlice := ([]driver.Char)(unsafe.Slice(uuidPtr, 16))
			uuid, err := deviceUUID.MarshalBinary()
			require.NoError(t, err)

			for i, b := range uuid {
				uuidSlice[i] = driver.Char(b)
			}

			limits := properties.FieldByName("limits")
			*(*uint32)(unsafe.Pointer(limits.FieldByName("maxUniformBufferRange").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(limits.FieldByName("maxVertexInputBindingStride").UnsafeAddr())) = uint32(11)
			workGroupCount := limits.FieldByName("maxComputeWorkGroupCount")
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(0).UnsafeAddr())) = uint32(13)
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(1).UnsafeAddr())) = uint32(17)
			*(*uint32)(unsafe.Pointer(workGroupCount.Index(2).UnsafeAddr())) = uint32(19)
			*(*float32)(unsafe.Pointer(limits.FieldByName("maxInterpolationOffset").UnsafeAddr())) = float32(23)
			*(*driver.VkBool32)(unsafe.Pointer(limits.FieldByName("strictLines").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkDeviceSize)(unsafe.Pointer(limits.FieldByName("optimalBufferCopyRowPitchAlignment").UnsafeAddr())) = driver.VkDeviceSize(29)

			sparseProperties := properties.FieldByName("sparseProperties")
			*(*driver.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard2DBlockShape").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard2DMultisampleBlockShape").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyStandard3DBlockShape").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyAlignedMipSize").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(sparseProperties.FieldByName("residencyNonResidentStrict").UnsafeAddr())) = driver.VkBool32(1)
		})

	outData := khr_get_physical_device_properties2.DevicePropertiesOutData{}
	err = extension.PhysicalDeviceProperties2(physicalDevice, &outData)
	require.NoError(t, err)

	require.Equal(t, common.Vulkan1_1, outData.Properties.APIVersion)
	require.Equal(t, common.CreateVersion(3, 2, 1), outData.Properties.DriverVersion)
	require.Equal(t, uint32(3), outData.Properties.VendorID)
	require.Equal(t, uint32(5), outData.Properties.DeviceID)
	require.Equal(t, core1_0.DeviceDiscreteGPU, outData.Properties.Type)
	require.Equal(t, "Some Device", outData.Properties.Name)
	require.Equal(t, deviceUUID, outData.Properties.PipelineCacheUUID)

	require.Equal(t, 7, outData.Properties.Limits.MaxUniformBufferRange)
	require.Equal(t, 11, outData.Properties.Limits.MaxVertexInputBindingStride)
	require.Equal(t, [3]int{13, 17, 19}, outData.Properties.Limits.MaxComputeWorkGroupCount)
	require.Equal(t, float32(23), outData.Properties.Limits.MaxInterpolationOffset)
	require.True(t, outData.Properties.Limits.StrictLines)
	require.Equal(t, 29, outData.Properties.Limits.OptimalBufferCopyRowPitchAlignment)

	require.True(t, outData.Properties.SparseProperties.ResidencyStandard2DBlockShape)
	require.False(t, outData.Properties.SparseProperties.ResidencyStandard2DMultisampleBlockShape)
	require.True(t, outData.Properties.SparseProperties.ResidencyStandard3DBlockShape)
	require.False(t, outData.Properties.SparseProperties.ResidencyAlignedMipSize)
	require.True(t, outData.Properties.SparseProperties.ResidencyNonResidentStrict)
}

func TestVulkanExtension_PhysicalDeviceQueueFamilyProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil()), nil).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pQueueFamilyPropertyCount *driver.Uint32, pQueueFamilyProperties *khr_get_physical_device_properties2_driver.VkQueueFamilyProperties2KHR) {
			*pQueueFamilyPropertyCount = 2
		})

	extDriver.EXPECT().VkGetPhysicalDeviceQueueFamilyProperties2KHR(physicalDevice.Handle(), gomock.Not(gomock.Nil()), gomock.Not(gomock.Nil())).
		DoAndReturn(func(physicalDevice driver.VkPhysicalDevice, pQueueFamilyPropertyCount *driver.Uint32, pQueueFamilyProperties *khr_get_physical_device_properties2_driver.VkQueueFamilyProperties2KHR) {
			require.Equal(t, driver.Uint32(2), *pQueueFamilyPropertyCount)

			propertySlice := ([]khr_get_physical_device_properties2_driver.VkQueueFamilyProperties2KHR)(unsafe.Slice(pQueueFamilyProperties, 2))
			val := reflect.ValueOf(propertySlice)
			property := val.Index(0)

			require.Equal(t, uint64(1000059005), property.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR
			require.True(t, property.FieldByName("pNext").IsNil())

			queueFamily := property.FieldByName("queueFamilyProperties")
			*(*driver.VkQueueFlags)(unsafe.Pointer(queueFamily.FieldByName("queueFlags").UnsafeAddr())) = driver.VkQueueFlags(8) // VK_QUEUE_SPARSE_BINDING_BIT
			*(*uint32)(unsafe.Pointer(queueFamily.FieldByName("queueCount").UnsafeAddr())) = uint32(3)
			*(*uint32)(unsafe.Pointer(queueFamily.FieldByName("timestampValidBits").UnsafeAddr())) = uint32(5)

			propertyExtent := queueFamily.FieldByName("minImageTransferGranularity")
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("width").UnsafeAddr())) = uint32(7)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("height").UnsafeAddr())) = uint32(11)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("depth").UnsafeAddr())) = uint32(13)

			property = val.Index(1)
			require.Equal(t, uint64(1000059005), property.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_QUEUE_FAMILY_PROPERTIES_2_KHR
			require.True(t, property.FieldByName("pNext").IsNil())

			queueFamily = property.FieldByName("queueFamilyProperties")
			*(*driver.VkQueueFlags)(unsafe.Pointer(queueFamily.FieldByName("queueFlags").UnsafeAddr())) = driver.VkQueueFlags(2) // VK_QUEUE_COMPUTE_BIT
			*(*uint32)(unsafe.Pointer(queueFamily.FieldByName("queueCount").UnsafeAddr())) = uint32(17)
			*(*uint32)(unsafe.Pointer(queueFamily.FieldByName("timestampValidBits").UnsafeAddr())) = uint32(19)

			propertyExtent = queueFamily.FieldByName("minImageTransferGranularity")
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("width").UnsafeAddr())) = uint32(23)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("height").UnsafeAddr())) = uint32(29)
			*(*uint32)(unsafe.Pointer(propertyExtent.FieldByName("depth").UnsafeAddr())) = uint32(31)
		})

	outData, err := extension.PhysicalDeviceQueueFamilyProperties2(physicalDevice, nil)
	require.NoError(t, err)

	require.Equal(t, []*khr_get_physical_device_properties2.QueueFamilyOutData{
		{
			QueueFamily: core1_0.QueueFamily{
				Flags:              core1_0.QueueSparseBinding,
				QueueCount:         3,
				TimestampValidBits: 5,
				MinImageTransferGranularity: core1_0.Extent3D{
					Width:  7,
					Height: 11,
					Depth:  13,
				},
			},
		},
		{
			QueueFamily: core1_0.QueueFamily{
				Flags:              core1_0.QueueCompute,
				QueueCount:         17,
				TimestampValidBits: 19,
				MinImageTransferGranularity: core1_0.Extent3D{
					Width:  23,
					Height: 29,
					Depth:  31,
				},
			},
		},
	}, outData)
}

func TestVulkanExtension_PhysicalDeviceSparseImageFormatProperties(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceSparseImageFormatProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pFormatInfo *khr_get_physical_device_properties2_driver.VkPhysicalDeviceSparseImageFormatInfo2KHR,
		pPropertyCount *driver.Uint32,
		pProperties *khr_get_physical_device_properties2_driver.VkSparseImageFormatProperties2KHR) {

		val := reflect.ValueOf(pFormatInfo).Elem()
		require.Equal(t, uint64(1000059008), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(66), val.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_USCALED_PACK32
		require.Equal(t, uint64(2), val.FieldByName("_type").Uint())          // VK_IMAGE_TYPE_3D
		require.Equal(t, uint64(32), val.FieldByName("samples").Uint())       // VK_SAMPLE_COUNT_32_BIT
		require.Equal(t, uint64(0x00000008), val.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_STORAGE_BIT
		require.Equal(t, uint64(1), val.FieldByName("tiling").Uint())         // VK_IMAGE_TILING_LINEAR

		*pPropertyCount = 1
	})

	extDriver.EXPECT().VkGetPhysicalDeviceSparseImageFormatProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pFormatInfo *khr_get_physical_device_properties2_driver.VkPhysicalDeviceSparseImageFormatInfo2KHR,
		pPropertyCount *driver.Uint32,
		pProperties *khr_get_physical_device_properties2_driver.VkSparseImageFormatProperties2KHR) {

		val := reflect.ValueOf(pFormatInfo).Elem()
		require.Equal(t, uint64(1000059008), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SPARSE_IMAGE_FORMAT_INFO_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(66), val.FieldByName("format").Uint())        // VK_FORMAT_A2B10G10R10_USCALED_PACK32
		require.Equal(t, uint64(2), val.FieldByName("_type").Uint())          // VK_IMAGE_TYPE_3D
		require.Equal(t, uint64(32), val.FieldByName("samples").Uint())       // VK_SAMPLE_COUNT_32_BIT
		require.Equal(t, uint64(0x00000008), val.FieldByName("usage").Uint()) // VK_IMAGE_USAGE_STORAGE_BIT
		require.Equal(t, uint64(1), val.FieldByName("tiling").Uint())         // VK_IMAGE_TILING_LINEAR

		require.Equal(t, driver.Uint32(1), *pPropertyCount)

		propertySlice := ([]khr_get_physical_device_properties2_driver.VkSparseImageFormatProperties2KHR)(unsafe.Slice(pProperties, 1))
		outData := reflect.ValueOf(propertySlice)
		prop := outData.Index(0)
		require.Equal(t, uint64(1000059007), prop.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SPARSE_IMAGE_FORMAT_PROPERTIES_2_KHR
		require.True(t, prop.FieldByName("pNext").IsNil())

		sparseProps := prop.FieldByName("properties")
		*(*uint32)(unsafe.Pointer(sparseProps.FieldByName("aspectMask").UnsafeAddr())) = uint32(1) // VK_IMAGE_ASPECT_COLOR_BIT
		*(*int32)(unsafe.Pointer(sparseProps.FieldByName("imageGranularity").FieldByName("width").UnsafeAddr())) = int32(1)
		*(*int32)(unsafe.Pointer(sparseProps.FieldByName("imageGranularity").FieldByName("height").UnsafeAddr())) = int32(3)
		*(*int32)(unsafe.Pointer(sparseProps.FieldByName("imageGranularity").FieldByName("depth").UnsafeAddr())) = int32(5)
		*(*uint32)(unsafe.Pointer(sparseProps.FieldByName("flags").UnsafeAddr())) = uint32(4) // VK_SPARSE_IMAGE_FORMAT_NONSTANDARD_BLOCK_SIZE_BIT
	})

	outData, err := extension.PhysicalDeviceSparseImageFormatProperties2(physicalDevice,
		khr_get_physical_device_properties2.SparseImageFormatOptions{
			Format:  core1_0.DataFormatA2B10G10R10UnsignedScaledPacked,
			Type:    core1_0.ImageType3D,
			Samples: core1_0.Samples32,
			Usage:   core1_0.ImageUsageStorage,
			Tiling:  core1_0.ImageTilingLinear,
		}, nil)
	require.NoError(t, err)
	require.Equal(t, []*khr_get_physical_device_properties2.SparseImageFormatPropertiesOutData{
		{
			SparseImageFormatProperties: core1_0.SparseImageFormatProperties{
				AspectMask: core1_0.AspectColor,
				ImageGranularity: core1_0.Extent3D{
					Width:  1,
					Height: 3,
					Depth:  5,
				},
				Flags: core1_0.SparseImageFormatNonstandardBlockSize,
			},
		},
	}, outData)
}
