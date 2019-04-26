package test

import (
	"context"
	"testing"

	pb "../api/mutation"
	"../handle"
)

func TestCloseOrder(t *testing.T) {
	t.Log("test close order")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, err := obj.CloseOrder(ctx, &pb.CloseRequest{OrderId:"cjuuwu7pa00070a10tyq4i2iz"})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.CloseResult)
}
