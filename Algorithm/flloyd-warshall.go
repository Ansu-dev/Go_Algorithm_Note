package main

import (
	"fmt"
)

// 플루이드-워셜(Flloyd-Warshall) 알고리즘은 그래프의 모든 정점 쌍 간의 최단 거리를 찾는 알고리즘입니다.
// 이 알고리즘은 다이나믹 프로그래밍을 기반으로 하며, 3중 반복문을 사용하여 구현할 수 있습니다.
func floydWarshall(graph [][]float64) [][]float64 {
	numVertices := len(graph)
	distances := make([][]float64, numVertices)

	// 거리 행렬 초기화
	for i := range distances {
		distances[i] = make([]float64, numVertices)
		for j := range distances[i] {
			distances[i][j] = graph[i][j]
		}
	}

	// 플로이드-워셜 알고리즘
	for k := 0; k < numVertices; k++ {
		for i := 0; i < numVertices; i++ {
			for j := 0; j < numVertices; j++ {
				if distances[i][j] > distances[i][k]+distances[k][j] {
					distances[i][j] = distances[i][k] + distances[k][j]
				}
			}
		}
	}

	return distances
}

func main() {
	// 그래프 예시 (2차원 배열로 표현, 0은 연결되지 않음)
	graph := [][]float64{
		{0, 5, 0, 8},
		{7, 0, 9, 0},
		{2, 0, 0, 4},
		{0, 0, 3, 0},
	}

	// 모든 정점 쌍 간의 최단 거리 계산
	shortestDistances := floydWarshall(graph)

	// 결과 출력
	for _, row := range shortestDistances {
		fmt.Println(row)
	}
}
