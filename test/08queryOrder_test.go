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
	ret,err:=obj.QueryOrder(ctx,&pb.QueryRequest{OrderId:"cju0sz0v2000v0976zgqtqrv3",Status:-1,Date:-1})		// 1554220800
	if err!=nil{
		t.Error(err)
	}
	t.Log(ret.Orders[0].OrderHotelModifies[0])
}
