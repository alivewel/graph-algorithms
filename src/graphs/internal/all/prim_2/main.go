package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
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
	g.adjacencyMatrix[from][to] = weight
}

// Item представляет элемент с приоритетом
type Item struct {
	value  int
	weight int
	index  int
}

// PriorityQueue для хранения рёбер с приоритетом
type PriorityQueue []*Item

// Len возвращает количество элементов в приоритетной очереди.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less сравнивает два элемента очереди по их весу.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

// Swap меняет местами два элемента в очереди по их индексам i и j.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push добавляет новый элемент в приоритетную очередь.
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

// Pop удаляет и возвращает элемент с наивысшим приоритетом (наименьшим весом).
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Prim реализует алгоритм Прима
func (g *Graph) Prim() (adjacencyMatrix [][]int) {
	n := len(g.adjacencyMatrix) // количество вершин в графе
	visited := make([]bool, n)
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Создание матрицы смежности для минимального остовного дерева
	adjacencyMatrix = make([][]int, n)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, n)
	}

	start := generateAndPrintRandomNumber(len(g.adjacencyMatrix)) // выбираем случайную вершину

	// Добавляем начальные рёбра у стартовой вершины
	for i, weight := range g.adjacencyMatrix[start] {
		if weight > 0 { // Проверяем, есть ли ребро
			heap.Push(pq, &Item{value: i, weight: weight})
		}
	}
	visited[start] = true // отмечаем стартовую вершину посещенной

	fmt.Println("Минимальное остовное дерево:")

	for pq.Len() > 0 {
		// Извлекаем рёбра с минимальным весом
		edge := heap.Pop(pq).(*Item)
		if visited[edge.value] { // пропусукаем вершину, если она уже песещенная
			continue
		}
		visited[edge.value] = true // отмечаем вершину посещенной
		fmt.Printf("Ребро: %d - %d, вес: %d\n", start, edge.value, edge.weight)
		adjacencyMatrix[start][edge.value] = edge.weight
		// Добавляем новые рёбра для вершины с минимальным весом
		for i, weight := range g.adjacencyMatrix[edge.value] {
			if !visited[i] && weight > 0 { // Проверяем, есть ли ребро
				heap.Push(pq, &Item{value: i, weight: weight})
			}
		}
		start = edge.value
	}
	return adjacencyMatrix // возврат матрицы смежности для минимального остовного дерева
}

// Функция для генерации и вывода случайного числа
func generateAndPrintRandomNumber(max int) int {
	// Создание нового генератора случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Генерация случайного числа
	randomNumber := r.Intn(max)
	fmt.Println("Случайное число:", randomNumber)
	return randomNumber
}

func PrintAdjacencyMatrix(adjacencyMatrix [][]int) {
	fmt.Println("Матрица смежности:")
	for _, row := range adjacencyMatrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func main() {
	// g := NewGraph(5)
	// g.AddEdge(0, 1, 2)
	// g.AddEdge(0, 3, 6)
	// g.AddEdge(1, 2, 3)
	// g.AddEdge(1, 3, 8)
	// g.AddEdge(1, 4, 5)
	// g.AddEdge(2, 4, 7)
	// g.AddEdge(3, 4, 9)

	// mat := g.Prim() // Запускаем алгоритм Прима

	// PrintAdjacencyMatrix(mat)

	vertexCount := 4
	g := NewGraph(vertexCount)

	edges := [][3]int{
		{0, 1, 10}, {0, 2, 6}, {0, 3, 5}, {1, 3, 15}, {2, 3, 4},
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1], edge[2])
	}

	result := g.Prim()

	fmt.Println(result)
}
