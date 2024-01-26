package sort

import "fmt"

// quickSort 快速排序
// 选择一个基数，找到所有的数左边都比这个数小，右边都比这个数大
// 递归左右两个数组，重复上述步骤
func quickSort(nums []int, start, end int) {
	if start >= end {
		//递归出口
		return
	}
	mid := partition2(nums, start, end)
	quickSort(nums, start, mid-1)
	quickSort(nums, mid+1, end)
}

// partition 普通分区
func partition(nums []int, start, end int) int {
	base := nums[start]
	left := start + 1
	right := end
	for left < right {
		for nums[left] <= base && left < right {
			left++
		}
		if left != right {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			right--
		}
	}
	if left == right && nums[right] > base {
		right--
	}
	if start != right {
		temp := nums[start]
		nums[start] = nums[right]
		nums[right] = temp
	}
	return right
}

//双指针分区
func partition2(nums []int, start, end int) int {
	base := nums[start]
	left := start + 1
	right := end
	for left < right {
		for nums[left] <= base && left < right {
			left++
		}
		for nums[right] > base && left < right {
			right--
		}
		if left != right {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left++
			right--
		}
	}
	if left == right && nums[right] > base {
		right--
	}
	temp := nums[start]
	nums[start] = nums[right]
	nums[right] = temp
	fmt.Println("p:",nums)
	return right
}
