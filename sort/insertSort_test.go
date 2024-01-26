package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{1,4,2,5,100,56,23}
	InsertSort(arr)
	fmt.Println(arr)
}
