package graph_algorithms_test

import (
	"graphs/internal/s21_graph"
	"graphs/internal/s21_graph_algorithms"
	"math"
	"testing"
)

func TestSolveTravelingSalesmanProblem(t *testing.T) {
	// Создаем граф с 4 вершинами
	graph := s21_graph.NewGraph(4)
	graph.SetAdjacencyMatrix([][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	})

	result, err := s21_graph_algorithms.SolveTravelingSalesmanProblem(graph)
	if err != nil {
		t.Fatalf("Ошибка при решении задачи коммивояжера: %v", err)
	}

	expectedDistance := 80.0             // Ожидаемое расстояние для кратчайшего пути
	expectedPath := []int{0, 1, 3, 2, 0} // Ожидаемый путь (возвращение в исходную точку)

	// Проверка расстояния
	if math.Abs(result.Distance-expectedDistance) > 1e-9 {
		t.Errorf("Неверное расстояние: ожидается %v, получено %v", expectedDistance, result.Distance)
	}

	// Проверка пути
	for i, v := range expectedPath {
		if result.Vertices[i] != v {
			t.Errorf("Неверный путь на позиции %d: ожидается %v, получено %v", i, v, result.Vertices[i])
		}
	}
}
