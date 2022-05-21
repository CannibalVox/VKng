package khr_portability_subset

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	khr_portability_subset_driver "github.com/CannibalVox/VKng/extensions/khr_portability_subset/driver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestPhysicalDevicePortabilitySubsetFeaturesOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extensionDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extensionDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extensionDriver.EXPECT().VkGetPhysicalDeviceFeatures2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pFeatures *khr_get_physical_device_properties2_driver.VkPhysicalDeviceFeatures2KHR) {

			val := reflect.ValueOf(pFeatures).Elem()
			require.Equal(t, uint64(1000059000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FEATURES_2_KHR

			next := (*khr_portability_subset_driver.VkPhysicalDevicePortabilitySubsetFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()
			require.Equal(t, uint64(1000163000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("constantAlphaColorBlendFactors").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("events").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("imageViewFormatReinterpretation").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("imageViewFormatSwizzle").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("imageView2DOn3DImage").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("multisampleArrayImage").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("mutableComparisonSamplers").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("pointPolygons").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("samplerMipLodBias").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("separateStencilMaskRef").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSampleRateInterpolationFunctions").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("tessellationIsolines").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("tessellationPointMode").UnsafeAddr())) = driver.VkBool32(1)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("triangleFans").UnsafeAddr())) = driver.VkBool32(0)
			*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("vertexAttributeAccessBeyondStride").UnsafeAddr())) = driver.VkBool32(1)
		})

	var subsetFeatures PhysicalDevicePortabilitySubsetFeaturesOutData
	err := extension.PhysicalDeviceFeatures(
		physicalDevice,
		&khr_get_physical_device_properties2.DeviceFeaturesOutData{
			HaveNext: common.HaveNext{&subsetFeatures},
		})
	require.NoError(t, err)
	require.Equal(t, PhysicalDevicePortabilitySubsetFeaturesOutData{
		ConstantAlphaColorBlendFactors:          true,
		ImageViewFormatReinterpretation:         true,
		ImageView2DOn3DImage:                    true,
		MutableComparisonSamplers:               true,
		SamplerMipLodBias:                       true,
		ShaderSamplerRateInterpolationFunctions: true,
		TessellationPointMode:                   true,
		VertexAttributeAccessBeyondStride:       true,
	}, subsetFeatures)
}

func TestPhysicalDevicePortabilitySubsetOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extensionDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extensionDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extensionDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice, pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {
			val := reflect.ValueOf(pProperties).Elem()
			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

			next := (*khr_portability_subset_driver.VkPhysicalDevicePortabilitySubsetPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000163001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_PROPERTIES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())
			*(*driver.Uint32)(unsafe.Pointer(val.FieldByName("minVertexInputBindingStrideAlignment").UnsafeAddr())) = driver.Uint32(3)
		})

	var subsetProperties PhysicalDevicePortabilitySubsetOutData
	err := extension.PhysicalDeviceProperties(
		physicalDevice,
		&khr_get_physical_device_properties2.DevicePropertiesOutData{
			HaveNext: common.HaveNext{&subsetProperties},
		})
	require.NoError(t, err)
	require.Equal(t, PhysicalDevicePortabilitySubsetOutData{
		MinVertexInputBindingStrideAlignment: 3,
	}, subsetProperties)
}

func TestPhysicalDevicePortabilitySubsetFeaturesOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	coreDriver.EXPECT().CreateDeviceDriver(gomock.Any()).Return(coreDriver, nil)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)
	mockDevice := mocks.EasyMockDevice(ctrl, coreDriver)
	loader, err := core.CreateLoaderFromDriver(coreDriver)
	require.NoError(t, err)

	coreDriver.EXPECT().VkCreateDevice(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pCreateInfo *driver.VkDeviceCreateInfo,
		pAllocator *driver.VkAllocationCallbacks,
		pDevice *driver.VkDevice) (common.VkResult, error) {

		*pDevice = mockDevice.Handle()

		val := reflect.ValueOf(pCreateInfo).Elem()
		require.Equal(t, uint64(3), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO

		next := (*khr_portability_subset_driver.VkPhysicalDevicePortabilitySubsetFeaturesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()

		require.Equal(t, uint64(1000163000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PORTABILITY_SUBSET_FEATURES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		require.Equal(t, uint64(0), val.FieldByName("constantAlphaColorBlendFactors").Uint())
		require.Equal(t, uint64(1), val.FieldByName("events").Uint())
		require.Equal(t, uint64(0), val.FieldByName("imageViewFormatReinterpretation").Uint())
		require.Equal(t, uint64(1), val.FieldByName("imageViewFormatSwizzle").Uint())
		require.Equal(t, uint64(0), val.FieldByName("imageView2DOn3DImage").Uint())
		require.Equal(t, uint64(1), val.FieldByName("multisampleArrayImage").Uint())
		require.Equal(t, uint64(0), val.FieldByName("mutableComparisonSamplers").Uint())
		require.Equal(t, uint64(1), val.FieldByName("pointPolygons").Uint())
		require.Equal(t, uint64(0), val.FieldByName("samplerMipLodBias").Uint())
		require.Equal(t, uint64(1), val.FieldByName("separateStencilMaskRef").Uint())
		require.Equal(t, uint64(0), val.FieldByName("shaderSampleRateInterpolationFunctions").Uint())
		require.Equal(t, uint64(1), val.FieldByName("tessellationIsolines").Uint())
		require.Equal(t, uint64(0), val.FieldByName("tessellationPointMode").Uint())
		require.Equal(t, uint64(1), val.FieldByName("triangleFans").Uint())
		require.Equal(t, uint64(0), val.FieldByName("vertexAttributeAccessBeyondStride").Uint())

		return core1_0.VKSuccess, nil
	})

	device, _, err := loader.CreateDevice(
		physicalDevice,
		nil,
		core1_0.DeviceCreateOptions{
			QueueFamilies: []core1_0.DeviceQueueCreateOptions{
				{
					CreatedQueuePriorities: []float32{0},
				},
			},

			HaveNext: common.HaveNext{PhysicalDevicePortabilitySubsetFeaturesOptions{
				Events:                 true,
				ImageViewFormatSwizzle: true,
				MultisampleArrayImage:  true,
				PointPolygons:          true,
				SeparateStencilMaskRef: true,
				TessellationIsolines:   true,
				TriangleFans:           true,
			}},
		})
	require.NoError(t, err)
	require.Equal(t, mockDevice.Handle(), device.Handle())
}
