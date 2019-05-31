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
		OrderId: "123456",
		AgentId: "qwertyui",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
