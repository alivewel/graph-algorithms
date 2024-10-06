package main

import (
	"container/list"
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
func (g *Graph) AddEdge(from, to int) {
	g.adjacencyMatrix[from][to] = 1
}

// BFS выполняет поиск в ширину
func BFS(graph *Graph, startVertex int) []int {
	vertexCount := len(graph.adjacencyMatrix)
	if vertexCount == 0 || startVertex >= vertexCount {
		return []int{}
	}

	enterOrder := []int{} // посещенные вершины
	visited := make([]bool, vertexCount)
	queue := list.New()

	startVertex-- // Приведение к нулевому индексу
	visited[startVertex] = true
	queue.PushBack(startVertex)
	enterOrder = append(enterOrder, startVertex+1)

	fmt.Printf("Начинаем BFS с вершины %d\n", startVertex+1)

	for queue.Len() > 0 {
		from := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		fmt.Printf("Извлекаем вершину %d из очереди\n", from+1)

		for to := 0; to < vertexCount; to++ {
			if !visited[to] && graph.adjacencyMatrix[from][to] != 0 {
				visited[to] = true
				queue.PushBack(to)
				enterOrder = append(enterOrder, to+1)
				fmt.Printf("Находим и добавляем в очередь вершину %d\n", to+1)
			}
		}
	}

	fmt.Println("Порядок посещения вершин:", enterOrder)
	return enterOrder
}

// DFS выполняет поиск в глубину
func DFS(graph *Graph, startVertex int) []int {
	vertexCount := len(graph.adjacencyMatrix)
	if vertexCount == 0 || startVertex >= vertexCount {
		return []int{}
	}

	enterOrder := []int{}
	visited := make([]bool, vertexCount)
	stack := list.New()

	startVertex-- // Приведение к нулевому индексу
	visited[startVertex] = true
	stack.PushBack(startVertex)
	enterOrder = append(enterOrder, startVertex+1)

	fmt.Printf("Начинаем DFS с вершины %d\n", startVertex+1)

	for stack.Len() > 0 {
		from := stack.Back().Value.(int)
		fmt.Printf("Извлекаем вершину %d из стека\n", from+1)
		isFound := false

		for to := 0; to < vertexCount; to++ {
			fmt.Printf("graph.adjacencyMatrix[from][to] %d | %d %d\n", graph.adjacencyMatrix[from][to], from, to)
			if !visited[to] && graph.adjacencyMatrix[from][to] != 0 {
				from = to
				isFound = true
				visited[to] = true
				stack.PushBack(to)
				enterOrder = append(enterOrder, to+1)
				fmt.Printf("Находим и добавляем в стек вершину %d\n", to+1)
			}
		}

		if !isFound {
			fmt.Printf("Возвращаемся назад от вершины %d\n", from+1)
			stack.Remove(stack.Back())
		}
	}

	fmt.Println("Порядок посещения вершин:", enterOrder)
	return enterOrder
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

func main() {
	graph := NewGraph(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 3)

	graph.PrintAdjacencyMatrix()

	// fmt.Println("Обход графа (BFS):", BFS(graph, 1)) 
	// fmt.Println("Обход графа (DFS):", DFS(graph, 1))
}
