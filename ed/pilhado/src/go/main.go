package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Top() T {
	if len(s.data) == 0 {
		panic("stack vazia")
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
*/
type Pos struct {
	l, c int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	line := scanner.Text()
	var linha, coluna int
	fmt.Sscanf(line, "%d %d", &linha, &coluna)

	laberinto := make([][]byte, linha)
	comeco := Pos{-1, -1}
	final := Pos{-1, -1}

	for i := 0; i < linha; i++ {
		scanner.Scan()
		laberinto[i] = []byte(scanner.Text())

		for j := 0; j < coluna; j++ {
			switch laberinto[i][j] {
			case 'I':
				comeco = Pos{i, j}
				laberinto[i][j] = ' '
			case 'F':
				final = Pos{i, j}
				laberinto[i][j] = ' '
			}
		}
	}
	visitado := make([][]bool, linha)
	for i := range visitado {
		visitado[i]  = make([]bool, coluna)
	}
	
	caminhoStack := NewStack[Pos]()
	visitado[comeco.l][comeco.c] = true
	dire := []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for !caminhoStack.IsEmpty() {
		agora := caminhoStack.Top()

		if agora.l == final.l && agora.c == final.c {
			break
		}

		var vizinho []Pos
		for _, d := range dire {
			nl, nc := agora.l+d.l, agora.c+d.c

			if nl >= 0 && nl < linha && nc >= 0 && nc < coluna && laberinto[nl][nc] != '#' && !visitado[nl][nc] {
				vizinho = append(vizinho, Pos{nl, nc})
			}
		}

		if len(vizinho) > 0 {
			proximo := vizinho[0]
			visitado[proximo.l][proximo.c] = true
			caminhoStack.Push(proximo)
		} else {
			caminhoStack.Pop()
		}
	}

	var rota []Pos

	for !caminhoStack.IsEmpty() {
		rota = append(rota, caminhoStack.Pop())
	}

	for i, j := 0, len(rota)-1; i < j; i, j = i+1, j-1 {
		rota[i], rota[j] = rota[j], rota[i]
	}

	for _, p := range rota {
		laberinto[p.l][p.c] = '.'
	}

	for i := 0; i < linha; i++ {
		fmt.Println(string(laberinto[i]))
	}
}