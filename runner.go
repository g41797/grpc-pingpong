package grpcpingpong

import (
	"fmt"
	"net"

	"github.com/g41797/grpc-pingpong/internal/server"
	"github.com/g41797/sputnik"
	"google.golang.org/grpc"
)

type GrpcConfig struct {
	LISTEN_NETWORK string
	LISTEN_ADDRESS string
}

const grpcConfigName = "grpcadapter"

// Runner functionality :
// Init,Run and Stop grpc server
type Runner struct {
	conf   GrpcConfig
	lis    net.Listener
	server *grpc.Server
}

// Prepares listener and creates grpc server
func (r *Runner) Init(cfact sputnik.ConfFactory) error {
	if err := cfact(grpcConfigName, &r.conf); err != nil {
		return err
	}

	lis, err := net.Listen(r.conf.LISTEN_NETWORK, r.conf.LISTEN_ADDRESS)
	if err != nil {
		return nil
	}
	r.lis = lis

	r.server = server.CreateGrpcServer()

	return nil
}

// Starts grpc server, returns stop server function
// for further usage during process finish
func (r *Runner) Run() (stop func(), err error) {
	if r.lis == nil {
		return nil, fmt.Errorf("nil listener")
	}

	if r.server == nil {
		return nil, fmt.Errorf("nil server")
	}

	if err = r.server.Serve(r.lis); err != nil {
		return nil, err
	}

	r.lis = nil

	return func() { r.server.Stop() }, nil
}
