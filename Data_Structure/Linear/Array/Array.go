package main

import "fmt"

func main() {
	// 빈 슬라이스 생성
	var mySlice []int

	// 원소를 포함한 슬라이스 생성
	mySliceWithElements := []int{1, 2, 3, 4, 5}

	// 다차원 슬라이스 생성
	_ = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	// 슬라이스의 길이(크기) 확인
	_ = len(mySliceWithElements)

	// 슬라이스의 특정 원소에 접근
	_ = mySliceWithElements[2] // 3번째 원소에 접근

	// 슬라이스의 원소 변경
	mySliceWithElements[0] = 10 // 첫 번째 원소를 10으로 변경

	// 슬라이스에 원소 추가
	mySlice = append(mySlice, 6) // 슬라이스에 6을 추가

	// 슬라이스의 원소 삭제
	indexToDelete := 2
	mySliceWithElements = append(mySliceWithElements[:indexToDelete], mySliceWithElements[indexToDelete+1:]...)

	// 슬라이스 순회(iteration)
	for _, element := range mySliceWithElements {
		fmt.Println(element)
	}

	// 슬라이스의 특정 원소 인덱스 찾기
	searchValue := 4
	index := -1
	for i, element := range mySliceWithElements {
		if element == searchValue {
			index = i
			break
		}
	}

	fmt.Println("Index of 4:", index)

	// 슬라이스 슬라이싱
	startIndex := 1
	endIndex := 4
	subSlice := mySliceWithElements[startIndex:endIndex] // 2번째부터 4번째 원소까지 슬라이싱
	fmt.Println(subSlice)
}
