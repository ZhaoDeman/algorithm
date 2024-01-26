package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	temp := []int{11,23,10,9,100,90,101,21}
	// 快速排序的过程 (0,7)
	// 以11 为基准 第一次分区 使得左边比11 小，右边比11 大 23 比11 9比11 小 交换11  9  10  23 100 90 101 21
	// 10 位置已经没有可以交换了，将10 与11 比较 交换10与11 的位置
	// 10 9 11 23 100 90 101 21    -- 第一次分区结束 11的左边都比他小右边都比他大，将11 的左右两边分别排序
	// [10 9] 11 [23 100 90 101 21]
	// 对10 和9 分区 选择10 为基准 9小于10 交换
	// 9 10 11 [23 100 90 101 21] 第二次分区结束，将右边排序 23 21 90 101 100
	// 以 23 为基准，100 与21 交换 23 21 90 101 100
	// 9 10 11 21 23 90 101 100]
	quickSort(temp,0,7)
	fmt.Println(temp)
}