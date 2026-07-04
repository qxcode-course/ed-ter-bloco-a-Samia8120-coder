package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	dirs := [][2]int{{-1,0}, {1,0}, {0,1}, {0,-1}}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if memo[i][j] > 0 {
			return memo[i][j]
		}

		best := 1
		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && matrix[ni][nj] > matrix[i][j] {
				candidate := 1 + dfs(ni, nj)
				if candidate > best {
					best = candidate
				}
			}
		}
		memo[i][j] = best
		return best
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			val := dfs(i, j)
			if val > ans {
				ans = val
			}
		}
	}
	return ans
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
