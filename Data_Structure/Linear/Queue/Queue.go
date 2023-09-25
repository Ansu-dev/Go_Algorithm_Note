package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 빈 큐 생성
	myQueue := list.New()

	// 큐에 원소 추가
	myQueue.PushBack(6)
	myQueue.PushBack(7)

	// 큐에서 원소 추출
	frontElement := myQueue.Front()
	if frontElement != nil {
		poppedElement := myQueue.Remove(frontElement)
		fmt.Println(poppedElement)
	}

	// 큐의 가장 앞의 원소 조회 (제거하지 않고)
	frontElement = myQueue.Front()
	if frontElement != nil {
		peekElement := frontElement.Value
		fmt.Println(peekElement)
	}

	// 큐가 비어있는지 확인
	isEmpty := myQueue.Len() == 0
	fmt.Println(isEmpty)

	// 큐 순회(iteration)
	for e := myQueue.Front(); e != nil; e = e.Next() {
		element := e.Value
		fmt.Println(element)
	}
}
