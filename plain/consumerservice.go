package plain

import "context"

type ConsumerService interface {
	Attach(send func(m *ConsumeResponse) error)
	Process(ctx context.Context, cms *ConsumeMessages) error
}
