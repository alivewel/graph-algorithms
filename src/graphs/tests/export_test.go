package graph_algorithms_test

import (
	"os"
	"testing"
	"graphs/internal/s21_graph"
)

func TestExportGraphToDot(t *testing.T) {
	// Создаем тестовый граф
	graph := s21_graph.NewGraph(4)
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(0, 2, 1)
	graph.AddEdge(1, 3, 1)

	// Определяем имя файла для экспорта
	filename := "test_graph.dot"
	defer os.Remove(filename) // Удаляем файл после теста

	// Экспортируем граф в файл DOT
	err := graph.ExportGraphToDot(filename)
	if err != nil {
		t.Fatalf("Ошибка при экспорте графа: %v", err)
	}

	// Читаем содержимое файла
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Ошибка при чтении файла: %v", err)
	}

	// Ожидаемое содержимое файла DOT
	expectedContent := "graph G {\n    1 -- 2;\n    1 -- 3;\n    2 -- 4;\n}\n"

	// Проверяем содержимое
	if string(content) != expectedContent {
		t.Errorf("Неверное содержимое файла: ожидается %q, получено %q", expectedContent, string(content))
	}
}
