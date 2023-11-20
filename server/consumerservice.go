package server

import (
	"io"

	"githib.com/g41797/grpcadapter/pb"
)

var _ pb.ConsumerServiceServer = (*ConsumerService)(nil)

type ConsumerService struct {
	pb.UnimplementedConsumerServiceServer
}

func (srv *ConsumerService) connect() error {
	return nil
}

func (srv *ConsumerService) disconnect() error {
	return nil
}

func (srv *ConsumerService) startConsume(start *pb.ConsumeRequest) error {

	if err := srv.connect(); err != nil {
		return err
	}

	// TODO Add consume activation

	return nil
}

func (srv *ConsumerService) abortConsume() error {
	return nil
}

func (srv *ConsumerService) Consume(c pb.ConsumerService_ConsumeServer) error {

	defer srv.abortConsume()
	defer srv.disconnect()

	for {
		next, err := c.Recv()
		resp := pb.ConsumeResponse{}

		if err == io.EOF {
			resp.Data = &pb.ConsumeResponse_Stop{}
			c.Send(&resp)
			return nil
		}

		if err != nil {
			status := pb.Status{}
			*status.Text = err.Error()
			resp.Data = &pb.ConsumeResponse_Error{&status}
			c.Send(&resp)
			return nil
		}

		if start := next.GetStart(); start != nil {
			if err := srv.startConsume(start); err != nil {
				status := pb.Status{}
				*status.Text = err.Error()
				resp.Data = &pb.ConsumeResponse_Error{&status}
				c.Send(&resp)
				return nil
			}
			continue
		}

		resp.Data = &pb.ConsumeResponse_Error{&pb.Status{}}
		c.Send(&resp)
		break
	}

	return nil
}
