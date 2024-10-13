package s21_graph_algorithms

import (
	"errors"
	"math"
	"math/rand"

	"graphs/internal/s21_graph"
)

const (
	alpha = 9.0 // регулировка алгоритма по расстоянию (чем больше значение, тем большее будет влияние расстояния на выбор направления, при alpha = 0.0 вообще не будет учитываться)
	beta  = 5.0 // регулировка алгоритма по количеству феромона (чем больше значение, тем большее будет влияние количества феромона на выбор направления, при beta = 0.0 вообще не будет учитываться)
	p     = 0.7 // скорость испарения феромона
)

var (
	n int     // количество вершин
	Q float64 // константа усиления феромона (отвечает за регулирование количества феромона, которое добавляется в длину пути)
)

// Матрица феромонов Wt
// initPheromones создает начальную матрицу Wt с одинаковым значением количества феромона - 0.1
func initPheromones() [][]float64 {
	pheromones := make([][]float64, n)
	for i := 0; i < n; i++ {
		pheromones[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			pheromones[i][j] = 0.1
		}
	}
	return pheromones
}

// Матрица приращений феромонов DWt
// initDeltaPheromones создает пустую начальную матрицу DWt
func initDeltaPheromones() [][]float64 {
	pheromones := make([][]float64, n)
	for i := 0; i < n; i++ {
		pheromones[i] = make([]float64, n)
	}
	return pheromones
}

// Каждой вершине - свой вес
// calculateProbabilities рассчитывает вероятность P по формуле (см. readme)
func calculateProbabilities(pheromones [][]float64, distanceMatrix [][]int, currentCity int, visited []bool) []float64 {
	probabilities := make([]float64, n)
	sum := 0.0

	for i := 0; i < n; i++ {
		if !visited[i] {
			distance := distanceMatrix[currentCity][i]
			if distance == 0 {
				// Обработка случая, когда расстояние равно нулю
				continue
			}
			probabilities[i] = math.Pow(pheromones[currentCity][i], alpha) * math.Pow(1.0/float64(distance), beta)
			sum += probabilities[i]
		}
	}

	if sum == 0 {
		// Обработка случая, когда сумма равна нулю
		return probabilities
	}

	for i := 0; i < n; i++ {
		probabilities[i] /= sum
	}

	return probabilities
}

// Выбор направления
// selectNextVertex принимает вероятности которые расчитываются calculateProbabilities
// и возвращает номер выбранной вершины, используя метод рулетки
func selectNextVertex(probabilities []float64) int {
	sm := 0.0                                  // сумма всех вероятностей
	beg := make([]float64, len(probabilities)) // начальные границы интервалов вероятностей для каждого города
	end := make([]float64, len(probabilities)) // конечные границы интервалов вероятностей для каждого города

	for _, p := range probabilities {
		sm += p // суммируем вероятности
	}

	c := 0.0 // отслеживание текущей позиции в интервале вероятностей
	for i := range probabilities {
		beg[i] = c
		end[i] = c + probabilities[i]/sm
		c = end[i]
	}

	m := rand.Float64()
	for i := range probabilities {
		if beg[i] <= m && end[i] >= m {
			return i // возвращаем номер выбранной вершины
		}
	}

	return len(probabilities) - 1 // редкий случай
}

// Добавление новых следов
func addNewTrails(pheromones, deltaPheromones [][]float64, p float64) {
	for i := range pheromones {
		for j := range pheromones[i] {
			pheromones[i][j] += deltaPheromones[i][j] * p
		}
	}
}

// Испарение феромона
func evaporatePheromones(pheromones [][]float64, p float64) {
	for i := range pheromones {
		for j := range pheromones[i] {
			pheromones[i][j] *= (1 - p)
		}
	}
}

// calculatePathLength расчитывает общую длину пути
// проходимся по всему пути и суммируем длины
func calculatePathLength(path []int, distanceMatrix [][]int) float64 {
	// func calculatePathLength(path []int, distanceMatrix [][]float64) float64 {
	L := 0.0
	for i := 0; i < len(path)-1; i++ {
		from := path[i]
		to := path[i+1]
		L += float64(distanceMatrix[from][to])
	}
	// Добавляем расстояние от последней вершины обратно к первой
	L += float64(distanceMatrix[path[len(path)-1]][path[0]])
	return L
}

// updateDeltaPheromones помечает дуги феромоном
func updateDeltaPheromones(deltaPheromones [][]float64, path []int, L float64) {
	for i := 0; i < len(path)-1; i++ {
		from := path[i]
		to := path[i+1]
		deltaPheromones[from][to] += Q / L
	}
}

type TsmResult struct {
	Vertices []int   // посещенные вершины (с порядком прохождения)
	Distance float64 // длина маршрута
}

// SolveTravelingSalesmanProblem решает задачу коммивояжера с использованием муравьиного алгоритма
func SolveTravelingSalesmanProblem(graph *s21_graph.Graph) (TsmResult, error) {
	distanceMatrix := graph.GetAdjacencyMatrix()
	n = len(distanceMatrix)
	Q = 5 * float64(n)

	if n < 3 {
		return TsmResult{}, errors.New("недостаточно вершин для решения задачи коммивояжера")
	}

	// Координаты вершин
	pheromones := initPheromones()           // Wt След
	deltaPheromones := initDeltaPheromones() // DWt Матрица приращений следа

	Lmin := math.MaxFloat64      // начальное значение длины - infinity
	bestPath := make([]int, n+1) // n+1 для возврата в исходную точку

	for k := 0; k < 100; k++ { // Основной цикл
		for ant := 0; ant < n; ant++ { // Цикл по муравьям
			visited := make([]bool, n) // Список непосещенных вершин
			currentCity := ant         // Начальная вершина для муравья ant
			visited[currentCity] = true
			path := []int{currentCity} // Начало пути

			for len(path) < n { // Цикл пострения пути
				probabilities := calculateProbabilities(pheromones, distanceMatrix, currentCity, visited) // Каждой вершине - свой вес
				nextCity := selectNextVertex(probabilities)                                                 // Выбор направления
				visited[nextCity] = true                                                                  // Tabu list увеличился
				path = append(path, nextCity)
				currentCity = nextCity // Начало дуги = конец предыдущей
			}
			path = append(path, path[0]) // Конец пути (возвращаемся в исходную точку)

			// проходимся по всему пути и суммируем длины
			L := calculatePathLength(path, distanceMatrix)

			// Проверка минимальной длины пути
			if L < Lmin {
				Lmin = L
				copy(bestPath, path)
			}

			// Пометка дуг феромоном
			updateDeltaPheromones(deltaPheromones, path, L)
		}

		evaporatePheromones(pheromones, p)           // Испарение феромона
		addNewTrails(pheromones, deltaPheromones, p) // Добавление новых следов
	}

	if len(bestPath) == 0 {
		return TsmResult{}, errors.New("не удалось найти решение задачи коммивояжера")
	}

	return TsmResult{
		Vertices: bestPath,
		Distance: Lmin,
	}, nil
}
