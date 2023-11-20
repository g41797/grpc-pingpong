package server

import (
	"context"
	"log"
	"net"
	"testing"

	"githib.com/g41797/grpcadapter/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestManagerService_CreateStation(t *testing.T) {

	conn, err := startServerConnectClient()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewManagerServiceClient(conn)

	for i := 0; i < 10; i++ {

		mreq := pb.ManageRequest{
			Data: &pb.ManageRequest_Createstation{
				Createstation: &pb.CreateStationRequest{Station: &pb.Station{Name: "FirstStation"}}}}

		ctx := context.Background()

		status, err := client.Manage(ctx, &mreq)
		if err != nil {
			t.Fatal(err)
		}

		if len(status.GetText()) > 0 {
			t.Error(status.GetText())
		}
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

func startServerConnectClient() (conn *grpc.ClientConn, err error) {
	ctx := context.Background()

	conn, err = grpc.DialContext(ctx, "",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(dialer()),
		grpc.WithBlock())

	return conn, err
}
