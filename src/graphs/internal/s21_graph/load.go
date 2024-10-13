package s21_graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Загружаем граф из файла
func (g *Graph) LoadGraphFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]int

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var row []int

		for _, val := range line {
			num, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("invalid number in file: %v", err)
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error while reading file: %v", err)
	}

	g.SetAdjacencyMatrix (matrix)

	return nil
}

