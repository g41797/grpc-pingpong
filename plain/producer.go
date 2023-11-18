package plain

type Producer struct {
	Name string
}

type ProducerOpions struct{}

type CreateProducerRequest struct {
	Station  *Station
	Producer *Producer
	Options  *ProducerOpions
}

type DestroyProducerRequest struct {
	Station  *Station
	Producer *Producer
}

type ProduceRequest struct {
	Station  *Station
	Producer *Producer
}

type ProduceMessages struct {
	//	*ProduceRequest
	//	*Msg
	//	*Stop
	Data any
}
