package test

import (
	pb "../api/mutation"
	"../handle"
	"context"
	"testing"
	"time"
)


func TestCreateOrder(t *testing.T) {
	t.Log("test create order ")

	ctx := context.TODO()

	obj := &handle.MutationServer{}
	ret, _ := obj.CreateOrder(ctx, &pb.CreateRequest{
		HotelId:     "001",
		AdviserId:   "002",
		Job:         "cleaning",
		Date:        int32(time.Now().Unix()),
		Duration:    4,
		Count:       12,
		CountMale:   8,
		CountFemale: 4,
		Mode:        0,
	})
	t.Log(ret)
}


