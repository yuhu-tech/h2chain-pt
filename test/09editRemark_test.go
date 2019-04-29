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
		OrderId:"cjv1xov8x004d09596mcj5l66",
		PtId:"95527",
		StartDate:int32(time.Now().Unix()),
		EndDate:int32(time.Now().Unix()+7200),
		RealSalary:350,
		IsWorked:1,
		Type:1,
	})
	t.Log(ret.EditResult)
 }
