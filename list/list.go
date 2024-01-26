package list

import "math"

type Node struct {
	Val  int
	Next *Node
}

// 反转链表
// 1 2 3 4
//     newHead
//     <- 反转
//      \ 断开
func reverseList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverse2(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverse2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverse3(head *Node) *Node {
	var prev *Node
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}

// 1 2 3 4
// prev cur
// reverseListTwo 反转链表迭代法
func reverseListTwo(head *Node) *Node {
	var prev *Node
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}

// 1 2 3 4
// 2 1 4 3
func swapPairs(head *Node) *Node {
	dummy := &Node{Val: math.MaxInt32, Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		n1 := cur.Next
		n2 := cur.Next.Next
		cur.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		cur = n1
	}
	return dummy.Next
}

// getListLen 获取链表长度
func getListLen(node *Node) int {
	c := 0
	for node != nil {
		node = node.Next
		c++
	}
	return c
}

// deleteNode 删除倒数n个节点
// dummy 虚拟头节点，针对头节点的删除，避免特殊处理
func deleteNode(root *Node, n int) *Node {
	m := getListLen(root)
	dummy := &Node{Val: -1, Next: root}
	cur := dummy
	for i := 0; i < n-m; i++ {
		cur = cur.Next
	}
	cur.Next.Next = cur.Next
	return dummy.Next
}

// addNode 新增
func addNode(tail *Node, m int) *Node {
	tail.Next = &Node{Val: m, Next: nil}
	tail = tail.Next
	return tail
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

type INode struct {
	node *TreeNode
	val  int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*INode{{node: root, val: 1}}
	res := 0
	for len(queue) > 0 {
		res = max(res, (queue[len(queue)-1].val-queue[0].val)+1)
		var item []*INode
		for _, v := range queue {
			if v.node.Left != nil {
				item = append(item, &INode{node: v.node.Left, val: v.val * 2})
			}
			if v.node.Right != nil {
				item = append(item, &INode{node: v.node.Right, val: 2*v.val + 1})
			}
		}
		queue = item
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
