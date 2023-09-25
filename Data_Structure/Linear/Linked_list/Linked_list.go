package main

import (
	"fmt"
)

// Node 구조체 정의
type Node struct {
	data int
	next *Node
}

// LinkedList 구조체 정의
type LinkedList struct {
	head *Node
}

// LinkedList에 노드 추가
func (ll *LinkedList) append(data int) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// LinkedList 순회
func (ll *LinkedList) traverse() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	// 빈 연결 리스트 생성
	myLinkedList := &LinkedList{}

	// 노드 추가
	myLinkedList.append(1)
	myLinkedList.append(2)
	myLinkedList.append(3)

	// 연결 리스트 순회
	myLinkedList.traverse()
}
