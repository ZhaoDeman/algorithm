package tree

import (
	"container/heap"
	"fmt"
)

type HuffmanNode struct {
	Char  byte
	Freq  int
	Left  *HuffmanNode
	Right *HuffmanNode
}

type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Freq < pq[j].Freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*HuffmanNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func buildHuffmanTree(data map[byte]int) *HuffmanNode {
	pq := make(PriorityQueue, len(data))
	i := 0
	for char, freq := range data {
		pq[i] = &HuffmanNode{Char: char, Freq: freq}
		i++
	}
	heap.Init(&pq)

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*HuffmanNode)
		right := heap.Pop(&pq).(*HuffmanNode)

		internalNode := &HuffmanNode{
			Char:  0, // Internal nodes have no character
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}

		heap.Push(&pq, internalNode)
	}

	return pq[0]
}

func printHuffmanCodes(root *HuffmanNode, code string) {
	if root == nil {
		return
	}

	if root.Char != 0 {
		fmt.Printf("Character: %c, Code: %s\n", root.Char, code)
	}

	printHuffmanCodes(root.Left, code+"0")
	printHuffmanCodes(root.Right, code+"1")
}



// w = {1,3,10,20,5,29}
// v = {a,b,c,d,e,f}
// 1 3
// [4,10,20,5,29]
//      (a,b,4)
//  (a,1)   (b,3)
// [9,10,20,29]
//        (a,b,e,9)
//    (a,b,4)  (e,5)
//  (a,1)(b,3)
// [19,20,29]
//            (a,b,e,c,19)
//       (a,b,e,9)    (c,10)
//    (a,b,4)  (e,5)
// (a,1)  (b,3)
// [39,29]                  68
//             (a,b,e,c,d,39)    29
//         (a,b,e,c,19)   (d,20)
//      (a,b,e,9)  (c,10)
//   (a,b,4) (e,5)
//  (a,1) (b,3)

// 20*2 + 29 + 10 *3