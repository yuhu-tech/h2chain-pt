package handle

import (
	"log"
	"reflect"
	"time"

	"golang.org/x/net/context"

	pb "../api/mutation"
	"../prisma"
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
		ApplyTime:           in.ApplyTime,
		SignInTime:          in.SignInTime,
		PtStatus:            in.PtStatus,
		RegistrationChannel: in.RegistrationChannel,
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
	var closeStatus int32 = 2

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
