package graph_algorithms

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"

	"graphs/internal/graph"
	"graphs/internal/priorityqueue"
)

// Prim реализует алгоритм Прима
func Prim(graph *graph.Graph) (adjacencyMatrix [][]int) {
	n := graph.GraphLength()
	visited := make([]bool, n)
	pq := &priorityqueue.PriorityQueue{}
	heap.Init(pq)

	// Создание матрицы смежности для минимального остовного дерева
	adjacencyMatrix = make([][]int, n)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, n)
	}

	start := generateAndPrintRandomNumber(n) // выбираем случайную вершину

	// Добавляем начальные рёбра у стартовой вершины
	for i := 0; i < n; i++ {
		weight := graph.GetValue(start, i)
		if weight > 0 { // Проверяем, есть ли ребро
			heap.Push(pq, &priorityqueue.Item{Value: i, Weight: weight})
		}
	}

	visited[start] = true // отмечаем стартовую вершину посещенной

	fmt.Println("Минимальное остовное дерево:")

	for pq.Len() > 0 {
		// Извлекаем рёбра с минимальным весом
		edge := heap.Pop(pq).(*priorityqueue.Item)
		if visited[edge.Value] { // пропусукаем вершину, если она уже песещенная
			continue
		}
		visited[edge.Value] = true // отмечаем вершину посещенной
		fmt.Printf("Ребро: %d - %d, вес: %d\n", start, edge.Value, edge.Weight)
		adjacencyMatrix[start][edge.Value] = edge.Weight
		// Добавляем новые рёбра для вершины с минимальным весом
		for i := 0; i < n; i++ {
			weight := graph.GetValue(edge.Value, i)
			if !visited[i] && weight > 0 { // Проверяем, есть ли ребро
				heap.Push(pq, &priorityqueue.Item{Value: i, Weight: weight})
			}
		}
		start = edge.Value
	}
	return adjacencyMatrix // возврат матрицы смежности для минимального остовного дерева
}

// Функция для генерации и вывода случайного числа
func generateAndPrintRandomNumber(max int) int {
	// Создание нового генератора случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Генерация случайного числа
	randomNumber := r.Intn(max)
	return randomNumber
}
