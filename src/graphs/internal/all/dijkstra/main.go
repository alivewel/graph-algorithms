package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge представляет ребро в графе
type Edge struct {
	to     int
	weight int
}

// WeightedDigraph представляет взвешенный ориентированный граф
type WeightedDigraph struct {
	adjacencyList [][]Edge
}

// NewWeightedDigraph создает новый взвешенный ориентированный граф
func NewWeightedDigraph(v int) *WeightedDigraph {
	return &WeightedDigraph{adjacencyList: make([][]Edge, v)}
}

// AddEdge добавляет ребро в граф
func (g *WeightedDigraph) AddEdge(from, to int, weight int) {
	g.adjacencyList[from] = append(g.adjacencyList[from], Edge{to, weight})
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

func (pq *PriorityQueue) update(item *PriorityQueueItem, dist int) {
	item.dist = dist
	heap.Fix(pq, item.index)
}

// Dijkstra находит кратчайшие пути из s в остальные вершины
func Dijkstra(g *WeightedDigraph, s int) ([]int, []Edge) {
	vertexCount := len(g.adjacencyList)
	dist := make([]int, vertexCount)
	for i := range dist {
		dist[i] = math.MaxInt32 - 1
	}
	pred := make([]Edge, vertexCount)
	marked := make([]bool, vertexCount)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	dist[s] = 0.0
	heap.Push(&pq, &PriorityQueueItem{vertex: s, dist: 0.0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PriorityQueueItem)
		v := item.vertex
		if marked[v] {
			continue
		}
		marked[v] = true

		for _, e := range g.adjacencyList[v] {
			w := e.to
			if dist[w] > dist[v]+e.weight {
				dist[w] = dist[v] + e.weight
				pred[w] = e
				heap.Push(&pq, &PriorityQueueItem{vertex: w, dist: dist[w]})
			}
		}
	}

	return dist, pred
}

func main() {
	g := NewWeightedDigraph(5)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 3)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 1, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 8)
	g.AddEdge(3, 4, 7)
	g.AddEdge(4, 0, 5)

	dist, pred := Dijkstra(g, 0)
	fmt.Println("Минимальные расстояния:", dist)
	fmt.Println("Предыдущие ребра:", pred)
}
