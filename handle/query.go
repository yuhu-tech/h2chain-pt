package handle

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"encoding/json"
	"golang.org/x/net/context"

	pb "../api/query"
	"../prisma"
)

type QueryServer struct {
}

func (s *QueryServer) QueryOrder(ctx context.Context, in *pb.QueryRequest) (*pb.QueryReply, error) {

	client := prisma.New(nil)
	where := ``
	isQuery := true

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
		arr := strings.Split(in.PtId, ",")
		if arr[0] == "none" {
			where = `orderCandidates_none:{ptId:"` + arr[1] + `"}`
		} else {
			where = `orderCandidates_some:{ptId:"` + arr[1] + `"}`
		}
		isQuery = false
	}

	// query by orderId
	orderId := reflect.ValueOf(in.OrderId)
	if orderId.Interface().(string) != "" {
		where = `id:"` + in.OrderId + `"`
		isQuery = false
	}

	// query by status
	if in.Status == 12 || in.Status == 1 || in.Status == 2 || in.Status == 3 {
		if in.Status == 12 {
			where = where + `status_not:3`
			isQuery = false
		} else {
			where = where + `status:` + strconv.Itoa(int(in.Status))
			isQuery = false
		}
	}

	// query by date
	date := reflect.ValueOf(in.Date)
	if date.Interface().(int32) != 0 {
		var date int32
		// check the date is day
		if (in.Date+28800)%86400 == 0 {
			date = in.Date
		} else {
			in.Date -= (in.Date + 28800) % 86400
			date = in.Date
		}
		dateMin := date
		dateMax := date + 86399
		where = where + `datetime_gte:` + strconv.Itoa(int(dateMin)) + `datetime_lte:` + strconv.Itoa(int(dateMax))
		isQuery = false
	}

	// check the query
	if isQuery {
		return &pb.QueryReply{OrderOrigins: ""}, nil
	}

	query := `
	  query{
		orderOrigins(where:{` + where + `}orderBy:datetime_DESC){
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

	// check the orderId
	orderId := reflect.ValueOf(in.OrderId)
	if orderId.Interface().(string) == "" {
		return nil, nil
	}

	where := &prisma.OrderCandidateWhereInput{}

	// query pts of order by orderId
	where.OrderOrigin = &prisma.OrderOriginWhereInput{ID: &in.OrderId}

	// query pts of order by ptId
	ptId := reflect.ValueOf(in.PtId)
	if ptId.Interface().(string) != "" {
		where.PtId = &in.PtId
	}

	// query pts of order by ptStatus
	ptStatus := reflect.ValueOf(in.PtStatus)
	if ptStatus.Interface().(int32) != 0 {
		if in.PtStatus == 13 {
			where.PtStatusIn = []int32{1, 3}
		}else {
			where.PtStatus = &in.PtStatus
		}
	}

	// query pts of order by registrationChannel
	registrationChannel := reflect.ValueOf(in.RegistrationChannel)
	if registrationChannel.Interface().(string) != "" {
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
		pt.ApplyTime = *queryRes[i].ApplyTime
		pt.SignTime = *queryRes[i].SignInTime
		pt.Id = queryRes[i].ID
		pt.PtStatus = queryRes[i].PtStatus
		result = append(result, &pt)
	}

	return &pb.QueryPTReply{OrderPts: result}, nil
}

func (s *QueryServer) QueryRemark(ctx context.Context, in *pb.QueryRemarkRequest) (*pb.QueryRemarkReply, error) {

	client := prisma.New(nil)

	orderId := reflect.ValueOf(in.OrderId)
	if orderId.Interface().(string) == "" {
		return nil, nil
	}

	ptId := reflect.ValueOf(in.PtId)
	if ptId.Interface().(string) == "" {
		return nil, nil
	}

	query := `
	  query{
	    orderCandidates(
	      where:{
	        orderOrigin:{id:"` + in.OrderId + `"}
	        ptId:"` + in.PtId + `"
	    }  
	    ){  
	      remark{
	        startDate  
	        endDate
	        realSalary
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
	return &pb.QueryRemarkReply{Remark: string(resByte)}, nil
}

func (s *QueryServer) QueryExperience(ctx context.Context, in *pb.QueryExperienceRequest) (*pb.QueryExperienceReply, error) {

	client := prisma.New(nil)

	ptId := reflect.ValueOf(in.PtId)
	if ptId.Interface().(string) == "" {
		return nil, nil
	}

	query := `
	  query{
	    orderOrigins(
	      where:{
	        orderCandidates_some:{
	          ptId:"` + in.PtId + `"
	          remark:{isWorked:1}
	        }
	      }
	      orderBy:datetime_DESC
	      ){	  
		  job
	      hotelId
	      orderCandidates{ 
	        remark{
			  ptId
	          startDate
	          endDate
	          isWorked
	        }    
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
	return &pb.QueryExperienceReply{WorkExperience: string(resByte)}, nil
}
