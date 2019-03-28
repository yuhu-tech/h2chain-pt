package main

import (
	"golang.org/x/net/context"
	"log"
	"time"

	mpb "../api/mutation"
	qpb "../api/query"
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
	mc := mpb.NewMutationClient(conn)
	qc := qpb.NewQueryOrderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// create order
	//
	createOrderRes, err := mc.CreateOrder(ctx, &mpb.CreateRequest{HotelId:"01",AdviserId:"9527",Job:"cleaning",Date:201903271300,Duration:2,Count:9,CountMale:6,CountFemale:3,Mode:1})
	if err != nil {
		log.Fatalf("could not create order: %v", err)
	}
	log.Printf("%v,%v/n",createOrderRes.OrderId,createOrderRes.CreateResult)


	//query order
	queryOrderRes,err:=qc.QueryOrder(ctx,&qpb.QueryRequest{OrderId:"10010",Date:nil,Status:nil,QueryFields:&qpb.QueryRequest_QueryWhat{HotelId:1,AdviseId:1}})
	if err != nil {
		log.Fatalf("could not query order: %v", err)
	}
	log.Printf("%v,%v/n",queryOrderRes.QueryResults[0].HotelId,queryOrderRes.QueryResults[0].AdviseId)
}


