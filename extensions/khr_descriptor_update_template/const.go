package khr_descriptor_update_template

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

type DescriptorUpdateTemplateType int32

var descriptorTemplateTypeMapping = make(map[DescriptorUpdateTemplateType]string)

func (e DescriptorUpdateTemplateType) Register(str string) {
	descriptorTemplateTypeMapping[e] = str
}

func (e DescriptorUpdateTemplateType) String() string {
	return descriptorTemplateTypeMapping[e]
}

////

type DescriptorUpdateTemplateFlags int32

var descriptorTemplateFlagsMapping = common.NewFlagStringMapping[DescriptorUpdateTemplateFlags]()

func (f DescriptorUpdateTemplateFlags) Register(str string) {
	descriptorTemplateFlagsMapping.Register(f, str)
}
func (f DescriptorUpdateTemplateFlags) String() string {
	return descriptorTemplateFlagsMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_DESCRIPTOR_UPDATE_TEMPLATE_EXTENSION_NAME

	DescriptorUpdateTemplateTypeDescriptorSet DescriptorUpdateTemplateType = C.VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET_KHR

	ObjectTypeDescriptorUpdateTemplate core1_0.ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR
)

func init() {
	DescriptorUpdateTemplateTypeDescriptorSet.Register("Descriptor Set")

	ObjectTypeDescriptorUpdateTemplate.Register("Descriptor Template")
}
