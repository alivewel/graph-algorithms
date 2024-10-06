package main

import (
	"fmt"
	"math"
)

func printValue(value int) {
	if value == math.MaxInt32 {
		fmt.Print("_ ")
	} else {
		fmt.Print(value, " ")
	}
}

func floydWarshall(graph [][]int) [][]int {
	// Количество вершин в графе
	V := len(graph)

	// Инициализация массива минимальных расстояний
	// Приводим матрицу из вида:
	// 	0 3 0
	//  0 0 1
	//  0 0 0
	// в вид, где '-' бесконечность:
	// 	0 3 -
	//  - 0 1
	//  - - 0
	dist := make([][]int, V)
	for i := range dist {
		dist[i] = make([]int, V)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else if graph[i][j] != 0 {
				dist[i][j] = graph[i][j]
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	// printDistMatrix(dist)
	// Основной цикл алгоритма
	for k := 0; k < V; k++ {
		for i := 0; i < V; i++ {
			for j := 0; j < V; j++ {
				// fmt.Println(i, j, "|", i, k, "|", k, j, dist[i][j] > dist[i][k]+dist[k][j])
				// fmt.Println(i, j, "|", i, k, "|", k, j, dist[i][j] > dist[i][k]+dist[k][j])
				if dist[i][j] > dist[i][k]+dist[k][j] {
					// fmt.Println("+")
					fmt.Println(i, j, "|", i, k, "|", k, j, dist[i][j] > dist[i][k]+dist[k][j])
					dist[i][j] = dist[i][k] + dist[k][j]
				} else {
					fmt.Println(i, j, "|", i, k, "|", k, j)
				}
			}
			fmt.Println()
		}
		fmt.Println("---------------")
		fmt.Println()
	}

	return dist
}

func printDistMatrix(dist [][]int) {
	for _, row := range dist {
		for _, val := range row {
			if val == math.MaxInt32 {
				fmt.Print("- ")
			} else {
				fmt.Printf("%d ", val)
			}
		}
		fmt.Println()
	}
}

func printMatrix(dist [][]int) {
	for _, row := range dist {
		for _, val := range row {
			fmt.Printf("%d,", val)
		}
		fmt.Println()
	}
}

func main() {
	// Пример графа в виде матрицы смежности
	// graph := [][]int{
	// 	{0, 3, 0, 0, 0, 0},
	// 	{0, 0, 1, 0, 0, 0},
	// 	{0, 0, 0, 7, 0, 2},
	// 	{0, 0, 0, 0, 0, 3},
	// 	{0, 0, 0, 2, 0, 0},
	// 	{0, 0, 0, 0, 1, 0},
	// }

	// graph := [][]int{
	// 	{0, 3, 0, 0, 0, 0},
	// 	{0, 0, 1, 0, 0, 0},
	// 	{0, 0, 0, 7, 0, 2},
	// 	{0, 0, 0, 0, 0, 3},
	// 	{0, 0, 0, 2, 0, 0},
	// 	{0, 0, 0, 0, 1, 0},
	// }

	graph := [][]int{
		{0, 8, 0, 1},
		{0, 0, 1, 0},
		{4, 0, 0, 0},
		{0, 2, 9, 0},
	}
	// printMatrix(graph)
	dist := floydWarshall(graph)

	fmt.Println("Матрица кратчайших путей:")
	printDistMatrix(dist)
	//	for _, row := range dist {
	//		fmt.Println(row)
	//	}
	// printMatrix(graph)
}

// Матрица смежности
// 0, 8, 0, 1,
// 0, 0, 1, 0,
// 0, 0, 0, 0,
// 0, 2, 9, 0,
