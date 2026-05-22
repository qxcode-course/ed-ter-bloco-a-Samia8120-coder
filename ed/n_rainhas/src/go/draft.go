package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    cols := make([]bool, n)
    diag1 := make([]bool, 2 * n-1)
    diag2 := make([]bool, 2 * n-1)
    solucoes := 0

    var backtrack func(int)
    backtrack = func(linha int) {
        if linha == n {
            solucoes++
            return
        }

        for coluna := 0; coluna < n; coluna++ {
            d1 := linha - coluna + n - 1
            d2 := linha + coluna

            if cols[coluna] || diag1[d1] || diag2[d2] {
                continue
            }

            cols[coluna] = true
            diag1[d1] = true
            diag2[d2] = true
            backtrack(linha + 1)
            cols[coluna] = false
            diag1[d1] = false
            diag2[d2] = false
        }
    }
    backtrack(0)
    fmt.Println(solucoes)
}
