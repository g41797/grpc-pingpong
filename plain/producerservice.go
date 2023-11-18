package plain

import "context"

type ProducerService interface {
	Attach(sendandclose func(*Status) error)
	Produce(ctx context.Context, prm *ProduceMessages) error
}
