package s21_graph

import (
	"errors"
	"fmt"
)

// Graph представляет граф с матрицей смежности
type Graph struct {
	adjacencyMatrix [][]int
}

// NewGraph создает новый граф
func NewGraph(vertexCount int) *Graph {
	matrix := make([][]int, vertexCount)
	for i := range matrix {
		matrix[i] = make([]int, vertexCount)
	}
	return &Graph{adjacencyMatrix: matrix}
}

// AddEdge добавляет ребро в граф
func (g *Graph) AddEdge(from, to, weight int) {
	if from < 0 || from >= len(g.adjacencyMatrix) || to < 0 || to >= len(g.adjacencyMatrix[from]) {
		// Индексы вне диапазона, ничего не делаем
		return
	}
	g.adjacencyMatrix[from][to] = weight
}

// GetAdjacencyMatrix возвращает матрицу смежности
func (g *Graph) GetAdjacencyMatrix() [][]int {
	return g.adjacencyMatrix
}

// SetAdjacencyMatrix устанавливает матрицу смежности
func (g *Graph) SetAdjacencyMatrix(matrix [][]int) error {
	// Проверка на инициализацию массива
	if len(matrix) == 0 {
		return errors.New("матрица смежности не может быть пустой")
	}

	// Проверка на равенство размеров строк
	for _, row := range matrix {
		if len(row) != len(matrix) {
			return errors.New("матрица смежности должна быть квадратной")
		}
	}

	g.adjacencyMatrix = matrix
	return nil
}

// GraphLength возвращает длину adjacencyMatrix
func (g *Graph) GraphLength() int {
	return len(g.adjacencyMatrix)
}

// GetValue возвращает значение из матрицы смежности
func (g *Graph) GetValue(row, col int) int {
	if row < 0 || row >= len(g.adjacencyMatrix) || col < 0 || col >= len(g.adjacencyMatrix[row]) {
		return -1
	}
	return g.adjacencyMatrix[row][col]
}

// PrintAdjacencyMatrix печатает матрицу смежности графа
func (g *Graph) PrintAdjacencyMatrix() {
	fmt.Println("Матрица смежности:")
	for _, row := range g.adjacencyMatrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
