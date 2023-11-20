package server

import (
	"io"

	"githib.com/g41797/grpcadapter/pb"
)

var _ pb.ProducerServiceServer = (*ProducerService)(nil)

type ProducerService struct {
	pb.UnimplementedProducerServiceServer
}

func (srv *ProducerService) connect() error {
	return nil
}

func (srv *ProducerService) disconnect() error {
	return nil
}

func (srv *ProducerService) produce(msg *pb.Msg) error {
	return nil
}

func (srv *ProducerService) start(Start *pb.ProduceRequest) error {

	if err := srv.connect(); err != nil {
		status := pb.Status{}
		*status.Text = err.Error()
		return err
	}

	return nil
}

func (srv *ProducerService) Produce(stream pb.ProducerService_ProduceServer) error {

	defer srv.disconnect()

	for {
		next, err := stream.Recv()

		if err == io.EOF {
			stream.SendAndClose(&pb.Status{})
			return nil
		}
		if err != nil {
			return err
		}

		if stop := next.GetStop(); stop != nil {
			stream.SendAndClose(&pb.Status{})
			return nil
		}

		if msg := next.GetMsg(); msg != nil {
			if err = srv.produce(msg); err != nil {
				status := pb.Status{}
				*status.Text = err.Error()
				stream.SendAndClose(&status)
				return nil
			}
			continue
		}

		if start := next.GetStart(); start != nil {
			if err = srv.start(start); err != nil {
				status := pb.Status{}
				*status.Text = err.Error()
				stream.SendAndClose(&status)
				return nil
			}
			continue
		}

		status := pb.Status{}
		*status.Text = "wrong client request"
		stream.SendAndClose(&status)
		break
	}

	return nil
}
