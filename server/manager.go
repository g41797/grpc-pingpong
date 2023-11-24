package server

import (
	"context"

	"githib.com/g41797/grpcadapter/pb"
	"github.com/memphisdev/memphis.go"
)

type manageRequest func(req *pb.ManageRequest) (responsible bool, status *pb.Status)

type manager struct {
	bc   *brokerConnector
	proc []manageRequest
	mc   *memphis.Conn
}

func newManager(bc *brokerConnector) *manager {
	manager := new(manager)
	manager.bc = bc
	manager.buildChainOfResp()
	return manager
}

func (srv *manager) Manage(ctx context.Context, mr *pb.ManageRequest) (*pb.Status, error) {
	status := pb.Status{}

	mc, err := srv.bc.connect()
	if err != nil {
		*status.Text = err.Error()
		return &status, err
	}

	srv.mc = mc

	for _, prc := range srv.proc {
		resp, status := prc(mr)
		if resp {
			return status, nil
		}
	}

	*status.Text = "not implemented"

	return &status, nil
}

func (srv *manager) buildChainOfResp() {
	if srv.proc != nil {
		return
	}
	srv.proc = append(srv.proc, srv.createStation)

	//
	// Rest of functions: delete. etc
	//
}

func (srv *manager) createStation(req *pb.ManageRequest) (responsible bool, status *pb.Status) {

	cmd := req.GetCreatestation()
	if cmd != nil {
		return true, &pb.Status{}
	}

	return false, nil
}

func (srv *manager) clean() {
	if srv == nil {
		return
	}

	if srv.mc != nil {
		srv.mc.Close()
	}
}
