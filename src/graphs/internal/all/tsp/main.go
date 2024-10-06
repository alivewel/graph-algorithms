package main

import (
	"fmt"
	"math"
)

// Функция для вычисления минимального пути и сохранения маршрута
func tsp(matrix [][]int, visited []bool, currPos, n, count, cost int, minCost *int, path []int, bestPath *[]int) {
	if count == n && matrix[currPos][0] > 0 {
		if cost+matrix[currPos][0] < *minCost {
			*minCost = cost + matrix[currPos][0]
			*bestPath = append([]int(nil), path...)
			*bestPath = append(*bestPath, 0) // возвращаемся в начальный город
		}
		return
	}

	for i := 0; i < n; i++ {
		if !visited[i] && matrix[currPos][i] > 0 {
			visited[i] = true
			path = append(path, i)
			tsp(matrix, visited, i, n, count+1, cost+matrix[currPos][i], minCost, path, bestPath)
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
}

func main() {
	// Пример матрицы смежности
	// matrix := [][]int{
	// 	{0, 10, 15, 20},
	// 	{10, 0, 35, 25},
	// 	{15, 35, 0, 30},
	// 	{20, 25, 30, 0},
	// }
	// matrix := [][]int{
	// 	{0, 9, 13, 15, 17, 19},
	// 	{9, 0, 7, 6, 9, 11},
	// 	{13, 7, 0, 9, 5, 6},
	// 	{15, 6, 9, 0, 7, 9},
	// 	{17, 9, 5, 7, 0, 2},
	// 	{19, 11, 6, 9, 2, 0},
	// }
	matrix := [][]int{
		{0, 14, 7, 16, 17, 9},
		{14, 0, 12, 12, 11, 9},
		{7, 12, 0, 19, 19, 12},
		{16, 12, 19, 0, 2, 8},
		{17, 11, 19, 2, 0, 9},
		{9, 9, 12, 8, 9, 0},
	}
	

	n := len(matrix)
	visited := make([]bool, n)
	visited[0] = true
	minCost := math.MaxInt32
	path := []int{0}
	bestPath := []int{}

	tsp(matrix, visited, 0, n, 1, 0, &minCost, path, &bestPath)

	fmt.Printf("Минимальная стоимость пути: %d\n", minCost)
	fmt.Printf("Лучший маршрут: %v\n", bestPath)
}

// 0, 10, 15, 20,
// 10, 0, 35, 25,
// 15, 35, 0, 30,
// 20, 25, 30, 0,

// 10 + 25 + 30 + 15 = 80
