package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	n     = 4     // количество городов
	Q     = 5 * n // константа усиления феромона (отвечает за регулирование количества феромона, которое добавляется в длину пути)
	alpha = 9.0   // регулировка алгоритма по расстоянию
	beta  = 5.0   // регулировка алгоритма по количеству феромона
	p     = 0.7   // скорость испарения феромона
)

// City хранит координаты вершины
type City struct {
	x, y float64
}

// # Координаты вершин
// initCities генерирует несколько (const n) точек (City) с случайными
// координатами по x и y в диапозоне от 0 до 19 включительно.
func initCities() []City {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	cities := make([]City, n)
	for i := 0; i < n; i++ {
		cities[i] = City{
			x: float64(r.Intn(20)),
			y: float64(r.Intn(20)),
		}
	}
	return cities
}

// # Матрица феромонов Wt
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

// # Матрица приращений феромонов DWt
// initDeltaPheromones создает пустую начальную матрицу DWt
func initDeltaPheromones() [][]float64 {
	pheromones := make([][]float64, n)
	for i := 0; i < n; i++ {
		pheromones[i] = make([]float64, n)
	}
	return pheromones
}

// # Каждой вершине - свой вес
// calculateProbabilities рассчитывает вероятность P по формуле (см. readme)
func calculateProbabilities(pheromones, distanceMatrix [][]float64, currentCity int, visited []bool) []float64 {
	probabilities := make([]float64, n)
	sum := 0.0

	for i := 0; i < n; i++ {
		if !visited[i] {
			probabilities[i] = math.Pow(pheromones[currentCity][i], alpha) * math.Pow(1.0/distanceMatrix[currentCity][i], beta)
			sum += probabilities[i]
		}
	}

	for i := 0; i < n; i++ {
		probabilities[i] /= sum
	}

	return probabilities
}

// # Выбор направления
// selectNextCity принимает вероятности которые расчитываются calculateProbabilities
// и возвращает номер выбранной вершины, используя метод рулетки
func selectNextCity(probabilities []float64) int {
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

// # Добавление новых следов
func addNewTrails(pheromones, deltaPheromones [][]float64, p float64) {
	for i := range pheromones {
		for j := range pheromones[i] {
			pheromones[i][j] += deltaPheromones[i][j] * p
		}
	}
}

// # Испарение феромонов
func evaporatePheromones(pheromones [][]float64, p float64) {
	for i := range pheromones {
		for j := range pheromones[i] {
			pheromones[i][j] *= (1 - p)
		}
	}
}

// calculatePathLength расчитывает общцю длину пути
// проходимся по всему пути и суммируем длины
func calculatePathLength(path []int, distanceMatrix [][]float64) float64 {
	L := 0.0
	for i := 0; i < len(path)-1; i++ {
		from := path[i]
		to := path[i+1]
		L += distanceMatrix[from][to]
	}
	// Добавляем расстояние от последней вершины обратно к первой
	L += distanceMatrix[path[len(path)-1]][path[0]]
	return L
}

func calculateDistanceMatrix(cities []City) [][]float64 {
	n := len(cities)
	W := make([][]float64, n)
	for i := range W {
		W[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				W[i][j] = 0
			} else {
				distance := math.Sqrt(math.Pow(cities[i].x-cities[j].x, 2) + math.Pow(cities[i].y-cities[j].y, 2))
				W[i][j] = distance
			}
		}
	}
	return W
}

// Пометка дуг феромоном
func updateDistanceMatrix(deltaPheromones [][]float64, path []int, L float64) {
	for i := 0; i < len(path)-1; i++ {
		from := path[i]
		to := path[i+1]
		deltaPheromones[from][to] += Q / L
	}
}

func main() {
	// Координаты вершин
	cities := initCities()
	pheromones := initPheromones()           // Wt След
	deltaPheromones := initDeltaPheromones() // DWt Матрица приращений следа

	// Симметричная матрица расстояний
	distanceMatrix := calculateDistanceMatrix(cities)

	Lmin := math.MaxFloat64 // начальное значение длины - infinity
	bestPath := make([]int, n+1)

	for k := 0; k < 100; k++ { // Основной цикл
		for ant := 0; ant < n; ant++ { // Цикл по муравьям
			visited := make([]bool, n) // Список непосещенных вершин
			currentCity := ant         // Начальная вершина для муравья ant
			visited[currentCity] = true
			path := []int{currentCity} // # Начало пути

			for len(path) < n { // Цикл j1
				probabilities := calculateProbabilities(pheromones, distanceMatrix, currentCity, visited) // Каждой вершине - свой вес
				nextCity := selectNextCity(probabilities)                                                 // Выбор направления
				visited[nextCity] = true                                                                  // Tabu list увеличился
				path = append(path, nextCity)
				currentCity = nextCity // # Начало дуги = конец предыдущей
			}
			path = append(path, path[0]) // # Конец пути (возвращаемся в исходную точку)
			// # Добавляем последнюю дугу
			L := calculatePathLength(path, distanceMatrix)

			// Проверка минимальной длины пути
			if L < Lmin {
				Lmin = L
				copy(bestPath, path)
			}

			// # Пометка дуг феромоном
			updateDistanceMatrix(deltaPheromones, path, L)
		}

		evaporatePheromones(pheromones, p)           // # Испарение феромонов
		addNewTrails(pheromones, deltaPheromones, p) // # Добавление новых следов
	}

	fmt.Println("Лучший путь:", bestPath)
	fmt.Println("Минимальная длина пути:", Lmin)
}
