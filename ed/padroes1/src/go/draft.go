package main
import "fmt"

func quadrados(n int)int {
    if n == 1 {
        return 20
    }
    return quadrados(n-1) + 8
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(quadrados(n))
}
