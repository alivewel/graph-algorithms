package graph_algorithms

import (
	"container/heap"
	"math"

	"graphs/internal/graph"
	"graphs/internal/priorityqueue"
)

// Dijkstra находит кратчайшие пути из s в остальные вершины
func Dijkstra(graph *graph.Graph, start int) []int {
	adjacencyMatrix := graph.GetAdjacencyMatrix()
	vertexCount := len(adjacencyMatrix)
	dist := make([]int, vertexCount)
	for i := range dist {
		dist[i] = math.MaxInt32 // Инициализация бесконечностью
	}
	dist[start] = 0

	visited := make([]bool, vertexCount)
	pq := &priorityqueue.PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &priorityqueue.Item{Value: start, Weight: 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(*priorityqueue.Item).Value
		if visited[u] {
			continue
		}
		visited[u] = true

		for v, weight := range adjacencyMatrix[u] {
			if weight > 0 && !visited[v] {
				newDist := dist[u] + weight
				if newDist < dist[v] {
					dist[v] = newDist
					heap.Push(pq, &priorityqueue.Item{Value: v, Weight: newDist})
				}
			}
		}
	}

	return dist
}
