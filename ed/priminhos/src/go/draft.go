package main
import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    var input string
    fmt.Scan(&input)

    n, err := strconv.Atoi(input)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Erro")
    }

    primos := make([]int, 0, n)
    for num := 2; len(primos) < n; num++ {
        if ePrimo(num) {
            primos = append(primos, num)
        }
    }

    var out strings.Builder
    out.WriteString("[")

    for i, p := range primos {
        if i > 0 {
            out.WriteString(", ")
        }
        fmt.Fprint(&out, p)
    }
    
    out.WriteString("]")
    fmt.Println(out.String())
}

func ePrimo(num int) bool {
    if num < 2 {
        return false
    }
    if num == 2 {
        return true
    }
    if num % 2 == 0 {
        return false 
    }

    for i := 3; i * i <= num; i += 2 {
        if num % i == 0 {
            return false
        }
    }
    return true    
}