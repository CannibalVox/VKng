package objects

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type Queue struct {
	handle C.VkQueue
}
