package graph_algorithms

import (
	"math"

	"graphs/internal/graph"
)

// алгоритм Флойда-Уоршелла
func FloydWarshall(graph *graph.Graph) [][]int {
	// Количество вершин в графе
	vertexCount := graph.GraphLength()

	// Инициализация массива минимальных расстояний
	// Приводим матрицу из вида:
	// 	0 3 0
	//  0 0 1
	//  0 0 0
	// в вид, где '-' бесконечность:
	// 	0 3 -
	//  - 0 1
	//  - - 0

	dist := make([][]int, vertexCount)
	for i := range dist {
		dist[i] = make([]int, vertexCount)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else if graph.GetValue(i, j) > 0 {
				dist[i][j] = graph.GetValue(i, j)
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	// Основной цикл алгоритма
	for k := 0; k < vertexCount; k++ {
		for i := 0; i < vertexCount; i++ {
			for j := 0; j < vertexCount; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}
