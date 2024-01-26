package tree

import (
	"fmt"
	"math"
)

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树遍历

func preOrder(root *Tree) {
	if root == nil {
		return
	}
	fmt.Println("先序遍历：", root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

func inOrder(root *Tree) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Println("中序遍历：", root.Val)
	inOrder(root.Right)
}

func postOrder(root *Tree) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Println("后续遍历：", root.Val)
}

func bfs(root *Tree) {
	queue := make([]*Tree, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		temp := make([]*Tree, 0)
		for _, v := range queue {
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
			fmt.Println("层次遍历：", v.Val)
		}
		queue = temp
	}
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(root *TreeNode, max int)
	dfs = func(root *TreeNode, max int) {
		if root == nil {
			return
		}
		if root.Val >= max {
			ans += 1
			max = root.Val
		}
		dfs(root.Left, max)
		dfs(root.Right, max)
	}
	dfs(root, math.MinInt64)
	return ans
}

// 重建二叉树
// 先序、中序求后序
//    A(10)
//  B(11)  C(12)
// D(13) E(14)
// 先序 ABDEC
// 中序 DBEAC
// 后序 DEBCA
// A 为根 (DBE) 左子树 (C) 右子树
// BDE DBE   B为根 DE  左子树(D) (E)右子树
// go 数组切片左闭右开

func BuildTree(preOrder, inOrder []*TreeNode) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}
	root := preOrder[0]
	i := 0
	for i = 0; i < len(inOrder); i++ {
		if root.Val == inOrder[i].Val {
			break
		}
	}
	// 将preOrder分为两个部分
	// inOrder分为两个部分
	return &TreeNode{
		Left:  BuildTree(preOrder[1:1+i], inOrder[0:i]),
		Right: BuildTree(inOrder[i+1:], preOrder[i+1:]),
		Val:   root.Val,
	}
}

func BuildTreeTwo(inOrder, postOrder []*TreeNode) *TreeNode {
	if len(postOrder) <= 0 || len(inOrder) <= 0 {
		return nil
	}
	root := postOrder[len(postOrder)-1]
	i := 0
	for i = 0; i < len(inOrder); i++ {
		if root.Val == inOrder[i].Val {
			break
		}
	}
	return &TreeNode{
		Left:  BuildTreeTwo(inOrder[0:i], postOrder[0:i]),
		Right: BuildTreeTwo(inOrder[i+1:], postOrder[i:len(postOrder)-1]),
		Val:   root.Val,
	}
}

func PostOrder(tree *TreeNode) (ans []*TreeNode) {
	if tree == nil {
		return []*TreeNode{}
	}
	ans = []*TreeNode{}
	ans = append(ans, PostOrder(tree.Left)...)
	ans = append(ans, PostOrder(tree.Right)...)
	ans = append(ans, tree)
	return
}

func PreOrder(tree *TreeNode) (ans []*TreeNode) {
	if tree == nil {
		return
	}
	ans = append(ans, tree)
	ans = append(ans, PreOrder(tree.Left)...)
	ans = append(ans, PreOrder(tree.Right)...)
	return
}

func PrintTreeNode(ans []*TreeNode) {
	for _, v := range ans {
		fmt.Print(v.Val, " ")
	}
	fmt.Println()
}

func LevelOrder(root *TreeNode) (ans [][]*TreeNode) {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		ans = append(ans,queue)
		var temp []*TreeNode
		for _,v := range queue {
			if v.Left != nil {
				temp = append(temp,v.Left)
			}
			if v.Right != nil {
				temp = append(temp,v.Right)
			}
		}
		queue = temp
	}
	for _,v := range ans {
		for _,vv := range v {
			fmt.Print(vv," ")
		}
		fmt.Println()
	}
	return ans
}
