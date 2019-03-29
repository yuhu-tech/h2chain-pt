package handle_test

import (
	pb "../api/mutation"
	"golang.org/x/net/context"
	"log"
)

type MutationServer struct {
}

func (s *MutationServer) CreateOrder(ctx context.Context, in *pb.CreateRequest) (*pb.CreateReply, error) {
	// check the create order request date
	log.Println(in.HotelId, in.AdviserId, in.Job, in.Date, in.Duration, in.Count, in.CountMale, in.CountFemale, in.Mode)

	// TODO

	// make dead data and return
	return &pb.CreateReply{OrderId:"10010",CreateResult:1},nil
}

func (s *MutationServer) PostOrder(ctx context.Context, in *pb.PostRequest) (*pb.PostReply, error) {
	// check the post order request date
	log.Println(in.OrderId, in.IsFloat, in.HourlySalary, in.WorkContent, in.Attention)

	// TODO

	// make dead data and return
	return &pb.PostReply{PostResult:1},nil
}

func (s *MutationServer) RegistryOrder(ctx context.Context, in *pb.RegistryRequest) (*pb.RegistryReply, error) {
	// get the registry order request
	log.Println(in.OrderId, in.AdviserId, in.PtId, in.ApplyTime, in.SignInTime, in.PtStatus)

	// TODO

	// make dead data and return
	return &pb.RegistryReply{RegistryResult:1},nil
}

func (s *MutationServer) ModifyOrder(ctx context.Context, in *pb.ModifyRequest) (*pb.ModifyReply, error) {
	// get the modify order request
	log.Println(in.OrderId, in.DateChanged, in.DurationChanged, in.CountChanged, in.CountMaleChanged, in.Mode)

	// TODO

	// make dead data and  return
	return &pb.ModifyReply{ModifyResult:1},nil
}

func (s *MutationServer) ModifyOrderPT(ctx context.Context, in *pb.ModifyPtRequest) (*pb.ModifyPtReply, error) {
	// check the modify order's pt request date
	log.Println(in.OrderId, in.PtId, in.TargetStatus, in.SourceStatus, in.PtPerformance, in.ObjectReason)

	// TODO

	// make dead data and return
	return &pb.ModifyPtReply{ModifyResult:1}, nil
}

func (s *MutationServer) CloseOrder(ctx context.Context, in *pb.CloseRequest) (*pb.CloseReply, error) {
	// check the close order request date
	log.Println(in.OrderId)

	// TODO

	// make dead data and return
	return &pb.CloseReply{CloseResult:1}, nil
}
