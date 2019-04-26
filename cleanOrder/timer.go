package cleanOrder

import "time"

// 定时调用清理订单函数
// 每次执行这个函数，立即执行清理函数，然后每天晚上00:00:00执行一次
func StartTimer() {
	go func() {
		for {
			CleanOrder()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

