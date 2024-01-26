package list

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	arr := []int{1, 3, 0, 9, 10, -1, 8, 99, 190, 78, 23}
	priority = arr
	q := &IHeap{make([]int, 10)}
	for i := 0; i < 10; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)
	fmt.Println("init:",q)
//	heap.Push(q,5)
	//heap.Push(q,6)
	//heap.Push(q,7)
	for q.Len() > 0 {
		fmt.Println("->",arr[q.IntSlice[0]])
		heap.Pop(q)
		fmt.Println(q.IntSlice)
	}
	//fmt.Println("优先级最高元素：",arr[q.IntSlice[0]]) // IntSlice 维护的大顶堆的id
	//fmt.Println("弹出之前优先级队列：",q.IntSlice)
	//fmt.Println("弹出元素：",arr[heap.Pop(q).(int)])
	//fmt.Println("弹出之后：",q.IntSlice)
	//fmt.Println("arr:",arr)
	//fmt.Println("priority:",priority)
}
var priority []int

type IHeap struct {
	sort.IntSlice
}

func (hp IHeap) Less(i, j int) bool {
	return priority[hp.IntSlice[i]] > priority[hp.IntSlice[j]]
}
func (hp *IHeap) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}
func (hp *IHeap) Pop() interface{} {
	p := hp.IntSlice
	v := p[len(p) - 1]
	hp.IntSlice = p[:len(p)-1]
	return v
}
