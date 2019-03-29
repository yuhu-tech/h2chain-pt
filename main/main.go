package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
	mpb "../api/mutation"
	qpb "../api/query"
	"../handle_test"
)

const (
	port = ":50051"
)

func main()  {
	lis,err:=net.Listen("tcp",port)
	if err!=nil{
		log.Fatalf("failed to listen: %v",err)
	}
	//TODO 是否需要注册两个grpc服务
	s:=grpc.NewServer()

	mpb.RegisterMutationServer(s,&handle_test.MutationServer{})
	qpb.RegisterQueryOrderServer(s,&handle_test.QueryServer{})

	reflection.Register(s)
	if err:=s.Serve(lis);err!=nil{
		log.Fatalf("failed to server: %v",err)
	}

}
