package plain

import "context"

type ProducerService interface {
	Service
	Produce(ctx context.Context, prm *ProduceMessages) error
	Close() *Status
}
