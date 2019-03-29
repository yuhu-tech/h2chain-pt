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

/*
	client:=prisma.New(nil)
	ctx=context.TODO()
	//3 insert data into prisma-mysql
	newOrder,err:=client.CreateOrderOrigin(prisma.OrderOriginCreateInput{}).Exec(ctx)
	if err!=nil{
		panic(err)
	}
 */

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
	ctx = context.TODO()

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
	ctx = context.TODO()

	// TODO how to achieve revision

	_, err := client.CreateOrderCandidate(prisma.OrderCandidateCreateInput{
		AdviserId:           in.AdviserId,
		PtId:                in.PtId,
		ApplyTime:           in.ApplyTime,
		SignInTime:          in.SignInTime,
		PtStatus:            1,
		RegistrationChannel: in.RegistrationChannel,
		OrderOrigin:         prisma.OrderOriginCreateOneInput{Connect: &prisma.OrderOriginWhereUniqueInput{ID: &in.OrderId}},
	}).Exec(ctx)
	if err != nil {
		log.Printf("create order candidate failed %v ", err)
		return &pb.RegistryReply{RegistryResult: 0}, nil
	}

	return &pb.RegistryReply{RegistryResult: 1}, nil
}

func (s *MutationServer) ModifyOrder(ctx context.Context, in *pb.ModifyRequest) (*pb.ModifyReply, error) {
	client := prisma.New(nil)
	ctx = context.TODO()

	// TODO
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
		return &pb.ModifyReply{ModifyResult: 0}, nil
	}

	return &pb.ModifyReply{ModifyResult: 1}, nil
}

func (s *MutationServer) ModifyOrderPT(ctx context.Context, in *pb.ModifyPtRequest) (*pb.ModifyPtReply, error) {

	// client:=prisma.New(nil)
	// ctx=context.TODO()

	// TODO

	return &pb.ModifyPtReply{ModifyResult: 1}, nil
}

func (s *MutationServer) CloseOrder(ctx context.Context, in *pb.CloseRequest) (*pb.CloseReply, error) {
	// check the close order request date
	log.Println(in.OrderId)

	// TODO

	// make dead data and return
	return &pb.CloseReply{CloseResult: 1}, nil
}
