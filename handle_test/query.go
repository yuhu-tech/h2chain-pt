package handle_test

import (
	pb "../api/query"
	"golang.org/x/net/context"
	"log"
	"strings"
)

type QueryServer struct {
}

func (s *QueryServer) QueryOrder(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {
	// check the query order request date
	log.Println(in.OrderId, in.Date, in.Status, in.QueryKey, in.QueryValue)

	queryKeys := strings.Split(in.QueryKey, ",")
	log.Println(queryKeys)

	//TODO

	// make dead data and return
	return &pb.QueryReply{
		Orders: []*pb.Order{{OrderId: "10010", HotelId: "001", AdviserId: "003", Date: 20190328, Duration: 2, Mode: 0, Count: 10, CountMale: 5, CountFemale: 5, Job: "cleaning", Status: 1, IsFloat: 0, HourlySalary: 10, WorkContent: "cleanRoom", Attention: "careful"}},
		OrdersHotelModifies:[]*pb.OrderHotelModifies{
			{OrderHotelModifies:[]*pb.OrderHotelModify{{Id: "001", Revision: 001, TimeStamp: 20190328, Date: 20190328, Duration: 2, Mode: 1, Count: 15, CountMale: 10}}},
	},
		OrdersAdviserModifies:[]*pb.OrderAdviserModifies{
			{OrderAdviserModifies:[]*pb.OrderAdviserModify{{Id: "001", Revision: 001, TimeStamp: 20190328, IsFloat: 0, HourlySalary: 20, WorkContent: "cleanRoom", Attention: "careful"}}},
		},
		OrdersCandidates:[]*pb.OrderCandidates{
			{OrderCandidates:[]*pb.OrderCandidate{
				{Id: "001", AdviserId: "003", PtId: "9527", ApplyTime: 2019032800, PtStatus: 1, RegistrationChannel: "WeChat"},
			},
	},
		},
	}, nil
}

func (s *QueryServer) QueryPTOfOrder(ctx context.Context, in *pb.QueryPTRequest) (*pb.QueryPTReply, error) {
	// check the query order's pt request
	log.Println(in.PtId, in.OrderId, in.RegistrationChannel, in.PtStatus)

	//TODO

	// make dead data and return
	return &pb.QueryPTReply{
		OrderPts:[]*pb.PT{
			{PtId: "001", AdviserId: "003", ApplyTime: 2019032800, OrderId: "10010"},
			{PtId: "002", AdviserId: "003", ApplyTime: 2019032811, OrderId: "10010"},
		},
	},nil
}
