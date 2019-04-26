package cleanOrder

import (
	"context"
	"fmt"
	"time"

	"../prisma"
)

//清除一周后的订单，修改订单状态
func CleanOrder() {
	clien := prisma.New(nil)
	ctx := context.TODO()
	var cleanDate = int32(time.Now().Unix() - 7*24*3600)
	var date int32

	if (cleanDate+28800)%86400 == 0 {
		date = cleanDate
	} else {
		cleanDate -= (cleanDate + 28800) % 86400
		date = cleanDate
	}
	dateMin := date
	dateMax := date + 86399
	var targetStatus int32 = 3
	_, err := clien.UpdateManyOrderOrigins(prisma.OrderOriginUpdateManyParams{
		Where: &prisma.OrderOriginWhereInput{
			DatetimeGte: &dateMin,
			DatetimeLte: &dateMax,
		},
		Data: prisma.OrderOriginUpdateManyMutationInput{
			Status: &targetStatus,
		},
	}).Exec(ctx)
	if err != nil {
		fmt.Println("clean order error:", err)
	}
}
