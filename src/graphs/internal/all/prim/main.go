package main

import (
	"container/heap"
	"fmt"
)

// Edge представляет рёбра графа
type Edge struct {
	to     int
	weight int
}

// Graph представляет граф с рёбрами
type Graph struct {
	vertices int
	edges    [][]Edge
}

// PriorityQueue для хранения рёбер с приоритетом
type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Edge))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// NewGraph создает новый граф
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make([][]Edge, vertices),
	}
}

// AddEdge добавляет рёбра в граф
func (g *Graph) AddEdge(from, to, weight int) {
	g.edges[from] = append(g.edges[from], Edge{to, weight})
	g.edges[to] = append(g.edges[to], Edge{from, weight}) // для неориентированного графа
}

// Prim реализует алгоритм Прима
func (g *Graph) Prim(start int) {
	visited := make([]bool, g.vertices)
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Добавляем начальные рёбра
	for _, edge := range g.edges[start] {
		heap.Push(pq, &edge)
	}
	visited[start] = true

	fmt.Println("Минимальное остовное дерево:")

	for pq.Len() > 0 {
		// Извлекаем рёбра с минимальным весом
		edge := heap.Pop(pq).(*Edge)
		if visited[edge.to] {
			continue
		}
		visited[edge.to] = true
		fmt.Printf("Ребро: %d - %d, вес: %d\n", start, edge.to, edge.weight)

		// Добавляем новые рёбра
		for _, nextEdge := range g.edges[edge.to] {
			if !visited[nextEdge.to] {
				heap.Push(pq, &nextEdge)
			}
		}
		start = edge.to
	}
}

func main() {
	g := NewGraph(5)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 3, 6)
	g.AddEdge(1, 2, 3)
	g.AddEdge(1, 3, 8)
	g.AddEdge(1, 4, 5)
	g.AddEdge(2, 4, 7)
	g.AddEdge(3, 4, 9)

	g.Prim(0) // Запускаем алгоритм Прима с вершины 0
}
