package main

import "fmt"

func matchingStrings (consulta []string, busca []string) []int{
	mapa := make(map[string]int)
	for _, palavra := range consulta {
		mapa[palavra]++
	}
	juncao := make([]int, len(busca))
	for i, palavra := range busca {
		juncao[i] = mapa[palavra]
	}
	return juncao
}

func main() {
	//entrada
	var t1, t2 int
		fmt.Scan(&t1)
	vetor_consulta := make([]string, t1)
	for i := 0; i < t1; i++ {
		fmt.Scan(&vetor_consulta[i])
	}
	fmt.Scan(&t2)
	vetor_busca := make([]string, t2)
	for i := 0; i < t2; i++ {
		fmt.Scan(&vetor_busca[i])
	}

	juncao := matchingStrings(vetor_consulta, vetor_busca)
	for i, qtd := range juncao {
		if i == len(juncao)-1{
			fmt.Printf("%d", qtd)
			break
		}
		fmt.Printf("%d ", qtd)
	}
	fmt.Println()
}
