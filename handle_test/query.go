package handle_test

import (
	"log"

	"golang.org/x/net/context"

	pb "../api/query"
)

type QueryServer struct {
}

func (s *QueryServer) QueryOrder(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {
	// check the query order request date
	log.Println(in.OrderId, in.Date, in.Status,in.HotelId,in.Adviser,in.HrId,in.PtId )

	//TODO

	// make dead data and return
	return &pb.QueryReply{OrderOrigins:"HelloWorld"}, nil
}

func (s *QueryServer) QueryPTOfOrder(ctx context.Context, in *pb.QueryPTRequest) (*pb.QueryPTReply, error) {
	// check the query order's pt request
	log.Println(in.PtId, in.OrderId, in.RegistrationChannel, in.PtStatus)

	//TODO

	// make dead data and return
	return &pb.QueryPTReply{
		OrderPts: []*pb.PT{
			{PtId: "001", AdviserId: "003", ApplyTime: 2019032800, OrderId: "10010"},
			{PtId: "002", AdviserId: "003", ApplyTime: 2019032811, OrderId: "10010"},
		},
	}, nil
}
