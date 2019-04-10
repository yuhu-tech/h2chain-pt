package handle

import (
	pb "../api/query"
	"../prisma"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"reflect"
	"strconv"
)

type QueryServer struct {
}

func (s *QueryServer) QueryOrder(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {

	client := prisma.New(nil)
	where := ``
	isQuery := true

	// query by orderId
	orderId := reflect.ValueOf(in.OrderId)
	fmt.Println(orderId.Interface().(string))
	if orderId.Interface().(string) != "" {
		where = `id:"` + in.OrderId + `"`
		isQuery = false
	}

	// query by status
	if in.Status == 1 || in.Status == 2 || in.Status == 3 {
		where = `status:` + strconv.Itoa(int(in.Status))
		isQuery = false
	}

	// query by date
	date := reflect.ValueOf(in.Date)
	if date.Interface().(int32) != 0 {
		var date int32
		// check the date is day
		if (in.Date+28800)%86400 == 0 {
			date = in.Date
		} else {
			in.Date -= (in.Date + 28800)%86400
			date = in.Date
		}
		dateMin := date
		dateMax := date + 86399
		where = `datetime_gte:` + strconv.Itoa(int(dateMin)) + `datetime_lte:` + strconv.Itoa(int(dateMax))
		isQuery = false
	}

	// query by hotelId
	hotelId := reflect.ValueOf(in.HotelId)
	if hotelId.Interface().(string) != "" {
		where = `hotelId:"` + in.HotelId + `"`
		isQuery = false
	}

	// query by adviserId
	adviserId := reflect.ValueOf(in.Adviser)
	if adviserId.Interface().(string) != "" {
		where = `adviserId:"` + in.Adviser + `"`
		isQuery = false
	}

	// query by ptId
	ptId := reflect.ValueOf(in.PtId)
	if ptId.Interface().(string) != "" {
		where = `orderCandidates_some:{ptId:"` + in.PtId + `"}`
		isQuery = false
	}

	// check the query
	if isQuery {
		return &pb.QueryReply{OrderOrigins: ""}, nil
	}

	query := `
	  query{
		orderOrigins(where:{` + where + `}){
	    id
	    hotelId
	    hrId
	    adviserId
	    datetime
	    duration
	    job
	    mode
	    count
	  	countMale
	    status
	    orderHotelModifies{
	      id
	      revision
	      timestamp
	      count
	      countMale
	      dateTime
	      duration
	      mode
	    }
	    orderAdviserModifies{
	      id
	      revision
	      timeStamp
	      isFloat
	      hourlySalary
	      workCount
	      attention
	    }
	    orderCandidates{
	      id
	      adviserId
	      agentId
	      ptId
	      applyTime
	      signInTime
	      ptStatus
	      ptPerformance
	      objectReason
	      registrationChannel
	    }
	  }
	}
 `
	variables := make(map[string]interface{})
	res, err := client.GraphQL(ctx, query, variables)
	if err != nil {
		return nil, err
	}
	resByte, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return &pb.QueryReply{OrderOrigins: string(resByte)}, nil
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
