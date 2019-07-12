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
		OrderId:          "cjxvdldng000l09380nyl6qw6",
		DateChanged:      int32(time.Now().Unix()),
		DurationChanged:  4,
		CountChanged:     12,
		CountMaleChanged: 10,
		Mode:             1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
