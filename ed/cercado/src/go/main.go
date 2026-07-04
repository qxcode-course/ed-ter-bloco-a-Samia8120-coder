package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	dirs := [][2]int{{1,0}, {-1,0}, {0,1}, {0,-1}}

	var bfs func(r, c int)
	bfs = func(r, c int) {
		queue := [][2]int{{r, c}}
		board[r][c] = 'T'

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			for _, d := range dirs {
				nr, nc := p[0]+d[0], p[1]+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && board[nr][nc] == 'O' {
					board[nr][nc] = 'T'
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			bfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			bfs(i, n-1)
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			bfs(0, j)
		}
		if board[m-1][j] == 'O' {
			bfs(m-1, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
