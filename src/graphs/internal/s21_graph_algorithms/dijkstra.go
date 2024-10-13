package s21_graph_algorithms

import (
	"container/heap"
	"errors"
	"math"

	"graphs/internal/priorityqueue"
	"graphs/internal/s21_graph"
)

// GetShortestPathBetweenVertices (Dijkstra) находит кратчайший путь из vertex1 в vertex2
func GetShortestPathBetweenVertices(graph *s21_graph.Graph, vertex1, vertex2 int) (int, []int, error) {
	vertexCount := graph.GraphLength() // количество вершин
	if vertex1 < 1 || vertex1 > vertexCount || vertex2 < 1 || vertex2 > vertexCount {
		return 0, nil, errors.New("vertex1 или vertex2 выходит за пределы количества вершин графа")
	}

	distances := make([]int, vertexCount)    // расстояния до всех остальных вершин
	visited := make([]bool, vertexCount)     // посещенные
	predecessors := make([]int, vertexCount) // предки

	for i := range distances {
		distances[i] = math.MaxInt32
		predecessors[i] = -1
	}

	distances[vertex1-1] = 0

	// Инициализируем приоритетную очередь
	pq := &priorityqueue.PriorityQueue{}
	heap.Push(pq, &priorityqueue.Item{Value: vertex1, Weight: 0})

	for pq.Len() > 0 {
		currentItem := heap.Pop(pq).(*priorityqueue.Item)
		currentVertex := currentItem.Value

		if visited[currentVertex-1] {
			continue
		}
		visited[currentVertex-1] = true

		// Обрабатываем всех соседей текущей вершины
		for i := 0; i < vertexCount; i++ {
			if graph.GetValue(currentVertex-1, i) != 0 && !visited[i] {
				newDist := distances[currentVertex-1] + graph.GetValue(currentVertex-1, i)
				if newDist < distances[i] {
					distances[i] = newDist
					predecessors[i] = currentVertex // Запоминаем предка
					heap.Push(pq, &priorityqueue.Item{Value: i + 1, Weight: newDist})
				}
			}
		}
	}

	// Восстановление пути
	path := []int{}
	if distances[vertex2-1] == math.MaxInt32 {
		return distances[vertex2-1], []int{vertex1}, nil // Если путь недостижим, возвращаем только начальную вершину
	}

	for v := vertex2; v != -1; v = predecessors[v-1] {
		path = append([]int{v}, path...)
	}

	return distances[vertex2-1], path, nil
}
