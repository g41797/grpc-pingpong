package server

import (
	"context"
	"fmt"

	"githib.com/g41797/grpcadapter/pb"
	"github.com/memphisdev/memphis.go"
)

var _ pb.AdapterServiceServer = (*AdapterService)(nil)

type AdapterService struct {
	pb.UnimplementedAdapterServiceServer
	bc *brokerConnector
}

func (srv *AdapterService) Manage(ctx context.Context, mr *pb.ManageRequest) (*pb.Status, error) {

	status := pb.Status{}

	err := srv.attachConnector()
	if err != nil {
		*status.Text = err.Error()
		return &status, err
	}

	mngr := newManager(srv.bc)

	defer mngr.clean()

	return mngr.Manage(ctx, mr)
}

func (srv *AdapterService) Produce(stream pb.AdapterService_ProduceServer) error {

	err := srv.attachConnector()
	if err != nil {
		return err
	}

	producer := newProducer(srv.bc)

	defer producer.clean()

	err = producer.Produce(stream)
	return err
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

func (srv *AdapterService) attachConnector() error {
	if srv.bc != nil {
		return nil
	}

	bc, err := newBrokerConnector()
	if err != nil {
		return err
	}
	srv.bc = bc
	return nil
}
