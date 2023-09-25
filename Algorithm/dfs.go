package main

import (
	"fmt"
)

// 방향 벡터
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func dfsStack(x, y int, visited [][]bool, grid [][]int) {
	stack := [][]int{{x, y}}
	// 현재 위치를 방문했다고 표시
	visited[x][y] = true

	for len(stack) > 0 {
		// 스택에서 현재 위치를 꺼냅니다.
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		x, y := current[0], current[1]
		fmt.Printf("방문 위치: (%d, %d)\n", x, y)

		// 현재 위치에서 인접한 위치 탐색
		for d := 0; d < 4; d++ {
			nx, ny := x+dx[d], y+dy[d]

			// 배열 범위 안에 있고 방문하지 않았다면 스택에 넣고 탐색
			if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[0]) && !visited[nx][ny] && grid[nx][ny] == 1 {
				stack = append(stack, []int{nx, ny})
				visited[nx][ny] = true
			}
		}
	}
}

func main() {
	// 2차원 배열 예제 (1은 갈 수 있는 경로, 0은 갈 수 없는 경로)
	grid := [][]int{
		{1, 1, 1, 1},
		{1, 0, 1, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
	}

	// 방문 정보
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	// 예시로 (0, 0)부터 탐색 시작
	dfsStack(0, 0, visited, grid)
}
