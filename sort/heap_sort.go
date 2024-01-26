package sort

import "fmt"

const maxHeap = "max"
const minHeap = "min"

func MaxHeapify(arr []int, heapifySize int, i int) {
	if i >= heapifySize {
		return
	}
	left := 2*i + 1
	right := 2*i + 2
	max := i
	if left <= heapifySize && arr[left] > arr[max] {
		max = left
	}
	if right <= heapifySize && arr[right] > arr[max] {
		max = right
	}
	if max != i {
		swap(arr, i, max)
		MaxHeapify(arr, heapifySize, max)
	}
}

func MinHeapify(arr []int, heapifySize int, i int) {
	if i > heapifySize {
		return
	}
	min := i
	left := 2*i + 1
	right := 2*i + 2
	if left <= heapifySize && arr[left] < arr[min] {
		min = left
	}
	if right <= heapifySize && arr[right] < arr[min] {
		min = right
	}
	if min != i {
		swap(arr, min, i)
		MinHeapify(arr, heapifySize, min)
	}
}

func BuildMinTree(arr []int) {
	heapify := len(arr) - 1
	for i := len(arr) / 2; i >= 0; i-- {
		MinHeapify(arr, heapify, i)
	}
}

// HeapSort 堆排序
// 构造大顶堆 大顶堆 每次移动最大值到最后，最终排序 从小到大
// 构造小顶堆 小顶堆 每次移动最小值到最后 最终排序 小大到小
func HeapSort(arr []int, str string) {
	if str == maxHeap {
		HeapMaxSort(arr)
	} else if str == minHeap {
		MinHeapSort(arr)
	}
}

func MinHeapSort(arr []int) {
	BuildMinTree(arr)
	n := len(arr) - 1
	for n >= 1 {
		swap(arr, 0, n)
		n -= 1
		MinHeapify(arr, n, 0)
	}
}

func BuildTree(arr []int) {
	heapifySize := len(arr) - 1
	for i := len(arr) / 2; i >= 0; i-- {
		MaxHeapify(arr, heapifySize, i)
	}
}

func HeapMaxSort(arr []int) {
	BuildTree(arr)
	fmt.Println("maxHeap:",arr)
	n := len(arr) - 1
	for n > 0 {
		swap(arr, n, 0)
		n -= 1
		MaxHeapify(arr, n, 0)
	}
}

func swap(a []int, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

