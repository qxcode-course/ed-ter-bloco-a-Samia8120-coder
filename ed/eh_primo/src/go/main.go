package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	_, _ = x, div
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0{
			return false
		}
	}
	return true
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
