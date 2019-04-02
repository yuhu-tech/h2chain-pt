package test

import (
	"../handle"
	"encoding/json"
	"testing"
)

func TestNativeQuery(t *testing.T) {
	t.Log("test native graphql query")
	res := handle.NativeGrapgql()
	res1,_:=json.Marshal(res)
	//var res2 pb.Order
	var res2 struct{
		OrderOrigins []struct{
			Id string
			AdviserId string
			HotelId string
		}
	}
	json.Unmarshal(res1,&res2)
	t.Log(res2.OrderOrigins[0].Id)
	t.Log(len(res2.OrderOrigins))
}
