package main

import (
	"bufio"
	"fmt"
	"graphs/internal/s21_graph"
	"graphs/internal/s21_graph_algorithms"
	"os"
	"strconv"
	"strings"
)

func main() {
	var g *s21_graph.Graph
	var loaded bool
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Загрузить граф из файла")
		fmt.Println("2. Обход графа в глубину (DFS)")
		fmt.Println("3. Обход графа в ширину (BFS)")
		fmt.Println("4. Поиск кратчайшего пути между двумя вершинами (Дейкстра)")
		fmt.Println("5. Поиск кратчайших путей между всеми вершинами (Флойд-Уоршелл)")
		fmt.Println("6. Поиск минимального остовного дерева (Прима)")
		fmt.Println("7. Решение задачи коммивояжера (муравьиный алгоритм)")
		fmt.Println("8. Выйти")
		fmt.Print("Выберите действие: ")
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			fmt.Print("Введите путь к файлу (нажать Enter для использования матрицы по умолчанию):")
			scanner.Scan()
			filepath := scanner.Text()
			filepath = strings.Trim(filepath, "'")
			g = s21_graph.NewGraph(0)
			if filepath == "" {
				filepath = "graphs/data/graph_input.txt"
			}
			err := g.LoadGraphFromFile(filepath)
			if err != nil {
				fmt.Printf("\033[31mОшибка загрузки графа: %v\033[0m\n", err)
				return
			} else {
				fmt.Println("\033[32mГраф успешно загружен.\033[0m")
				loaded = true
			}
			// Экспорт графа в файл в формате DOT
			err = g.ExportGraphToDot("graphs/data/graph_output.dot")
			if err != nil {
				fmt.Println("Ошибка экспорта графа:", err)
			} else {
				fmt.Println("Граф успешно экспортирован в graph_output.dot")
			}
		case "2":
			if !loaded {
				fmt.Println("\033[31mСначала загрузите граф.\033[0m")
				continue
			}
			startVertex := getValidVertex(scanner, g.GraphLength(), "Введите начальную вершину для поиска в глубину: ")
			dfsResult := s21_graph_algorithms.DepthFirstSearch(g, startVertex)
			fmt.Printf("\033[32mРезультат поиска в глубину (DFS): %v\033[0m\n", dfsResult)
		case "3":
			if !loaded {
				fmt.Println("\033[31mСначала загрузите граф.\033[0m")
				continue
			}
			startVertex := getValidVertex(scanner, g.GraphLength(), "Введите начальную вершину для поиска в ширину: ")
			bfsResult := s21_graph_algorithms.BreadthFirstSearch(g, startVertex)
			fmt.Printf("\033[32mРезультат поиска в ширину (BFS): %v\033[0m\n", bfsResult)
		case "4":
			if !loaded {
				fmt.Println("\033[31mСначала загрузите граф.\033[0m")
				continue
			}
			vertex1 := getValidVertex(scanner, g.GraphLength(), "Введите первую вершину: ")
			vertex2 := getValidVertex(scanner, g.GraphLength(), "Введите вторую вершину: ")
			shortestPath, path, err := s21_graph_algorithms.GetShortestPathBetweenVertices(g, vertex1, vertex2)
			if err != nil {
				fmt.Printf("\033[31mОшибка при нахождении кратчайшего пути между вершинами: %v\033[0m\n", err)
			} else {
				fmt.Printf("\033[32mКратчайший путь между вершинами %d и %d: %d\nМаршрут: %v\033[0m\n", vertex1, vertex2, shortestPath, path)
			}
		case "5":
			if !loaded {
				fmt.Println("\033[31mСначала загрузите граф.\033[0m")
				continue
			}
			shortestPaths := s21_graph_algorithms.GetShortestPathsBetweenAllVertices(g)
			fmt.Println("\033[32mМатрица кратчайших путей (Флойд-Уоршелл):\033[0m")
			for _, row := range shortestPaths {
				fmt.Println(row)
			}
		case "6":
			if !loaded {
				fmt.Println("\033[31mСначала загрузите граф.\033[0m")
				continue
			}
			mstMatrix := s21_graph_algorithms.GetLeastSpanningTree(g)
			fmt.Println("\033[32mМинимальное остовное дерево (Прима):\033[0m")
			for _, row := range mstMatrix {
				fmt.Println(row)
			}
		case "7":
			if !loaded {
				fmt.Println("\033[31mСначала загрузите граф.\033[0m")
				continue
			}
			tspResult, err := s21_graph_algorithms.SolveTravelingSalesmanProblem(g)
			if err != nil {
				fmt.Printf("\033[31mОшибка при решении задачи коммивояжера: %v\033[0m\n", err)
			} else {
				fmt.Printf("\033[32mРешение задачи коммивояжера (муравьиный алгоритм):\nМаршрут: %v\nДлина маршрута: %.2f\033[0m\n", tspResult.Vertices, tspResult.Distance)
			}
		case "8":
			fmt.Println("\033[32mВыход из программы.\033[0m")
			return
		default:
			fmt.Println("\033[31mНеверный выбор, попробуйте ещё раз.\033[0m")
		}
	}
}

// Функция для валидации ввода вершины
func getValidVertex(scanner *bufio.Scanner, maxVertex int, prompt string) int {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if err == nil && input > 0 && input <= maxVertex {
			return input
		}
		fmt.Printf("\033[31mОшибка: введена некорректная вершина. Введите значение от 1 до %d\033[0m\n", maxVertex)
	}
}
