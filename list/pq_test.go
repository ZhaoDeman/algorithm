package list

import (
	"container/heap"
	"container/list"
	"fmt"
	"testing"
)

// TestPq 测试优先级队列
func TestPq(t *testing.T) {
	p := []int{1,3,4,0,90,30,21,10}
	pq := make(PriorityQueue,0)
	for _,v := range p {
		heap.Push(&pq,v)
	}
	heap.Init(&pq)
	fmt.Println(pq)
	for pq.Len() > 0 {
		//fmt.Println("pop:",pq[0])
		fmt.Println("Heap Pop:",heap.Pop(&pq))
	}
}

// Queue 测试
// 队列的作用：
// 1.用来模拟
// 2.图的遍历BFS
// 3.cpu调度

func TestQueue(t *testing.T) {
	l := list.New()
	l.PushBack([]int{1})
	l.PushBack([]int{2,3,4})
	l.PushBack([]int{4,5,6})
	pop := l.Back()
	for pop != nil {
		fmt.Println(pop.Value)
		pop = pop.Prev()
	}
}
