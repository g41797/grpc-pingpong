package server

import (
	"context"
	"testing"

	"githib.com/g41797/grpcadapter/pb"
)

const stName = "AdapterStation"

func TestAdapterService_CreateDestroyStation(t *testing.T) {

	conn, err := startServerConnectClient(createGrpcServer)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewAdapterServiceClient(conn)

	createStation(t, client, stName)
	destroyStation(t, client, stName)
}

func createStation(t *testing.T, client pb.AdapterServiceClient, sname string) {

	mreq := createStationRequest(sname)

	manage(t, client, mreq)
}

func destroyStation(t *testing.T, client pb.AdapterServiceClient, sname string) {

	mreq := destroyStationRequest(sname)

	manage(t, client, mreq)
}

func createStationRequest(sname string) (mreq *pb.ManageRequest) {

	station := &pb.Station{Name: sname}
	abret := &pb.RetentionOpt_Abret{Abret: &pb.AckBasedRet{AckBased: false}}
	retopt := &pb.RetentionOpt{Retentions: abret}
	stopt := &pb.StorageOpt{StorageType: pb.StorageOpt_Disk}
	partopt := &pb.PartitionOpt{Number: 1}
	sopts := &pb.StationOpions{Part: partopt, Storage: stopt, Retention: retopt}

	mreq = &pb.ManageRequest{
		Data: &pb.ManageRequest_Createstation{
			Createstation: &pb.CreateStationRequest{Station: station, Options: sopts}}}

	return mreq
}

func destroyStationRequest(sname string) (mreq *pb.ManageRequest) {

	station := &pb.Station{Name: sname}

	destst := &pb.DestroyStationRequest{Station: station}

	mreq = &pb.ManageRequest{
		Data: &pb.ManageRequest_Destroystation{
			Destroystation: destst}}

	return mreq
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
