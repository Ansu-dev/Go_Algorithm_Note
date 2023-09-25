package main

import (
	"fmt"
)

// MinHeap 구조체 정의
type MinHeap []int

// 최소 힙에 원소 추가
func (h *MinHeap) Push(value int) {
	*h = append(*h, value)
	n := len(*h) - 1
	// 새로운 원소를 힙의 제자리로 이동하여 최소 힙 속성 유지
	for n > 0 && (*h)[n] < (*h)[(n-1)/2] {
		(*h)[n], (*h)[(n-1)/2] = (*h)[(n-1)/2], (*h)[n]
		n = (n - 1) / 2
	}
}

// 최소 힙에서 최솟값 추출
func (h *MinHeap) ExtractMin() int {
	if len(*h) == 0 {
		return -1 // 힙이 비어있을 경우 -1 반환 또는 에러 처리
	}
	min := (*h)[0]
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	if len(*h) > 0 {
		(*h)[0] = last
		// 힙의 루트에서부터 아래로 이동하여 최소 힙 속성 유지
		n := 0
		for {
			left := 2*n + 1
			right := 2*n + 2
			smallest := n
			if left < len(*h) && (*h)[left] < (*h)[smallest] {
				smallest = left
			}
			if right < len(*h) && (*h)[right] < (*h)[smallest] {
				smallest = right
			}
			if smallest == n {
				break
			}
			(*h)[n], (*h)[smallest] = (*h)[smallest], (*h)[n]
			n = smallest
		}
	}
	return min
}

// 최소 힙 출력
func (h MinHeap) Display() {
	fmt.Println(h)
}

func main() {
	// 최소 힙 생성
	myHeap := &MinHeap{}

	// 데이터 삽입
	myHeap.Push(5)
	myHeap.Push(3)
	myHeap.Push(8)
	myHeap.Push(1)
	myHeap.Push(10)

	// 힙 출력
	myHeap.Display()

	// 최소값 추출
	minVal := myHeap.ExtractMin()
	fmt.Println("Extracted min value:", minVal)

	// 힙 출력
	myHeap.Display()
}
