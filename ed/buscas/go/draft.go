package main

import "fmt"

type str struct {
}

func main() {
	var t1, t2 int
	fmt.Scan(&t1)
	vetor_consulta := [...]string{}
	for i := 0; i < t1; i++ {
		fmt.Scan(&vetor_consulta[i])
	}

	fmt.Scan(&t2)
	vetor_busca := [...]string{}
	for i := 0; i < t2; i++ {
		fmt.Scan(&vetor_busca[i])
	}
}
