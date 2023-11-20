package server

import (
	"context"
	"log"
	"net"
	"testing"

	"githib.com/g41797/grpcadapter/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestManagerService_CreateStation(t *testing.T) {

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewManagerServiceClient(conn)

	csr := pb.CreateStationRequest{Station: &pb.Station{Name: "FirstStation"}}
	mreq := pb.ManageRequest{}
	mreq.Data = &pb.ManageRequest_Createstation{Createstation: &csr}

	ctx = context.Background()

	status, err := client.Manage(ctx, &mreq)
	if err != nil {
		t.Fatal(err)
	}

	if len(status.GetText()) > 0 {
		t.Error(status.GetText())
	}
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	pb.RegisterManagerServiceServer(server, &ManagerService{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}
