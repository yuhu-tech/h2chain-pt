package main

import (
	"net"
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
	pb "../api"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type Server struct {

}

func (s *Server)HotelCreateOrder(ctx context.Context, in *pb.HotelCreateOrderRequest) (*pb.HotelCreateOrderReply,error) {
	log.Printf("request:%v",in)
	log.Printf("Received: %v", in.Name)
	return &pb.HotelCreateOrderReply{OrderId:01,CreateResult:"create order successful!"},nil
}

func main()  {
	lis,err:=net.Listen("tcp",port)
	if err!=nil{
		log.Fatalf("failed to listen: %v",err)
	}
	s:=grpc.NewServer()
	pb.RegisterMutationManagerServer(s,&Server{})
	reflection.Register(s)
	if err:=s.Serve(lis);err!=nil{
		log.Fatalf("failed to server: %v",err)
	}

}
