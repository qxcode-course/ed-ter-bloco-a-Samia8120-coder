package main
import (
    "fmt"
    "math"
)

func main() {
    var n int
    var c int
    fmt.Scan(&n, &c)

    coisa := math.Pow(2, float64(n))
    resultado := float64(c) - float64(c)/coisa
    fmt.Printf("%.2f\n" ,resultado)
}
