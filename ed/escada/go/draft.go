package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    if n == 1 || n == 2 {
        fmt.Print("1\n")
        return
    }
    if n == 3{
        fmt.Print("2\n")
        return
    }

    a, b, c := 1, 1, 2 
    var res int
    for i := 4; i <= n; i++ {
        res = a + c
        a, b, c = b, c, res
    }
    fmt.Print(res, "\n")
}
