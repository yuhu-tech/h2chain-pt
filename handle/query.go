package handle

import (
	pb "../api/query"
	"../prisma"
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
		Order: &pb.Order{OrderId: "10010", HotelId: "001", AdviserId: "003", Date: 20190328, Duration: 2, Mode: 0, Count: 10, CountMale: 5, CountFemale: 5, Job: "cleaning", Status: 1, IsFloat: 0, HourlySalary: 10, WorkContent: "cleanRoom", Attention: "careful"},
		OrderHotelModifies: []*pb.OrderHotelModify{
			{Id: "001", Revision: 001, TimeStamp: 20190328, Date: 20190328, Duration: 2, Mode: 1, Count: 15, CountMale: 10},
		},
		OrderAdviserModifies: []*pb.OrderAdviserModify{
			{Id: "001", Revision: 001, TimeStamp: 20190328, IsFloat: 0, Count: 10, CountMale: 5, HourlySalary: 20, WorkContent: "cleanRoom", Attention: "careful"},
		},
		OrderCandidates: []*pb.OrderCandidate{
			{Id: "001", AdviserId: "003", PtId: "9527", ApplyTime: 20190328, PtStatus: 1, RegistrationChannel: "WeChat"},
			{Id: "002", AdviserId: "003", PtId: "8527", ApplyTime: 20190328, PtStatus: 1, RegistrationChannel: "WeChat"},
		},
	}, nil
}

func (s *QueryServer) QueryOrderPT(ctx context.Context, in *pb.QueryPTRequest) (*pb.QueryPTReply, error) {

	client := prisma.New(nil)
	ctx = context.TODO()


	where := &prisma.OrderCandidateWhereInput{}
	where.OrderOrigin = &prisma.OrderOriginWhereInput{ID: &in.OrderId}
	where.PtStatus = &in.PtStatus
	where.PtId = &in.PtId
	where.RegistrationChannel = &in.RegistrationChannel
	//where.And = []prisma.OrderCandidateWhereInput{}

	queryRes, err := client.OrderCandidates(&prisma.OrderCandidatesParams{Where: where}).Exec(ctx)
	if err != nil {
		log.Printf(" query order's pt filed %v ", err)
		return nil, err
	}

	var result []*pb.PT
	for i := 0; i < len(queryRes); i++ {
		var pt pb.PT
		pt.OrderId = in.OrderId
		pt.PtId = queryRes[i].PtId
		pt.AdviserId = queryRes[i].AdviserId
		pt.ApplyTime = queryRes[i].ApplyTime
		pt.SignTime = queryRes[i].SignInTime
		result = append(result, &pt)
	}

	return &pb.QueryPTReply{OrderPts: result}, nil
}
