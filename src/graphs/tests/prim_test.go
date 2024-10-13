package graph_algorithms_test

import (
	"graphs/internal/s21_graph"
	"graphs/internal/s21_graph_algorithms"
	"testing"
)

func TestPrim(t *testing.T) {
	tests := []struct {
		name        string
		edges       [][3]int
		vertexCount int
		expectedMST [][]int
	}{
		{
			name:        "Simple graph",
			vertexCount: 4,
			edges: [][3]int{
				{0, 1, 10}, {0, 2, 6}, {0, 3, 5}, {1, 3, 15}, {2, 3, 4},
			},
			expectedMST: [][]int{
				{0, 0, 6, 5},
				{0, 0, 0, 0},
				{0, 0, 0, 4},
				{0, 0, 0, 0},
			},
		},
		{
			name:        "Disconnected graph",
			vertexCount: 4,
			edges: [][3]int{
				{0, 1, 5},
			},
			expectedMST: [][]int{
				{0, 5, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := s21_graph.NewGraph(tt.vertexCount)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1], edge[2])
			}

			result := s21_graph_algorithms.GetLeastSpanningTree(g)

			// Проверка размера результирующей матрицы
			if len(result) != len(tt.expectedMST) {
				t.Errorf("Expected matrix with length %d, but got %d", len(tt.expectedMST), len(result))
				return
			}

			for i := range result {
				if len(result[i]) != len(tt.expectedMST[i]) {
					t.Errorf("Expected row %d with length %d, but got %d", i, len(tt.expectedMST[i]), len(result[i]))
					return
				}
			}
		})
	}
}
