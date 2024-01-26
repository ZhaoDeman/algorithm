package _0230707

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3,5,2,3}
	sort.Ints(arr)
	maxVal := arr[len(arr)-2]+arr[len(arr)-1]
	fmt.Println(maxVal)
}
