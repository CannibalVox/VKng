package internal1_0_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	internal_mocks "github.com/CannibalVox/VKng/core/internal/dummies"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestVulkanLoader1_0_CreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := internal_mocks.EasyDummyDevice(mockDriver)
	eventHandle := mocks.NewFakeEventHandle()

	mockDriver.EXPECT().VkCreateEvent(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device driver.VkDevice, pCreateInfo *driver.VkEventCreateInfo, pAllocator *driver.VkAllocationCallbacks, pEvent *driver.VkEvent) (common.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(10), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EVENT_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())

			*pEvent = eventHandle
			return core1_0.VKSuccess, nil
		})

	event, _, err := device.CreateEvent(nil, core1_0.EventCreateOptions{
		Flags: 0,
	})
	require.NoError(t, err)
	require.NotNil(t, event)
	require.Equal(t, eventHandle, event.Handle())
}

func TestVulkanEvent_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, driver)
	event := internal_mocks.EasyDummyEvent(driver, device)

	driver.EXPECT().VkSetEvent(device.Handle(), event.Handle()).Return(core1_0.VKSuccess, nil)

	_, err := event.Set()
	require.NoError(t, err)
}

func TestVulkanEvent_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, driver)
	event := internal_mocks.EasyDummyEvent(driver, device)

	driver.EXPECT().VkResetEvent(device.Handle(), event.Handle()).Return(core1_0.VKSuccess, nil)

	_, err := event.Reset()
	require.NoError(t, err)
}

func TestVulkanEvent_Status(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, driver)
	event := internal_mocks.EasyDummyEvent(driver, device)

	driver.EXPECT().VkGetEventStatus(device.Handle(), event.Handle()).Return(core1_0.VKEventReset, nil)

	res, err := event.Status()
	require.NoError(t, err)
	require.Equal(t, core1_0.VKEventReset, res)
}
