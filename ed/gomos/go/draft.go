package main

import "fmt"

func main() {
	var Q int
	var D string

	fmt.Scan(&Q)
	fmt.Scan(&D)

	vetor_x := make([]int, Q)
	vetor_y := make([]int, Q)

	for i := 0; i < Q; i++ {
		fmt.Scan(&vetor_x[i], &vetor_y[i])
	}

	//Logica

	if D == "L" {
		for i := 0; i < Q; i++ {
			if i == 0 {
				fmt.Print(vetor_x[i]-1, vetor_y[i], "\n")
			} else {
				fmt.Print(vetor_x[i-1], vetor_y[i-1], "\n")
			}
		}
	}
	if D == "R" {
		for i := 0; i < Q; i++ {
			if i == 0 {
				fmt.Print(vetor_x[i]+1, vetor_y[i], "\n")
			} else {
				fmt.Print(vetor_x[i-1], vetor_y[i-1], "\n")
			}
		}
	}
	if D == "U" {
		for i := 0; i < Q; i++ {
			if i == 0 {
				fmt.Print(vetor_x[i], vetor_y[i]-1, "\n")
			} else {
				fmt.Print(vetor_x[i-1], vetor_y[i-1], "\n")
			}
		}
	}
	if D == "D" {
		for i := 0; i < Q; i++ {
			if i == 0 {
				fmt.Print(vetor_x[i], vetor_y[i]+1, "\n")
			} else {
				fmt.Print(vetor_x[i-1], vetor_y[i-1], "\n")
			}
		}
	}
}
