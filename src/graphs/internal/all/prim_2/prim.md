### Алгоритм Прима

В данной статье я бы хотел объяснить работу алгоритма Прима. Алгоритм используется для нахождения минимального остовного дерева. Сам алгоритм очень прост, в статье хотел бы поделиться своей реализации на языке Go. 

#### Начальные термины

- **Граф** — это структура данных в которой хранятся вершины и связи между ними. Удобнее всего представлять графы в виде матрицы смежности.

- **Матрица смежности** — эта квадратная матрица, размер матрицы равен количеству вершин в графе. В ней хранится информация о соседях вершин графа.

- **Минимальное остовное дерево** — это поиск минимального количества ребер, чтобы из одной любой вершины графа попасть в любую другую вершину графа. Не стоит забывать про отсутствие циклов в дереве.

- **Приоритетная очередь** — по структура данных, внутри которой, в отличие от обычной очереди элементы отсортирированы по приоритету (в нашем случае по весу). Элементы с наивысшим приоритетом извлекаются первыми.

#### Описание работы алгоритма

- **Структура Graph**: Структура содержит внутри себя двумерный срез (матрица смежности).
- **NewGraph**: Конструктор для создания Graph.
- **AddEdge**: Функция для добавления соседей вершины.

<details>
<summary>Структура Graph </summary>

```
// Graph представляет граф с матрицей смежности
type Graph struct {
	adjacencyMatrix [][]int
}
```

```
// NewGraph создает новый граф
func NewGraph(vertexCount int) *Graph {
	matrix := make([][]int, vertexCount)
	for i := range matrix {
		matrix[i] = make([]int, vertexCount)
	}
	return &Graph{adjacencyMatrix: matrix}
}
```

```
// AddEdge добавляет ребро в граф
func (g *Graph) AddEdge(from, to, weight int) {
	g.adjacencyMatrix[from][to] = weight
}
```

</details>
<br>

Для реализации приоритетной очереди мы используем структуру данных под названием куча (heap) из пакета стандартной библиотеки container/heap. Для этого нам нужно реализовать методы Len(), Less(), Swap(), Push(), Pop().

<details>
<summary> Реализация PriorityQueue </summary>

```
// Item представляет элемент с приоритетом
type Item struct {
	value  int
	weight int
	index  int
}
```

```
// PriorityQueue для хранения рёбер с приоритетом
type PriorityQueue []*Item
```

```
// Len возвращает количество элементов в приоритетной очереди.
func (pq PriorityQueue) Len() int { return len(pq) }
```

```
// Less сравнивает два элемента очереди по их весу.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}
```

```
// Swap меняет местами два элемента в очереди по их индексам i и j.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
```

```
// Push добавляет новый элемент в приоритетную очередь.
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}
```

```
// Pop удаляет и возвращает элемент с наивысшим приоритетом (наименьшим весом).
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
```

</details>

#### Алгоритм Прима

**Начальные условия**:  
- Срез вершин `unvisited`, в котором хранятся все вершины. 
- Приоритетная очередь, которая отвечает за поиск соседа с минимальным весом.
- Матрица смежности, в которой мы записываем получившееся минимальное остовное дерево.

**Шаг 1**:  
Из слайса `unvisited` выбираем случайную вершину. У этой вершины просматриваем всех соседей и выбираем соседа с минимальным весом ребра. Добавляем выбранную вершину в слайс `visited`.

**Шаг 2 и последующие**:  
У вершин из слайса `unvisited` просматриваем всех соседей и выбираем соседа с минимальным весом ребра. Добавляем выбранную вершину в слайс `visited`.  
Проделываем шаги до тех пор, пока слайс `unvisited` не окажется пустым.

**Результатом работы алгоритма** является остовное дерево минимальной стоимости, а точнее её  матрица смежности.

<details>
<summary> Алгоритм Прима </summary>

```
// Prim реализует алгоритм Прима
func (g *Graph) Prim() (adjacencyMatrix [][]int) {
	n := len(g.adjacencyMatrix) // количество вершин в графе
	visited := make([]bool, n)
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Создание матрицы смежности для минимального остовного дерева
	adjacencyMatrix = make([][]int, n)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, n)
	}

	start := generateAndPrintRandomNumber(len(g.adjacencyMatrix)) // выбираем случайную вершину

	// Добавляем начальные рёбра у стартовой вершины
	for i, weight := range g.adjacencyMatrix[start] {
		if weight > 0 { // Проверяем, есть ли ребро
			heap.Push(pq, &Item{value: i, weight: weight})
		}
	}
	visited[start] = true // отмечаем стартовую вершину посещенной

	fmt.Println("Минимальное остовное дерево:")

	for pq.Len() > 0 {
		// Извлекаем рёбра с минимальным весом
		edge := heap.Pop(pq).(*Item)
		if visited[edge.value] { // пропусукаем вершину, если она уже песещенная
			continue
		}
		visited[edge.value] = true // отмечаем вершину посещенной
		fmt.Printf("Ребро: %d - %d, вес: %d\n", start, edge.value, edge.weight)
		adjacencyMatrix[start][edge.value] = edge.weight
		// Добавляем новые рёбра для вершины с минимальным весом
		for i, weight := range g.adjacencyMatrix[edge.value] {
			if !visited[i] && weight > 0 { // Проверяем, есть ли ребро
				heap.Push(pq, &Item{value: i, weight: weight})
			}
		}
		start = edge.value
	}
	return adjacencyMatrix // возврат матрицы смежности для минимального остовного дерева
}
```

```
// Функция для генерации и вывода случайного числа
func generateAndPrintRandomNumber(max int) int {
	// Создание нового генератора случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Генерация случайного числа
	randomNumber := r.Intn(max)
	fmt.Println("Случайное число:", randomNumber)
	return randomNumber
}
```

</details>

<br>

Полное и подробное описание работы алгоритма можно найти в моем [GitHub](https://github.com/alivewel/prim).
