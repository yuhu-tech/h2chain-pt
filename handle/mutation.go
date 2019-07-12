package handle

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"encoding/json"

	"golang.org/x/net/context"

	pb "../api/mutation"
	"../prisma"
)

type MutationServer struct {
}

type Order struct {
	OrderOrigin struct {
		Datetime           int32  `json:"datetime"`
		Duration           int32  `json:"duration"`
		ID                 string `json:"id"`
		OrderHotelModifies []struct {
			DateTime int32 `json:"dateTime"`
			Duration int32 `json:"duration"`
		} `json:"orderHotelModifies"`
	} `json:"orderOrigin"`
}

type Orders struct {
	OrderOrigins []struct {
		Datetime           int32 `json:"datetime"`
		Duration           int32 `json:"duration"`
		OrderHotelModifies []struct {
			DateTime int32 `json:"dateTime"`
			Duration int32 `json:"duration"`
		} `json:"orderHotelModifies"`
	} `json:"orderOrigins"`
}

func (s *MutationServer) CreateOrder(ctx context.Context, in *pb.CreateRequest) (*pb.CreateReply, error) {

	client := prisma.New(nil)
	ctx = context.TODO()

	createResult, err := client.CreateOrderOrigin(
		prisma.OrderOriginCreateInput{
			HotelId:   in.HotelId,
			AdviserId: in.AdviserId,
			Datetime:  in.Date,
			Duration:  in.Duration * 3600,
			Job:       in.Job,
			Count:     in.Count,
			CountMale: in.CountMale,
			Status:    1,
			Mode:      in.Mode,
		}).Exec(ctx)
	if err != nil {
		log.Fatalf(" create order err: %v", err)
		return &pb.CreateReply{CreateResult: 0}, err
	}

	return &pb.CreateReply{OrderId: createResult.ID, CreateResult: 1}, nil
}

func (s *MutationServer) PostOrder(ctx context.Context, in *pb.PostRequest) (*pb.PostReply, error) {

	client := prisma.New(nil)
	// TODO how to achieve revision

	// create orderAdviserModify
	_, err := client.CreateOrderAdviserModify(prisma.OrderAdviserModifyCreateInput{
		Revision:     int32(time.Now().Unix()),
		TimeStamp:    int32(time.Now().Unix()),
		IsFloat:      &in.IsFloat,
		HourlySalary: &in.HourlySalary,
		WorkCount:    &in.WorkContent,
		Attention:    &in.Attention,
		OrderOrigin:  prisma.OrderOriginCreateOneWithoutOrderAdviserModifiesInput{Connect: &prisma.OrderOriginWhereUniqueInput{ID: &in.OrderId}},
	}).Exec(ctx)
	if err != nil {
		log.Printf("post order filed %v", err)
		return &pb.PostReply{PostResult: 0}, err
	}

	// modify orderOrigin status
	var status int32 = 2
	_, err = client.UpdateOrderOrigin(prisma.OrderOriginUpdateParams{
		Data:  prisma.OrderOriginUpdateInput{Status: &status},
		Where: prisma.OrderOriginWhereUniqueInput{ID: &in.OrderId},
	}).Exec(ctx)

	return &pb.PostReply{PostResult: 1}, nil
}

func (s *MutationServer) RegistryOrder(ctx context.Context, in *pb.RegistryRequest) (*pb.RegistryReply, error) {

	client := prisma.New(nil)

	res, err := RegistryConflict(in.OrderId, in.PtId)
	if err != nil {
		return &pb.RegistryReply{RegistryResult: 0}, err
	}
	if res == true {
		return &pb.RegistryReply{RegistryResult: 2}, err
	}

	_, err = client.CreateOrderCandidate(prisma.OrderCandidateCreateInput{
		AdviserId:   in.AdviserId,
		PtId:        in.PtId,
		ApplyTime:   &in.ApplyTime,
		SignInTime:  &in.SignInTime,
		PtStatus:    in.PtStatus,
		Type:        in.Type,
		InviterId:   &in.InviterId,
		OrderOrigin: prisma.OrderOriginCreateOneWithoutOrderCandidatesInput{Connect: &prisma.OrderOriginWhereUniqueInput{ID: &in.OrderId}},
	}).Exec(ctx)
	if err != nil {
		log.Printf("create order candidate failed %v ", err)
		return &pb.RegistryReply{RegistryResult: 0}, err
	}

	return &pb.RegistryReply{RegistryResult: 1}, nil
}

func (s *MutationServer) ModifyOrder(ctx context.Context, in *pb.ModifyRequest) (*pb.ModifyReply, error) {

	client := prisma.New(nil)
	var durationChanged = in.DurationChanged * 3600

	_, err := client.CreateOrderHotelModify(prisma.OrderHotelModifyCreateInput{
		Revision:    int32(time.Now().Unix()),
		Timestamp:   int32(time.Now().Unix()),
		Count:       &in.CountChanged,
		CountMale:   &in.CountMaleChanged,
		DateTime:    &in.DateChanged,
		Duration:    &durationChanged,
		Mode:        &in.Mode,
		OrderOrigin: prisma.OrderOriginCreateOneWithoutOrderHotelModifiesInput{Connect: &prisma.OrderOriginWhereUniqueInput{ID: &in.OrderId}},
	}).Exec(ctx)
	if err != nil {
		log.Printf("create order candidate failed %v ", err)
		return &pb.ModifyReply{ModifyResult: 0}, err
	}

	return &pb.ModifyReply{ModifyResult: 1}, nil
}

func (s *MutationServer) ModifyPTOfOrder(ctx context.Context, in *pb.ModifyPtRequest) (*pb.ModifyPtReply, error) {
	client := prisma.New(nil)

	orderId := reflect.ValueOf(in.OrderId)
	if orderId.Interface().(string) == "" {
		return nil, nil
	}

	data := prisma.OrderCandidateUpdateManyMutationInput{}
	targetStatus := reflect.ValueOf(in.TargetStatus)
	if targetStatus.Interface().(int32) != 0 {
		if in.TargetStatus == 3 {
			result, err := RegistryConflict(in.OrderId, in.PtId)
			if err!= nil {
				return &pb.ModifyPtReply{ModifyResult: 0}, err
			}
			if result == true {
				return &pb.ModifyPtReply{ModifyResult: 2}, nil
			}
		}
		data.PtStatus = &in.TargetStatus
	}

	ptPerformance := reflect.ValueOf(in.PtPerformance)
	if ptPerformance.Interface().(int32) != 0 {
		data.PtPerformance = &in.PtPerformance
	}

	objectReason := reflect.ValueOf(in.ObjectReason)
	if objectReason.Interface().(int32) != 0 {
		data.ObjectReason = &in.ObjectReason
	}

	where := &prisma.OrderCandidateWhereInput{}
	where.OrderOrigin = &prisma.OrderOriginWhereInput{ID: &in.OrderId}

	ptId := reflect.ValueOf(in.PtId)
	if ptId.Interface().(string) != "" {
		where.PtId = &in.PtId
	}

	sourceStatus := reflect.ValueOf(in.SourceStatus)
	if sourceStatus.Interface().(int32) != 0 {
		where.PtStatus = &in.SourceStatus
	}

	_, err := client.UpdateManyOrderCandidates(prisma.OrderCandidateUpdateManyParams{
		Data:  data,
		Where: where,
	}).Exec(ctx)
	if err != nil {
		log.Printf("create order candidate failed %v ", err)
		return &pb.ModifyPtReply{ModifyResult: 0}, err
	}

	return &pb.ModifyPtReply{ModifyResult: 1}, nil
}

func (s *MutationServer) CloseOrder(ctx context.Context, in *pb.CloseRequest) (*pb.CloseReply, error) {

	client := prisma.New(nil)
	var closeStatus int32 = 3

	_, err := client.UpdateOrderOrigin(prisma.OrderOriginUpdateParams{
		Data:  prisma.OrderOriginUpdateInput{Status: &closeStatus},
		Where: prisma.OrderOriginWhereUniqueInput{ID: &in.OrderId},
	}).Exec(ctx)
	if err != nil {
		log.Printf("close order failed err : %v", err)
		return &pb.CloseReply{CloseResult: 0}, err
	}

	return &pb.CloseReply{CloseResult: 1}, nil
}

func (s *MutationServer) EditRemark(ctx context.Context, in *pb.EditRequest) (*pb.EditReply, error) {

	client := prisma.New(nil)
	res, err := client.OrderCandidates(&prisma.OrderCandidatesParams{
		Where: &prisma.OrderCandidateWhereInput{
			OrderOrigin: &prisma.OrderOriginWhereInput{ID: &in.OrderId},
			PtId:        &in.PtId,
		},
	}).Exec(ctx)
	if err != nil {
		return &pb.EditReply{EditResult: 0}, err
	}

	remark := &prisma.RemarkUpdateOneWithoutOrderCandidateInput{}
	typeValue := reflect.ValueOf(in.Type)
	if typeValue.Interface().(int32) != 0 {
		if in.Type == 1 {
			remark.Create = &prisma.RemarkCreateWithoutOrderCandidateInput{
				PtId:       in.PtId,
				StartDate:  &in.StartDate,
				EndDate:    &in.EndDate,
				RealSalary: &in.RealSalary,
				IsWorked:   in.IsWorked,
			}
		} else if in.Type == 2 {
			remark.Update = &prisma.RemarkUpdateWithoutOrderCandidateDataInput{
				PtId:       &in.PtId,
				StartDate:  &in.StartDate,
				EndDate:    &in.EndDate,
				RealSalary: &in.RealSalary,
				IsWorked:   &in.IsWorked,
			}
		} else {
			return &pb.EditReply{EditResult: 0}, err
		}
	}

	_, err = client.UpdateOrderCandidate(prisma.OrderCandidateUpdateParams{
		Where: prisma.OrderCandidateWhereUniqueInput{ID: &res[0].ID},
		Data: prisma.OrderCandidateUpdateInput{
			Remark: remark,
		},
	}).Exec(ctx)
	if err != nil {
		return &pb.EditReply{EditResult: 0}, err
	}

	return &pb.EditReply{EditResult: 1}, nil
}

func (s *MutationServer) CleanOrder(ctx context.Context, in *pb.CleanRequest) (*pb.CleanReply, error) {
	client := prisma.New(nil)

	// 获取传入时间，转化为三天前时间，更新当天订单状态
	inDate := reflect.ValueOf(in.Date)
	if inDate.Interface().(int32) == 0 {
		return nil, fmt.Errorf("clean order failed! the date is nil")
	}

	var cleanDate = int32(in.Date - 3*24*3600)
	//var cleanDate = int32(in.Date)
	var date int32

	if (cleanDate+28800)%86400 == 0 {
		date = cleanDate
	} else {
		cleanDate -= (cleanDate + 28800) % 86400
		date = cleanDate
	}
	dateMin := date
	dateMax := date + 86399

	// 将更新的订单信息返回
	where := `status_not:3`
	where = where + `datetime_gte:` + strconv.Itoa(int(dateMin)) + `datetime_lte:` + strconv.Itoa(int(dateMax))
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
	      count
	      countMale
	      dateTime
	      duration
	      mode
	    }
	    orderCandidates{
	      id
	      ptId
	      ptStatus
          remark{
            isWorked
          }
	    }
	  }
	}
 `

	variables := make(map[string]interface{})
	res, err := client.GraphQL(ctx, query, variables)
	if err != nil {
		return nil, fmt.Errorf("clean order failed! qeury order err: %s", err)
	}

	resByte, err := json.Marshal(res)
	if err != nil {
		return nil, fmt.Errorf("clean order failed! qeury order err: %s", err)
	}

	// 更新订单状态
	var targetStatus int32 = 3
	_, err = client.UpdateManyOrderOrigins(prisma.OrderOriginUpdateManyParams{
		Where: &prisma.OrderOriginWhereInput{
			DatetimeGte: &dateMin,
			DatetimeLte: &dateMax,
		},
		Data: prisma.OrderOriginUpdateManyMutationInput{
			Status: &targetStatus,
		},
	}).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("clean order failed! update order err: %s", err)
	}

	return &pb.CleanReply{OrderOrigins: string(resByte)}, nil
}

func (s *MutationServer) TransmitOrder(ctx context.Context, in *pb.TransmitRequest) (*pb.TransmitReply, error) {
	client := prisma.New(nil)

	// 防重保护
	res, err := client.OrderAgents(&prisma.OrderAgentsParams{
		Where: &prisma.OrderAgentWhereInput{
			AgentId: &in.AgentId,
			OrderId: &in.OrderId,
		},
	}).Exec(ctx)
	if err != nil {
		return &pb.TransmitReply{TransmitResult: 0}, fmt.Errorf("query the record of order of agent")
	}
	if len(res) != 0 {
		return &pb.TransmitReply{TransmitResult: 0}, fmt.Errorf("the agent has bind the order")
	}

	_, err = client.CreateOrderAgent(prisma.OrderAgentCreateInput{
		OrderId: in.OrderId,
		AgentId: in.AgentId,
	}).Exec(ctx)
	if err != nil {
		return &pb.TransmitReply{TransmitResult: 0}, err
	}
	return &pb.TransmitReply{TransmitResult: 1}, nil
}

func RegistryConflict(orderId, ptId string) (bool, error) {
	client := prisma.New(nil)
	ctx := context.TODO()

	// 获取待报名订单时间信息
	query := `
	  query{
		orderOrigin(where:{id:"` + orderId + `"}){
		  id
		  datetime
		  duration
		  orderHotelModifies{
			dateTime
			duration
		  }
		}
	  }
	`
	variables := make(map[string]interface{})
	res, err := client.GraphQL(ctx, query, variables)
	if err != nil {
		return true, err
	}
	resByte, err := json.Marshal(res)
	if err != nil {
		return true, err
	}
	var targetOrder Order
	json.Unmarshal(resByte, &targetOrder)

	// 获取 pt 已报名且未关闭订单的时间信息
	query = `
	  query{
		orderOrigins(
		  where:{
			  status:2
			orderCandidates_some:{
			  ptId:"` + ptId + `"
			}
		  }
		){
		  datetime
		  duration
		  orderHotelModifies{
			dateTime
			duration
		  }
		}
	  }
	`

	result, err := client.GraphQL(ctx, query, variables)
	if err != nil {
		return true, err
	}
	resultByte, err := json.Marshal(result)
	if err != nil {
		return true, err
	}
	var registeredOrders Orders
	json.Unmarshal(resultByte, &registeredOrders)

	// fmt.Println("targetOrder:", targetOrder)
	// fmt.Println("registeredOrders:", registeredOrders)

	// 判断是否产生时间冲突
	var tDatetime int32
	var tDuration int32
	if len(targetOrder.OrderOrigin.OrderHotelModifies) != 0 {
		tDatetime = targetOrder.OrderOrigin.OrderHotelModifies[0].DateTime
		tDuration = targetOrder.OrderOrigin.OrderHotelModifies[0].Duration
	} else {
		tDatetime = targetOrder.OrderOrigin.Datetime
		tDuration = targetOrder.OrderOrigin.Duration
	}
	tarMax := tDatetime + tDuration

	for _, ele := range registeredOrders.OrderOrigins {
		var da int32
		var du int32
		da = ele.Datetime
		du = ele.Duration
		if len(ele.OrderHotelModifies) != 0 {
			da = ele.OrderHotelModifies[0].DateTime
			du = ele.OrderHotelModifies[0].Duration
		}
		regMax := da + du

		if tDatetime <= regMax && tarMax >= da {
			return true, nil
		}

	}

	return false, nil
}
