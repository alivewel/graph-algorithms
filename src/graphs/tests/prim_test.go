package graph_algorithms_test

import (
	"graphs/internal/graph"
	"graphs/internal/graph_algorithms"
	"testing"
)

func TestPrim(t *testing.T) {
	tests := []struct {
		name        string
		edges       [][3]int
		vertexCount int
		expectedMST [][]int
	}{
		// {
		// 	name:        "Simple graph",
		// 	vertexCount: 4,
		// 	edges: [][3]int{
		// 		{0, 1, 10}, {0, 2, 6}, {0, 3, 5}, {1, 3, 15}, {2, 3, 4},
		// 	},
		// 	expectedMST: [][]int{
		// 		{0, 0, 6, 5},
		// 		{0, 0, 0, 0},
		// 		{0, 0, 0, 4},
		// 		{0, 0, 0, 0},
		// 	},
		// },
		// {
		// 	name:        "Disconnected graph",
		// 	vertexCount: 4,
		// 	edges: [][3]int{
		// 		{0, 1, 5},
		// 	},
		// 	expectedMST: [][]int{
		// 		{0, 5, 0, 0},
		// 		{0, 0, 0, 0},
		// 		{0, 0, 0, 0},
		// 		{0, 0, 0, 0},
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := graph.NewGraph(tt.vertexCount)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1], edge[2])
				// g.AddEdge(edge[1], edge[0], edge[2]) // Добавляем обратное ребро для неориентированного графа
			}

			result := graph_algorithms.Prim(g)

			if !equal2D(result, tt.expectedMST) {
				t.Errorf("Expected %v, but got %v", tt.expectedMST, result)
			}
		})
	}
}
