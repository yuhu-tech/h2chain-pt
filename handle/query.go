package handle

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
	log.Println(in.OrderId,in.Date,in.Status,in.QueryKey,in.QueryValue)

	queryKeys:=strings.Split(in.QueryKey,",")
	log.Println(queryKeys)


	//TODO

	// make dead data and return
	var queryReply pb.QueryReply
	queryReply.Order = &pb.Order{OrderId:"10010",HotelId:"001",AdviserId:"003",Date:"201903281300",Duration:2,Mode:0,Count:10,CountMale:5,CountFemale:5,Job:"cleaning",Status:1,IsFloat:0,HourlySalary:10,WorkContent:"cleanRoom",Attention:"careful"}
	queryReply.OrderHotelModifies[0] = &pb.OrderHotelModify{Id:"001",Revision:001,TimeStamp:"201903281400",Date:"201903281300",Duration:2,Mode:1,Count:15,CountMale:10}
	queryReply.OrderAdviserModifies[0] = &pb.OrderAdviserModify{Id:"001",Revision:001,TimeStamp:"201903281500",IsFloat:0,Count:10,CountMale:5,HourlySalary:20,WorkContent:"cleanRoom",Attention:"careful"}
	queryReply.OrderCandidates[0] = &pb.OrderCandidate{Id:"001",AdviserId:"003",PtId:"9527",ApplyTime:"2019032800",PtStatus:1,RegistrationChannel:"WeChat"}
	queryReply.OrderCandidates[1] = &pb.OrderCandidate{Id:"002",AdviserId:"003",PtId:"8527",ApplyTime:"2019032811",PtStatus:1,RegistrationChannel:"WeChat"}

	// return the result
	return &queryReply, nil
}

func (s *QueryServer) QueryOrderPT(ctx context.Context, in *pb.QueryPTRequest) (*pb.QueryPTReply, error) {
	// check the query order's pt request
	log.Println(in.PtId,in.OrderId,in.RegistrationChannel,in.PtStatus)

	//TODO

	// make dead data and return
	var queryPtReply pb.QueryPTReply
	queryPtReply.OrderPts[0] = &pb.PT{PtId:"001",AdviserId:"003",ApplyTime:"2019032800",OrderId:"10010"}
	queryPtReply.OrderPts[1] = &pb.PT{PtId:"002",AdviserId:"003",ApplyTime:"2019032811",OrderId:"10010"}

	// return the result
	return &queryPtReply, nil
}
