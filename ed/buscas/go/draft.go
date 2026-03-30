package main

import "fmt"

func matchingStrings (t1 int, t2 int) {
	fmt.Scan(&t1)
	//palavras := [][]any{}
	vetor_consulta := make(map[int]string)
	for i := 0; i < t1; i++ {
		var palavra string
		fmt.Scan(&palavra)
		vetor_consulta[i] = palavra
	}

	fmt.Scan(&t2)
	vetor_busca := make([]string, t2)
	for i := 0; i < t2; i++ {
		fmt.Scan(&vetor_busca[i])
	}

}

func main() {
	//entrada
	var t1, t2 int
	matchingStrings(t1, t2)

}
