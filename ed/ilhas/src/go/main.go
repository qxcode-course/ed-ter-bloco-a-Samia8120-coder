package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	
	linhas, colunas := len(grid), len(grid[0])
	ilhas := 0
	busca := func(r, c int) {
		fila := [][2]int{{r, c}}
		grid[r][c] = '0'
	

		for len(fila) > 0 {
			primeiro := fila[0]
			fila = fila[1:]
			lin, col := primeiro[0], primeiro[1]

			if lin-1 >= 0 && grid[lin-1][col] == '1' {
				grid[lin-1][col] = '0'
				fila = append(fila, [2]int{lin-1, col})
			}
			if lin+1 < linhas && grid[lin+1][col] == '1' {
				grid[lin+1][col] = '0'
				fila = append(fila, [2]int{lin+1, col})
			}
			if col-1 >= '0' && grid[lin][col-1] == '1' {
				grid[lin][col-1] = '0'
				fila = append(fila, [2]int{lin, col-1})
			}
			if col+1 < colunas && grid[lin][col+1] == '1' {
				grid[lin][col+1] = '0'
				fila = append(fila, [2]int{lin, col+1})
			}
		}
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			if grid[i][j] == '1' {
				ilhas++
				busca(i, j)
			}
		}
	}
	return ilhas
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
