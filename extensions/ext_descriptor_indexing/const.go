package ext_descriptor_indexing

/*
#include <stdlib.h>
#include "../vulkan/vulkan.h"
*/
import "C"
import (
	"github.com/CannibalVox/VKng/core/common"
	"github.com/CannibalVox/VKng/core/core1_0"
)

type DescriptorBindingFlags int32

var descriptorBindingFlagsMapping = common.NewFlagStringMapping[DescriptorBindingFlags]()

func (f DescriptorBindingFlags) Register(str string) {
	descriptorBindingFlagsMapping.Register(f, str)
}
func (f DescriptorBindingFlags) String() string {
	return descriptorBindingFlagsMapping.FlagsToString(f)
}

////

const (
	ExtensionName string = C.VK_EXT_DESCRIPTOR_INDEXING_EXTENSION_NAME

	DescriptorBindingPartiallyBound           DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT
	DescriptorBindingUpdateAfterBind          DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT
	DescriptorBindingUpdateUnusedWhilePending DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT_EXT
	DescriptorBindingVariableDescriptorCount  DescriptorBindingFlags = C.VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT

	DescriptorPoolCreateUpdateAfterBind core1_0.DescriptorPoolCreateFlags = C.VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT

	DescriptorSetLayoutCreateUpdateAfterBindPool core1_0.DescriptorSetLayoutCreateFlags = C.VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT

	VkErrorFragmentation common.VkResult = C.VK_ERROR_FRAGMENTATION_EXT
)

func init() {
	DescriptorBindingPartiallyBound.Register("Partially-Bound")
	DescriptorBindingUpdateAfterBind.Register("Update After Bind")
	DescriptorBindingUpdateUnusedWhilePending.Register("Update Unused While Pending")
	DescriptorBindingVariableDescriptorCount.Register("Variable Descriptor Count")

	DescriptorPoolCreateUpdateAfterBind.Register("Update After Bind")

	DescriptorSetLayoutCreateUpdateAfterBindPool.Register("Update After Bind Pool")

	VkErrorFragmentation.Register("fragmentation")
}
