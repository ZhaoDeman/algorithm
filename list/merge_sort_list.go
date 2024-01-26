package list

func sortList(head *Node) *Node {
	return merge_sort(head, nil)
}

func merge(h1 *Node, h2 *Node) *Node {
	dummy := &Node{}
	cur := dummy
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			cur.Next = h1
			h1 = h1.Next
		} else {
			cur.Next = h2
			h2 = h2.Next
		}
		cur = cur.Next
	}
	for h1 != nil {
		cur.Next = h1
		h1 = h1.Next
		cur = cur.Next
	}
	for h2 != nil {
		cur.Next = h2
		h2 = h2.Next
		cur = cur.Next
	}
	return dummy.Next
}

func merge_sort(head, tail *Node) *Node {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow := head
	fast := head
	// 找出mid
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(merge_sort(head, mid), merge_sort(mid, tail))
}
