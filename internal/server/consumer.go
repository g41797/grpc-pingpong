package server

import (
	"context"
	"fmt"
	"io"

	"githib.com/g41797/memphisgrpc/pb"
	"github.com/g41797/kissngoqueue"
	"github.com/memphisdev/memphis.go"
)

type consumer struct {
	bc      *brokerConnector
	started bool
	mc      *memphis.Conn
	stream  pb.AdapterService_ConsumeServer
	cons    *memphis.Consumer
	q       *kissngoqueue.Queue[any]
	done    chan struct{}
}

type wakeup struct{}

func newConsumer(bc *brokerConnector) *consumer {
	consumer := new(consumer)
	consumer.bc = bc
	consumer.q = kissngoqueue.NewQueue[any]()
	consumer.done = make(chan struct{}, 1)
	return consumer
}

// check implementation:
// - https://github.com/omri86/longlived-grpc/blob/master/server/server.go

func (srv *consumer) Consume(c pb.AdapterService_ConsumeServer) error {

	srv.stream = c

	mc, err := srv.bc.connect()
	if err != nil {
		return err
	}

	go srv.send()

	var reterr error

	srv.mc = mc

	for {
		next, err := c.Recv()

		if err != nil {
			if err == io.EOF {
				err = nil
			}

			reterr = err
			break
		}

		if stop := next.GetStop(); stop != nil {
			reterr = nil
			break
		}

		start := next.GetStart()

		if start == nil {
			reterr = fmt.Errorf("wrong request")
			break
		}

		if srv.started {
			reterr = fmt.Errorf("already started")
			break
		}

		err = srv.startConsume(start)
		if err != nil {
			reterr = err
			break
		}

		srv.started = true
		continue
	}

	srv.abortConsume(reterr)
	return reterr
}

func (srv *consumer) startConsume(start *pb.CreateConsumerRequest) error {

	station := start.Station.GetName()

	st, err := srv.mc.CreateStation(station)
	if err != nil {
		return err
	}

	consumer := start.Consumer.GetName()

	cons, err := st.CreateConsumer(consumer)
	if err != nil {
		return err
	}

	srv.cons = cons

	err = srv.cons.Consume(srv.processMessages)

	return err
}

func (srv *consumer) processMessages(msgs []*memphis.Msg, err error, ctx context.Context) {

	if err != nil {
		srv.q.PutMT(err)
		return
	}

	if msgs == nil {
		srv.q.PutMT(wakeup{})
		return
	}

	for _, msg := range msgs {
		srv.q.PutMT(msg)
	}
	return
}

func (srv *consumer) send() {

	defer close(srv.done)

	for {

		m, ok := srv.q.Get()
		if !ok {
			break
		}

		msg, ok := m.(*memphis.Msg)
		if ok {
			headers := &pb.Headers{Headers: msg.GetHeaders()}
			respmsg := &pb.ConsumeResponse_Msg{Msg: &pb.Msg{Headers: headers, Body: msg.Data()}}
			resp := &pb.ConsumeResponse{Data: respmsg}
			err := srv.stream.Send(resp)
			if err != nil {
				break
			}
			msg.Ack()
			continue
		}

		_, ok = m.(wakeup)
		if ok {
			resp := &pb.ConsumeResponse{Data: &pb.ConsumeResponse_Wakeup{Wakeup: &pb.Wakeup{}}}
			_ = srv.stream.Send(resp)
			continue
		}

		err, ok := m.(error)
		var textptr *string

		if ok {
			text := err.Error()
			textptr = &text
		}

		resp := &pb.ConsumeResponse{Data: &pb.ConsumeResponse_Status{Status: &pb.Status{Text: textptr}}}
		_ = srv.stream.Send(resp)
		break
	}

	if srv.cons != nil {
		srv.cons.StopConsume()
	}

	return
}

func (srv *consumer) abortConsume(err error) error {

	for {

		if !srv.started || srv.cons == nil {
			srv.q.Cancel()
			break
		}

		if err != nil {
			srv.q.PutMT(err)
		} else {
			srv.q.PutMT("stop")
		}

		break
	}

	<-srv.done

	return nil
}

func (srv *consumer) clean() {

	if srv == nil {
		return
	}

	if srv.cons != nil {
		srv.cons.Destroy()
	}

	if srv.mc != nil {
		srv.mc.Close()
	}
}
