package server

import (
	"io"

	"githib.com/g41797/grpcadapter/pb"
	"github.com/memphisdev/memphis.go"
)

type producer struct {
	bc       *brokerConnector
	mc       *memphis.Conn
	station  string
	producer string
	started  bool
}

func newProducer(bc *brokerConnector) *producer {
	producer := new(producer)
	producer.bc = bc
	return producer
}

func (srv *producer) Produce(stream pb.AdapterService_ProduceServer) error {

	mc, err := srv.bc.connect()
	if err != nil {
		status := pb.Status{}
		*status.Text = err.Error()
		stream.SendAndClose(&status)
		return nil
	}

	srv.mc = mc

	for {
		next, err := stream.Recv()

		if err == io.EOF {
			stream.SendAndClose(&pb.Status{})
			return nil
		}
		if err != nil {
			return err
		}

		if stop := next.GetStop(); stop != nil {
			stream.SendAndClose(&pb.Status{})
			return nil
		}

		start := next.GetStart()

		if start != nil {

			if srv.started {
				status := pb.Status{}
				*status.Text = "already started"
				stream.SendAndClose(&status)
				return nil
			}

			if status := srv.createProducer(start); status != nil {

				stream.SendAndClose(status)
				return nil
			}
			srv.started = true
			continue
		}

		msg := next.GetMsg()

		if msg != nil {

			if !srv.started {
				status := pb.Status{}
				*status.Text = "not started yet"
				stream.SendAndClose(&status)
				return nil
			}

			if status := srv.produce(msg); status != nil {
				stream.SendAndClose(status)
				return nil
			}
			continue
		}

		status := pb.Status{}
		*status.Text = "wrong client request"
		stream.SendAndClose(&status)
		break
	}

	return nil
}

func (srv *producer) createProducer(start *pb.CreateProducerRequest) *pb.Status {
	status := pb.Status{}
	*status.Text = "start not implemented"

	return &status
}

func (srv *producer) produce(msg *pb.Msg) *pb.Status {
	status := pb.Status{}
	*status.Text = "produce not implemented"

	return &status
}

func (srv *producer) clean() {
	if srv == nil {
		return
	}

	if srv.mc != nil {
		srv.mc.Close()
	}
}
