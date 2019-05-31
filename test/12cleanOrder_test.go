package test

import (
	"../handle"
	"context"
	"testing"

	pb "../api/mutation"
)

func TestCleanOrder(t *testing.T) {
	t.Log("test clean order ")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, err := obj.CleanOrder(ctx, &pb.CleanRequest{
		Date: 1558573200,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.OrderOrigins)
}
