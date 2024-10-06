package main

import (
	"fmt"
	"graphs/internal/graph"
	"graphs/internal/graph_algorithms"
)


func main() {
	// matrix := [][]int{
	// 	{0, 14, 7, 16, 17, 9},
	// 	{14, 0, 12, 12, 11, 9},
	// 	{7, 12, 0, 19, 19, 12},
	// 	{16, 12, 19, 0, 2, 8},
	// 	{17, 11, 19, 2, 0, 9},
	// 	{9, 9, 12, 8, 9, 0},
	// }
	// graph := graph.NewGraph(len(matrix))

	// for i := 0; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix[i]); j++ {
	// 		graph.AddEdge(i, j, matrix[i][j])
	// 	}
	// }

	// res := graph_algorithms.SolveTravelingSalesmanProblem(graph)
	// fmt.Println(res)

	vertexCount := 4
	g := graph.NewGraph(vertexCount)

	// edges := [][3]int{
	// 	{0, 1, 10}, {0, 2, 6}, {0, 3, 5}, {1, 3, 15}, {2, 3, 4},
	// }

	edges := [][3]int{
		{0, 1, 10}, {0, 2, 6}, {0, 3, 5}, {1, 3, 15}, {2, 3, 4},
		{1, 2, 8},  // Добавлено ребро между вершинами 1 и 2
		{2, 0, 6},  // Добавлено ребро между вершинами 2 и 0 (если граф ориентированный)
		{3, 1, 15}, // Добавлено ребро между вершинами 3 и 1 (если граф ориентированный)
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1], edge[2])
	}
	g.PrintAdjacencyMatrix()
	result := graph_algorithms.Prim(g)

	fmt.Println(result)
}
