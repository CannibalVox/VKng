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

type DescriptorTemplateType int32

var descriptorTemplateTypeMapping = make(map[DescriptorTemplateType]string)

func (e DescriptorTemplateType) Register(str string) {
	descriptorTemplateTypeMapping[e] = str
}

func (e DescriptorTemplateType) String() string {
	return descriptorTemplateTypeMapping[e]
}

////

type DescriptorTemplateFlags int32

var descriptorTemplateFlagsMapping = common.NewFlagStringMapping[DescriptorTemplateFlags]()

func (f DescriptorTemplateFlags) Register(str string) {
	descriptorTemplateFlagsMapping.Register(f, str)
}
func (f DescriptorTemplateFlags) String() string {
	return descriptorTemplateFlagsMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_KHR_DESCRIPTOR_UPDATE_TEMPLATE_EXTENSION_NAME

	DescriptorTemplateTypeDescriptorSet DescriptorTemplateType = C.VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET_KHR

	ObjectTypeDescriptorTemplate core1_0.ObjectType = C.VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR
)

func init() {
	DescriptorTemplateTypeDescriptorSet.Register("Descriptor Set")

	ObjectTypeDescriptorTemplate.Register("Descriptor Template")
}
