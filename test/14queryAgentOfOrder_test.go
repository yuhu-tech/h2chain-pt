package test

import (
	"context"
	"testing"

	pb "../api/query"
	"../handle"
)

func TestQueryAgentOfOrder(t *testing.T) {
	t.Log("test query agent of order")
	ctx := context.TODO()
	obj := &handle.QueryServer{}
	ret, err := obj.QueryAgentOfOrder(ctx, &pb.QueryAgentRequest{
		OrderId: "",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
