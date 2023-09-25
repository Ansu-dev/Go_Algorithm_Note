package main

import (
	"fmt"
)

func main() {
	// 빈 스택 생성
	myStack := []int{}

	// 스택에 원소 추가
	myStack = append(myStack, 6)
	myStack = append(myStack, 7)

	// 스택에서 원소 추출
	stackSize := len(myStack)
	if stackSize > 0 {
		poppedElement := myStack[stackSize-1]
		myStack = myStack[:stackSize-1]
		fmt.Println(poppedElement)
	}

	// 스택의 가장 위의 원소 조회 (제거하지 않고)
	stackSize = len(myStack)
	if stackSize > 0 {
		peekElement := myStack[stackSize-1]
		fmt.Println(peekElement)
	}

	// 스택이 비어있는지 확인
	isEmpty := len(myStack) == 0
	fmt.Println(isEmpty)

	// 스택 순회(iteration)
	for _, element := range myStack {
		fmt.Println(element)
	}
}
