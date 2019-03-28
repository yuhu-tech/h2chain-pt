package handle

import (
	pb "../api/query"
	"golang.org/x/net/context"
	"log"
)

type QueryServer struct {
}

/*
//2 create prisma client
	client:=prisma.New(nil)
	ctx=context.TODO()

	//3 insert data into prisma-mysql
	newOrder,err:=client.CreateOrderOrigin(prisma.OrderOriginCreateInput{
												HotelId:in.HotelId,
												HrId:"",
												AdviserId:in.AdviserId,
												Datetime:in.Date,
												Duration:in.Duration,
												Mode:0,
												Count:in.Count,
												CountMale:in.CountMale,
												Status:0,
											}).Exec(ctx)
	if err!=nil{
		panic(err)
	}
 */

func (s *QueryServer) QueryOrder(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {
	// get create order request
	log.Printf("request:%v", in)

	//TODO

	// return the result
	return &pb.QueryReply{}, nil
}

func (s *QueryServer) QueryOrderPT(ctx context.Context, in *pb.QueryPTRequest) (*pb.QueryPTReply, error) {
	// get create order request
	log.Printf("request:%v", in)

	//TODO

	// return the result
	return &pb.QueryPTReply{}, nil
}
