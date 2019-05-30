package test

import (
	"context"
	"testing"

	pb "../api/query"
	"../handle"
)

func TestQueryOrderOfAgent(t *testing.T) {
	t.Log("test query order of agent")
	ctx := context.TODO()
	obj := &handle.QueryServer{}
	ret, err := obj.QueryOrderOfAgent(ctx, &pb.QueryOOARequest{
		AgentId: "",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(ret)
}
