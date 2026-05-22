package main
import "fmt"
func main() {
    var n, k int
    fmt.Scan(&n, &k)

    if n == 1 || n == 2 {
        fmt.Println(1)
        return 
    }

    f1 := int64(1)
    f2 := int64(1)
    var f3 int64

    for i := 3; i <= n; i++ {
        f3 = f2 + int64(k)*f1
        f1, f2 = f2, f3
    }

    fmt.Println(f3)
}
