package handler

import (
	"context"
	"io"
	"time"

	log "go-micro.dev/v4/logger"

	pb "github.com/awe76/sagaproc/proto"
)

type Sagaproc struct{}

func (e *Sagaproc) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received Sagaproc.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Sagaproc) ClientStream(ctx context.Context, stream pb.Sagaproc_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *Sagaproc) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.Sagaproc_ServerStreamStream) error {
	log.Infof("Received Sagaproc.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *Sagaproc) BidiStream(ctx context.Context, stream pb.Sagaproc_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
