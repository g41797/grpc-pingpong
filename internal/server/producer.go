package server

import (
	"io"

	"github.com/g41797/memphisgrpc/pb"
	"github.com/memphisdev/memphis.go"
)

type producer struct {
	bc      *brokerConnector
	started bool
	stream  pb.AdapterService_ProduceServer
	mc      *memphis.Conn
	mpr     *memphis.Producer
}

func newProducer(bc *brokerConnector) *producer {
	producer := new(producer)
	producer.bc = bc
	return producer
}

func (srv *producer) Produce(stream pb.AdapterService_ProduceServer) error {

	srv.stream = stream
	status := pb.Status{}

	mc, err := srv.bc.connect()
	if err != nil {
		text := err.Error()
		status.Text = &text
		srv.finish(&status)
		return nil
	}

	srv.mc = mc

	for {
		next, err := stream.Recv()

		if err == io.EOF {
			srv.finish(&pb.Status{})
			return nil
		}
		if err != nil {
			return err
		}

		stop := next.GetStop()
		if stop != nil {
			srv.finish(&pb.Status{})
			return nil
		}

		start := next.GetStart()
		if start != nil {

			if srv.started {
				status := pb.Status{}
				text := "already started"
				status.Text = &text
				srv.finish(&status)
				return nil
			}

			if err := srv.createProducer(start); err != nil {
				text := err.Error()
				status.Text = &text
				srv.finish(&status)
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
				srv.finish(&status)
				return nil
			}

			if err := srv.produce(msg); err != nil {
				text := err.Error()
				status.Text = &text
				srv.finish(&status)
				return nil
			}
			continue
		}

		status := pb.Status{}
		text := "wrong client request"
		status.Text = &text
		srv.finish(&status)
		break
	}

	return nil
}

func (srv *producer) createProducer(start *pb.CreateProducerRequest) error {

	station := start.Station.GetName()

	st, err := srv.mc.CreateStation(station)
	if err != nil {
		return err
	}

	producer := start.Producer.GetName()

	prod, err := st.CreateProducer(producer)
	if err != nil {
		return err
	}

	srv.mpr = prod

	return nil
}

func (srv *producer) produce(msg *pb.Msg) error {

	hdrs := memphis.Headers{}
	hdrs.New()

	h := msg.GetHeaders()

	if h != nil {
		hh := h.GetHeaders()
		for k, v := range hh {
			hdrs.Add(k, v)
		}
	}

	body := msg.GetBody()

	err := srv.mpr.Produce(body, memphis.MsgHeaders(hdrs), memphis.AsyncProduce())

	return err
}

func (srv *producer) finish(status *pb.Status) {
	if srv == nil {
		return
	}

	if srv.stream == nil {
		return
	}

	srv.stream.SendAndClose(status)
}

func (srv *producer) clean() {
	if srv == nil {
		return
	}

	if srv.mpr != nil {
		srv.mpr.Destroy()
	}

	if srv.mc != nil {
		srv.mc.Close()
	}
}
