package test

import (
	"context"
	"testing"

	pb "../api/query"
	"../handle"
)

func TestQueryOrderPT(t *testing.T){
	t.Log("query order's pt")
	ctx := context.TODO()
	obj := &handle.QueryServer{}
	ret, _ :=obj.QueryPTOfOrder(ctx,&pb.QueryPTRequest{
		OrderId:"cju9jjub0000g0a10levs5wqa",
	})

	t.Log(len(ret.OrderPts))
	t.Log(ret.OrderPts)
}
