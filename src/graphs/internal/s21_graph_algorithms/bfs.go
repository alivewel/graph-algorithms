package s21_graph_algorithms

import (
	"container/list"
	"fmt"

	"graphs/internal/s21_graph"
)

// BreadthFirstSearch выполняет поиск в ширину
func BreadthFirstSearch(graph *s21_graph.Graph, startVertex int) []int {
	vertexCount := graph.GraphLength()
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
			if !visited[to] && graph.GetValue(from, to) > 0 {
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
