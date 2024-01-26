package sort

// InsertSort 插入排序
func InsertSort(arr []int) {
	n := len(arr)
	i := 1
	for i < n {
		temp := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > temp {
				arr[j+1] = arr[j]
			} else {
				arr[j+1] = temp
				break
			}
		}
		i++
	}
}
