package main

import (
	"container/heap"
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
func (g *Graph) AddEdge(from, to, value int) {
	g.adjacencyMatrix[from][to] = value
}

// PriorityQueueItem представляет элемент приоритетной очереди
type PriorityQueueItem struct {
	vertex int
	dist   int
	index  int
}

// PriorityQueue реализует приоритетную очередь
type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// func (pq *PriorityQueue) update(item *PriorityQueueItem, dist float64) {
// 	item.dist = dist
// 	heap.Fix(pq, item.index)
// }

// Dijkstra находит кратчайшие пути из s в остальные вершины
func Dijkstra(graph Graph, start int) []int {
	vertexCount := len(graph.adjacencyMatrix)
	dist := make([]int, vertexCount)
	for i := range dist {
		dist[i] = int(^uint(0) >> 1) // Инициализация бесконечностью
	}
	dist[start] = 0

	visited := make([]bool, vertexCount)
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &PriorityQueueItem{vertex: start, dist: 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(*PriorityQueueItem).vertex
		if visited[u] {
			continue
		}
		visited[u] = true

		for v, weight := range graph.adjacencyMatrix[u] {
			if weight > 0 && !visited[v] {
				newDist := dist[u] + weight
				if newDist < dist[v] {
					dist[v] = newDist
					heap.Push(pq, &PriorityQueueItem{vertex: v, dist: newDist})
				}
			}
		}
	}

	return dist
}

func main() {
	g := NewGraph(5)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 3)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 1, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 8)
	g.AddEdge(3, 4, 7)
	g.AddEdge(4, 0, 5)

	dist := Dijkstra(*g, 0)
	fmt.Println("Минимальные расстояния:", dist)
}
