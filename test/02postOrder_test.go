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
		OrderId:      "cjv2au7qj001x09515r1n6bam",
		IsFloat:      1, // 0 represent don't float
		HourlySalary: 200,
		WorkContent:  "cleaning",
		Attention:    "carefully",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
