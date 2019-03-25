package main

import (
	pb "../api"
	"../prisma"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

const (
	port = ":50051"
)

type Server struct {

}

func (s *Server)HotelCreateOrder(ctx context.Context, in *pb.HotelCreateOrderRequest) (*pb.HotelCreateOrderReply,error) {
	//1 get create order request
	log.Printf("request:%v",in)

	//2 create prisma client
	client:=prisma.New(nil)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//ctx:=context.TODO()

	//3 insert data into prisma-mysql
	newOrder,err:=client.CreateOrderOrigin(prisma.OrderOriginCreateInput{
												HotelId:in.HotelId,
												HrId:"",
												AdviserId:in.AdviserId,
												Datetime:in.Date,
												Duration:in.Duration,
												Mode:0,				//TODO the create order request in .proto file  haven't set
												Count:in.Count,
												CountMale:in.CountMale,
											}).Exec(ctx)
	if err!=nil{
		panic(err)
	}

	//4 return the result
	return &pb.HotelCreateOrderReply{OrderId:newOrder.ID,CreateResult:"create order successful!"},nil
}

//TODO the way to query OrderOrigin by order_id is something wrong
func (s *Server)QueryOrder(ctx context.Context, in *pb.QueryOrderRequest) (*pb.QueryOrderReply,error) {

	log.Printf("request:%v",in)

	client:=prisma.New(nil)
	ctx=context.TODO()

	// now just for single order query by order_id
	order,err := client.OrderOrigin(prisma.OrderOriginWhereUniqueInput{ID:&in.OrderId}).Exec(ctx)
	if err!=nil{
		log.Fatalf("failed to query :%v",err)
	}

	//orders,err:=client.OrderOrigins(&prisma.OrderOriginsParams{
	//	Where:&prisma.OrderOriginWhereInput{
	//		ID:&in.OrderId,
	//	},
	//}).Exec(ctx)
	//if err!=nil{
	//	log.Fatalf("failed to query :%v",err)
	//}


	fmt.Println(order)
	return &pb.QueryOrderReply{HotelId:order.HotelId,Job:"cleaning",Date:order.Datetime,Duration:order.Duration,Count:order.Count,CountMale:order.CountMale,CountFemale:order.Count-order.CountMale,AdviserId:order.AdviserId,Mode:order.Mode},nil
	//return &pb.QueryOrderReply{HotelId:orders[0].HotelId,Job:"cleaning",Date:orders[0].DateTime,Duration:orders[0].Duration,Count:orders[0].Count,CountMale:orders[0].CountMale,CountFemale:orders[0].Count-orders[0].CountMale,AdviserId:orders[0].AdviserId,Mode:orders[0].Mode},nil

}

func main()  {
	lis,err:=net.Listen("tcp",port)
	if err!=nil{
		log.Fatalf("failed to listen: %v",err)
	}
	s:=grpc.NewServer()
	pb.RegisterOrderOperationServer(s,&Server{})
	reflection.Register(s)
	if err:=s.Serve(lis);err!=nil{
		log.Fatalf("failed to server: %v",err)
	}

}
