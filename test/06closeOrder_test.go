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
	ret, err := obj.CloseOrder(ctx, &pb.CloseRequest{OrderId:"cjtuwf9k700100a93fpcjaf1i"})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.CloseResult)
}
