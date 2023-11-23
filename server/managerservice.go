package server

import (
	"context"

	"githib.com/g41797/grpcadapter/pb"
)

var _ pb.ManagerServiceServer = (*ManagerService)(nil)

func (srv *ManagerService) connect() error {
	return nil
}

func (srv *ManagerService) disconnect() error {
	return nil
}

func (srv *ManagerService) createStation(req *pb.CreateStationRequest) error {
	return nil
}

type ManagerService struct {
	pb.UnimplementedManagerServiceServer
	bc *brokerConnector
}

func (srv *ManagerService) Manage(ctx context.Context, mr *pb.ManageRequest) (*pb.Status, error) {

	status := pb.Status{}

	if srv.bc == nil {
		bc, err := newBrokerConnector()
		if err != nil {
			*status.Text = err.Error()
			return &status, err
		}
		srv.bc = bc
	}

	mc, err := srv.bc.connect()
	if err != nil {
		*status.Text = err.Error()
		return &status, err
	}

	defer mc.Close()

	if req := mr.GetCreatestation(); req != nil {
		if err := srv.createStation(req); err != nil {
			*status.Text = err.Error()
		}
		return &status, nil
	}

	return &pb.Status{}, nil
}
