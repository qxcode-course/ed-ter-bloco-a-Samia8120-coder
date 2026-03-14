package main

import "fmt"

func comparacao(a string, i int) {
	if i < 12 {
		fmt.Print(a, " eh crianca\n")
	} else if i >= 12 && i < 18 {
		fmt.Print(a, " eh jovem\n")
	} else if i >= 18 && i < 65 {
		fmt.Print(a, " eh adulto\n")
	} else if i >= 65 && i < 1000 {
		fmt.Print(a, " eh idoso\n")
	} else {
		fmt.Print(a, " eh mumia\n")
	}
}

func main() {
	var nome string
	var idade int

	fmt.Scanln(&nome)
	fmt.Scan(&idade)

	comparacao(nome, idade)
}
