package graph

import (
	"fmt"
	"testing"
)

func TestErgodic(t *testing.T) {
	ergodic([][]int{{0, 1}, {1, 2}, {3, 4}}, 5)
}

func TestBFS(t *testing.T) {
	fmt.Println(BFS([][]int{{1, 2}, {3, 4}, {}, {}, {}}))
	fmt.Println(DFS([][]int{{1, 2}, {3, 4}, {}, {}, {}}))
}

func TestDijkstra(t *testing.T) {
	dijkstra([][]int{{1,2,3},{2,3,4}},1,4)
}

// 创建图，遍历图
// 0 -> 1
// 1 -> 2
// 3 -> 4
// 0 ----> 1 ------> 2
// 3 ----> 4
