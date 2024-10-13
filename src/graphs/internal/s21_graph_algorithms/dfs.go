package s21_graph_algorithms

import (
	"container/list"
	"fmt"

	"graphs/internal/s21_graph"
)

// DepthFirstSearch выполняет поиск в глубину
func DepthFirstSearch(graph *s21_graph.Graph, startVertex int) []int {
	vertexCount := graph.GraphLength()
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
			if !visited[to] && graph.GetValue(from, to) > 0 {
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
