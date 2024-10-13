package graph_algorithms_test

import (
	"graphs/internal/s21_graph"
	"graphs/internal/s21_graph_algorithms"
	"math"
	"testing"
)

func TestGetShortestPathBetweenVertices(t *testing.T) {
	// Создаем граф с 5 вершинами
	graph := s21_graph.NewGraph(5)
	graph.AddEdge(0, 1, 10)
	graph.AddEdge(0, 2, 3)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 2)
	graph.AddEdge(2, 1, 4)
	graph.AddEdge(2, 3, 8)
	graph.AddEdge(2, 4, 2)
	graph.AddEdge(3, 4, 7)
	graph.AddEdge(4, 3, 9)

	graph.PrintAdjacencyMatrix()
	
	tests := []struct {
		vertex1      int
		vertex2      int
		expectedDist int
		expectedPath []int
		expectError  bool
	}{
		{1, 4, 9, []int{1, 3, 2, 4}, false}, // Путь 1 -> 3 -> 2 -> 4 с расстоянием 9
		{1, 3, 3, []int{1, 3}, false},       // Путь 1 -> 3 с расстоянием 3
		{2, 3, 1, []int{2, 3}, false},       // Путь 2 -> 3 с расстоянием 1
		{3, 1, math.MaxInt32, []int{3}, false}, // Путь недостижим, возвращаем только начальную вершину
		{0, 4, 0, nil, true},                // Ошибка: vertex1 < 1
		{6, 4, 0, nil, true},                // Ошибка: vertex1 > vertexCount
		{2, 6, 0, nil, true},                // Ошибка: vertex2 > vertexCount
	}

	for _, test := range tests {
		dist, path, err := s21_graph_algorithms.GetShortestPathBetweenVertices(graph, test.vertex1, test.vertex2)

		if test.expectError {
			if err == nil {
				t.Errorf("Ожидалась ошибка для вершин %d -> %d, но ошибка не была получена", test.vertex1, test.vertex2)
			}
			continue
		}

		if err != nil {
			t.Errorf("Ошибка при вызове функции: %v", err)
		}

		if dist != test.expectedDist {
			t.Errorf("Неверное расстояние для %d -> %d: ожидается %d, получено %d", test.vertex1, test.vertex2, test.expectedDist, dist)
		}

		if !equal(path, test.expectedPath) {
			t.Errorf("Неверный путь для %d -> %d: ожидается %v, получено %v", test.vertex1, test.vertex2, test.expectedPath, path)
		}
	}
}
