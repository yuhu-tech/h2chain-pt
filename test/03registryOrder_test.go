package test

import (
	pb "../api/mutation"
	"../handle"
	"context"
	"testing"
	"time"
)

func TestRegistryOrder(t *testing.T) {
	t.Log("test registry order")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, err := obj.RegistryOrder(ctx, &pb.RegistryRequest{
		OrderId:             "cju9jk6de000l0a10kmuubt7l",
		PtId:                "001",
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
