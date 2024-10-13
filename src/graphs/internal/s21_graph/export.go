package s21_graph

import (
	"fmt"
	"os"
)

// Экспортируем граф в файл формата DOT
func (g *Graph) ExportGraphToDot(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot create file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString("graph G {\n"); err != nil {
		return fmt.Errorf("cannot write to file: %v", err)
	}

	vertexCount := g.GraphLength() // Получаем количество вершин

	for i := 0; i < vertexCount; i++ {
		for j := 0; j < vertexCount; j++ {
			if g.GetValue(i, j) != 0 { // Проверяем наличие ребра
				if _, err := file.WriteString(fmt.Sprintf("    %d -- %d;\n", i+1, j+1)); err != nil {
					return fmt.Errorf("cannot write edge: %v", err)
				}
			}
		}
	}

	if _, err := file.WriteString("}\n"); err != nil {
		return fmt.Errorf("cannot close graph: %v", err)
	}

	return nil
}
