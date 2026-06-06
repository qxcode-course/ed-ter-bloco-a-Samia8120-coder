package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processa(vet []int) {
	_ = vet

	if len(vet) == 1 {
		fmt.Print("[ ")
		fmt.Print(Join(vet, " "))
		fmt.Println(" ]")
		return
	}

	aux := make([]int, len(vet)-1)

	for i := 0; i < len(vet)-1; i++ {
		aux[i] = vet[i] + vet[i+1]
	}

	processa(aux)

	fmt.Print("[ ")
	fmt.Print(Join(vet, " "))
	fmt.Println(" ]")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}

	linha := scanner.Text()
	partes := strings.Fields(linha)
	vet := []int{}

	for _, parte := range partes {
		if value, err := strconv.Atoi(parte); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
}

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := ""
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	return s
}