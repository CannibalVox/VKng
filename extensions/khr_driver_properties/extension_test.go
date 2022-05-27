package khr_driver_properties

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	khr_driver_properties_driver "github.com/CannibalVox/VKng/extensions/khr_driver_properties/driver"
	"github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2"
	khr_get_physical_device_properties2_driver "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/driver"
	mock_get_physical_device_properties2 "github.com/CannibalVox/VKng/extensions/khr_get_physical_device_properties2/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"unsafe"
)

func TestPhysicalDeviceDriverOutData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	extDriver := mock_get_physical_device_properties2.NewMockDriver(ctrl)
	extension := khr_get_physical_device_properties2.CreateExtensionFromDriver(extDriver)

	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	physicalDevice := mocks.EasyMockPhysicalDevice(ctrl, coreDriver)

	extDriver.EXPECT().VkGetPhysicalDeviceProperties2KHR(
		physicalDevice.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(physicalDevice driver.VkPhysicalDevice,
			pProperties *khr_get_physical_device_properties2_driver.VkPhysicalDeviceProperties2KHR) {

			val := reflect.ValueOf(pProperties).Elem()
			require.Equal(t, uint64(1000059001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2_KHR

			next := (*khr_driver_properties_driver.VkPhysicalDeviceDriverPropertiesKHR)(val.FieldByName("pNext").UnsafePointer())
			val = reflect.ValueOf(next).Elem()

			require.Equal(t, uint64(1000196000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES_KHR
			require.True(t, val.FieldByName("pNext").IsNil())

			*(*uint32)(unsafe.Pointer(val.FieldByName("driverID").UnsafeAddr())) = uint32(10) // VK_DRIVER_ID_GOOGLE_SWIFTSHADER_KHR
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("major").UnsafeAddr())) = uint8(1)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("minor").UnsafeAddr())) = uint8(3)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("subminor").UnsafeAddr())) = uint8(5)
			*(*uint8)(unsafe.Pointer(val.FieldByName("conformanceVersion").FieldByName("patch").UnsafeAddr())) = uint8(7)

			driverNamePtr := (*driver.Char)(unsafe.Pointer(val.FieldByName("driverName").UnsafeAddr()))
			driverNameSlice := ([]driver.Char)(unsafe.Slice(driverNamePtr, 256))
			driverName := "Some Driver"
			for i, r := range []byte(driverName) {
				driverNameSlice[i] = driver.Char(r)
			}
			driverNameSlice[len(driverName)] = 0

			driverInfoPtr := (*driver.Char)(unsafe.Pointer(val.FieldByName("driverInfo").UnsafeAddr()))
			driverInfoSlice := ([]driver.Char)(unsafe.Slice(driverInfoPtr, 256))
			driverInfo := "Whooo Info"
			for i, r := range []byte(driverInfo) {
				driverInfoSlice[i] = driver.Char(r)
			}
			driverInfoSlice[len(driverInfo)] = 0
		})

	var driverOutData PhysicalDeviceDriverOutData
	err := extension.PhysicalDeviceProperties2(
		physicalDevice,
		&khr_get_physical_device_properties2.DevicePropertiesOutData{
			HaveNext: common.HaveNext{&driverOutData},
		})
	require.NoError(t, err)
	require.Equal(t, PhysicalDeviceDriverOutData{
		DriverID:           DriverIDGoogleSwiftshader,
		DriverName:         "Some Driver",
		DriverInfo:         "Whooo Info",
		ConformanceVersion: ConformanceVersion{Major: 1, Minor: 3, Subminor: 5, Patch: 7},
	}, driverOutData)
}
