package test

import (
	"context"
	"testing"
	"../handle"
	pb "../api/mutation"
)

func TestCreateOrder(t *testing.T) {
	t.Log("test create order ")

	ctx := context.TODO()

	obj:=&handle.MutationServer{}
	ret,_:=obj.CreateOrder(ctx,&pb.CreateRequest{})
	t.Log(ret)
}
