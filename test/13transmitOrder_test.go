package test

import (
	"context"
	"testing"

	pb "../api/mutation"
	"../handle"
)

func TestTransmitOrder(t *testing.T) {
	t.Log("test transmit order ")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, err := obj.TransmitOrder(ctx, &pb.TransmitRequest{
		OrderId: "cjwae3d9c19j509103h9fcmuy",
		AgentId: "cjwa4p7eo001v095406dcozn0",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
