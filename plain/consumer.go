package plain

type Consumer struct {
	Name string
}

type ConsumerOpions struct{}
type ConsumingOpions struct{}

type CreateConsumerRequest struct {
	Station  *Station
	Consumer *Consumer
	Options  *ConsumerOpions
}

type DestroyConsumerRequest struct {
	Station  *Station
	Consumer *Consumer
}

type ConsumeRequest struct {
	Station  *Station
	Consumer *Consumer
	Options  *ConsumingOpions
}

type ConsumeMessages struct {
	//	*ConsumeRequest
	//	*Status
	//	*Stop
	Data any
}

type ConsumeResponse struct {
	//	*Msg
	//	*Status
	//	*Stop
	Data any
}
