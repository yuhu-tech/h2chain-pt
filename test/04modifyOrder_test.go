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
		OrderId:          "cjtu22y0u000v0a938magq208",
		DateChanged:      int32(time.Now().Unix()),
		DurationChanged:  3,
		CountChanged:     12,
		CountMaleChanged: 8,
		Mode:             1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
