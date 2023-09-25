package main

import (
	"fmt"
)

// TrieNode 구조체 정의
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie 구조체 정의
type Trie struct {
	root *TrieNode
}

// Trie 초기화
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// 단어 삽입
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// 단어 검색
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func main() {
	// Trie 생성
	myTrie := NewTrie()

	// 단어 삽입
	myTrie.Insert("apple")
	myTrie.Insert("app")
	myTrie.Insert("banana")

	// 단어 검색
	fmt.Println("Search 'apple':", myTrie.Search("apple"))
	fmt.Println("Search 'app':", myTrie.Search("app"))
	fmt.Println("Search 'banana':", myTrie.Search("banana"))
	fmt.Println("Search 'grape':", myTrie.Search("grape"))
}
