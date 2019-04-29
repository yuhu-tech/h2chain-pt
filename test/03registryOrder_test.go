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
		OrderId:             "cjv1xov8x004d09596mcj5l66",
		PtId:                "95527",
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
