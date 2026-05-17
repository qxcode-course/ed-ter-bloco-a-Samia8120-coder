package main
import (
    "fmt"
    "sort"
)

func main() {
    var s string
    fmt.Scan(&s)

    perms := gerarPermutacoes(s)
    sort.Strings(perms)

    for _, p := range perms {
        fmt.Println(p)
    }
}

func gerarPermutacoes(s string) []string {
    n := len(s)
    if n == 0 {
        return []string{}
    }

    usado := make([]bool, n)
    agora := make([]byte, n)
    resultado := []string{}

    backtrack(s, usado, agora, 0, &resultado)
    return resultado
}

func backtrack(s string, usado []bool, agora []byte, pos int, resultado *[]string) {
    if pos == len(s) {
        *resultado = append(*resultado, string(agora))
        return
    }
    for i := 0; i < len(s); i++ {
        if !usado[i] {
            usado[i] = true
            agora[pos] = s[i]
            backtrack(s, usado, agora, pos+1, resultado)
            usado[i] = false
        }
    }    
}