package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
