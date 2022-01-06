package core_test

import (
	"github.com/CannibalVox/VKng/core"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestVulkanLoader1_0_CreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyDummyDevice(t, ctrl, loader)
	eventHandle := mocks.NewFakeEventHandle()

	driver.EXPECT().VkCreateEvent(device.Handle(), gomock.Not(nil), nil, gomock.Not(nil)).DoAndReturn(
		func(device core.VkDevice, pCreateInfo *core.VkEventCreateInfo, pAllocator *core.VkAllocationCallbacks, pEvent *core.VkEvent) (core.VkResult, error) {
			val := reflect.ValueOf(*pCreateInfo)

			require.Equal(t, uint64(10), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_EVENT_CREATE_INFO
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(1), val.FieldByName("flags").Uint()) // VK_EVENT_CREATE_DEVICE_ONLY_BIT_KHR

			*pEvent = eventHandle
			return core.VKSuccess, nil
		})

	event, _, err := loader.CreateEvent(device, &core.EventOptions{
		Flags: core.EventDeviceOnlyKHR,
	})
	require.NoError(t, err)
	require.NotNil(t, event)
	require.Equal(t, eventHandle, event.Handle())
}

func TestVulkanEvent_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	event := mocks.EasyDummyEvent(t, loader, device)

	driver.EXPECT().VkSetEvent(device.Handle(), event.Handle()).Return(core.VKSuccess, nil)

	_, err = event.Set()
	require.NoError(t, err)
}

func TestVulkanEvent_Reset(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	event := mocks.EasyDummyEvent(t, loader, device)

	driver.EXPECT().VkResetEvent(device.Handle(), event.Handle()).Return(core.VKSuccess, nil)

	_, err = event.Reset()
	require.NoError(t, err)
}

func TestVulkanEvent_Status(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	driver := mocks.NewMockDriver(ctrl)
	loader, err := core.CreateLoaderFromDriver(driver)
	require.NoError(t, err)

	device := mocks.EasyMockDevice(ctrl, driver)
	event := mocks.EasyDummyEvent(t, loader, device)

	driver.EXPECT().VkGetEventStatus(device.Handle(), event.Handle()).Return(core.VKEventReset, nil)

	res, err := event.Status()
	require.NoError(t, err)
	require.Equal(t, core.VKEventReset, res)
}