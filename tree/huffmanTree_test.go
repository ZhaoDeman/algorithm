package tree

import (
	"fmt"
	"testing"
)

func TestHuf(t *testing.T) {
	data := map[byte]int{
		'A': 3,
		'B': 2,
		'C': 6,
		'D': 8,
	}

	root := buildHuffmanTree(data)

	fmt.Println("Huffman Codes:")
	printHuffmanCodes(root, "")
}
//        O
//  0(D)    0
//        0   0(C)
//      0(B) 0(A)