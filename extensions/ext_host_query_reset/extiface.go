package ext_host_query_reset

import "github.com/CannibalVox/VKng/core/core1_0"

//go:generate mockgen -source extiface.go -destination ./mocks/extension.go -package mock_host_query_reset

type Extension interface {
	ResetQueryPool(device core1_0.Device, queryPool core1_0.QueryPool, firstQuery, queryCount int)
}
