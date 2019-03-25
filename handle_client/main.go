package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"time"

	pb "../api"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderOperationClient(conn)

	// Contact the server and print out its response.

	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//hotel create order request
	//createOrderRes, err := c.HotelCreateOrder(ctx, &pb.HotelCreateOrderRequest{HotelId:"01",Job:"",Date:20190324,Duration:2,Count:10,CountMale:5,CountFemale:5,AdviserId:"9527"})
	//if err != nil {
	//	log.Fatalf("could not create order: %v", err)
	//}
	//log.Printf("%s,%v", createOrderRes.CreateResult,createOrderRes.OrderId)

	//query order
	queryOrderRes,err:=c.QueryOrder(ctx,&pb.QueryOrderRequest{OrderId:"cjtobqrsz000g0963haqp9fmk"})
	if err != nil {
		log.Fatalf("could not query order: %v", err)
	}
	fmt.Printf("hotel_id:%v,adviser_id:%v,date:%v,duration:%v,mode:%v,count:%v,countmale:%v,countFemale:%v",queryOrderRes.HotelId,queryOrderRes.AdviserId,queryOrderRes.Date,queryOrderRes.Duration,queryOrderRes.Mode,queryOrderRes.Count,queryOrderRes.CountMale,queryOrderRes.CountFemale)
}


