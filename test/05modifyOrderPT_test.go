package test

import (
	"context"
	"testing"

	pb "../api/mutation"
	"../handle"
)

func TestModifyOrderPT(t *testing.T) {
	t.Log("test modify order's pt")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, err := obj.ModifyPTOfOrder(ctx, &pb.ModifyPtRequest{
		OrderId:       "cju9jjub0000g0a10levs5wqa",
		SourceStatus:  1,
		TargetStatus:  2,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret.ModifyResult)
}
