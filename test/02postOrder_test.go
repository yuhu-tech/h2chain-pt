package test

import (
	pb "../api/mutation"
	"../handle"
	"context"
	"testing"
)

func TestPostOrder(t *testing.T) {
	t.Log("test post order ")

	ctx := context.TODO()

	obj := &handle.MutationServer{}
	ret, err := obj.PostOrder(ctx, &pb.PostRequest{
		OrderId:      "cju0sz0v2000v0976zgqtqrv3",
		IsFloat:      1, // 0 represent don't float
		HourlySalary: 18,
		WorkContent:  "cleaningHotel",
		Attention:    "patience",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
