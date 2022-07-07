package ext_debug_utils

import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/VKng/core/driver"
	ext_driver "github.com/CannibalVox/VKng/extensions/ext_debug_utils/driver"
	"github.com/CannibalVox/cgoparam"
)

//go:generate mockgen -source extension.go -destination ./mocks/extension.go -package mock_debugutils

type VulkanExtension struct {
	driver ext_driver.Driver
}

type Extension interface {
	CreateDebugUtilsMessenger(instance core1_0.Instance, allocation *driver.AllocationCallbacks, o DebugUtilsMessengerCreateInfo) (Messenger, common.VkResult, error)

	CmdBeginDebugUtilsLabel(commandBuffer core1_0.CommandBuffer, label DebugUtilsLabel) error
	CmdEndDebugUtilsLabel(commandBuffer core1_0.CommandBuffer)
	CmdInsertDebugUtilsLabel(commandBuffer core1_0.CommandBuffer, label DebugUtilsLabel) error

	QueueBeginDebugUtilsLabel(queue core1_0.Queue, label DebugUtilsLabel) error
	QueueEndDebugUtilsLabel(queue core1_0.Queue)
	QueueInsertDebugUtilsLabel(queue core1_0.Queue, label DebugUtilsLabel) error

	SetDebugUtilsObjectName(instance core1_0.Device, name DebugUtilsObjectNameInfo) (common.VkResult, error)
	SetDebugUtilsObjectTag(instance core1_0.Device, tag DebugUtilsObjectTagInfo) (common.VkResult, error)

	SubmitDebugUtilsMessage(instance core1_0.Instance, severity MessageSeverities, types MessageTypes, data DebugUtilsMessengerCallbackData) error
}

func CreateExtensionFromInstance(instance core1_0.Instance) *VulkanExtension {
	driver := ext_driver.CreateDriverFromCore(instance.Driver())

	if !instance.IsInstanceExtensionActive(ExtensionName) {
		return nil
	}

	return CreateExtensionFromDriver(driver)
}

func CreateExtensionFromDriver(driver ext_driver.Driver) *VulkanExtension {
	return &VulkanExtension{
		driver: driver,
	}
}

func (l *VulkanExtension) CreateDebugUtilsMessenger(instance core1_0.Instance, allocation *driver.AllocationCallbacks, o DebugUtilsMessengerCreateInfo) (Messenger, common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	createInfo, err := common.AllocOptions(arena, o)
	if err != nil {
		return nil, core1_0.VKErrorUnknown, err
	}

	var messenger ext_driver.VkDebugUtilsMessengerEXT
	res, err := l.driver.VkCreateDebugUtilsMessengerEXT(instance.Handle(), (*ext_driver.VkDebugUtilsMessengerCreateInfoEXT)(createInfo), allocation.Handle(), &messenger)

	if err != nil {
		return nil, res, err
	}

	coreDriver := instance.Driver()
	newMessenger := coreDriver.ObjectStore().GetOrCreate(driver.VulkanHandle(messenger), driver.Core1_0, func() any {
		return &vulkanMessenger{
			coreDriver: coreDriver,
			handle:     messenger,
			instance:   instance.Handle(),
			driver:     l.driver,
		}
	}).(*vulkanMessenger)

	return newMessenger, res, nil
}

func (l *VulkanExtension) CmdBeginDebugUtilsLabel(commandBuffer core1_0.CommandBuffer, label DebugUtilsLabel) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.driver.VKCmdBeginDebugUtilsLabelEXT(commandBuffer.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtension) CmdEndDebugUtilsLabel(buffer core1_0.CommandBuffer) {
	l.driver.VkCmdEndDebugUtilsLabelEXT(buffer.Handle())
}

func (l *VulkanExtension) CmdInsertDebugUtilsLabel(buffer core1_0.CommandBuffer, label DebugUtilsLabel) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.driver.VkCmdInsertDebugUtilsLabelEXT(buffer.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtension) QueueBeginDebugUtilsLabel(queue core1_0.Queue, label DebugUtilsLabel) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.driver.VkQueueBeginDebugUtilsLabelEXT(queue.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtension) QueueEndDebugUtilsLabel(queue core1_0.Queue) {
	l.driver.VkQueueEndDebugUtilsLabelEXT(queue.Handle())
}

func (l *VulkanExtension) QueueInsertDebugUtilsLabel(queue core1_0.Queue, label DebugUtilsLabel) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	labelPtr, err := common.AllocOptions(arena, label)
	if err != nil {
		return err
	}

	l.driver.VkQueueInsertDebugUtilsLabelEXT(queue.Handle(), (*ext_driver.VkDebugUtilsLabelEXT)(labelPtr))

	return nil
}

func (l *VulkanExtension) SetDebugUtilsObjectName(device core1_0.Device, name DebugUtilsObjectNameInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	namePtr, err := common.AllocOptions(arena, name)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return l.driver.VkSetDebugUtilsObjectNameEXT(device.Handle(), (*ext_driver.VkDebugUtilsObjectNameInfoEXT)(namePtr))
}

func (l *VulkanExtension) SetDebugUtilsObjectTag(device core1_0.Device, tag DebugUtilsObjectTagInfo) (common.VkResult, error) {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	tagPtr, err := common.AllocOptions(arena, tag)
	if err != nil {
		return core1_0.VKErrorUnknown, err
	}

	return l.driver.VkSetDebugUtilsObjectTagEXT(device.Handle(), (*ext_driver.VkDebugUtilsObjectTagInfoEXT)(tagPtr))
}

func (l *VulkanExtension) SubmitDebugUtilsMessage(instance core1_0.Instance, severity MessageSeverities, types MessageTypes, data DebugUtilsMessengerCallbackData) error {
	arena := cgoparam.GetAlloc()
	defer cgoparam.ReturnAlloc(arena)

	callbackPtr, err := common.AllocOptions(arena, &data)
	if err != nil {
		return err
	}

	l.driver.VkSubmitDebugUtilsMessageEXT(instance.Handle(),
		ext_driver.VkDebugUtilsMessageSeverityFlagBitsEXT(severity),
		ext_driver.VkDebugUtilsMessageTypeFlagsEXT(types),
		(*ext_driver.VkDebugUtilsMessengerCallbackDataEXT)(callbackPtr))

	return nil
}
