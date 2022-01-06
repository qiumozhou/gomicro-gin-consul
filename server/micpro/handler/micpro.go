package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	micpro "micpro/proto/micpro"
)

type Micpro struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Micpro) Call(ctx context.Context, req *micpro.Request, rsp *micpro.Response) error {
	log.Log("Received Micpro.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Micpro) Stream(ctx context.Context, req *micpro.StreamingRequest, stream micpro.Micpro_StreamStream) error {
	log.Logf("Received Micpro.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&micpro.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Micpro) PingPong(ctx context.Context, stream micpro.Micpro_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&micpro.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
