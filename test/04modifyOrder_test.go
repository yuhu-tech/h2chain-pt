package test

import (
	pb "../api/mutation"
	"../handle"
	"context"
	"testing"
	"time"
)

func TestModifyOrder(t *testing.T) {
	t.Log("test hotel modify order")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, err := obj.ModifyOrder(ctx, &pb.ModifyRequest{
		OrderId:          "cju0sz0v2000v0976zgqtqrv3",
		DateChanged:      int32(time.Now().Unix()),
		DurationChanged:  2,
		CountChanged:     12,
		CountMaleChanged: 10,
		Mode:             1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
