package plain

type Station struct {
	Name string
}

type StorageType int32

const (
	Disk   StorageType = 0
	Memory StorageType = 1
)

type MaxMessageAgeSecondsRet struct {
	Seconds *int32
}

type MessagesRet struct {
	Messages *int32
}

type BytesRet struct {
	Bytes *int32
}

type AckBasedRet struct {
	AckBased *bool
}

type RetentionOpt struct {
	// *MaxMessageAgeSecondsRet
	// *MessagesRet
	// *BytesRet
	// *AckBasedRet
	Data any
}

type StorageOpt struct {
	StorageType *StorageType
}

type PartitionOpt struct {
	Number *int32
}

type StationOpions struct {
	PartitionOpt *PartitionOpt
	StorageOpt   *StorageOpt
	RetentionOpt *RetentionOpt
}

type CreateStationRequest struct {
	Station *Station
	Opions  *StationOpions
}

type DestroyStationRequest struct {
	Station *Station
}
