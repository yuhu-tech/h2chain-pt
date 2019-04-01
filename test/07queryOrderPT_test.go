package test

import (
	"context"
	"testing"
	"../handle"
	pb "../api/query"
)



func TestQueryOrderPT(t *testing.T){
	t.Log("query order's pt")

	ctx := context.TODO()

	obj := &handle.QueryServer{}
	ret, _ :=obj.QueryPTOfOrder(ctx,&pb.QueryPTRequest{
		OrderId:"cjtu22y0u000v0a938magq208",
		PtStatus:1,
		//PtId:"001",
		RegistrationChannel:"QQ",
	})

	t.Log(len(ret.OrderPts))
	t.Log(ret.OrderPts)
}
