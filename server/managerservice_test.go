package server

import (
	"context"
	"testing"

	"githib.com/g41797/grpcadapter/pb"
	"google.golang.org/grpc"
)

func TestManagerService_CreateStation(t *testing.T) {

	conn, err := startServerConnectClient(createManagerServiceServer)
	if err != nil {
		t.Fatal(err)
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

func createManagerServiceServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterManagerServiceServer(server, &ManagerService{})
	return server
}
