package main

import "fmt"

func main() {
	visit := make([]bool,3)
	dfs([]int{1,2,3},0,[]int{},visit)
	fmt.Println(result)
}

var result [][]int

// 1 2 3 全排列
// depth == length
// copy item
// visit 已经访问的标识
func dfs(nums []int, depth int, item []int, visit []bool) {
	if depth == len(nums) {
		dst := make([]int, len(item))
		copy(dst, item)
		result = append(result, item)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !visit[i] {
			visit[i] = true
			item = append(item, nums[i])
			fmt.Println("递归前--->",item)
			dfs(nums, depth+1, item, visit)
			visit[i] = false
			fmt.Println("递归后--->",item)
			item = item[:len(item)-1]
		}
	}
}
