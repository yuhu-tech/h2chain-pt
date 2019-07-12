package test

import (
	"context"
	"testing"

	pb "../api/mutation"
	"../handle"
)

func TestPostOrder(t *testing.T) {
	t.Log("test post order ")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, err := obj.PostOrder(ctx, &pb.PostRequest{
		OrderId:      "cjxvhf5rc000f0a38pik26d7c",
		IsFloat:      1, // 0 represent don't float
		HourlySalary: 230,
		WorkContent:  "cleaning",
		Attention:    "carefully",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
