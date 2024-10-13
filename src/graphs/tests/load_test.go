package graph_algorithms_test

import (
	"graphs/internal/s21_graph"
	"os"
	"testing"
)

func TestLoadGraphFromFile(t *testing.T) {
	// Создаем временный файл для теста
	tempFile, err := os.CreateTemp("", "test_graph_*.txt")
	if err != nil {
		t.Fatalf("Ошибка при создании временного файла: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Удаляем файл после теста

	// Записываем тестовые данные в файл
	testData := "0 1 2\n1 0 3\n2 3 0\n"
	if _, err := tempFile.WriteString(testData); err != nil {
		t.Fatalf("Ошибка при записи в файл: %v", err)
	}
	tempFile.Close() // Закрываем файл после записи

	// Создаем граф и загружаем его из файла
	graph := s21_graph.NewGraph(3)
	err = graph.LoadGraphFromFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Ошибка при загрузке графа: %v", err)
	}

	// Проверяем, что матрица смежности загружена правильно
	expectedMatrix := [][]int{
		{0, 1, 2},
		{1, 0, 3},
		{2, 3, 0},
	}

	for i := range expectedMatrix {
		for j := range expectedMatrix[i] {
			if graph.GetValue(i, j) != expectedMatrix[i][j] {
				t.Errorf("Неверное значение в матрице: ожидается %d, получено %d", expectedMatrix[i][j], graph.GetValue(i, j))
			}
		}
	}
}
