package test

import (
	"context"
	"testing"
	"time"

	pb "../api/mutation"
	"../handle"
)

func TestModifyOrder(t *testing.T) {
	t.Log("test hotel modify order")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, err := obj.ModifyOrder(ctx, &pb.ModifyRequest{
		OrderId:          "cju9jjub0000g0a10levs5wqa",
		DateChanged:      int32(time.Now().Unix()),
		DurationChanged:  3,
		CountChanged:     12,
		CountMaleChanged: 10,
		Mode:             1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
