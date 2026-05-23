package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	resultado := []int{}

	for _, v := range vet {
		if v > 0 {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

func getCalmWomen(vet []int) []int {
	resultado := []int{}

	for _, v := range vet {
		if v < 0 && -v < 10 {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

func sortVet(vet []int) []int {
	sort.Ints(vet)
	return vet
}

func sortStress(vet []int) []int {
	sort.SliceStable(vet, func(i, j int) bool {
		moduloI := vet[i]
		if moduloI < 0 {
			moduloI = -moduloI
		}

		moduloJ := vet[j]
		if moduloJ < 0 {
			moduloJ = -moduloJ
		}

		return moduloI < moduloJ
	})
	return vet
}

func reverse(vet []int) []int {
	n := len(vet)
	rever := make([]int, n)

	for i, v := range vet {
		rever[n-1-i] = v
	}
	return rever
}

func unique(vet []int) []int {
	visto := make(map[int]bool)
	resultado := []int{}

	for _, v := range vet {
		if !visto[v] {
			visto[v] = true
			resultado = append(resultado, v)
		}
	}
	return resultado
}

func repeated(vet []int) []int {
	count := make(map[int]int)
	resultado := []int{}

	for _, v := range vet {
		if count[v] >= 1 {
			resultado = append(resultado, v)
		}
		count[v]++
	}
	return resultado
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

