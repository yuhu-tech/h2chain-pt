package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
	pb "../api"
	"../handle"
)

const (
	port = ":50051"
)

func main()  {
	lis,err:=net.Listen("tcp",port)
	if err!=nil{
		log.Fatalf("failed to listen: %v",err)
	}
	s:=grpc.NewServer()
	pb.RegisterOrderOperationServer(s,&handle.Server{})
	reflection.Register(s)
	if err:=s.Serve(lis);err!=nil{
		log.Fatalf("failed to server: %v",err)
	}

}
