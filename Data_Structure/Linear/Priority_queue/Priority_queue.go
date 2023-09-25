package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// 빈 우선순위 큐 생성
	myPriorityQueue := &IntHeap{}

	// 우선순위 큐에 원소 추가
	heap.Push(myPriorityQueue, 5)
	heap.Push(myPriorityQueue, 2)
	heap.Push(myPriorityQueue, 8)

	// 우선순위 큐에서 최솟값 추출
	_ = heap.Pop(myPriorityQueue).(int)

	// 우선순위 큐에서 최솟값 조회 (제거하지 않고)
	_ = (*myPriorityQueue)[0]

	// 원소를 추가하지 않고 최솟값 조회
	_ = (*myPriorityQueue)[0]

	// 우선순위 큐 순회(iteration)
	for myPriorityQueue.Len() > 0 {
		element := heap.Pop(myPriorityQueue).(int)
		fmt.Println(element)
	}
}

// IntHeap은 정수를 저장하는 최소 힙을 나타냅니다.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push는 힙에 요소를 추가합니다.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop은 힙에서 요소를 추출합니다.
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
