package main
import "fmt"
func main() {
    var N int
    var cont_pares int
    especie := make([]int, N)
    
    fmt.Scan(&N)
    for i := 0; i < N; i++{
        fmt.Scan(&especie[i])
    }
}
