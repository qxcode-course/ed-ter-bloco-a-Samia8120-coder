package main
import "fmt"

func pontos(n, m int) int {
    if m == 1 {
        return 1
    }
    
    return pontos(n, m-1) + (n-2)*(m-1) + 1
}

func main() {
    var n, m int
    fmt.Scan(&n, &m)
    fmt.Println(pontos(n, m))
}
