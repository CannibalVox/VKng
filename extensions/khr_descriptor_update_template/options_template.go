package khr_descriptor_update_template

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
	"github.com/CannibalVox/cgoparam"
	"unsafe"
)

type DescriptorUpdateTemplateCreateOptions struct {
	Flags        DescriptorTemplateFlags
	Entries      []DescriptorUpdateTemplateEntry
	TemplateType DescriptorTemplateType

	DescriptorSetLayout core1_0.DescriptorSetLayout

	PipelineBindPoint common.PipelineBindPoint
	PipelineLayout    core1_0.PipelineLayout
	Set               int

	common.HaveNext
}

func (o DescriptorUpdateTemplateCreateOptions) PopulateCPointer(allocator *cgoparam.Allocator, preallocatedPointer unsafe.Pointer, next unsafe.Pointer) (unsafe.Pointer, error) {
	if preallocatedPointer == nil {
		preallocatedPointer = allocator.Malloc(int(unsafe.Sizeof(C.VkDescriptorUpdateTemplateCreateInfoKHR{})))
	}

	createInfo := (*C.VkDescriptorUpdateTemplateCreateInfoKHR)(preallocatedPointer)
	createInfo.sType = C.VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO_KHR
	createInfo.pNext = next
	createInfo.flags = C.VkDescriptorUpdateTemplateCreateFlags(o.Flags)

	entryCount := len(o.Entries)
	createInfo.descriptorUpdateEntryCount = C.uint32_t(entryCount)

	var err error
	createInfo.pDescriptorUpdateEntries, err = common.AllocSlice[C.VkDescriptorUpdateTemplateEntryKHR, DescriptorUpdateTemplateEntry](allocator, o.Entries)
	if err != nil {
		return nil, err
	}

	createInfo.templateType = C.VkDescriptorUpdateTemplateType(o.TemplateType)
	createInfo.descriptorSetLayout = nil
	createInfo.pipelineLayout = nil

	if o.DescriptorSetLayout != nil {
		createInfo.descriptorSetLayout = C.VkDescriptorSetLayout(unsafe.Pointer(o.DescriptorSetLayout.Handle()))
	}

	if o.PipelineLayout != nil {
		createInfo.pipelineLayout = C.VkPipelineLayout(unsafe.Pointer(o.PipelineLayout.Handle()))
	}

	createInfo.pipelineBindPoint = C.VkPipelineBindPoint(o.PipelineBindPoint)
	createInfo.set = C.uint32_t(o.Set)

	return preallocatedPointer, nil
}

func (o DescriptorUpdateTemplateCreateOptions) PopulateOutData(cDataPointer unsafe.Pointer, helpers ...any) (next unsafe.Pointer, err error) {
	createInfo := (*C.VkDescriptorUpdateTemplateCreateInfoKHR)(cDataPointer)
	return createInfo.pNext, nil
}
