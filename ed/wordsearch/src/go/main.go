package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	m := len(grid)
	if m == 0 {
		return false
	}
	n := len(grid[0])
	visited := make([][]bool, m)

	for i := range visited {
		visited[i] = make([]bool, n)
	}
	
	var dfs func(i, j, idx int) bool 
	dfs = func(i, j, idx int) bool {
		if idx == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if visited[i][j] || grid[i][j] != word[idx] {
			return false
		}
		visited[i][j] = true

		if dfs(i+1, j, idx+1) || dfs(i-1, j, idx+1) || dfs(i, j+1, idx+1) || dfs(i, j-1, idx+1) {
			return true
		}
		visited[i][j] = false
		return false	
	}

	for i := 0 ; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
