package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    gas := make([]int, n)
    valor := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&gas[i], &valor[i])
    }

    totalGas := 0
    totalValor := 0
    tanque := 0
    comeco := 0

    for i := 0; i < n; i++ {
        totalGas += gas[i]
        totalValor += valor[i]
        tanque += gas[i] - valor[i]

        if tanque < 0 {
            comeco = i + 1
            tanque = 0
        }
    }

    if totalGas < totalValor {
        fmt.Print(-1)
    } else {
        fmt.Println(comeco)
    }
}