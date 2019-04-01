package handle

import (
	pb "../api/mutation"
	"../prisma"
	"golang.org/x/net/context"
	"log"
	"time"
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
			Job:       in.Job, Mode: in.Mode,
			Count:     in.Count,
			CountMale: in.CountMale,
			Status:    1,
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
		Revision:     01,
		TimeStamp:    int32(time.Now().Unix()),
		IsFloat:      &in.IsFloat,
		HourlySalary: &in.HourlySalary,
		WorkCount:    &in.WorkContent,
		Attention:    &in.Attention,
		OrderOrigin:  prisma.OrderOriginCreateOneInput{Connect: &prisma.OrderOriginWhereUniqueInput{ID: &in.OrderId}},
	}).Exec(ctx)
	if err != nil {
		log.Printf("post order filed %v", err)
		return &pb.PostReply{PostResult: 0}, err
	}

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
		OrderOrigin:         prisma.OrderOriginCreateOneInput{Connect: &prisma.OrderOriginWhereUniqueInput{ID: &in.OrderId}},
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
		Revision:    01,
		Timestamp:   int32(time.Now().Unix()),
		Count:       &in.CountChanged,
		CountMale:   &in.CountMaleChanged,
		DateTime:    &in.DateChanged,
		Duration:    &durationChanged,
		OrderOrigin: prisma.OrderOriginCreateOneInput{Connect: &prisma.OrderOriginWhereUniqueInput{ID: &in.OrderId}},
	}).Exec(ctx)
	if err != nil {
		log.Printf("create order candidate failed %v ", err)
		return &pb.ModifyReply{ModifyResult: 0}, err
	}

	return &pb.ModifyReply{ModifyResult: 1}, nil
}

func (s *MutationServer) ModifyPTOfOrder(ctx context.Context, in *pb.ModifyPtRequest) (*pb.ModifyPtReply, error) {

	client := prisma.New(nil)

	data := prisma.OrderCandidateUpdateInput{}
	if in.TargetStatus != -1 {
		data.PtStatus = &in.TargetStatus
	}
	if in.PtPerformance != -1 {
		data.PtPerformance = &in.PtPerformance
	}
	if in.ObjectReason != -1 {
		data.ObjectReason = &in.ObjectReason
	}
	_, err := client.UpdateOrderCandidate(prisma.OrderCandidateUpdateParams{
		Data:  data,
		Where: prisma.OrderCandidateWhereUniqueInput{ID: &in.Id},
	}).Exec(ctx)
	if err != nil {
		log.Printf("create order candidate failed %v ", err)
		return &pb.ModifyPtReply{ModifyResult: 0}, err
	}

	return &pb.ModifyPtReply{ModifyResult: 1}, nil
}

func (s *MutationServer) CloseOrder(ctx context.Context, in *pb.CloseRequest) (*pb.CloseReply, error) {

	client := prisma.New(nil)
	var closeStatus int32 = 0

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
