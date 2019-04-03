package test

import (
	pb "../api/mutation"
	"../handle"
	"context"
	"testing"
)

func TestCloseOrder(t *testing.T) {
	t.Log("test close order")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, err := obj.CloseOrder(ctx, &pb.CloseRequest{OrderId:"cju0szhh1001009767nik93ry"})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.CloseResult)
}
