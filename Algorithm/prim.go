package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node   int
	weight int
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].weight < pq[j].weight }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Edge)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// 프림(Prim) 알고리즘은 주어진 연결 그래프에서 최소 신장 트리를 찾는 알고리즘입니다.
// 최소 신장 트리는 그래프의 모든 정점을 연결하는 간선의 가중치 합이 최소인 트리를 의미합니다.
func prim(graph [][]int) int {
	numVertices := len(graph)
	visited := make([]bool, numVertices)
	pq := make(PriorityQueue, 0)
	result := 0

	heap.Push(&pq, Edge{0, 0})

	for len(pq) > 0 {
		e := heap.Pop(&pq).(Edge)
		node := e.node
		weight := e.weight

		if visited[node] {
			continue
		}

		visited[node] = true
		result += weight

		for adjacent, adjWeight := range graph[node] {
			if !visited[adjacent] && adjWeight > 0 {
				heap.Push(&pq, Edge{adjacent, adjWeight})
			}
		}
	}

	return result
}

func main() {
	// 그래프 예시 (2차원 배열로 표현)
	graph := [][]int{
		{0, 29, 0, 0, 10},
		{29, 0, 16, 0, 0},
		{0, 16, 0, 12, 0},
		{0, 0, 12, 0, 18},
		{10, 0, 0, 18, 0},
	}

	// 최소 신장 트리의 가중치 합
	mstWeight := prim(graph)
	fmt.Println(mstWeight) // 출력: 65
}
