package handle

import (
	pb "../api/query"
	"../prisma"
	"fmt"
	"golang.org/x/net/context"
	"log"
)

type QueryServer struct {
}

func (s *QueryServer) QueryOrder(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {
	client:=prisma.New(nil)
	orderOrigin:=&pb.Order{}
	orderHotelModifies:=[]*pb.OrderHotelModify{}
	orderAdviserModifies:=[]*pb.OrderAdviserModify{}
	orderCandidates:=[]*pb.OrderCandidate{}

	if in.OrderId!=""{
		queryRes,err:=client.OrderOrigin(prisma.OrderOriginWhereUniqueInput{ID:&in.OrderId}).Exec(ctx)
		if err!=nil{
			return &pb.QueryReply{},err
		}
		fmt.Println(queryRes)
	}

	// TODO





	return &pb.QueryReply{
		Order:orderOrigin,
		OrderHotelModifies:orderHotelModifies,
		OrderAdviserModifies:orderAdviserModifies,
		OrderCandidates:orderCandidates,
	}, nil
}

func (s *QueryServer) QueryPTOfOrder(ctx context.Context, in *pb.QueryPTRequest) (*pb.QueryPTReply, error) {

	client := prisma.New(nil)

	if in.OrderId == "" {
		return nil, nil
	}

	where := &prisma.OrderCandidateWhereInput{}
	where.OrderOrigin = &prisma.OrderOriginWhereInput{ID: &in.OrderId}

	if in.PtId != "" {
		where.PtId = &in.PtId
	}
	if  in.PtStatus != -1 {
		where.PtStatus = &in.PtStatus
	}
	if  in.RegistrationChannel != "" {
		where.RegistrationChannel = &in.RegistrationChannel
	}

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
		pt.Id = queryRes[i].ID
		result = append(result, &pt)
	}

	return &pb.QueryPTReply{OrderPts: result}, nil
}
