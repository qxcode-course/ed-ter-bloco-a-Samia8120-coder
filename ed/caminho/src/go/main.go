package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return nil
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	prev := make([][]Pos, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
		prev[i] = make([]Pos, cols)
		for j := 0; j < cols; j++ {
			prev[i][j] = Pos{-1, -1}
		}
	}

	queue := NewQueue[Pos]()
	queue.Enqueue(startPos)
	visited[startPos.l][startPos.c] = true

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	found := false

	for !queue.IsEmpty() {
		cur, _ := queue.Dequeue()

		if cur == endPos {
			found = true
			break
		}
		for _, d := range dirs {
			nl, nc := cur.l+d[0], cur.c+d[1]
			np := Pos{nl, nc}

			if inside(grid, np) && !visited[nl][nc] && grid[nl][nc] != '#' {
				visited[nl][nc] = true
				prev[nl][nc] = cur
				queue.Enqueue(np)
			}
		}
	}
	if found {
		voltar(grid, startPos, endPos, prev)
	}
}

func voltar(grid [][]rune, start Pos, end Pos, prev [][]Pos) {
	cur := end
	for cur != start {
		grid[cur.l][cur.c] = '.'
		cur = prev[cur.l][cur.c]
	}
	grid[start.l][start.c] = '.'
} 

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
