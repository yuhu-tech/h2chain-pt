package test

import (
	"context"
	"testing"
	"../handle"
	pb "../api/query"
)



func TestQueryOrderPT(t *testing.T){
	t.Log("abc")

	ctx := context.TODO()

	obj := &handle.QueryServer{}
	ret, _ :=obj.QueryOrderPT(ctx,&pb.QueryPTRequest{
		OrderId:"1",
		PtId:"2",
	})
	t.Log(ret)
}
