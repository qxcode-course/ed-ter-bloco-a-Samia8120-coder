package main
import "fmt"

func pecas(n int) int {
    if n == 1 {
        return 3
    }
    return pecas(n-1) + 2*n + 1
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(pecas(n))
}
