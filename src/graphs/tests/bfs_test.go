package graph_algorithms_test

import (
	"graphs/internal/graph"
	"graphs/internal/graph_algorithms"
	"testing"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name          string
		edges         [][3]int
		startVertex   int
		expectedOrder []int
	}{
		{
			name: "Simple graph",
			edges: [][3]int{
				{0, 1, 1}, {0, 2, 1}, {1, 2, 1}, {2, 0, 1}, {2, 3, 1}, {3, 3, 1},
			},
			startVertex:   1,
			expectedOrder: []int{1, 2, 3, 4},
		},
		{
			name: "Disconnected graph",
			edges: [][3]int{
				{0, 1, 1}, {2, 3, 1},
			},
			startVertex:   1,
			expectedOrder: []int{1, 2},
		},
		{
			name: "Single node graph",
			edges: [][3]int{
				// No edges
			},
			startVertex:   1,
			expectedOrder: []int{1},
		},
		{
			name: "Graph with self-loop",
			edges: [][3]int{
				{0, 0, 1}, {0, 1, 1},
			},
			startVertex:   1,
			expectedOrder: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := graph.NewGraph(4)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1], edge[2])
			}

			result := graph_algorithms.BFS(g, tt.startVertex)

			if !equal(result, tt.expectedOrder) {
				t.Errorf("Expected %v, but got %v", tt.expectedOrder, result)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
