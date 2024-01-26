package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*func main() {
	root := &TreeNode{1,&TreeNode{Val: 2},nil}
	fmt.Println(treeDfs(root,0))
}*/

// 二叉树根节点之和
func treeDfs(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	fmt.Println(node.Val)
	sum = 10*sum + node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}
	return treeDfs(node.Left, sum) + treeDfs(node.Right, sum)
}
