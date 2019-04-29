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
		OrderId:      "cjv1xov8x004d09596mcj5l66",
		IsFloat:      0, // 0 represent don't float
		HourlySalary: 300,
		WorkContent:  "driveCar",
		Attention:    "carefully",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
