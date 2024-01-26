package sort

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestHeapify(t *testing.T) {
	//11 3 10 19 20 100 21 1
	arr := []int{11, 3, 10, 19, 20, 100, 21, 1}
	// build Max heap 过程
	// 开始index = 8/2 - 1 = 3  left 7 right 8
	// 11 3 10 19 20 100 21 1
	// index (2 5 6)
	// 11 3 100 19 20 10 21 1
	// index (1 3 4)
	// 11 20 100 19 3 10 21 1
	// 100 20 11 19 3 10 21 1
	// 100 20 21 19 3 10 11 1
	HeapSort(arr, maxHeap)
	fmt.Println(arr)
}

//[3,2,3,1,2,4,5,5,6]
//4
func TestTopK(t *testing.T) {
	fmt.Println(findKthLargest([]int{2, 1}, 2))
}

func findKthLargest(nums []int, k int) int {
	buildTree(nums)
	n := len(nums) - 1
	index := 0
	for n >= 1 {
		swap(nums, 0, n)
		index++
		if index == k {
			return nums[n]
		}
		n -= 1
		heapify(nums, n, 0)
	}
	return nums[len(nums)-k]
}

func buildTree(arr []int) {
	n := len(arr) - 1
	for i := n / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

// 小
func heapify(arr []int, n, i int) {
	if i > n {
		return
	}
	left := 2*i + 1
	right := 2*i + 2
	max := i
	if left <= n && arr[max] < arr[left] {
		max = left
	}
	if right <= n && arr[max] < arr[right] {
		max = right
	}
	if max != i {
		swap(arr, i, max)
		heapify(arr, n, max)
	}
}

func robotSim(commands []int, obstacles [][]int) int {
	obstaclesXMap := make(map[int][]int, 0)
	obstaclesYMap := make(map[int][]int, 0)
	for _, v := range obstacles {
		obstaclesXMap[v[0]] = append(obstaclesXMap[v[0]], v[1]) // x点对应的y点
		obstaclesYMap[v[1]] = append(obstaclesYMap[v[1]], v[1]) // y点对应的x 点
	}
	index := 1 // 1y轴正方向 2 x轴正方向 3 y 轴负方向 4x 负
	// 遇到-1 index+1 遇到 -2 index - 1 （index mod 5） 即可
	x, y := 0, 0
	res := 0
	for i := 0; i < len(commands); i++ {
		if commands[i] == -1 {
			index += 1
			continue
		}
		if commands[i] == -2 {
			index -= 1
			continue
		}
		if index%5 == 1 {
			//y 正方向 y+操作
			// 判断 x 相同 y ~ y+commands[i] 在不在 obstacles 之间
			yList := obstaclesXMap[x]
			sort.Ints(yList)
			exist := false
			for _, v := range yList {
				if v >= y && v <= y+commands[i] {
					y = v
					exist = true
					break
				}
			}
			if !exist {
				y += commands[i]
			}
		}
		if index%5 == 2 {
			// x 轴正方向
			xList := obstaclesYMap[y]
			sort.Ints(xList)
			exist := false
			for _, v := range xList {
				if v >= x && v <= x+commands[i] {
					exist = true
					x = v
					break
				}
			}
			if !exist {
				x += commands[i]
			}
		}
		if index%5 == 3 {
			// x 轴正方向
			yList := obstaclesXMap[x]
			sort.Ints(yList)
			exist := false
			for _, v := range yList {
				if v >= y && v <= y-commands[i] {
					y = v
					exist = true
					break
				}
			}
			if !exist {
				y -= commands[i]
			}
		}
		if index%5 == 2 {
			// x 轴负方向
			xList := obstaclesYMap[y]
			sort.Ints(xList)
			exist := false
			for _, v := range xList {
				if v >= x && v <= x-commands[i] {
					exist = true
					x = v
					break
				}
			}
			if !exist {
				x -= commands[i]
			}
		}
		res = max(res, square(x, y))
	}
	return res
}

// 返回当前位置平方数
func square(x, y int) int {
	x = int(math.Abs(float64(x - 0)))
	y = int(math.Abs(float64(y - 0)))
	return x*x + y*y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestSquare(t *testing.T) {
	fmt.Println(square(3, 4))

	arr := []int{1, 2, 3}
	for i := 0; i < 3; i++ {
		arr = append(arr[0:i], arr[i+1:]...)
	}
}
func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	left := 0
	right := len(intervals)-1
	count := 0
	for left < right && right >= 0 {
		rightVal := intervals[left][1]
		leftVal := intervals[left][0]
		for right >= 0 && (intervals[right][0] > rightVal) || (intervals[right][1] < leftVal){
			rightVal = intervals[right][1]
			leftVal = intervals[right][0]
			right--
			fmt.Println(right)
		}
		left++
		count++
	}
	return count
}