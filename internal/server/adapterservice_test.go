package server

import (
	"context"
	"testing"

	"githib.com/g41797/grpcadapter/pb"
)

const (
	stName   = "AdapterStation"
	prodName = "AdapterProducer"
	consName = "AdapterConsumer"
)

func TestAdapterService_CreateDestroyStation(t *testing.T) {

	server := CreateGrpcServer()
	t.Cleanup(func() { server.Stop() })

	conn, err := startServerConnectClient(server)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { conn.Close() })

	client := pb.NewAdapterServiceClient(conn)

	createStation(t, client, stName)
	destroyStation(t, client, stName)
}

func TestAdapterService_Produce(t *testing.T) {

	server := CreateGrpcServer()
	t.Cleanup(
		func() {
			server.Stop()
		})

	conn, err := startServerConnectClient(server)

	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { conn.Close() })

	client := pb.NewAdapterServiceClient(conn)

	createStation(t, client, stName)

	t.Cleanup(func() { destroyStation(t, client, stName) })

	produceToStation(t, client, stName, prodName)
}

func createStation(t *testing.T, client pb.AdapterServiceClient, sname string) {

	mreq := createStationRequest(sname)

	manage(t, client, mreq)
}

func destroyStation(t *testing.T, client pb.AdapterServiceClient, sname string) {

	mreq := destroyStationRequest(sname)

	manage(t, client, mreq)
}

func produceToStation(t *testing.T, client pb.AdapterServiceClient, sname, pname string) {
	ctx := context.Background()

	pstr, err := client.Produce(ctx)
	if err != nil {
		t.Errorf("Produce error %v", err)
	}

	startProduce := startProduceRequest(sname, pname)

	err = pstr.Send(startProduce)
	if err != nil {
		t.Errorf("Create producer error %v", err)
	}

	status, err := pstr.CloseAndRecv()
	if err != nil {
		t.Errorf("Failed CloseAndRecv %v", err)
	}

	if len(status.GetText()) != 0 {
		t.Errorf("Failed produce status %s", status.GetText())
	}

}

func manage(t *testing.T, client pb.AdapterServiceClient, mreq *pb.ManageRequest) {

	ctx := context.Background()

	status, err := client.Manage(ctx, mreq)
	if err != nil {
		t.Fatal(err)
	}

	if len(status.GetText()) > 0 {
		t.Error(status.GetText())
	}
}
