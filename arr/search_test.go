package arr

import (
	"container/heap"
	"fmt"
	"strings"
	"testing"
)

// 有序数组中找出target下标
func search(nums []int, target int) {
	arr := []int{2, 3, 4, 5}
	// target 4 || 9
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		}
	}
	fmt.Println(left)
}

func TestSearch(t *testing.T) {
	search(nil, 4)
	strings.Contains("aa","a")
}

func halveArray(nums []int) int {
	sum := float64(0)
	hp := &hp{}
	for _, v := range nums {
		sum += float64(v)
		heap.Push(hp, float64(v))
	}
	halveSum := sum / 2
	ans := float64(0)
	count := 0
	for ans < halveSum {
		ans += hp.pq[0] / 2
		hp.pq[0] /= 2
		heap.Fix(hp, 0)
		count++
	}
	return count
}

type hp struct {
	pq []float64
}

func (hp hp) Less(i, j int) bool {
	return hp.pq[i] > hp.pq[j]
}
func (hp hp) Swap(i, j int) {
	hp.pq[i], hp.pq[j] = hp.pq[j], hp.pq[i]
}
func (hp hp) Len() int {
	return len(hp.pq)
}
func (hp *hp) Push(x interface{}) {
	hp.pq = append(hp.pq, x.(float64))
}
func (hp *hp) Pop() interface{} {
	a := hp.pq
	x := a[len(a)-1]
	hp.pq = a[0 : len(a)-1]
	return x
}

