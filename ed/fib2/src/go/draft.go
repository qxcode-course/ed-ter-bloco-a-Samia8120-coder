package main
import "fmt"

var valor [41]int

func f(n int) int {
    if n == 1 || n == 2 {
        return 1
    }
    if n == 3 {
        return 2
    }
    if valor[n] != 0 {
        return valor[n]
    }

    valor[n] = f(n-2) + f(n-3)
    return valor[n]
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(f(n))
}
