package khr_sampler_ycbcr_conversion_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/common/extensions"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_bind_memory2"
	khr_bind_memory2_driver "github.com/CannibalVox/VKng/extensions/khr_bind_memory2/driver"
	mock_bind_memory2 "github.com/CannibalVox/VKng/extensions/khr_bind_memory2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_memory_requirements2"
	khr_get_memory_requirements2_driver "github.com/CannibalVox/VKng/extensions/khr_get_memory_requirements2/driver"
	mock_get_memory_requirements2 "github.com/CannibalVox/VKng/extensions/khr_get_memory_requirements2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_sampler_ycbcr_conversion"
	khr_sampler_ycbcr_conversion_driver "github.com/CannibalVox/VKng/extensions/khr_sampler_ycbcr_conversion/driver"
	mock_sampler_ycbcr_conversion "github.com/CannibalVox/VKng/extensions/khr_sampler_ycbcr_conversion/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestVulkanExtension_CreateSamplerYcbcrConversion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_sampler_ycbcr_conversion.NewMockDriver(ctrl)
	extension := khr_sampler_ycbcr_conversion.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	mockYcbcr := mock_sampler_ycbcr_conversion.EasyMockSamplerYcbcrConversion(ctrl)

	extDriver.EXPECT().VkCreateSamplerYcbcrConversionKHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice,
			pCreateInfo *khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionCreateInfoKHR,
			pAllocator *driver.VkAllocationCallbacks,
			pYcbcrConversion *khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR,
		) (common.VkResult, error) {
			*pYcbcrConversion = mockYcbcr.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(1000156000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1000156021), val.FieldByName("format").Uint())             // VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR
			require.Equal(t, uint64(2), val.FieldByName("ycbcrModel").Uint())                  // VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR
			require.Equal(t, uint64(1), val.FieldByName("ycbcrRange").Uint())                  // VK_SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR
			require.Equal(t, uint64(4), val.FieldByName("components").FieldByName("r").Uint()) // VK_COMPONENT_SWIZZLE_G
			require.Equal(t, uint64(6), val.FieldByName("components").FieldByName("g").Uint()) // VK_COMPONENT_SWIZZLE_A
			require.Equal(t, uint64(0), val.FieldByName("components").FieldByName("b").Uint()) // VK_COMPONENT_SWIZZLE_IDENTITY
			require.Equal(t, uint64(2), val.FieldByName("components").FieldByName("a").Uint()) // VK_COMPONENT_SWIZZLE_ONE
			require.Equal(t, uint64(0), val.FieldByName("yChromaOffset").Uint())               // VK_CHROMA_LOCATION_COSITED_EVEN_KHR
			require.Equal(t, uint64(1), val.FieldByName("xChromaOffset").Uint())               // VK_CHROMA_LOCATION_MIDPOINT_KHR
			require.Equal(t, uint64(1), val.FieldByName("forceExplicitReconstruction").Uint())

			return core1_0.VKSuccess, nil
		})

	ycbcr, _, err := extension.CreateSamplerYcbcrConversion(device,
		khr_sampler_ycbcr_conversion.SamplerYcbcrConversionCreateOptions{
			Format:     khr_sampler_ycbcr_conversion.DataFormatB12X4G12X4R12X4G12X4HorizontalChromaComponentPacked,
			YcbcrModel: khr_sampler_ycbcr_conversion.SamplerYcbcrModelConversionYcbcr709,
			YcbcrRange: khr_sampler_ycbcr_conversion.SamplerYcbcrRangeITUNarrow,
			Components: core1_0.ComponentMapping{
				R: core1_0.SwizzleGreen,
				G: core1_0.SwizzleAlpha,
				B: core1_0.SwizzleIdentity,
				A: core1_0.SwizzleOne,
			},
			ChromaOffsetY:               khr_sampler_ycbcr_conversion.ChromaLocationCositedEven,
			ChromaOffsetX:               khr_sampler_ycbcr_conversion.ChromaLocationMidpoint,
			ChromaFilter:                core1_0.FilterLinear,
			ForceExplicitReconstruction: true,
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, mockYcbcr.Handle(), ycbcr.Handle())

	extDriver.EXPECT().VkDestroySamplerYcbcrConversionKHR(
		device.Handle(),
		ycbcr.Handle(),
		gomock.Nil(),
	)

	ycbcr.Destroy(nil)
}
func TestBindImagePlaneMemoryOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_bind_memory2.NewMockDriver(ctrl)
	extension := khr_bind_memory2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	image := mocks.EasyMockImage(ctrl)
	memory := mocks.EasyMockDeviceMemory(ctrl)

	extDriver.EXPECT().VkBindImageMemory2KHR(
		device.Handle(),
		driver.Uint32(1),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		bindInfoCount driver.Uint32,
		pBindInfos *khr_bind_memory2_driver.VkBindImageMemoryInfoKHR,
	) (common.VkResult, error) {
		val := reflect.ValueOf(pBindInfos).Elem()

		require.Equal(t, uint64(1000157001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO_KHR
		require.Equal(t, image.Handle(), driver.VkImage(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, memory.Handle(), driver.VkDeviceMemory(val.FieldByName("memory").UnsafePointer()))

		next := (*khr_sampler_ycbcr_conversion_driver.VkBindImagePlaneMemoryInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000156002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x40), val.FieldByName("planeAspect").Uint()) // VK_IMAGE_ASPECT_PLANE_2_BIT_KHR

		return core1_0.VKSuccess, nil
	})

	_, err := extension.BindImageMemory(device,
		[]khr_bind_memory2.BindImageMemoryOptions{
			{
				Image:  image,
				Memory: memory,

				HaveNext: common.HaveNext{
					khr_sampler_ycbcr_conversion.BindImagePlaneMemoryOptions{
						PlaneAspect: khr_sampler_ycbcr_conversion.ImageAspectPlane2,
					},
				},
			},
		})
	require.NoError(t, err)
}

func TestImagePlaneMemoryRequirementsOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_memory_requirements2.NewMockDriver(ctrl)
	extension := khr_get_memory_requirements2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)

	image := mocks.EasyMockImage(ctrl)

	extDriver.EXPECT().VkGetImageMemoryRequirements2KHR(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pInfo *khr_get_memory_requirements2_driver.VkImageMemoryRequirementsInfo2KHR,
		pMemoryRequirements *khr_get_memory_requirements2_driver.VkMemoryRequirements2KHR,
	) {
		val := reflect.ValueOf(pInfo).Elem()
		require.Equal(t, uint64(1000146001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_MEMORY_REQUIREMENTS_INFO_2_KHR
		require.Equal(t, image.Handle(), driver.VkImage(val.FieldByName("image").UnsafePointer()))

		next := (*khr_sampler_ycbcr_conversion_driver.VkImagePlaneMemoryRequirementsInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000156003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0x20), val.FieldByName("planeAspect").Uint()) // VK_IMAGE_ASPECT_PLANE_1_BIT_KHR

		val = reflect.ValueOf(pMemoryRequirements).Elem()
		require.Equal(t, uint64(1000146003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_MEMORY_REQUIREMENTS_2_KHR
		require.True(t, val.FieldByName("pNext").IsNil())

		*(*uint32)(unsafe.Pointer(val.FieldByName("memoryRequirements").FieldByName("size").UnsafeAddr())) = uint32(17)
		*(*uint32)(unsafe.Pointer(val.FieldByName("memoryRequirements").FieldByName("alignment").UnsafeAddr())) = uint32(19)
		*(*uint32)(unsafe.Pointer(val.FieldByName("memoryRequirements").FieldByName("memoryTypeBits").UnsafeAddr())) = uint32(7)
	})

	var outData khr_get_memory_requirements2.MemoryRequirementsOutData
	err := extension.ImageMemoryRequirements(
		device,
		khr_get_memory_requirements2.ImageMemoryRequirementsOptions{
			Image: image,
			HaveNext: common.HaveNext{
				khr_sampler_ycbcr_conversion.ImagePlaneMemoryRequirementsOptions{
					PlaneAspect: khr_sampler_ycbcr_conversion.ImageAspectPlane1,
				},
			},
		},
		&outData)
	require.NoError(t, err)
	require.Equal(t, khr_get_memory_requirements2.MemoryRequirementsOutData{
		MemoryRequirements: core1_0.MemoryRequirements{
			Size:       17,
			Alignment:  19,
			MemoryType: 7,
		},
	}, outData)
}

func TestSamplerYcbcrConversionOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := extensions.CreateDeviceObject(coreDriver, mocks.NewFakeDeviceHandle(), common.Vulkan1_0)
	image := mocks.EasyMockImage(ctrl)
	ycbcr := mock_sampler_ycbcr_conversion.EasyMockSamplerYcbcrConversion(ctrl)
	mockImageView := mocks.EasyMockImageView(ctrl)

	coreDriver.EXPECT().VkCreateImageView(
		device.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(device driver.VkDevice,
		pCreateInfo *driver.VkImageViewCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pView *driver.VkImageView,
	) (common.VkResult, error) {
		*pView = mockImageView.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(15), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_VIEW_CREATE_INFO
		require.Equal(t, image.Handle(), driver.VkImage(val.FieldByName("image").UnsafePointer()))
		require.Equal(t, uint64(1000156028), val.FieldByName("format").Uint()) // VK_FORMAT_B16G16R16G16_422_UNORM_KHR

		next := (*khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionInfoKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000156001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, ycbcr.Handle(), khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionKHR(val.FieldByName("conversion").UnsafePointer()))

		return core1_0.VKSuccess, nil
	})

	imageView, _, err := device.CreateImageView(
		nil,
		core1_0.ImageViewCreateOptions{
			Image:  image,
			Format: khr_sampler_ycbcr_conversion.DataFormatB16G16R16G16HorizontalChroma,

			HaveNext: common.HaveNext{
				khr_sampler_ycbcr_conversion.SamplerYcbcrConversionOptions{
					Conversion: ycbcr,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockImageView.Handle(), imageView.Handle())
}

func TestSamplerYcbcrFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	physicalDevice := extensions.CreatePhysicalDeviceObject(coreDriver, instance.Handle(), mocks.NewFakePhysicalDeviceHandle(), common.Vulkan1_0, common.Vulkan1_0)
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pCreateInfo *driver.VkDeviceCreateInfo,
			pAllocator *driver.VkAllocationCallbacks,
			pDevice *driver.VkDevice,
		) (common.VkResult, error) {
			*pDevice = mockDevice.Handle()

			val := reflect.ValueOf(pCreateInfo).Elem()
			require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

			next := (*khr_sampler_ycbcr_conversion_driver.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000156004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("samplerYcbcrConversion").Uint())

			return core1_0.VKSuccess, nil
		})

	device, _, err := physicalDevice.CreateDevice(
		nil,
		core1_0.DeviceCreateOptions{
			QueueFamilies: []core1_0.DeviceQueueCreateOptions{
				{
					CreatedQueuePriorities: []float32{0},
				},
			},

			HaveNext: common.HaveNext{
				khr_sampler_ycbcr_conversion.PhysicalDeviceSamplerYcbcrFeaturesOptions{
					SamplerYcbcrConversion: true,
				},
			},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}

func TestSamplerYcbcrFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR,
		) {
			val := reflect.ValueOf(pFeatures).Elem()
			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			next := (*khr_sampler_ycbcr_conversion_driver.VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000156004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("samplerYcbcrConversion").UnsafeAddr())) = driver.VkBool32(1)
		})

	var outData khr_sampler_ycbcr_conversion.PhysicalDeviceSamplerYcbcrFeaturesOutData

	err := extension.PhysicalDeviceFeatures(
		physicalDevice,
		&khr_get_physical_device_properties2.DeviceFeaturesOutData{
			HaveNext: common.HaveNext{
				&outData,
			},
		})
	require.NoError(t, err)
	require.Equal(t, khr_sampler_ycbcr_conversion.PhysicalDeviceSamplerYcbcrFeaturesOutData{
		SamplerYcbcrConversion: true,
	}, outData)
}

func TestSamplerYcbcrImageFormatOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceImageFormatProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pImageFormatInfo *khr_get_physical_device_properties2_driver.VkPhysicalDeviceImageFormatInfo2KHR,
			pImageFormatProperties *khr_get_physical_device_properties2_driver.VkImageFormatProperties2KHR,
		) (common.VkResult, error) {
			val := reflect.ValueOf(pImageFormatInfo).Elem()
			require.Equal(t, uint64(1000059004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_FORMAT_INFO_2_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			val = reflect.ValueOf(pImageFormatProperties).Elem()
			require.Equal(t, uint64(1000059003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_IMAGE_FORMAT_PROPERTIES_2_KHR

			next := (*khr_sampler_ycbcr_conversion_driver.VkSamplerYcbcrConversionImageFormatPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000156005), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			*(*uint32)(unsafe.Pointer(val.FieldByName("combinedImageSamplerDescriptorCount").UnsafeAddr())) = uint32(7)

			return core1_0.VKSuccess, nil
		})

	var outData khr_sampler_ycbcr_conversion.SamplerYcbcrImageFormatOutData
	_, err := extension.PhysicalDeviceImageFormatProperties(
		physicalDevice,
		khr_get_physical_device_properties2.ImageFormatOptions{},
		&khr_get_physical_device_properties2.ImageFormatPropertiesOutData{
			HaveNext: common.HaveNext{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, khr_sampler_ycbcr_conversion.SamplerYcbcrImageFormatOutData{
		CombinedImageSamplerDescriptorCount: 7,
	}, outData)
}
