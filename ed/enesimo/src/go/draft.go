package main
import "fmt"

func ehPrimo(num, divisor int) bool {
    if num <= 1 {
        return false
    }
    if divisor * divisor > num {
        return true
    }
    if num % divisor == 0 {
        return false
    }

    return ehPrimo(num, divisor+1)
}

func enePrimo(n, atual, cont int) int {
    if cont == n {
        return atual
    }

    proximo := atual + 1
    if ehPrimo(proximo, 2) {
        return enePrimo(n, proximo, cont+1)
    }
    return enePrimo(n, proximo, cont)
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(enePrimo(n, 2, 1))
}
