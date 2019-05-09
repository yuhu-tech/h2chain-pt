package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	mpb "../api/mutation"
	qpb "../api/query"
	"../handle"
	"../cleanOrder"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	mpb.RegisterMutationServer(s, &handle.MutationServer{})
	qpb.RegisterQueryOrderServer(s, &handle.QueryServer{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

	// 启动定时清理订单服务
	cleanOrder.StartTimer()
}
