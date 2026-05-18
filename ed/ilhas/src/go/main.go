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
	
	rows, cols := len(grid), len(grid[0])
	islands := 0
	bfs := func(r, c int) {
		queue := [][2]int{{r, c}}
		grid[r][c] = '0'
	

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			row, col := curr[0], curr[1]

			if row-1 >= 0 && grid[row-1][col] == '1' {
				grid[row-1][col] = '0'
				queue = append(queue, [2]int{row-1, col})
			}
			if row+1 < rows && grid[row+1][col] == '1' {
				grid[row+1][col] = '0'
				queue = append(queue, [2]int{row+1, col})
			}
			if col-1 >= '0' && grid[row][col-1] == '1' {
				grid[row][col-1] = '0'
				queue = append(queue, [2]int{row, col-1})
			}
			if col+1 < cols && grid[row][col+1] == '1' {
				grid[row][col+1] = '0'
				queue = append(queue, [2]int{row, col+1})
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				islands++
				bfs(i, j)
			}
		}
	}
	return islands
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
