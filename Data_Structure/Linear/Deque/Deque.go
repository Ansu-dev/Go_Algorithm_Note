package main

import (
	"fmt"
)

func main() {
	// 빈 슬라이스 생성
	var mySlice []int

	// 원소를 포함한 슬라이스 생성
	mySliceWithElements := []int{1, 2, 3, 4, 5}

	// 슬라이스의 길이(크기) 확인
	_ = len(mySliceWithElements)

	// 슬라이스의 앞쪽에 원소 추가 (덱의 appendleft와 유사)
	mySlice = append([]int{0}, mySlice...)

	// 슬라이스의 뒤쪽에 원소 추가 (덱의 append와 유사)
	mySlice = append(mySlice, 6)

	// 슬라이스의 앞쪽 원소 추출 (덱의 popleft와 유사)
	_ = mySlice[0]
	mySlice = mySlice[1:]

	// 슬라이스의 뒤쪽 원소 추출 (덱의 pop과 유사)
	rearIndex := len(mySlice) - 1
	_ = mySlice[rearIndex]
	mySlice = mySlice[:rearIndex]

	// 슬라이스 순회(iteration)
	for _, element := range mySliceWithElements {
		fmt.Println(element)
	}

	// 슬라이스의 특정 원소 인덱스 찾기
	searchValue := 3
	index := -1
	for i, element := range mySliceWithElements {
		if element == searchValue {
			index = i
			break
		}
	}

	fmt.Println("Index of 3:", index)
}
