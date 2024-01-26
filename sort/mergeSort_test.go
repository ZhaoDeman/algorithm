package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	r := MergeSort([]int{12,2,3,4,11,34,56})
	fmt.Println(r)
}
