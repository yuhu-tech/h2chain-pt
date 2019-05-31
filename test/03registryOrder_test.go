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
		OrderId:    "cjw8xo2a400uy0910f01qicfa",
		PtId:       "3001",
		AdviserId:  "2001",
		ApplyTime:  int32(time.Now().Unix()),
		SignInTime: int32(time.Now().Unix() + 100),
		PtStatus:   1,
		Type:       3,
		InviterId:  "InviterId001",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
