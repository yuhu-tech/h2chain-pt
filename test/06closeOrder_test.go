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
	ret, err := obj.CloseOrder(ctx, &pb.CloseRequest{OrderId:"cjuc3dxhn007a0a10dct4o4kn"})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.CloseResult)
}
