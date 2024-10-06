package priorityqueue

// Item представляет элемент с приоритетом
type Item struct {
	Value  int
	Weight int
	Index  int
}

// PriorityQueue для хранения рёбер с приоритетом
type PriorityQueue []*Item

// Len возвращает количество элементов в приоритетной очереди.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less сравнивает два элемента очереди по их весу.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
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
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
