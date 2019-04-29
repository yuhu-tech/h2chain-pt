package test

import (
	"context"
	"testing"

	pb "../api/query"
	"../handle"
)

func TestQueryOrder(t *testing.T) {
	t.Log("test query order by orderId date status")
	ctx := context.TODO()
	obj := &handle.QueryServer{}
	ret, err := obj.QueryOrder(ctx, &pb.QueryRequest{OrderId:"cjv1oyuqo000j0959wrvjh113"})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.OrderOrigins)
}
