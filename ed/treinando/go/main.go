package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	strs := make([]string, len(vet))
	for i, v := range vet {
		strs[i] = strconv.Itoa(v)
	}
	return "[" + strings.Join(strs, ", ") + "]"
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	strs := []string{}
	for i := len(vet) - 1; i >= 0; i-- {
		strs = append(strs, strconv.Itoa(vet[i]))
	}
	return "[" + strings.Join(strs, ", ") + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	for i, j := 0, len(vet)-1; i < j; i, j = i+1, j-1 {
		vet[i], vet[j] = vet[j], vet[i]
	}
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	total := 0
	for _, v := range vet {
		total += v
	}
	return total
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice do menor valor
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}

	var rec func(v []int) (int, int)
	rec = func(v []int) (int, int) {
		if len(v) == 1 {
			return 0, v[0]
		}
		idx, val := rec(v[1:])
		if v[0] <= val {
			return 0, v[0]
		}
		return idx + 1, val
	}

	idx, _ := rec(vet)
	return idx
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}