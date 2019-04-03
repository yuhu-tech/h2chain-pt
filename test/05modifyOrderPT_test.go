package test

import (
	pb "../api/mutation"
	"../handle"
	"context"
	"testing"
)

func TestModifyOrderPT(t *testing.T) {
	t.Log("test modify order's pt")
	ctx := context.TODO()

	obj := &handle.MutationServer{}
	ret, err := obj.ModifyPTOfOrder(ctx, &pb.ModifyPtRequest{
		Id:            "cju0tmh6400400976c26d5eet",
		TargetStatus:  2,
		PtPerformance: -1,
		ObjectReason:  -1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.ModifyResult)
}
