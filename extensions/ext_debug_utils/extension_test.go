package ext_debug_utils_test

import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	mock_driver "github.com/CannibalVox/VKng/core/driver/mocks"
	"github.com/CannibalVox/VKng/core/mocks"
	"github.com/CannibalVox/VKng/extensions/ext_debug_utils"
	ext_debug_utils_driver "github.com/CannibalVox/VKng/extensions/ext_debug_utils/driver"
	mock_debugutils "github.com/CannibalVox/VKng/extensions/ext_debug_utils/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/image/colornames"
	"image/color"
	"reflect"
	"runtime/cgo"
	"testing"
	"unsafe"
)

func strToCharSlice(text string) []driver.Char {
	slice := make([]driver.Char, len(text)+1)
	byteSlice := []byte(text)
	for idx, b := range byteSlice {
		slice[idx] = driver.Char(b)
	}
	slice[len(byteSlice)] = driver.Char(0)

	return slice
}

func cStringToCharSlice(strPtr unsafe.Pointer) []driver.Char {
	var slice []driver.Char

	var iterPtr *driver.Char
	for doWhile := true; doWhile; doWhile = (*iterPtr != 0) {
		iterPtr = (*driver.Char)(strPtr)
		slice = append(slice, *iterPtr)

		strPtr = unsafe.Add(strPtr, 1)
	}

	return slice
}

func TestVulkanExtension_CreateMessenger(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	debugDriver := mock_debugutils.NewMockDriver(ctrl)
	extension := ext_debug_utils.CreateExtensionFromDriver(debugDriver)
	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)

	calledRightFunction := false
	var cb ext_debug_utils.CallbackFunction = func(msgType ext_debug_utils.MessageTypes, severity ext_debug_utils.MessageSeverities, data *ext_debug_utils.DebugUtilsMessengerCallbackData) bool {
		calledRightFunction = true
		return true
	}

	expectedMessenger := mock_debugutils.EasyMockMessenger(ctrl)

	debugDriver.EXPECT().VkCreateDebugUtilsMessengerEXT(
		instance.Handle(),
		gomock.Not(gomock.Nil()),
		gomock.Nil(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(instance driver.VkInstance, pCreateInfo *ext_debug_utils_driver.VkDebugUtilsMessengerCreateInfoEXT, pAllocator *driver.VkAllocationCallbacks, pDebugMessenger *ext_debug_utils_driver.VkDebugUtilsMessengerEXT) (common.VkResult, error) {
			*pDebugMessenger = expectedMessenger.Handle()

			val := reflect.ValueOf(*pCreateInfo)
			require.Equal(t, uint64(1000128004), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, uint64(0x00000100), val.FieldByName("messageSeverity").Uint()) // VK_DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT
			require.Equal(t, uint64(0x00000002), val.FieldByName("messageType").Uint())     // VK_DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT
			require.False(t, val.FieldByName("pfnUserCallback").IsNil())

			callbackHandle := cgo.Handle(val.FieldByName("pUserData").UnsafePointer())
			callback := callbackHandle.Value().(ext_debug_utils.CallbackFunction)
			require.True(t, callback(0, 0, nil))
			require.True(t, calledRightFunction)

			return core1_0.VKSuccess, nil
		})

	messenger, _, err := extension.CreateDebugUtilsMessenger(instance, nil, ext_debug_utils.DebugUtilsMessengerCreateInfo{
		MessageSeverity: ext_debug_utils.SeverityWarning,
		MessageType:     ext_debug_utils.TypeValidation,
		UserCallback:    cb,
	})
	require.NoError(t, err)
	require.NotNil(t, messenger)
	require.Equal(t, expectedMessenger.Handle(), messenger.Handle())
}

func TestVulkanExtension_CmdBeginLabel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	debugDriver := mock_debugutils.NewMockDriver(ctrl)
	extension := ext_debug_utils.CreateExtensionFromDriver(debugDriver)
	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	label := ext_debug_utils.DebugUtilsLabel{
		LabelName: "someLabel",
		Color:     colornames.Blue,
	}

	debugDriver.EXPECT().VKCmdBeginDebugUtilsLabelEXT(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, pLabel *ext_debug_utils_driver.VkDebugUtilsLabelEXT) {
			val := reflect.ValueOf(*pLabel)
			require.Equal(t, uint64(1000128002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
			require.True(t, val.FieldByName("pNext").IsNil())

			expectedStr := strToCharSlice("someLabel")
			actualStr := cStringToCharSlice(unsafe.Pointer(val.FieldByName("pLabelName").Elem().UnsafeAddr()))
			require.Equal(t, expectedStr, actualStr)

			require.InDelta(t, float32(0), val.FieldByName("color").Index(0).Float(), 0.002)
			require.InDelta(t, float32(0), val.FieldByName("color").Index(1).Float(), 0.002)
			require.InDelta(t, float32(1), val.FieldByName("color").Index(2).Float(), 0.002)
			require.InDelta(t, float32(1), val.FieldByName("color").Index(3).Float(), 0.002)
		})

	err := extension.CmdBeginDebugUtilsLabel(commandBuffer, label)
	require.NoError(t, err)
}

func TestVulkanExtension_CmdEndLabel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	debugDriver := mock_debugutils.NewMockDriver(ctrl)
	extension := ext_debug_utils.CreateExtensionFromDriver(debugDriver)
	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	debugDriver.EXPECT().VkCmdEndDebugUtilsLabelEXT(commandBuffer.Handle())

	extension.CmdEndDebugUtilsLabel(commandBuffer)
}

func TestVulkanExtension_CmdInsertLabel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	debugDriver := mock_debugutils.NewMockDriver(ctrl)
	extension := ext_debug_utils.CreateExtensionFromDriver(debugDriver)
	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	label := ext_debug_utils.DebugUtilsLabel{
		LabelName: "otherLabel",
		Color:     colornames.Gray,
	}

	debugDriver.EXPECT().VkCmdInsertDebugUtilsLabelEXT(
		commandBuffer.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(commandBuffer driver.VkCommandBuffer, pLabel *ext_debug_utils_driver.VkDebugUtilsLabelEXT) {
			val := reflect.ValueOf(*pLabel)
			require.Equal(t, uint64(1000128002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
			require.True(t, val.FieldByName("pNext").IsNil())

			expectedStr := strToCharSlice("otherLabel")
			actualStr := cStringToCharSlice(unsafe.Pointer(val.FieldByName("pLabelName").Elem().UnsafeAddr()))
			require.Equal(t, expectedStr, actualStr)

			require.InDelta(t, float32(0.5), val.FieldByName("color").Index(0).Float(), 0.002)
			require.InDelta(t, float32(0.5), val.FieldByName("color").Index(1).Float(), 0.002)
			require.InDelta(t, float32(0.5), val.FieldByName("color").Index(2).Float(), 0.002)
			require.InDelta(t, float32(1), val.FieldByName("color").Index(3).Float(), 0.002)
		})

	err := extension.CmdInsertDebugUtilsLabel(commandBuffer, label)
	require.NoError(t, err)
}

func TestVulkanExtension_QueueBeginLabel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	debugDriver := mock_debugutils.NewMockDriver(ctrl)
	extension := ext_debug_utils.CreateExtensionFromDriver(debugDriver)
	queue := mocks.EasyMockQueue(ctrl)

	label := ext_debug_utils.DebugUtilsLabel{
		LabelName: "someLabel",
		Color:     colornames.Saddlebrown,
	}

	debugDriver.EXPECT().VkQueueBeginDebugUtilsLabelEXT(
		queue.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(queue driver.VkQueue, pLabel *ext_debug_utils_driver.VkDebugUtilsLabelEXT) {
			val := reflect.ValueOf(*pLabel)
			require.Equal(t, uint64(1000128002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
			require.True(t, val.FieldByName("pNext").IsNil())

			expectedStr := strToCharSlice("someLabel")
			actualStr := cStringToCharSlice(unsafe.Pointer(val.FieldByName("pLabelName").Elem().UnsafeAddr()))
			require.Equal(t, expectedStr, actualStr)

			require.InDelta(t, float32(0.545), val.FieldByName("color").Index(0).Float(), 0.002)
			require.InDelta(t, float32(0.271), val.FieldByName("color").Index(1).Float(), 0.002)
			require.InDelta(t, float32(0.075), val.FieldByName("color").Index(2).Float(), 0.002)
			require.InDelta(t, float32(1), val.FieldByName("color").Index(3).Float(), 0.002)
		})

	err := extension.QueueBeginDebugUtilsLabel(queue, label)
	require.NoError(t, err)
}

func TestVulkanExtension_QueueEndLabel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	debugDriver := mock_debugutils.NewMockDriver(ctrl)
	extension := ext_debug_utils.CreateExtensionFromDriver(debugDriver)
	queue := mocks.EasyMockQueue(ctrl)

	debugDriver.EXPECT().VkQueueEndDebugUtilsLabelEXT(queue.Handle())

	extension.QueueEndDebugUtilsLabel(queue)
}

func TestVulkanExtension_QueueInsertLabel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	debugDriver := mock_debugutils.NewMockDriver(ctrl)
	extension := ext_debug_utils.CreateExtensionFromDriver(debugDriver)
	queue := mocks.EasyMockQueue(ctrl)

	label := ext_debug_utils.DebugUtilsLabel{
		LabelName: "",
		Color:     color.RGBA{A: 0, B: 0, G: 0, R: 0},
	}

	debugDriver.EXPECT().VkQueueInsertDebugUtilsLabelEXT(
		queue.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(queue driver.VkQueue, pLabel *ext_debug_utils_driver.VkDebugUtilsLabelEXT) {
			val := reflect.ValueOf(*pLabel)
			require.Equal(t, uint64(1000128002), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
			require.True(t, val.FieldByName("pNext").IsNil())

			expectedStr := strToCharSlice("")
			actualStr := cStringToCharSlice(unsafe.Pointer(val.FieldByName("pLabelName").Elem().UnsafeAddr()))
			require.Equal(t, expectedStr, actualStr)

			require.InDelta(t, float32(0), val.FieldByName("color").Index(0).Float(), 0.002)
			require.InDelta(t, float32(0), val.FieldByName("color").Index(1).Float(), 0.002)
			require.InDelta(t, float32(0), val.FieldByName("color").Index(2).Float(), 0.002)
			require.InDelta(t, float32(0), val.FieldByName("color").Index(3).Float(), 0.002)
		})

	err := extension.QueueInsertDebugUtilsLabel(queue, label)
	require.NoError(t, err)
}

func TestVulkanExtension_SetObjectName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	debugDriver := mock_debugutils.NewMockDriver(ctrl)
	extension := ext_debug_utils.CreateExtensionFromDriver(debugDriver)
	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	commandBuffer := mocks.EasyMockCommandBuffer(ctrl)

	objectName := ext_debug_utils.DebugUtilsObjectNameInfo{
		ObjectName:   "someCommandBuffer",
		ObjectHandle: uintptr(unsafe.Pointer(commandBuffer.Handle())),
		ObjectType:   core1_0.ObjectTypeCommandBuffer,
	}

	debugDriver.EXPECT().VkSetDebugUtilsObjectNameEXT(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice, pNameInfo *ext_debug_utils_driver.VkDebugUtilsObjectNameInfoEXT) (common.VkResult, error) {
			val := reflect.ValueOf(*pNameInfo)

			require.Equal(t, uint64(1000128000), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(6), val.FieldByName("objectType").Uint()) // VK_OBJECT_TYPE_COMMAND_BUFFER
			require.Equal(t, uint64(uintptr(unsafe.Pointer(commandBuffer.Handle()))), val.FieldByName("objectHandle").Uint())

			expectedStr := strToCharSlice("someCommandBuffer")
			actualStr := cStringToCharSlice(unsafe.Pointer(val.FieldByName("pObjectName").Elem().UnsafeAddr()))
			require.Equal(t, expectedStr, actualStr)

			return core1_0.VKSuccess, nil
		})

	_, err := extension.SetDebugUtilsObjectName(device, objectName)
	require.NoError(t, err)
}

func TestVulkanExtension_SetObjectTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	debugDriver := mock_debugutils.NewMockDriver(ctrl)
	extension := ext_debug_utils.CreateExtensionFromDriver(debugDriver)
	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	device := mocks.EasyMockDevice(ctrl, coreDriver)
	queryPool := mocks.EasyMockQueryPool(ctrl)

	objectTag := ext_debug_utils.DebugUtilsObjectTagInfo{
		ObjectType:   core1_0.ObjectTypeQueryPool,
		ObjectHandle: uintptr(unsafe.Pointer(queryPool.Handle())),
		TagName:      53,
		Tag:          []byte("some tag data"),
	}

	debugDriver.EXPECT().VkSetDebugUtilsObjectTagEXT(
		device.Handle(),
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(device driver.VkDevice, pTagInfo *ext_debug_utils_driver.VkDebugUtilsObjectTagInfoEXT) (common.VkResult, error) {
			val := reflect.ValueOf(*pTagInfo)

			require.Equal(t, uint64(1000128001), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_TAG_INFO_EXT
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(12), val.FieldByName("objectType").Uint()) // VK_OBJECT_TYPE_QUERY_POOL
			require.Equal(t, uint64(uintptr(unsafe.Pointer(queryPool.Handle()))), val.FieldByName("objectHandle").Uint())
			require.Equal(t, uint64(53), val.FieldByName("tagName").Uint())
			require.Equal(t, uint64(13), val.FieldByName("tagSize").Uint())

			tagPtr := val.FieldByName("pTag").UnsafePointer()
			tagBytes := ([]byte)(unsafe.Slice((*byte)(tagPtr), 13))
			require.Equal(t, []byte("some tag data"), tagBytes)

			return core1_0.VKSuccess, nil
		})

	_, err := extension.SetDebugUtilsObjectTag(device, objectTag)
	require.NoError(t, err)
}

func TestVulkanExtension_SubmitMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	debugDriver := mock_debugutils.NewMockDriver(ctrl)
	extension := ext_debug_utils.CreateExtensionFromDriver(debugDriver)
	coreDriver := mock_driver.DriverForVersion(ctrl, common.Vulkan1_0)
	instance := mocks.EasyMockInstance(ctrl, coreDriver)
	pipeline := mocks.EasyMockPipeline(ctrl)

	callbackData := ext_debug_utils.DebugUtilsMessengerCallbackData{
		MessageIDNumber: 3,
		MessageIDName:   "some message id",
		Message:         "a cool message",
		QueueLabels: []ext_debug_utils.DebugUtilsLabel{
			{LabelName: "queue label 1", Color: colornames.Red},
			{LabelName: "queue label 2", Color: colornames.Blue},
		},
		CmdBufLabels: []ext_debug_utils.DebugUtilsLabel{
			{LabelName: "cmd label 1", Color: colornames.Green},
			{LabelName: "cmd label 2", Color: colornames.Yellow},
			{LabelName: "cmd label 3", Color: colornames.White},
		},
		Objects: []ext_debug_utils.DebugUtilsObjectNameInfo{
			{
				ObjectName:   "a object",
				ObjectType:   core1_0.ObjectTypePipeline,
				ObjectHandle: uintptr(unsafe.Pointer(pipeline.Handle())),
			},
		},
	}

	debugDriver.EXPECT().VkSubmitDebugUtilsMessageEXT(
		instance.Handle(),
		ext_debug_utils_driver.VkDebugUtilsMessageSeverityFlagBitsEXT(0x00001000), // VK_DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT
		ext_debug_utils_driver.VkDebugUtilsMessageTypeFlagsEXT(0x00000002),        // VK_DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT
		gomock.Not(gomock.Nil()),
	).DoAndReturn(
		func(instance driver.VkInstance,
			severity ext_debug_utils_driver.VkDebugUtilsMessageSeverityFlagBitsEXT,
			messageType ext_debug_utils_driver.VkDebugUtilsMessageTypeFlagsEXT,
			pCallbackData *ext_debug_utils_driver.VkDebugUtilsMessengerCallbackDataEXT) {

			val := reflect.ValueOf(*pCallbackData)
			require.Equal(t, uint64(1000128003), val.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CALLBACK_DATA_EXT
			require.True(t, val.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(0), val.FieldByName("flags").Uint())
			require.Equal(t, int64(3), val.FieldByName("messageIdNumber").Int())

			expectedMessageId := strToCharSlice("some message id")
			actualMessageId := cStringToCharSlice(unsafe.Pointer(val.FieldByName("pMessageIdName").Elem().UnsafeAddr()))
			require.Equal(t, expectedMessageId, actualMessageId)

			expectedMessage := strToCharSlice("a cool message")
			actualMessage := cStringToCharSlice(unsafe.Pointer(val.FieldByName("pMessage").Elem().UnsafeAddr()))
			require.Equal(t, expectedMessage, actualMessage)

			require.Equal(t, uint64(2), val.FieldByName("queueLabelCount").Uint())
			require.Equal(t, uint64(3), val.FieldByName("cmdBufLabelCount").Uint())
			require.Equal(t, uint64(1), val.FieldByName("objectCount").Uint())

			// Queue Labels
			queueLabelsPtr := (*ext_debug_utils_driver.VkDebugUtilsLabelEXT)(unsafe.Pointer(val.FieldByName("pQueueLabels").Elem().UnsafeAddr()))
			queueLabelSlice := ([]ext_debug_utils_driver.VkDebugUtilsLabelEXT)(unsafe.Slice(queueLabelsPtr, 2))
			labels := reflect.ValueOf(queueLabelSlice)

			// Queue Label 1
			label := labels.Index(0)
			require.Equal(t, uint64(1000128002), label.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
			require.True(t, label.FieldByName("pNext").IsNil())

			expectedLabelName := strToCharSlice("queue label 1")
			actualLabelName := cStringToCharSlice(unsafe.Pointer(label.FieldByName("pLabelName").Elem().UnsafeAddr()))
			require.Equal(t, expectedLabelName, actualLabelName)

			require.InDelta(t, float64(1), label.FieldByName("color").Index(0).Float(), 0.002)
			require.InDelta(t, float64(0), label.FieldByName("color").Index(1).Float(), 0.002)
			require.InDelta(t, float64(0), label.FieldByName("color").Index(2).Float(), 0.002)
			require.InDelta(t, float64(1), label.FieldByName("color").Index(3).Float(), 0.002)

			// Queue Label 2
			label = labels.Index(1)
			require.Equal(t, uint64(1000128002), label.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
			require.True(t, label.FieldByName("pNext").IsNil())

			expectedLabelName = strToCharSlice("queue label 2")
			actualLabelName = cStringToCharSlice(unsafe.Pointer(label.FieldByName("pLabelName").Elem().UnsafeAddr()))
			require.Equal(t, expectedLabelName, actualLabelName)

			require.InDelta(t, float64(0), label.FieldByName("color").Index(0).Float(), 0.002)
			require.InDelta(t, float64(0), label.FieldByName("color").Index(1).Float(), 0.002)
			require.InDelta(t, float64(1), label.FieldByName("color").Index(2).Float(), 0.002)
			require.InDelta(t, float64(1), label.FieldByName("color").Index(3).Float(), 0.002)

			// CmdBufLabels
			cmdBufLabelsPtr := (*ext_debug_utils_driver.VkDebugUtilsLabelEXT)(unsafe.Pointer(val.FieldByName("pCmdBufLabels").Elem().UnsafeAddr()))
			cmdBufLabelSlice := ([]ext_debug_utils_driver.VkDebugUtilsLabelEXT)(unsafe.Slice(cmdBufLabelsPtr, 3))
			labels = reflect.ValueOf(cmdBufLabelSlice)

			// Cmd Label 1
			label = labels.Index(0)
			require.Equal(t, uint64(1000128002), label.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
			require.True(t, label.FieldByName("pNext").IsNil())

			expectedLabelName = strToCharSlice("cmd label 1")
			actualLabelName = cStringToCharSlice(unsafe.Pointer(label.FieldByName("pLabelName").Elem().UnsafeAddr()))
			require.Equal(t, expectedLabelName, actualLabelName)

			require.InDelta(t, float64(0), label.FieldByName("color").Index(0).Float(), 0.002)
			require.InDelta(t, 0.5, label.FieldByName("color").Index(1).Float(), 0.002)
			require.InDelta(t, float64(0), label.FieldByName("color").Index(2).Float(), 0.002)
			require.InDelta(t, float64(1), label.FieldByName("color").Index(3).Float(), 0.002)

			// Cmd Label 2
			label = labels.Index(1)
			require.Equal(t, uint64(1000128002), label.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
			require.True(t, label.FieldByName("pNext").IsNil())

			expectedLabelName = strToCharSlice("cmd label 2")
			actualLabelName = cStringToCharSlice(unsafe.Pointer(label.FieldByName("pLabelName").Elem().UnsafeAddr()))
			require.Equal(t, expectedLabelName, actualLabelName)

			require.InDelta(t, float64(1), label.FieldByName("color").Index(0).Float(), 0.002)
			require.InDelta(t, float64(1), label.FieldByName("color").Index(1).Float(), 0.002)
			require.InDelta(t, float64(0), label.FieldByName("color").Index(2).Float(), 0.002)
			require.InDelta(t, float64(1), label.FieldByName("color").Index(3).Float(), 0.002)

			// Cmd Label 3
			label = labels.Index(2)
			require.Equal(t, uint64(1000128002), label.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_LABEL_EXT
			require.True(t, label.FieldByName("pNext").IsNil())

			expectedLabelName = strToCharSlice("cmd label 3")
			actualLabelName = cStringToCharSlice(unsafe.Pointer(label.FieldByName("pLabelName").Elem().UnsafeAddr()))
			require.Equal(t, expectedLabelName, actualLabelName)

			require.InDelta(t, float64(1), label.FieldByName("color").Index(0).Float(), 0.002)
			require.InDelta(t, float64(1), label.FieldByName("color").Index(1).Float(), 0.002)
			require.InDelta(t, float64(1), label.FieldByName("color").Index(2).Float(), 0.002)
			require.InDelta(t, float64(1), label.FieldByName("color").Index(3).Float(), 0.002)

			// Objects
			objectNamePtr := (*ext_debug_utils_driver.VkDebugUtilsObjectNameInfoEXT)(unsafe.Pointer(val.FieldByName("pObjects").Elem().UnsafeAddr()))
			objectName := reflect.ValueOf(*objectNamePtr)

			require.Equal(t, uint64(1000128000), objectName.FieldByName("sType").Uint()) // VK_STRUCTURE_TYPE_DEBUG_UTILS_OBJECT_NAME_INFO_EXT
			require.True(t, objectName.FieldByName("pNext").IsNil())
			require.Equal(t, uint64(19), objectName.FieldByName("objectType").Uint()) // VK_OBJECT_TYPE_PIPELINE
			require.Equal(t, uint64(uintptr(unsafe.Pointer(pipeline.Handle()))), objectName.FieldByName("objectHandle").Uint())

			expectedObjectName := strToCharSlice("a object")
			actualObjectName := cStringToCharSlice(unsafe.Pointer(objectName.FieldByName("pObjectName").Elem().UnsafeAddr()))
			require.Equal(t, expectedObjectName, actualObjectName)
		})

	err := extension.SubmitDebugUtilsMessage(instance, ext_debug_utils.SeverityError, ext_debug_utils.TypeValidation, callbackData)
	require.NoError(t, err)
}
