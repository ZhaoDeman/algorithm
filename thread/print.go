package thread

import (
	"fmt"
	"sync"
	"time"
)

func Print() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go func() {
		for i := 0; i < 50; i++ {
			<-ch1
			fmt.Println(2*i + 1)
			ch2 <- true
		}
	}()
	go func() {
		for i := 0; i < 50; i++ {
			<-ch2
			fmt.Println(2*i + 2)
			ch1 <- true
		}
	}()
	ch1 <- true
	time.Sleep(1 * time.Second)
}



func Work() {

	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go Produce(&wg, ch)
	wg.Add(1)
	go Consume(&wg, ch)
	wg.Wait()
	fmt.Println("end")
}

func Produce(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	defer close(ch)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
}

func Consume(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

func PrintCom() {
	// 利用两个chan
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch2)
		for i := 0; i < 50; i++ {
			<-ch1
			fmt.Println(2*i + 1)
			ch2 <- 1
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch1)
		for i := 0; i < 50; i++ {
			<-ch2
			fmt.Println(2*i + 2)
			ch1 <- 1
		}
	}()
	ch1 <- 1
	wg.Wait()
}

func PrintCom2() {
	ch := make(chan int) // 读写对称
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- 1
			if i % 2 == 1 {
				fmt.Println(i)
			}

		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			<-ch
			if i % 2 == 0 {
				fmt.Println(i)
			}
		}
	}()
	wg.Wait()
}

//用一个chan 怎么实现
// g1 ---> 写ch 打印
// g2 ---> 读ch 打印
// for i := 0;i<
