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
}

func (srv *ManagerService) Manage(ctx context.Context, mr *pb.ManageRequest) (*Status, error) {

	status := pb.Status{}

	if err := srv.connect(); err != nil {
		*status.Text = err.Error()
		return status
	}

	defer srv.disconnect()

	if req := mr.GetCreatestation(); req != nil {
		if err := srv.createStation(req); err != nil {
			*status.Text = err.Error()
		}
		return status
	}

	return pb.Status{}
}
