package test

import (
	"context"
	"testing"
	"time"

	pb "../api/mutation"
	"../handle"
)

func TestRegistryOrder(t *testing.T) {
	t.Log("test registry order")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, err := obj.RegistryOrder(ctx, &pb.RegistryRequest{
		OrderId:             "cjv2au7qj001x09515r1n6bam",
		PtId:                "95566",
		AdviserId:           "002",
		ApplyTime:           int32(time.Now().Unix()),
		SignInTime:          int32(time.Now().Unix() + 100),
		PtStatus:            1,
		RegistrationChannel: "WeChat",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
