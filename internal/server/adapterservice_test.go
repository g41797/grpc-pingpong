package server

import (
	"context"
	"io"
	"reflect"
	"testing"

	"githib.com/g41797/memphisgrpc/pb"
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

var headersG = map[string]string{"1": "1", "2": "2"}
var bodyG = []byte("first grpc produce")

func TestAdapterService_ProduceConsume(t *testing.T) {

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

	consumeFromStation(t, client, stName, consName)
}

func createStation(t *testing.T, client pb.AdapterServiceClient, sname string) {

	req := createStationRequest(sname)

	ctx := context.Background()

	status, err := client.CreateStation(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	if len(status.GetText()) > 0 {
		t.Error(status.GetText())
	}

}

func destroyStation(t *testing.T, client pb.AdapterServiceClient, sname string) {

	req := destroyStationRequest(sname)

	ctx := context.Background()

	status, err := client.DestroyStation(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	if len(status.GetText()) > 0 {
		t.Error(status.GetText())
	}
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

	produceMessage := produceRequest(headersG, bodyG)
	err = pstr.Send(produceMessage)
	if err != nil {
		t.Errorf("Produce error %v", err)
	}

	err = pstr.Send(stopProduceRequest())
	if err != nil {
		t.Errorf("Stop produce error %v", err)
	}

	status, err := pstr.CloseAndRecv()
	if err != nil {
		t.Errorf("Failed CloseAndRecv %v", err)
	}

	if len(status.GetText()) != 0 {
		t.Errorf("Failed produce status %s", status.GetText())
	}

}

func consumeFromStation(t *testing.T, client pb.AdapterServiceClient, sname, name string) {
	ctx := context.Background()

	stream, err := client.Consume(ctx)
	if err != nil {
		t.Errorf("Consume error %v", err)
	}

	start := startConsumeRequest(sname, name)

	err = stream.Send(start)
	if err != nil {
		t.Errorf("Create consumer error %v", err)
	}

	respmsg, _, err := skipWakeup(stream)
	if err != nil {
		t.Errorf("Recv consumed message error %v", err)
	}

	if respmsg == nil {
		t.Errorf("Recv consumed error - expected message")
	}

	h := respmsg.GetHeaders()
	if h == nil {
		t.Errorf("Recv consumed error - expected headers")
	}

	headers := h.GetHeaders()
	if headers == nil {
		t.Errorf("Recv consumed error - expected map")
	}

	if !reflect.DeepEqual(headers, headersG) {
		t.Errorf("Recv consumed headers mismatch")
	}

	body := respmsg.GetBody()
	if body == nil {
		t.Errorf("Recv consumed error - nil body")
	}

	if !reflect.DeepEqual(body, bodyG) {
		t.Errorf("Recv consumed body mismatch")
	}

	stop := stopConsumeRequest()

	err = stream.Send(stop)
	if err != nil {
		t.Errorf("Stop consume error %v", err)
	}

	_, respstat, err := skipWakeup(stream)
	if err != nil {
		t.Errorf("Recv consumed message error %v", err)
	}
	if respstat == nil {
		t.Errorf("Recv consumed error - expected status")
	}

	text := respstat.GetText()
	if len(text) != 0 {
		t.Errorf("Recv status error - received status %s", text)
	}

	_, _, err = skipWakeup(stream)
	if (err == io.EOF) || (err == nil) {
		return
	}

	t.Errorf("Recv consumed error %v", err)
}

func skipWakeup(stream pb.AdapterService_ConsumeClient) (msg *pb.Msg, status *pb.Status, err error) {

	for {

		resp, err := stream.Recv()
		if err != nil {
			break
		}

		msg = resp.GetMsg()
		if msg != nil {
			break
		}

		status = resp.GetStatus()
		if status != nil {
			break
		}

		continue // for break
	}

	return
}
