package test

import (
	"context"
	"testing"
	"time"

	pb "../api/mutation"
	"../handle"
)

func TestCreateOrder(t *testing.T) {
	t.Log("test create order ")
	ctx := context.TODO()
	obj := &handle.MutationServer{}
	ret, _ := obj.CreateOrder(ctx, &pb.CreateRequest{
		HotelId:     "001",
		AdviserId:   "002",
		Job:         "driver",
		Date:        int32(time.Now().Unix()),
		Duration:    4,
		Count:       5,
		CountFemale: 0,
		Mode:        0,
	})
	t.Log(ret)
}
