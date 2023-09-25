package main

import (
	"fmt"
)

// TreeNode 구조체 정의
type TreeNode struct {
	Value    string
	Children []*TreeNode
}

// 새로운 TreeNode 생성
func NewTreeNode(value string) *TreeNode {
	return &TreeNode{Value: value}
}

// 자식 노드 추가
func (node *TreeNode) AddChild(child *TreeNode) {
	node.Children = append(node.Children, child)
}

// 트리 출력
func (node *TreeNode) String() string {
	return node.StringWithIndent(0)
}

// 들여쓰기 포함하여 트리 출력
func (node *TreeNode) StringWithIndent(indent int) string {
	indentStr := "\t"
	indentation := ""
	for i := 0; i < indent; i++ {
		indentation += indentStr
	}

	result := fmt.Sprintf("%s%s\n", indentation, node.Value)
	for _, child := range node.Children {
		result += child.StringWithIndent(indent + 1)
	}

	return result
}

func main() {
	// 트리 생성
	root := NewTreeNode("Root")

	// 노드 추가
	child1 := NewTreeNode("Child 1")
	child2 := NewTreeNode("Child 2")
	child3 := NewTreeNode("Child 3")

	root.AddChild(child1)
	root.AddChild(child2)
	root.AddChild(child3)

	child11 := NewTreeNode("Child 1.1")
	child12 := NewTreeNode("Child 1.2")

	child1.AddChild(child11)
	child1.AddChild(child12)

	// 트리 출력
	fmt.Println(root)
}
