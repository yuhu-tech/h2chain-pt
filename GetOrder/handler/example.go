package handler

import (
	"context"

	example "H2chainProject/GetOrder/proto/example"


)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) Call(ctx context.Context, req *example.Request, rsp *example.Response) error {
	rsp.Errno ="001"
	rsp.Errmsg = "获取成功"
	order := example.ResponseOrder{Orderid:001,Orderuser:"mark"}
	rsp.Data=append(rsp.Data,&order)
	return nil
}

