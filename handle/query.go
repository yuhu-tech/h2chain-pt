package handle

import (
	pb "../api/query"
	"../prisma"
	"golang.org/x/net/context"
	"log"
)

type QueryServer struct {
}

func QueryHMOfOrder(ctx context.Context, orderId string) ([]*pb.OrderHotelModify, error) {
	client := prisma.New(nil)
	orderHotelModifies := []*pb.OrderHotelModify{}

	queryHMRes, err := client.OrderHotelModifies(&prisma.OrderHotelModifiesParams{
		Where: &prisma.OrderHotelModifyWhereInput{OrderOrigin: &prisma.OrderOriginWhereInput{ID: &orderId}},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(queryHMRes); i++ {
		var oHM pb.OrderHotelModify
		oHM.Id = queryHMRes[i].ID
		oHM.Revision = queryHMRes[i].Revision
		oHM.TimeStamp = queryHMRes[i].Timestamp
		oHM.Count = *queryHMRes[i].Count
		oHM.CountMale = *queryHMRes[i].CountMale
		oHM.CountYet = *queryHMRes[i].CountYet
		oHM.CountMaleYet = *queryHMRes[i].CountMaleYet
		oHM.Date = *queryHMRes[i].DateTime
		oHM.Duration = *queryHMRes[i].Duration
		oHM.Mode = *queryHMRes[i].Mode
		orderHotelModifies = append(orderHotelModifies, &oHM)
	}
	return orderHotelModifies, nil
}

func QueryAMOfOrder(ctx context.Context, orderId string) ([]*pb.OrderAdviserModify, error) {
	client := prisma.New(nil)
	orderAdviserModifies := []*pb.OrderAdviserModify{}
	queryAMRes, err := client.OrderAdviserModifies(&prisma.OrderAdviserModifiesParams{
		Where: &prisma.OrderAdviserModifyWhereInput{OrderOrigin: &prisma.OrderOriginWhereInput{ID: &orderId}},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(queryAMRes); i++ {
		var oAM pb.OrderAdviserModify
		oAM.Id = queryAMRes[i].ID
		oAM.Revision = queryAMRes[i].Revision
		oAM.TimeStamp = queryAMRes[i].TimeStamp
		oAM.IsFloat = *queryAMRes[i].IsFloat
		oAM.Count = *queryAMRes[i].Count
		oAM.CountMale = *queryAMRes[i].CountMale
		oAM.HourlySalary = *queryAMRes[i].HourlySalary
		oAM.WorkContent = *queryAMRes[i].WorkCount
		oAM.Attention = *queryAMRes[i].Attention
		orderAdviserModifies = append(orderAdviserModifies, &oAM)
	}
	return orderAdviserModifies, nil
}

func QueryCOfOrder(ctx context.Context, orderId string) ([]*pb.OrderCandidate, error) {
	client := prisma.New(nil)
	orderCandidates := []*pb.OrderCandidate{}
	queryOCRes, err := client.OrderCandidates(&prisma.OrderCandidatesParams{
		Where: &prisma.OrderCandidateWhereInput{OrderOrigin: &prisma.OrderOriginWhereInput{ID: &orderId}},
	}).Exec(ctx)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(queryOCRes); i++ {
		var oC pb.OrderCandidate
		oC.Id = queryOCRes[i].ID
		oC.AdviserId = queryOCRes[i].AdviserId
		oC.AgentId = queryOCRes[i].AgentId
		oC.PtId = queryOCRes[i].PtId
		oC.ApplyTime = queryOCRes[i].ApplyTime
		oC.SignInTime = queryOCRes[i].SignInTime
		oC.PtStatus = queryOCRes[i].PtStatus
		oC.PtPerformance = queryOCRes[i].PtPerformance
		oC.ObjectReason = queryOCRes[i].ObjectReason
		oC.RegistrationChannel = queryOCRes[i].RegistrationChannel
		orderCandidates = append(orderCandidates, &oC)
	}
	return orderCandidates, nil
}

func (s *QueryServer) QueryOrder(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {
	client := prisma.New(nil)
	orderOrigins := []*pb.Order{}


	where := &prisma.OrderOriginWhereInput{}
	if in.OrderId != "" {
		where.ID = &in.OrderId
	}
	if in.Date != -1 {
		dateMin := in.Date
		dateMax := in.Date + 86399
		where.DatetimeLte = &dateMin
		where.DatetimeGte = &dateMax
	}
	if in.Status != -1 {
		where.Status = &in.Status
	}

	queryRes, err := client.OrderOrigins(&prisma.OrderOriginsParams{
		Where: where,
	}).Exec(ctx)
	if err != nil {
		return &pb.QueryReply{}, err
	}
	for i := 0; i < len(queryRes); i++ {
		var orderOrigin pb.Order
		orderOrigin.OrderId = queryRes[i].ID
		orderOrigin.HotelId = queryRes[i].HotelId
		orderOrigin.AdviserId = queryRes[i].AdviserId
		orderOrigin.Date = queryRes[i].Datetime
		orderOrigin.Duration = queryRes[i].Duration
		orderOrigin.Count = queryRes[i].Count
		orderOrigin.CountMale = queryRes[i].CountMale
		orderOrigin.CountFemale = queryRes[i].Count - queryRes[i].CountMale
		orderOrigin.Status = queryRes[i].Status
		orderOrigin.Job = queryRes[i].Job
		orderOrigin.Mode = queryRes[i].Mode
		orderOrigins = append(orderOrigins, &orderOrigin)
	}

	// get the orderHotelModify orderAdviserMOdify candidate of order by orderId
	orderHotelModifies, _ := QueryHMOfOrder(ctx, orderOrigins[0].OrderId)
	orderAdviserModifies, _ := QueryAMOfOrder(ctx, orderOrigins[0].OrderId)
	orderCandidates, _ := QueryCOfOrder(ctx, orderOrigins[0].OrderId)

	return &pb.QueryReply{
		Orders: orderOrigins,
		OrderHotelModifies:   orderHotelModifies,
		OrderAdviserModifies: orderAdviserModifies,
		OrderCandidates:      orderCandidates,
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
	if in.PtStatus != -1 {
		where.PtStatus = &in.PtStatus
	}
	if in.RegistrationChannel != "" {
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

func NativeGrapgql() map[string]interface{} {
	client := prisma.New(nil)
	ctx := context.TODO()
	query := `
	  query{
		orderOrigins(where:{status:1}){
    	  id
    	  hotelId
    	  adviserId
      }
    }
  `
	variables := make(map[string]interface{})
	res, _ := client.GraphQL(ctx, query, variables)
	return res
}
