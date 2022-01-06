package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	micpro "micpro/proto/micpro"
)

type Micpro struct{}

func (e *Micpro) Handle(ctx context.Context, msg *micpro.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *micpro.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
