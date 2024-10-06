package graph_algorithms_test

import (
	"graphs/internal/graph"
	"graphs/internal/graph_algorithms"
	"math"
	"testing"
)

func TestFloydWarshall(t *testing.T) {
	tests := []struct {
		name         string
		edges        [][3]int
		vertexCount  int
		expectedDist [][]int
	}{
		{
			name:        "Simple graph",
			vertexCount: 3,
			edges: [][3]int{
				{0, 1, 3}, {1, 2, 1},
			},
			expectedDist: [][]int{
				{0, 3, 4},
				{math.MaxInt32, 0, 1},
				{math.MaxInt32, math.MaxInt32, 0},
			},
		},
		{
			name:        "Disconnected graph",
			vertexCount: 4,
			edges: [][3]int{
				{0, 1, 5},
			},
			expectedDist: [][]int{
				{0, 5, math.MaxInt32, math.MaxInt32},
				{math.MaxInt32, 0, math.MaxInt32, math.MaxInt32},
				{math.MaxInt32, math.MaxInt32, 0, math.MaxInt32},
				{math.MaxInt32, math.MaxInt32, math.MaxInt32, 0},
			},
		},
		{
			name:        "Graph with self-loop",
			vertexCount: 3,
			edges: [][3]int{
				{0, 0, 2}, {0, 1, 3}, {1, 2, 1},
			},
			expectedDist: [][]int{
				{0, 3, 4},
				{math.MaxInt32, 0, 1},
				{math.MaxInt32, math.MaxInt32, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := graph.NewGraph(tt.vertexCount)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1], edge[2])
			}

			result := graph_algorithms.FloydWarshall(g)

			if !equal2D(result, tt.expectedDist) {
				t.Errorf("Expected %v, but got %v", tt.expectedDist, result)
			}
		})
	}
}

// equal2D compares two 2D slices for equality
func equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
