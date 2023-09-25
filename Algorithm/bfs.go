package main

import (
	"container/list"
	"fmt"
)

// 상, 하, 좌, 우 탐색을 위한 방향 배열
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func bfs(x, y int, visited [][]bool, grid [][]int) {
	queue := list.New()
	queue.PushBack([2]int{x, y})
	visited[x][y] = true

	for queue.Len() > 0 {
		pos := queue.Front()
		queue.Remove(pos)
		xy := pos.Value.([2]int)
		x, y := xy[0], xy[1]
		fmt.Printf("방문 위치: (%d, %d)\n", x, y)

		// 현재 위치에서 인접한 위치 탐색
		for d := 0; d < 4; d++ {
			nx, ny := x+dx[d], y+dy[d]

			// 배열 범위 안에 있고 방문하지 않았다면 큐에 넣고 탐색
			if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[0]) && !visited[nx][ny] && grid[nx][ny] == 1 {
				queue.PushBack([2]int{nx, ny})
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
	bfs(0, 0, visited, grid)
}
