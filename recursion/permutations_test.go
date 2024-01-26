package main

import (
	"fmt"
	"testing"
)

func TestPermutations(t *testing.T) {
	visit := make([]bool,3)
	dfs([]int{1,2,2},0,[]int{},visit)
	fmt.Println(result)
}
