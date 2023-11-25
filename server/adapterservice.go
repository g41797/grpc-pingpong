package server

import (
	"context"

	"githib.com/g41797/grpcadapter/pb"
	"google.golang.org/grpc"
)

var _ pb.AdapterServiceServer = (*AdapterService)(nil)

type AdapterService struct {
	pb.UnimplementedAdapterServiceServer
	bc *brokerConnector
}

func createGrpcServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterAdapterServiceServer(server, &AdapterService{})
	return server
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

func (srv *AdapterService) Consume(c pb.AdapterService_ConsumeServer) error {
	err := srv.attachConnector()
	if err != nil {
		return err
	}

	consumer := newConsumer(srv.bc)

	defer consumer.clean()

	err = consumer.Consume(c)

	return err
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
