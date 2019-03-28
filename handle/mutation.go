package handle

import (
	pb "../api/mutation"
	"golang.org/x/net/context"
	"log"
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
	// check the create order request date
	log.Println(in.HotelId, in.AdviserId, in.Job, in.Date, in.Date, in.Duration, in.Count, in.CountMale, in.CountFemale, in.Mode)

	// TODO

	// make dead data and return
	var createReply pb.CreateReply
	createReply.OrderId = "10010"
	createReply.CreateResult = " Create Order Successful! "

	// return the result
	return &createReply, nil
}

func (s *MutationServer) PostOrder(ctx context.Context, in *pb.PostRequest) (*pb.PostReply, error) {
	// check the post order request date
	log.Println(in.OrderId, in.IsFloat, in.HourlySalary, in.WorkContent, in.Attention)

	// TODO

	// make dead data and return
	var postReply pb.PostReply
	postReply.PostResult = " Post Order Successful "

	// return the result
	return &postReply, nil
}

func (s *MutationServer) RegisterOrder(ctx context.Context, in *pb.RegistryRequest) (*pb.RegistryReply, error) {
	// get the registry order request
	log.Println(in.OrderId, in.AdviserId, in.PtId, in.ApplyTime, in.SignInTime, in.PtStatus)

	// TODO

	// make dead data and return
	var registryReply pb.RegistryReply
	registryReply.RegistryResult = " registry order successful "

	// return the result
	return &registryReply, nil
}

func (s *MutationServer) ModifyOrder(ctx context.Context, in *pb.ModifyRequest) (*pb.ModifyReply, error) {
	// get the modify order request
	log.Println(in.OrderId, in.DateChanged, in.DurationChanged, in.CountChanged, in.CountMaleChanged, in.Mode)

	// TODO

	// make dead data and  return
	var modifyReply pb.ModifyReply
	modifyReply.ModifyResult = " modify order successful "

	// return the result
	return &modifyReply, nil
}

func (s *MutationServer) ModifyOrderPT(ctx context.Context, in *pb.ModifyPtRequest) (*pb.ModifyPtReply, error) {
	// check the modify order's pt request date
	log.Println(in.OrderId, in.PtId, in.TargetStatus, in.SourceStatus, in.PtPerformance, in.ObjectReason)

	// TODO

	// make dead data and return
	var modifyPtReply pb.ModifyPtReply
	modifyPtReply.ModifyResult = " modify order's pt successful "

	// return the result
	return &modifyPtReply, nil
}

func (s *MutationServer) CloseOrder(ctx context.Context, in *pb.CloseRequest) (*pb.CloseReply, error) {
	// check the close order request date
	log.Println(in.OrderId)

	// TODO

	// make dead data and return
	var closeReply pb.CloseReply
	closeReply.CloseResult = " close order successful "

	// return the result
	return &closeReply, nil
}
