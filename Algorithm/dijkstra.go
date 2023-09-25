package main

import (
	"container/heap"
	"fmt"
)

// 그래프를 인접 리스트로 표현
var graph = map[string]map[string]int{
	"A": {"B": 8, "C": 1, "D": 2},
	"B": {},
	"C": {"B": 5, "D": 2},
	"D": {"E": 3, "F": 5},
	"E": {"F": 1},
	"F": {"A": 5},
}

// 다익스트라 알고리즘은 가중치가 있는 그래프에서 주어진 시작 정점에서
// 다른 모든 정점까지의 최단 거리를 찾는 알고리즘입니다.
func dijkstra(graph map[string]map[string]int, start string) map[string]int {
	distances := make(map[string]int)
	for node := range graph {
		distances[node] = int(^uint(0) >> 1) // 무한대 값으로 초기화
	}
	distances[start] = 0

	// 우선순위 큐를 사용하여 최단 거리를 관리
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: start, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		currentNode := current.value
		currentDistance := current.priority

		if currentDistance > distances[currentNode] {
			continue
		}

		for neighbor, weight := range graph[currentNode] {
			distance := currentDistance + weight
			if distance < distances[neighbor] {
				distances[neighbor] = distance
				heap.Push(pq, &Item{value: neighbor, priority: distance})
			}
		}
	}

	return distances
}

// 우선순위 큐 구조체 정의
type Item struct {
	value    string
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	start := "A"
	distances := dijkstra(graph, start)
	fmt.Printf("최단 거리 결과: %v\n", distances)
}
