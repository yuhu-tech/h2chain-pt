package test

import (
	"context"
	"testing"

	pb "../api/query"
	"../handle"
)

func TestQueryExperience(t *testing.T) {
	t.Log("test query ")
	ctx := context.TODO()

	obj := handle.QueryServer{}
	ret, err := obj.QueryExperience(ctx, &pb.QueryExperienceRequest{
		PtId:"95527",
	})
	if err!=nil{
		t.Error(err)
	}
	t.Log(ret.WorkExperience)
}
