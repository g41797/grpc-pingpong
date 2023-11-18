package plain

import "context"

type ProducerService interface {
	Service
	Attach(sendandclose func(*Status) error)
	Produce(ctx context.Context, prm *ProduceMessages) error
}
