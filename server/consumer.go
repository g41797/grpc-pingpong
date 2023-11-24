package server

import (
	"io"

	"githib.com/g41797/grpcadapter/pb"
	"github.com/memphisdev/memphis.go"
)

type consumer struct {
	bc       *brokerConnector
	mc       *memphis.Conn
	station  string
	producer string
	options  struct{}
	started  bool
}

func newConsumer(bc *brokerConnector) *consumer {
	consumer := new(consumer)
	consumer.bc = bc
	return consumer
}

func (srv *consumer) Consume(c pb.ConsumerService_ConsumeServer) error {

	for {
		next, err := c.Recv()
		resp := pb.ConsumeResponse{}

		if err == io.EOF {
			resp.Data = &pb.ConsumeResponse_Stop{}
			srv.abortConsume()
			c.Send(&resp)
			return nil
		}

		if err != nil {
			status := pb.Status{}
			*status.Text = err.Error()
			resp.Data = &pb.ConsumeResponse_Error{Error: &status}
			srv.abortConsume()
			c.Send(&resp)
			return nil
		}

		if !srv.started {

		}

		if start := next.GetStart(); start != nil {
			if err := srv.startConsume(start); err != nil {
				status := pb.Status{}
				*status.Text = err.Error()
				resp.Data = &pb.ConsumeResponse_Error{Error: &status}
				c.Send(&resp)
				return nil
			}
			continue
		}

		resp.Data = &pb.ConsumeResponse_Error{Error: &pb.Status{}}
		c.Send(&resp)
		break
	}

	return nil
}

func (srv *consumer) startConsume(start *pb.ConsumeRequest) error {

	// TODO Add consume activation

	return nil
}

func (srv *consumer) abortConsume() error {
	if !srv.started {
		return nil
	}

	return nil
}

func (srv *consumer) clean() {

	srv.abortConsume()

	if srv == nil {
		return
	}

	if srv.mc != nil {
		srv.mc.Close()
	}

}
