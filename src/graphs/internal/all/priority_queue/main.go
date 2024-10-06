package main

import (
	"container/heap"
	"fmt"
)

// Item представляет элемент с приоритетом
type Item struct {
	value    string
	priority int
	index    int
}

// PriorityQueue реализует приоритетную очередь
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
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

func main() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Item{value: "Task A", priority: 3})
	heap.Push(&pq, &Item{value: "Task B", priority: 1})
	heap.Push(&pq, &Item{value: "Task C", priority: 2})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Извлечена задача: %s с приоритетом %d\n", item.value, item.priority)
	}
}
