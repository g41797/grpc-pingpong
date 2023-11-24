package server

import (
	"fmt"
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

// check implementation:
// - https://github.com/omri86/longlived-grpc/blob/master/server/server.go

func (srv *consumer) Consume(c pb.AdapterService_ConsumeServer) error {

	mc, err := srv.bc.connect()
	if err != nil {
		return err
	}
	srv.mc = mc

	for {
		next, err := c.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		if stop := next.GetStop(); stop != nil {
			srv.abortConsume()
			return nil
		}

		start := next.GetStart()

		if start == nil {
			return fmt.Errorf("wrong request")
		}

		if srv.started {
			return fmt.Errorf("already started")
		}

		err = srv.startConsume(start)
		if err != nil {
			return err
		}

		srv.started = true
		continue
	}
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
