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
		OrderId:"cjv2au7qj001x09515r1n6bam",
		PtId:"95555",
	})
	t.Log(ret.Remark)
}
