package graph_algorithms_test

import (
	"graphs/internal/graph"
	"graphs/internal/graph_algorithms"
	"math"
	"testing"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name         string
		edges        [][3]int
		startVertex  int
		expectedDist []int
	}{
		{
			name: "Simple graph",
			edges: [][3]int{
				{0, 1, 10}, {0, 2, 3}, {1, 2, 1}, {2, 1, 4},
				{1, 3, 2}, {2, 3, 8}, {3, 4, 7}, {4, 0, 5},
			},
			startVertex:  0,
			expectedDist: []int{0, 7, 3, 9, 16},
		},
		{
			name: "Disconnected graph",
			edges: [][3]int{
				{0, 1, 1}, {2, 3, 1},
			},
			startVertex:  0,
			expectedDist: []int{0, 1, math.MaxInt32, math.MaxInt32, math.MaxInt32},
		},
		{
			name:  "Single node graph",
			edges: [][3]int{
				// No edges
			},
			startVertex: 0,
			expectedDist: []int{0, math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := graph.NewGraph(5)
			for _, edge := range tt.edges {
				g.AddEdge(edge[0], edge[1], edge[2])
			}

			result := graph_algorithms.Dijkstra(g, tt.startVertex)

			if !equal(result, tt.expectedDist) {
				t.Errorf("Expected %v, but got %v", tt.expectedDist, result)
			}
		})
	}
}
