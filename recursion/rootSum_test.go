package main

import (
	"fmt"
	"testing"
)

func TestRootSum(t *testing.T) {
	root := &TreeNode{1,&TreeNode{Val: 2},nil}
	fmt.Println(treeDfs(root,0))
}
