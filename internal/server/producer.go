package server

import (
	"io"

	"githib.com/g41797/grpcadapter/pb"
	"github.com/memphisdev/memphis.go"
)

type producer struct {
	bc      *brokerConnector
	mc      *memphis.Conn
	started bool
}

func newProducer(bc *brokerConnector) *producer {
	producer := new(producer)
	producer.bc = bc
	return producer
}

func (srv *producer) Produce(stream pb.AdapterService_ProduceServer) error {

	pm := prmemento{stream: stream}

	mc, err := srv.bc.connect()
	if err != nil {
		status := pb.Status{}
		text := err.Error()
		status.Text = &text
		pm.finish(&status)
		return nil
	}

	srv.mc = mc

	for {
		next, err := stream.Recv()

		if err == io.EOF {
			pm.finish(&pb.Status{})
			return nil
		}
		if err != nil {
			return err
		}

		start := next.GetStart()

		if start != nil {

			if srv.started {
				status := pb.Status{}
				text := "already started"
				status.Text = &text
				pm.finish(&status)
				return nil
			}

			if status := srv.createProducer(start); status != nil {

				pm.finish(status)
				return nil
			}
			srv.started = true
			continue
		}

		msg := next.GetMsg()

		if msg != nil {

			if !srv.started {
				status := pb.Status{}
				text := "not started yet"
				status.Text = &text
				pm.finish(&status)
				return nil
			}

			if status := pm.produce(msg); status != nil {
				pm.finish(status)
				return nil
			}
			continue
		}

		status := pb.Status{}
		text := "wrong client request"
		status.Text = &text
		pm.finish(&status)
		break
	}

	return nil
}

func (srv *producer) createProducer(start *pb.CreateProducerRequest) *pb.Status {
	status := pb.Status{}
	text := "start not implemented"
	status.Text = &text

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

type prmemento struct {
	stream pb.AdapterService_ProduceServer
	mpr    *memphis.Producer
}

func (pm *prmemento) produce(msg *pb.Msg) *pb.Status {
	status := pb.Status{}
	text := "produce not implemented"
	status.Text = &text

	return &status
}

func (pm *prmemento) finish(status *pb.Status) {
	if pm == nil {
		return
	}

	if pm.stream == nil {
		return
	}

	pm.stream.SendAndClose(status)

	if pm.mpr != nil {
		pm.mpr.Destroy()
	}
}
