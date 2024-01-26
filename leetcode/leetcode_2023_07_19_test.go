package leetcode

import (
	"math"
	"sort"
	"testing"
)

func TestMinInterval(t *testing.T) {
	minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5})
}
func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := []int{}
	for _, v := range queries {
		temp := math.MaxInt32
		for _, vv := range intervals {
			if v >= vv[0] && v <= vv[1] {
				temp = min(temp, vv[1]-vv[0]+1)
			}
		}
		if temp == math.MaxInt32 {
			temp = -1
		}
		res = append(res, temp)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

