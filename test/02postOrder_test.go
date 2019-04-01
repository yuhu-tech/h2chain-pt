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
		OrderId:      "cjtu22y0u000v0a938magq208",
		IsFloat:      0, // 0 represent don't float
		HourlySalary: 18,
		WorkContent:  "cleaningHotel",
		Attention:    "patience",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
