package khr_shader_float_controls_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/CannibalVox/VKng/extensions/khr_shader_float_controls"
	khr_shader_float_controls_driver "github.com/CannibalVox/VKng/extensions/khr_shader_float_controls/driver"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestPhysicalDeviceFloatControlsOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(func(physicalDevice driver.VkPhysicalDevice,
		pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {

		val := reflect.ValueOf(pProperties).Elem()
		require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

		next := (*khr_shader_float_controls_driver.VkPhysicalDeviceFloatControlsPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
		val = reflect.ValueOf(next).Elem()
		require.Equal(t, uint64(1000197000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FLOAT_CONTROLS_PROPERTIES_KHR
		require.True(t, val.FieldByName("pNext").IsNil())
		denormBehavior := (*khr_shader_float_controls_driver.VkShaderFloatControlsIndependenceKHR)(unsafe.Pointer(val.FieldByName("denormBehaviorIndependence").UnsafeAddr()))
		*denormBehavior = khr_shader_float_controls_driver.VkShaderFloatControlsIndependenceKHR(0) // VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR
		roundingMode := (*khr_shader_float_controls_driver.VkShaderFloatControlsIndependenceKHR)(unsafe.Pointer(val.FieldByName("roundingModeIndependence").UnsafeAddr()))
		*roundingMode = khr_shader_float_controls_driver.VkShaderFloatControlsIndependenceKHR(1) // VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR

		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat16").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat32").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderSignedZeroInfNanPreserveFloat64").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat16").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat32").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormPreserveFloat64").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat16").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat32").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderDenormFlushToZeroFloat64").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat16").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat32").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTEFloat64").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat16").UnsafeAddr())) = driver.VkBool32(1)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat32").UnsafeAddr())) = driver.VkBool32(0)
		*(*driver.VkBool32)(unsafe.Pointer(val.FieldByName("shaderRoundingModeRTZFloat64").UnsafeAddr())) = driver.VkBool32(1)
	})

	var outData khr_shader_float_controls.PhysicalDeviceFloatControlsProperties
	err := extension.PhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.PhysicalDeviceProperties2{
			NextOutData: common.NextOutData{&outData},
		})
	require.NoError(t, err)
	require.Equal(t, khr_shader_float_controls.PhysicalDeviceFloatControlsProperties{
		DenormBehaviorIndependence: khr_shader_float_controls.ShaderFloatControlsIndependence32BitOnly,
		RoundingMoundIndependence:  khr_shader_float_controls.ShaderFloatControlsIndependenceAll,

		ShaderSignedZeroInfNanPreserveFloat16: true,
		ShaderSignedZeroInfNanPreserveFloat32: false,
		ShaderSignedZeroInfNanPreserveFloat64: true,
		ShaderDenormPreserveFloat16:           false,
		ShaderDenormPreserveFloat32:           true,
		ShaderDenormPreserveFloat64:           false,
		ShaderDenormFlushToZeroFloat16:        true,
		ShaderDenormFlushToZeroFloat32:        false,
		ShaderDenormFlushToZeroFloat64:        true,
		ShaderRoundingModeRTEFloat16:          false,
		ShaderRoundingModeRTEFloat32:          true,
		ShaderRoundingModeRTEFloat64:          false,
		ShaderRoundingModeRTZFloat16:          true,
		ShaderRoundingModeRTZFloat32:          false,
		ShaderRoundingModeRTZFloat64:          true,
	}, outData)
}
