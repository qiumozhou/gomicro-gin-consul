package handler

import (
	"context"
	micpro "micpro/proto/micpro"
)

type Micpro struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Micpro) Call(ctx context.Context, req *micpro.Request, rsp *micpro.Response) error {

	//此处编写需要处理的业务逻辑，
	rsp.Msg = "wqw"
	return nil
}

