package arr

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"testing"
)

//  原数组d
//  差分数组 f  f[0] = d[0]
//  性质1：
//  f[i] = d[i] - d[i-1] (i > 0)
//  d[i] = f[1]+..f[i] 其中 (i > 0)

func nthUglyNumber(n int) int {
	hash := map[int]bool{}
	hash[1] = true
	temp := []int{2, 3, 5}
	hp := &hp{}
	heap.Push(hp, 1)
	for i := 1; i <= 1690; i++ {
		x := hp.Pop().(int)
		if i == n {
			return x
		}
		for _, vv := range temp {
			if !hash[vv*x] {
				heap.Push(hp, vv*x)
				hash[vv*x] = true
			}
		}
	}
	return 0
}

func TestDeleteGreatestValue(t *testing.T) {
	fmt.Println(deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}}))
}

func deleteGreatestValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for k, v := range grid {
		sort.Ints(v)
		grid[k] = v
	}
	res := 0
	for j := n - 1; j >= 0; j-- {
		ans := math.MinInt32
		for i := 0; i < m; i++ {
			ans = max(ans, grid[i][j])
		}
		res += ans
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestMaxDistToClosest(t *testing.T) {
	fmt.Println(maxDistToClosest([]int{1,0,0,0}))
}

func maxDistToClosest(seats []int) int {
	if len(seats) == 0 {
		return 0
	}
	startIndex := 0
	ans := 0
	if seats[0] == 1 {
		start := startIndex
		startIndex += 1
		for startIndex < len(seats) {
			if seats[startIndex] == 1 {
				ans = max((startIndex-start)/2, ans)
				start = startIndex
			}
			startIndex++
		}
		if seats[len(seats)-1] != 1 {
			ans = max(startIndex-start-1, ans)
		}
	} else {
		end := startIndex
		startIndex += 1
		for startIndex < len(seats) {
			if seats[startIndex] == 1 {
				if end == 0 {
					ans = max(ans,startIndex-end)
				}else {
					ans = max(ans,(startIndex-end)/2)
				}
				end = startIndex
			}
			startIndex++
		}
		if seats[len(seats)-1] != seats[end] {
			ans = max(ans,startIndex-end-1)
		}
	}
	return ans
}
