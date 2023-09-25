package main

import (
	"container/heap"
	"fmt"
	"math"
)

// 그래프를 나타내는 구조체
type Graph struct {
	nodes map[string]bool
	edges map[string][]Edge
}

// 간선을 나타내는 구조체
type Edge struct {
	to     string
	weight int
}

// 그래프 초기화
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]bool),
		edges: make(map[string][]Edge),
	}
}

// 노드 추가
func (g *Graph) AddNode(node string) {
	g.nodes[node] = true
	g.edges[node] = []Edge{}
}

// 간선 추가 (양방향 그래프)
func (g *Graph) AddEdge(from, to string, weight int) {
	g.edges[from] = append(g.edges[from], Edge{to, weight})
	g.edges[to] = append(g.edges[to], Edge{from, weight})
}

// 다익스트라 알고리즘
func (g *Graph) Dijkstra(start string) map[string]int {
	distances := make(map[string]int)
	for node := range g.nodes {
		distances[node] = math.MaxInt
	}
	distances[start] = 0

	priorityQueue := make(PriorityQueue, 0)
	heap.Init(&priorityQueue)
	heap.Push(&priorityQueue, Node{start, 0})

	for len(priorityQueue) > 0 {
		current := heap.Pop(&priorityQueue).(Node)

		if current.Weight > distances[current.Name] {
			continue
		}

		for _, edge := range g.edges[current.Name] {
			distance := current.Weight + edge.weight

			if distance < distances[edge.to] {
				distances[edge.to] = distance
				heap.Push(&priorityQueue, Node{edge.to, distance})
			}
		}
	}

	return distances
}

// 우선순위 큐를 위한 Node 구조체
type Node struct {
	Name   string
	Weight int
}

// 우선순위 큐
type PriorityQueue []Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Weight < pq[j].Weight }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Node)
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
	// 그래프 생성
	myGraph := NewGraph()
	myGraph.AddNode("A")
	myGraph.AddNode("B")
	myGraph.AddNode("C")
	myGraph.AddNode("D")
	myGraph.AddNode("E")

	myGraph.AddEdge("A", "B", 5)
	myGraph.AddEdge("B", "C", 3)
	myGraph.AddEdge("C", "D", 1)
	myGraph.AddEdge("D", "E", 2)
	myGraph.AddEdge("A", "E", 8)

	// 최단 경로 계산
	startNode := "A"
	shortestDistances := myGraph.Dijkstra(startNode)

	// 결과 출력
	for node, distance := range shortestDistances {
		fmt.Printf("Shortest distance from %s to %s: %d\n", startNode, node, distance)
	}
}
