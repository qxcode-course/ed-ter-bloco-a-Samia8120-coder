package main

import (
	"fmt"
	"math"
)

func main() {
	var n1, n2, n3 float64
	fmt.Scanln(n1, n2, n3)
	//fmt.Scanln(n2)
	//fmt.Scanln(n3)

	p := (n1 + n2 + n3) / 2
	//temp := p * (p - n1) * (p - n2) * (p - n3)
	res := math.Sqrt((p * (p - n1*(p-n2)*(p-n3))))

	fmt.Print(res)
}
