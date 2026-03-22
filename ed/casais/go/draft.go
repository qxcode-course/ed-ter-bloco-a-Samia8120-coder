package main
import "fmt"
func main() {
    N := 0
    fmt.Scan(&N)
    animais := make([]int, N)
    
    solteiros := make(map[int]int)
    for i := range animais{
        fmt.Scan(&animais[i])
    }
    n_pares := 0
    for _, animal := range animais{
        if solteiros[-animal] > 0 {
            solteiros[-animal]--
            n_pares++
        } else {
            solteiros[animal]++
        }
    }
    fmt.Println(n_pares)
}
