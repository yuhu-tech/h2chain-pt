package test

import (
	pb "../api/mutation"
	"../handle"
	"context"
	"testing"
	"time"
)

func TestEditRemark(t *testing.T) {
	t.Log("test edit remark")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret,_:=obj.EditRemark(ctx,&pb.EditRequest{
		OrderId:"cjv2au7qj001x09515r1n6bam",
		PtId:"95566",
		StartDate:int32(time.Now().Unix()),
		EndDate:int32(time.Now().Unix()+7200),
		RealSalary:200,
		IsWorked:2,
		Type:1,
	})
	t.Log(ret.EditResult)
 }
