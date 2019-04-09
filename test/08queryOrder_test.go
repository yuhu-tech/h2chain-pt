package test

import (
	"context"
	"testing"
	"../handle"
	pb "../api/query"
)

func TestQueryOrder(t *testing.T)  {
	t.Log("test query order by orderId date status")
	ctx := context.TODO()

	obj:=&handle.QueryServer{}
	ret,err:=obj.QueryOrder(ctx,&pb.QueryRequest{OrderId:"",Status:-1,Date:-1,HotelId:"001",Adviser:"",PtId:""})		// 1554220800
	if err!=nil{
		t.Error(err)
	}
	t.Log(ret.Orders)
	t.Log(len(ret.Orders))
}
