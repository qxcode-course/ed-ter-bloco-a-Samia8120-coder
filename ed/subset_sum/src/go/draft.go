package main
import "fmt"

func main() {
    var n, k int
    fmt.Scan(&n, &k)

    nums := make([]int, n)
    for i := 0; i < n; i++{
        fmt.Scan(&nums[i])
    }

    dp := make([]bool, k+1)
    dp[0] = true

    for _, val := range nums {
        for s := k; s >= val; s-- {
            if dp[s-val] {
                dp[s] = true
            }
        } 
    }

    if dp[k] {
        fmt.Print("true\n")
    } else {
        fmt.Print("false\n")
    }
}
