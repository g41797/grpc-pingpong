package server

import (
	"context"
	"testing"

	"githib.com/g41797/grpcadapter/pb"
)

const stName = "AdapterStation"

func TestAdapterService_CreateStation(t *testing.T) {

	conn, err := startServerConnectClient(createGrpcServer)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewAdapterServiceClient(conn)

	for i := 0; i < 1; i++ {
		createStation(t, client, stName)
	}
}

func createStation(t *testing.T, client pb.AdapterServiceClient, sname string) {

	station := &pb.Station{Name: sname}
	abret := &pb.RetentionOpt_Abret{Abret: &pb.AckBasedRet{AckBased: false}}
	retopt := &pb.RetentionOpt{Retentions: abret}
	stopt := &pb.StorageOpt{StorageType: pb.StorageOpt_Disk}
	partopt := &pb.PartitionOpt{Number: 1}
	sopts := &pb.StationOpions{Part: partopt, Storage: stopt, Retention: retopt}

	mreq := pb.ManageRequest{
		Data: &pb.ManageRequest_Createstation{
			Createstation: &pb.CreateStationRequest{Station: station, Options: sopts}}}

	ctx := context.Background()

	status, err := client.Manage(ctx, &mreq)
	if err != nil {
		t.Fatal(err)
	}

	if len(status.GetText()) > 0 {
		t.Error(status.GetText())
	}

}
