package VKng


type QueueFlags uint

const (
	Graphics QueueFlags = 0x01
	Compute QueueFlags = 0x02
	Transfer QueueFlags = 0x04
	SparseBinding QueueFlags = 0x08
)

type Extent3D struct {
	Width uint32
	Height uint32
	Depth uint32
}

type QueueFamily struct {
	Flags QueueFlags
	QueueCount uint32
	TimestampValidBits uint32
	MinImageTransferGranularity Extent3D
}
