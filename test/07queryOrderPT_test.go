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
	ret, err :=obj.QueryPTOfOrder(ctx,&pb.QueryPTRequest{
		OrderId:"cjw8z8k4v00xn0910gd2b36yg",
		Type:3,
		InviterId:"InviterId001",
	})
	if err!=nil{
		t.Error(err)
	}
	t.Log(len(ret.OrderPts))
	t.Log(ret.OrderPts)
}
