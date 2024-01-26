package tree

import (
	"fmt"
	"testing"
)

// Node 二叉树
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func TestNode(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(goodNodes(root))
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

func TestBuildTree(t *testing.T) {
	/*	preOrder := []*TreeNode{
		{Left: nil,Right: nil,Val: 10},
		{Left: nil,Right: nil,Val: 11},
		{Left: nil,Right: nil,Val: 13},
		{Left: nil,Right: nil,Val: 14},
		{Left: nil,Right: nil,Val: 12},
	}*/
	inOrder := []*TreeNode{
		{Left: nil, Right: nil, Val: 13},
		{Left: nil, Right: nil, Val: 11},
		{Left: nil, Right: nil, Val: 14},
		{Left: nil, Right: nil, Val: 10},
		{Left: nil, Right: nil, Val: 12},
	}
	postOrder := []*TreeNode{
		{Left: nil, Right: nil, Val: 13},
		{Left: nil, Right: nil, Val: 14},
		{Left: nil, Right: nil, Val: 11},
		{Left: nil, Right: nil, Val: 12},
		{Left: nil, Right: nil, Val: 10},
	}

	/*	root := BuildTree(preOrder,inOrder)
		PrintTreeNode(PreOrder(root))*/

	root := BuildTreeTwo(inOrder, postOrder)
	PrintTreeNode(PreOrder(root))
}
