package main

import (
	"fmt"
)

type HashTable struct {
	size  int
	table map[string]int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:  size,
		table: make(map[string]int, size),
	}
}

func (ht *HashTable) _hashFunction(key string) int {
	return len(key) % ht.size
}

func (ht *HashTable) Insert(key string, value int) {
	// _ = ht._hashFunction(key)
	// Go에서는 이미 해시 충돌 처리를 내장하고 있으므로 별도로 리스트 등을 사용할 필요가 없습니다.
	ht.table[key] = value
}

func (ht *HashTable) Get(key string) (int, bool) {
	value, found := ht.table[key]
	return value, found
}

func (ht *HashTable) Remove(key string) {
	delete(ht.table, key)
}

func (ht *HashTable) Display() {
	for key, value := range ht.table {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

func main() {
	// 해시 테이블 생성
	myHashTable := NewHashTable(10)

	// 데이터 삽입
	myHashTable.Insert("apple", 5)
	myHashTable.Insert("banana", 7)
	myHashTable.Insert("cherry", 3)

	// 데이터 조회
	if value, found := myHashTable.Get("apple"); found {
		fmt.Println("apple:", value)
	}
	if value, found := myHashTable.Get("banana"); found {
		fmt.Println("banana:", value)
	}
	if value, found := myHashTable.Get("grape"); found {
		fmt.Println("grape:", value)
	}

	// 데이터 삭제
	myHashTable.Remove("apple")

	// 해시 테이블 출력
	myHashTable.Display()
}
