package sort

// MergeSort 归并排序 分为2步
// 1 分 2分就行
// 2 并 依赖于合并两个有序数组
func MergeSort(arr []int) (res []int) {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	return MergerArr(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}

func MergerArr(arr1, arr2 []int) (res []int) {
	res = []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	res = append(res, arr1[i:]...)
	res = append(res, arr2[j:]...)
	return res
}