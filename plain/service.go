package plain

type Service interface {
	IsAllowedToConnect() bool
	Connect() error
	Disconnect()
}
