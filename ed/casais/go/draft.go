package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    solteiros := make(map[int]int)
    pares := 0

    for i := 0; i < n; i++ {
        var animal int
        fmt.Scan(&animal)

        if solteiros[-animal] > 0 {
            pares++
            solteiros[-animal]--
        } else {
            solteiros[animal]++
        }
    }
    fmt.Println(pares)
}
