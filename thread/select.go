package thread

import (
	"fmt"
	"sync"
)

//NoChan 无缓冲和有缓冲的chan
func NoChan() {
	ch := make(chan bool)
	go func() {
		t := <-ch
		fmt.Println("recv:", t)
	}()
	ch <- true
}

// PrintSwap 交替打印
func PrintSwap() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			<-ch1
			fmt.Println("线程1：", i*2+1)
			ch2 <- 1
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			<-ch2
			fmt.Println("线程2：", i*2+2)
			ch1 <- 1
		}
	}()
	ch1 <- 1
}

// WaitGroup 处理协程组
func WaitGroup() {
	chan1 := make(chan int, 1)
	chan2 := make(chan int, 2)
	chan3 := make(chan int, 3)
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		fmt.Println("协程开始处理...")
		chan1 <- 1 // 写入1
	}()

	go func() {
		defer wg.Done()
		recv1 := <-chan1
		fmt.Println("接受到go1：", recv1)
		chan2 <- 2 // 写入2
	}()

	go func() {
		defer wg.Done()
		recv2 := <-chan2
		fmt.Println("接受到go2：", recv2)
		chan3 <- 3
	}()

	go func() {
		defer wg.Done()
		recv := <-chan3
		fmt.Println("接受go3:", recv)
	}()
	wg.Wait()

}

func SelectWaitGroup() {
	ch := make(chan int, 1)
	res := 0

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			select {
			case value := <-ch:
				res += value
			default:
			}
		}
		/*	for v := range ch {
			res += v
		}*/
	}()
	wg.Wait()
	fmt.Println(res)
	return
}
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	start := nums[0]
	end := start
	ans := []string{}
	for i := 1; i < len(nums); i++ {
		if nums[i]-end > 1 || i == len(nums)-1 {
			if start == end {
				ans = append(ans, fmt.Sprintf("%d", start))
			} else {
				ans = append(ans, fmt.Sprintf("%d-->%d", start, end))
			}
			start = nums[i]
			end = start
		}
	}
	return ans
}
