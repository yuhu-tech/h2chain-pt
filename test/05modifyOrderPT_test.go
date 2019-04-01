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
		Id:            "cjtuy3xnh001r0a93pzrz2sb7",
		TargetStatus:  2,
		PtPerformance: -1,
		ObjectReason:  -1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.ModifyResult)
}
