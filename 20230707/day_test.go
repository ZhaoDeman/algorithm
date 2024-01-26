package _0230707

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"sync"
	"testing"
)

func Test01(t *testing.T) {
	arr := []int{3, 5, 4, 6, 4, 2}
	sort.Ints(arr)
	left := 0
	right := len(arr) - 1
	maxVal := math.MinInt32
	for left < len(arr)/2 {
		maxVal = max(maxVal, arr[left]+arr[right])
		left++
		right--
	}
	fmt.Println(maxVal)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Person struct {
	Age  int
	Name string
}

func TestPersonList(t *testing.T) {
	// 值传递和引用传递
	// 基本数据类型
}

func t3(person Person) {
	person.Age = 19
}
func t4(person *Person) {
	person.Age = 19
}

func t1(p []Person) {
	for _,v := range p {
		v.Age = 19
	}
}

func t2(p []*Person) {
	for _,v := range p {
		v.Age = 19
	}
}


/*type MyMutex struct{
	count int
	mu *sync.Mutex
}

func TestMutex(t *testing.T) {
	var mu *MyMutex
	mu = &MyMutex{count: 1,mu: &sync.Mutex{}}
	mu.mu.Lock()
	var mu1= mu
	mu.count++
	mu.mu.Unlock()
	mu1.mu.Lock()
	mu1.count++
	mu1.mu.Unlock()
	fmt.Println(mu.count,mu1.count)
}*/

type MyMutex struct{
	count int
	mu sync.Mutex
}

func TestMutex(t *testing.T) {
	var mu MyMutex
	mu = MyMutex{count: 1,mu: sync.Mutex{}}
	mu.mu.Lock()
	var mu1=mu // 获得了一个锁的副本
	mu.count++
	mu.mu.Unlock() // mu1 已经加了一个锁
	mu1.mu.Unlock() // 释放
	mu1.mu.Lock() // 加锁
	mu1.count++
	mu1.mu.Unlock()

	fmt.Println(mu.count,mu1.count)
}


func TestRe(t *testing.T) {
	matched, err := regexp.MatchString("^[0-9a-zA-Z\\-_\\.]{3,64}$", "com.moguo.redPacketldiom")
	if !matched{
		fmt.Println(err.Error())
	}
	fmt.Println(matched)
}