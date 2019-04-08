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
		OrderId:       "cju0sz0v2000v0976zgqtqrv3",
		PtId:          "",
		SourceStatus:  0,
		TargetStatus:  1,
		PtPerformance: -1,
		ObjectReason:  -1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.ModifyResult)
}
