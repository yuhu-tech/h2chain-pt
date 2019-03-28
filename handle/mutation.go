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
	log.Println(in.HotelId, in.AdviserId, in.Job, in.Date, in.Date, in.Duration, in.Count, in.CountMale, in.CountFemale)

	//TODO

	// make dead data and return
	var createReply pb.CreateReply
	createReply.OrderId = "10010"
	createReply.CreateResult = "Create Order Successful!"

	// return the result
	return &pb.CreateReply{OrderId: "10010",}, nil
}

func (s *MutationServer) PostOrder(ctx context.Context, in *pb.PostRequest) (*pb.PostReply, error) {
	// check the create order request date
	log.Println(in.OrderId, in.IsFloat, in.HourlySalary, in.WorkContent, in.Attention)

	//TODO

	// make dead data and return
	var postOrder pb.PostReply
	postOrder.PostResult = " Post Order Successful "

	// return the result
	return &pb.PostReply{}, nil
}

func (s *MutationServer) RegisterOrder(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	// get create order request
	log.Printf("request:%v", in)

	//TODO

	// return the result
	return &pb.RegisterReply{}, nil
}

func (s *MutationServer) ModifyOrder(ctx context.Context, in *pb.ModifyRequest) (*pb.ModifyReply, error) {
	// get create order request
	log.Printf("request:%v", in)

	//TODO

	// return the result
	return &pb.ModifyReply{}, nil
}

func (s *MutationServer) CloseOrder(ctx context.Context, in *pb.CloseRequest) (*pb.CloseReply, error) {
	// get create order request
	log.Printf("request:%v", in)

	//TODO

	// return the result
	return &pb.CloseReply{}, nil
}
