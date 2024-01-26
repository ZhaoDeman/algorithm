package graph

import (
	"fmt"
	"math"
)

// 图的表示：数组表示

// 图的表示：领接矩阵
// 遍历整个图
func ergodic(edge [][]int, n int) {
	graph := make([][]int, n)
	// 初始化邻接矩阵
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, v := range edge {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	fmt.Println("g:", graph)
	visit := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visit[i] {
			dfs(graph, i, visit)
			item = []int{}
		}
	}
}

var item []int

func dfs(graph [][]int, v int, visit []bool) {
	if !visit[v] {
		item = append(item, v)
	}
	visit[v] = true
	list := graph[v]
	for i := 0; i < len(list); i++ {
		if !visit[list[i]] {
			fmt.Println("节点：", list[i])
			visit[list[i]] = true
			item = append(item, list[i])
			fmt.Println("递归前->", item)
			dfs(graph, list[i], visit)
			fmt.Println("递归后->", item)
		}
	}
}

// BFS 广度优先遍历
// 队列不断扩展下一个可以访问的点
//          2     4
//          ^     ^
//          |     |
// [0] 节点  0 - > 1 -> 3
// [0,2] [0,1] [0,1,3]
//
// 1,2
// 3,4
func BFS(graph [][]int) [][]int {
	paths := make([][]int, 0)
	if len(graph) == 0 {
		return paths
	}
	queue := make([]interface{}, 0)
	path := []int{0}
	queue = append(queue, path)
	for len(queue) > 0 {
		currentPath := queue[len(queue)-1].([]int)
		node := currentPath[len(currentPath)-1]
		queue = queue[:len(queue)-1]
		for _, v := range graph[node] {
			tempPath := make([]int, len(currentPath))
			copy(tempPath, currentPath)
			tempPath = append(tempPath, v)
			fmt.Println("tempPath:", tempPath)
			if v == len(graph)-1 {
				paths = append(paths, tempPath)
			} else {
				queue = append(queue, tempPath)
			}
		}
	}
	return paths
}

// DFS 深度优先遍历
func DFS(graph [][]int) [][]int {
	var Dfs func(level int, temp []int)
	var paths [][]int
	Dfs = func(node int, temp []int) {
		if len(graph[node]) == 0 { // 递归到最后一个不能再向下找的路径
			newArr := make([]int, len(temp))
			copy(newArr, temp)
			paths = append(paths, newArr)
			return
		}
		for _, v := range graph[node] {
			temp = append(temp, v)
			Dfs(v, temp)
			temp = temp[:len(temp)-1]
		}
	}
	Dfs(0, []int{0})
	return paths
}

// 单源点最短路径:迪杰斯特拉（dijkstra）
// 邻接矩阵 graph graph[i][j] 表示点i到j存在路径 graph[i][j]的值为边的权重 math.MaxInt64 表示i不能到j
// dist数组 dist[i] 表示k(源)到i 点的最短路径。
// visit数组 已经访问的顶点集
// 1.初始化dist 数组，dist[k] = 0 visit[k] = true
func dijkstra(edge [][]int, k, n int) {
	// k点到余下点的最短路径，edge []int{u,v,w} // u 起始点 v终止点 w权重
	graph := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			graph[i][j] = math.MaxInt64
		}
	}
	for _,v := range edge {
		graph[v[0]][v[1]] = v[2]
	}
	for _,v := range graph {
		fmt.Println("graph:",v)
	}
}

// 无向图 有向图 带权图
// 图的概念 联通性 图上的两个点是否是联通的
// 联通性判断，采用并查集
// 1. quickFind
// 2. quickUnion
// 3. quickUnion 优化，使得root 数组查找时更快
// 4. 按秩合并
// 5. 路径压缩，每次find 函数查找时，会从当前节点一直往上查找，直到找到根节点，由于这个过程会是重复的，我们可以把查找的过程，更新父节点。

type UnionFind struct {
	root []int // 父节点
	rank []int // 当前为父节点的高度，优化union，使find 更快
}

func (u *UnionFind) UnionFind(n int) *UnionFind {
	ans := &UnionFind{
		root: make([]int, n),
	}
	rank := make([]int, n)
	for k := range ans.root {
		ans.root[k] = k
		rank[k] = 1
	}
	return ans
}

func (u *UnionFind) QuickFind(x int) int {
	return u.root[x]
}

// Union x y 建立关系
func (u *UnionFind) Union(x, y int) {
	xRoot := u.QuickFind(x)
	yRoot := u.QuickFind(y)
	if xRoot != yRoot {
		for i := 0; i < len(u.root); i++ {
			if u.root[i] == yRoot {
				u.root[i] = xRoot
			}
		}
	}
}

// Connected 判断x y 是否联通
func (u *UnionFind) Connected(x, y int) bool {
	return u.QuickFind(x) == u.QuickFind(y)
}

func (u *UnionFind) Find(x int) int {
	// 不断地找根节点
	if u.root[x] == x {
		return x
	}
	u.root[x] = u.Find(u.root[x]) // 递归直到根节点，直到找到根，回溯时，全部赋值根节点
	return u.root[x]
}

// QuickUnion 时间复杂度，取决Find Find 函数最差O(n)
func (u *UnionFind) QuickUnion(x, y int) {
	xRoot := u.Find(x)
	yRoot := u.Find(y)
	if xRoot != yRoot {
		if u.rank[xRoot] < u.rank[yRoot] { // yRoot 高于 xRoot xRoot连接到 yRoot上
			u.root[xRoot] = u.root[yRoot]
		} else if u.rank[xRoot] > u.rank[yRoot] { // xRoot 高于 yRoot，yRoot 连接到xRoot 上
			u.root[yRoot] = u.root[xRoot]
		} else { // 相等时，xRoot 连接到yRoot ，更新yRoot的rank
			u.root[xRoot] = u.root[yRoot]
			u.rank[yRoot] += 1
		}
	}
}
