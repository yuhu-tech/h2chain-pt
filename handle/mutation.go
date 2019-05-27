package handle

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	pb "../api/mutation"
	"../prisma"
	"golang.org/x/net/context"
)

type MutationServer struct {
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
	// TODO add transaction to bind create and update

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

	_, err := client.CreateOrderCandidate(prisma.OrderCandidateCreateInput{
		AdviserId:           in.AdviserId,
		PtId:                in.PtId,
		ApplyTime:           &in.ApplyTime,
		SignInTime:          &in.SignInTime,
		PtStatus:            in.PtStatus,
		RegistrationChannel: &in.RegistrationChannel,
		OrderOrigin:         prisma.OrderOriginCreateOneWithoutOrderCandidatesInput{Connect: &prisma.OrderOriginWhereUniqueInput{ID: &in.OrderId}},
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
	//TODO 测试环境 修改当天的订单
	//var cleanDate = int32(in.Date - 3*24*3600)
	var cleanDate = int32(in.Date)
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
