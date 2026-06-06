package main

import (
	"bufio"
	//"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fila := NewQueue[string]()
	for _, letra := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"} {
		fila.Enqueue(letra)
	}
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 15; i++ {
		scanner.Scan()
		
		linha := scanner.Text()
		partes := strings.Split(linha, " ")
		
		m, err1 := strconv.Atoi(partes[0])
		n, err2 := strconv.Atoi(partes[1])
		if err1 != nil || err2 != nil {
			fmt.Fprintln(os.Stderr, "Erro: gols nao numericos")
			return
		}

		time1 := fila.Dequeue()
		time2 := fila.Dequeue()

		var vencedor string
		if m > n {
			vencedor = time1			
		} else {
			vencedor = time2
		}
		fila.Enqueue(vencedor)
	}
	campeao := fila.Dequeue()
	fmt.Println(campeao)
}
