package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int{
	esquerda, direita := 0, len(slice)-1
	achado := -1

	for esquerda <= direita {
		meio := (esquerda + direita) / 2

		if slice[meio] == value{
			achado = meio
			break
		} else if slice[meio] < value {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}

	if achado != -1 {
		i := achado

		for i+1 < len(slice) && slice[i+1] == value {
			i++
		}
		return i
	}
	return esquerda
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	partes := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)

	for _, elem := range partes[1 : len(partes)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	resultado := MagicSearch(slice, value)
	fmt.Println(resultado)
}
/*
func MagicSearch(slice []int, value int) int {
	left, right := 0, len(slice)-1
	found := -1

	for left <= right {
		mid := (left + right) / 2

		if slice[mid] == value {
			found = mid
			break
		} else if slice[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if found != -1 {
		i := found
		for i+1 < len(slice) && slice[i+1] == value {
			i++
		}
		return i
	}

	return left
}
*/