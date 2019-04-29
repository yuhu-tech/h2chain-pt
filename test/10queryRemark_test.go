package test

import (
	"context"
	"testing"

	"../handle"
	pb "../api/query"
)

func TestQueryRemark(t *testing.T) {
	t.Log("test query remark")
	ctx := context.TODO()

	obj := &handle.QueryServer{}
	ret,_:=obj.QueryRemark(ctx,&pb.QueryRemarkRequest{
		OrderId:"cjv1oyuqo000j0959wrvjh113",
		PtId:"95527",
	})
	t.Log(ret.Remark)
}
