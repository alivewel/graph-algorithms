package main

import (
	"fmt"
	"graphs/internal/graph"
	"graphs/internal/graph_algorithms"
)


func main() {
	vertexCount := 4
	g := graph.NewGraph(vertexCount)

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
