package VKng

/*
#include <stdlib.h>
#include "vulkan/vulkan.h"
*/
import "C"

type QueueHandle C.VkQueue
type Queue struct {
	handle QueueHandle
}

func (q *Queue) Handle() QueueHandle {
	return q.handle
}
