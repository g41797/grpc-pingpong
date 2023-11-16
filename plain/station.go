package plain

type Station struct {
	WasSetFlag
	Name string
}

type StorageType int32

const (
	Disk   StorageType = 0
	Memory StorageType = 1
)

type  MaxMessageAgeSecondsRet struct {
	WasSetFlag
	Seconds int32
}

type MessagesRet struct {
	WasSetFlag
	Messages int32
}

type BytesRet struct {
	WasSetFlag
	Bytes int32
}

type AckBasedRet {
	WasSetFlag
	AckBased bool
}

type RetentionOpt {
	WasSetFlag
	// *MaxMessageAgeSecondsRet
	// *MessagesRet
	// *BytesRet
	// *AckBasedRet
	Option any
}

type StorageOpt struct{
	WasSetFlag
	StorageType StorageType
}

// func (so *StorageOpt) Set(st int32) bool {
// 	if (st < 0) || (st > 1) {return false}
// 	return true
// }

type PartitionOpt struct {
	WasSetFlag
	Number			int32
}

type StationOpions  struct {
	WasSetFlag
	PartitionOpt	*PartitionOpt
    StorageOpt		*StorageOpt
    RetentionOpt	*RetentionOpt
}

type CreateStationRequest struct {
	WasSetFlag
	Station 		Station
    StationOpions	*StationOpions
}

type DestroyStationRequest struct {
	WasSetFlag
	Station 		Station
}
