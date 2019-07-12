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
		OrderId:    "cjxvdldng000l09380nyl6qw6",
		PtId:       "3005",
		AdviserId:  "2001",
		ApplyTime:  int32(time.Now().Unix()),
		SignInTime: int32(time.Now().Unix() + 100),
		PtStatus:   1,
		Type:       2,
		InviterId:  "InviterId003",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
