package main

import "fmt"

func posicao(x int, y int, d string) {
	if d == "L" {
		x--
		fmt.Print(x, y, "\n")
	} else if d == "R" {
		x++
		fmt.Print(x, y, "\n")
	} else if d == "U" {
		y--
		fmt.Print(x, y, "\n")
	} else if d == "D" {
		y++
		fmt.Print(x, y, "\n")
	}
}
func main() {
	var Q int
	var D string

	fmt.Scan(&Q)
	fmt.Scan(&D)

	var x int
	var y int

	for i := 0; i < Q; i++ {
		fmt.Scan(&x, &y)
		posicao(x, y, D)
	}

	//for i := 0; i < Q; i++ {
	//	posicao(x, y, D)
	//}
}
