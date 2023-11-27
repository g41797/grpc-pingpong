package server

import "githib.com/g41797/grpcadapter/pb"

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

func startProduceRequest(sname, pname string) *pb.ProduceMessages {

	station := &pb.Station{Name: sname}
	producer := &pb.Producer{Name: pname}
	cpreq := &pb.CreateProducerRequest{Station: station, Producer: producer}
	start := &pb.ProduceMessages_Start{Start: cpreq}

	request := &pb.ProduceMessages{Data: start}

	return request
}

func produceRequest(hdrs map[string]string, body []byte) *pb.ProduceMessages {
	h := &pb.Headers{Headers: hdrs}
	msg := &pb.ProduceMessages_Msg{Msg: &pb.Msg{Headers: h, Body: body}}

	request := &pb.ProduceMessages{Data: msg}

	return request
}
