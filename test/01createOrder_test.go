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
		HotelId:     "1002",
		AdviserId:   "2002",
		Job:         "clean",
		Date:        int32(time.Now().Unix()),
		Duration:    4,
		Count:       10,
		CountFemale: 5,
		Mode:        0,
	})
	t.Log(ret)
}
