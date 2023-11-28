package server

import (
	"context"

	"githib.com/g41797/grpcadapter/pb"
	"github.com/gogo/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ pb.AdapterServiceServer = (*AdapterService)(nil)

type AdapterService struct {
	pb.UnimplementedAdapterServiceServer
	bc *brokerConnector
}

func CreateGrpcServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterAdapterServiceServer(server, &AdapterService{})
	return server
}

func (srv *AdapterService) Manage(ctx context.Context, mr *pb.ManageRequest) (*pb.Status, error) {

	status := pb.Status{}

	err := srv.attachConnector()
	if err != nil {
		text := err.Error()
		status.Text = &text
		return &status, err
	}

	mngr := newManager(srv.bc)

	defer mngr.clean()

	return mngr.Manage(ctx, mr)
}

func (srv *AdapterService) CreateStation(ctx context.Context, req *pb.CreateStationRequest) (*pb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStation not implemented")
}
func (srv *AdapterService) DestroyStation(ctx context.Context, req *pb.DestroyStationRequest) (*pb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyStation not implemented")
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
