package limiter

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func Limit() {
	//1.初始化 limiter 每秒10个令牌，令牌桶容量为20
	//第一个参数r Limit代表每秒可以向桶中产生多少令牌,Limit 实际上是 float64 的别名
	//第二个参数b int代表桶的容量大小,当前为20,如果桶中有20个令牌,
	//可以立即获取20个令牌,不需要等待直接执行,也就是最大并发数
	limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 20)

	//2.获取1个令牌,获取到返回true,否则false
	bo := limiter.Allow()
	if bo {
		fmt.Println("获取令牌成功")
	}

	//2.获取指定时间内指定个数令牌,获取到返回true,
	//实际上方的Allow()内部调用的就是AllowN()
	limiter.AllowN(time.Now(), 2)

	//3.阻塞直到获取足够的令牌或者上下文取消
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	limiter.Wait(ctx)
	err := limiter.WaitN(ctx, 20)
	if err != nil {
		fmt.Println("error", err)
	}

	//4.可以理解为预定一个令牌,当调用Reserve后不管是否存在有效令牌都会返回一个Reservation指针对象
	//接下来可以通过返回的Reservation进行指定操作
	reservation := limiter.Reserve()
	if 0 == reservation.Delay() {
		fmt.Println("获取令牌成功")
	}

	//5.指定实际内预定指定个数令牌
	//上面Reserve()内部实际就是调用的ReserveN()
	limiter.ReserveN(time.Now(), 1)

	//6.修改令牌生成速率
	limiter.SetLimit(rate.Every(time.Millisecond * 100))
	limiter.SetLimitAt(time.Now(), rate.Every(time.Millisecond*100))

	//7.修改令牌桶大小,也就是生成令牌的最大数量限制
	limiter.SetBurst(50)
	limiter.SetBurstAt(time.Now(), 50)

	//8.获取限流的速率即结构体中limit的值,每秒允许处理的事件数量,即每秒处理事件的频率
	l := limiter.Limit()
	fmt.Printf("每秒允许处理的事件数量,即每秒处理事件的频率为: %v", l)
	//9.返回桶的最大容量
	limiter.Burst()

}
