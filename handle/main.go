package handle

import (
	cpb "../api/mutation"
	qpb "../api/query"
	"../prisma"
	"golang.org/x/net/context"
	"log"
	"strconv"
)

type Server struct {

}

func (s *Server)HotelCreateOrder(ctx context.Context, in *cpb.HotelCreateOrderRequest) (*cpb.HotelCreateOrderReply,error) {
	//1 get create order request
	log.Printf("request:%v",in)

	//2 create prisma client
	client:=prisma.New(nil)
	ctx=context.TODO()

	//3 insert data into prisma-mysql
	newOrder,err:=client.CreateOrderOrigin(prisma.OrderOriginCreateInput{
												HotelId:in.HotelId,
												HrId:"",
												AdviserId:in.AdviserId,
												Datetime:in.Date,
												Duration:in.Duration,
												Mode:0,				//TODO the create order request in .proto file  haven't set
												Count:in.Count,
												CountMale:in.CountMale,
												Status:0,
											}).Exec(ctx)
	if err!=nil{
		panic(err)
	}

	//4 return the result
	return &cpb.HotelCreateOrderReply{OrderId:newOrder.ID,CreateResult:"create order successful!"},nil
}

//TODO the way to query OrderOrigin by order_id is something wrong
//func (s *Server)QueryOrder(ctx context.Context, in *pb.QueryOrderRequest) (*pb.QueryOrderReply,error) {
//
//	log.Printf("request:%v",in)
//
//	client:=prisma.New(nil)
//	ctx=context.TODO()
//
//	// now just for single order query by order_id
//	order,err := client.OrderOrigin(prisma.OrderOriginWhereUniqueInput{ID:&in.OrderId}).Exec(ctx)
//	if err!=nil{
//		log.Fatalf("failed to query :%v",err)
//	}
//
//	return &pb.QueryOrderReply{HotelId:order.HotelId,Job:"cleaning",Date:order.Datetime,Duration:order.Duration,Count:order.Count,CountMale:order.CountMale,CountFemale:order.Count-order.CountMale,AdviserId:order.AdviserId,Mode:order.Mode},nil
//
//}

//query by order_id date status
func (s *Server)QueryOrder(ctx context.Context, in *qpb.QueryRequest) (*qpb.QueryReply,error) {
	log.Printf("request:%v",in)

	client:=prisma.New(nil)
	ctx=context.TODO()

	if in.Who.Hotel!=""{
		var orderCount int32 =1
		var order qpb.QueryReply_OrderOrigin
		var orderm qpb.QueryReply_OrderHotelModify

		orderOrigin,err:=client.OrderOrigin(prisma.OrderOriginWhereUniqueInput{ID:&in.OrderId}).Exec(ctx)
		if err!=nil{
			log.Fatalf("failed to query :%v",err)
		}
		order.Id=orderOrigin.ID
		order.HotelId= orderOrigin.HotelId
		order.HrId=orderOrigin.HrId
		order.AdviserId=orderOrigin.AdviserId
		order.Date= strconv.Itoa(int(orderOrigin.Datetime))
		order.Duration=orderOrigin.Duration
		order.Mode=orderOrigin.Mode
		order.Count=orderOrigin.Count
		order.CountMale=orderOrigin.CountMale
		order.Status= orderOrigin.Status


		orderModify,err:= client.OrderOrigin(prisma.OrderOriginWhereUniqueInput{ID:&in.OrderId}).OrderHotelModify(&prisma.OrderHotelModifyParamsExec{Last:&orderCount}).Exec(ctx)
		if err!=nil{
			log.Fatalf("failed to query :%v",err)
		}
		//TODO database.graphql need to modify the Int type to Int!
		orderm.Id=orderModify[0].ID
		orderm.Revision=orderModify[0].Revision
		orderm.TimeStamp=string(int(orderModify[0].Timestamp))
		orderm.Count=*orderModify[0].Count
		orderm.CountMale=*orderModify[0].CountMale
		orderm.Id=orderModify[0].ID
		

		return &qpb.QueryReply{OrderOrigin:&order,OrderHotelModify:[]*qpb.QueryReply_OrderHotelModify{&orderm}},nil
	}else if (in.Who.Adviser!=""){

	}


	return &qpb.QueryReply{},nil
}
