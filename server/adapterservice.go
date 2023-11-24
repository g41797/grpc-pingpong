package server

import (
	"context"
	"fmt"
	"io"

	"githib.com/g41797/grpcadapter/pb"
	"github.com/memphisdev/memphis.go"
)

var _ pb.AdapterServiceServer = (*AdapterService)(nil)

type AdapterService struct {
	pb.UnimplementedAdapterServiceServer
	bc   *brokerConnector
	proc []manageRequest
}

//---------------
// MANAGE SERVICE
//---------------

func (srv *AdapterService) Manage(ctx context.Context, mr *pb.ManageRequest) (*pb.Status, error) {

	srv.buildChainOfResp()

	status := pb.Status{}

	mc, err := srv.connect()
	if err != nil {
		*status.Text = err.Error()
		return &status, err
	}

	defer mc.Close()

	for _, prc := range srv.proc {
		resp, status := prc(mr, mc)
		if resp {
			return status, nil
		}
	}

	*status.Text = "not implemented"

	return &status, nil
}

type manageRequest func(req *pb.ManageRequest, mc *memphis.Conn) (responsible bool, status *pb.Status)

func (srv *AdapterService) buildChainOfResp() {
	if srv.proc != nil {
		return
	}
	srv.proc = append(srv.proc, srv.createStation)

	//
	// Rest of functions: delete. etc
	//
}

func (srv *AdapterService) createStation(req *pb.ManageRequest, mc *memphis.Conn) (responsible bool, status *pb.Status) {

	cmd := req.GetCreatestation()
	if cmd != nil {
		return true, &pb.Status{}
	}

	return false, nil
}

//----------------
// PRODUCE SERVICE
//----------------

type producerMemo struct {
	mc       *memphis.Conn
	station  string
	producer string
}

func (pm *producerMemo) clean() {
	if pm == nil {
		return
	}

	if pm.mc != nil {
		pm.mc.Close()
	}
}

func (srv *AdapterService) Produce(stream pb.AdapterService_ProduceServer) error {

	var pm producerMemo

	defer pm.clean()

	mc, err := srv.connect()
	if err != nil {
		status := pb.Status{}
		*status.Text = err.Error()
		stream.SendAndClose(&status)
		return nil
	}

	pm.mc = mc

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

		if msg := next.GetMsg(); msg != nil {
			if err = srv.produce(msg, &pm); err != nil {
				status := pb.Status{}
				*status.Text = err.Error()
				stream.SendAndClose(&status)
				return nil
			}
			continue
		}

		if start := next.GetStart(); start != nil {
			if err = srv.createProducer(start, &pm); err != nil {
				status := pb.Status{}
				*status.Text = err.Error()
				stream.SendAndClose(&status)
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

func (srv *AdapterService) createProducer(start *pb.ProduceRequest, pm *producerMemo) error {
	return fmt.Errorf("start not implemented")
}

func (srv *AdapterService) produce(msg *pb.Msg, pm *producerMemo) error {
	return fmt.Errorf("produce not implemented")
}

//----------------
// CONSUME SERVICE
//----------------

type consumerMemo struct {
	mc       *memphis.Conn
	station  string
	producer string
	options  struct{}
}

func (cm *consumerMemo) clean() {
	if cm == nil {
		return
	}

	if cm.mc != nil {
		cm.mc.Close()
	}
}

func (srv *AdapterService) Consume(c pb.AdapterService_ConsumeServer) error {

	return nil
}

func (srv *AdapterService) startConsume(start *pb.ConsumeRequest, cm *consumerMemo) error {
	return fmt.Errorf("startConsume not implemented")
}

func (srv *AdapterService) abortConsume(cm *consumerMemo) error {
	return fmt.Errorf("abortConsume not implemented")
}

//----------------
//
//----------------

func (srv *AdapterService) connect() (*memphis.Conn, error) {
	if srv.bc == nil {
		bc, err := newBrokerConnector()
		if err != nil {
			return nil, err
		}
		srv.bc = bc
	}

	mc, err := srv.bc.connect()
	return mc, err
}
