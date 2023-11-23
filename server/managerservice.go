package server

import (
	"context"

	"githib.com/g41797/grpcadapter/pb"
	"github.com/memphisdev/memphis.go"
)

var _ pb.ManagerServiceServer = (*ManagerService)(nil)

type ManagerService struct {
	pb.UnimplementedManagerServiceServer
	bc   *brokerConnector
	proc []processRequest
}

func (srv *ManagerService) Manage(ctx context.Context, mr *pb.ManageRequest) (*pb.Status, error) {

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

type processRequest func(req *pb.ManageRequest, mc *memphis.Conn) (responsible bool, status *pb.Status)

func (srv *ManagerService) connect() (*memphis.Conn, error) {
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

func (srv *ManagerService) buildChainOfResp() {
	if srv.proc != nil {
		return
	}
	srv.proc = append(srv.proc, srv.createStation)

	//
	// Rest of functions: delete. etc
	//
}

func (srv *ManagerService) createStation(req *pb.ManageRequest, mc *memphis.Conn) (responsible bool, status *pb.Status) {

	cmd := req.GetCreatestation()
	if cmd != nil {
		return true, &pb.Status{}
	}

	return false, nil
}
