package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	m := make(map[int]int)
	for _, v := range vet {
		if v < 0 {
			v = -v
		}
		m[v]++
	}

	pares := make([]Pair, 0, len(m))
	for k, v := range m {
		pares = append(pares, Pair{One: k, Two: v})
	}
	sort.Slice(pares, func(i, j int) bool {
		return pares[i].One < pares[j].One
	})
	return pares
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}

	pares := []Pair{}
	atual := vet[0]
	count := 1

	for i := 1; i < len(vet); i++ {
		if vet[i] == atual {
			count++
		} else {
			pares = append(pares, Pair{One: atual, Two: count})
			atual = vet[i]
			count = 1
		}
	}
	pares = append(pares, Pair{One: atual, Two: count})
	return pares
}

func mnext(vet []int) []int {
	n := len(vet)
	resultado := make([]int, n)

	for i := 0; i < n; i++ {
		if vet[i] > 0 {
			mulherE := i > 0 && vet[i-1] < 0
			mulherD := i <  n-1 && vet[i+1] < 0

			if mulherE || mulherD {
				resultado[i] = 1
			}
		}
	}
	return resultado
}

func alone(vet []int) []int {
	n := len(vet)
	resultado := make([]int, n)

	for i := 0; i < n; i++ {
		if vet[i] > 0 {
			esqdOk := i == 0 || vet[i-1] >= 0
			dirtOk := i == n-1 || vet[i+1] >= 0

			if esqdOk && dirtOk {
				resultado[i] = 1
			}
		}
	}
	return resultado
}

func couple(vet []int) int {
	posCount := make(map[int]int)
	negCount := make(map[int]int)

	for _, v := range vet {
		abso := v
		
		if abso < 0 {
			abso = -abso
			negCount[abso]++
		} else {
			posCount[abso]++
		}
	}

	total := 0
	for k, p := range posCount {
		n := negCount[k]
		
		if p < n {
			total += p
		} else {
			total += n
		}
	}
	return total
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos+len(seq) > len(vet) {
		return false
	}

	for i := 0; i < len(seq); i++ {
		if vet[pos+i] != seq[i] {
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	if len(seq) > len(vet) {
		return -1
	}

	for i := 0; i <= len(vet)-len(seq); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	remove := make(map[int]bool)
	for _, pos := range posList {
		remove[pos] = true
	}

	resultado := []int{}
	for i, v := range vet {
		if !remove[i] {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

func clear(vet []int, value int) []int {
	res := []int{}
	for _, v := range vet {
		if v != value {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
