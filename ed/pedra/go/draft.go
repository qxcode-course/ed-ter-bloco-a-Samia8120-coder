package main
import (
    "fmt"
)

type Jogada struct{
    pa, pb int
}

func pontuacao(j Jogada) int {
    if j.pa > j.pb{
        return j.pa - j.pb
    }
    return j.pb - j.pa
}

func main() {
    qtd := 0
    fmt.Scan(&qtd)
    jogadas := make([]Jogada, qtd)
    for i := range qtd{
        fmt.Scan(&jogadas[i].pa, &jogadas[i].pb)
    }
    ind_melhor := -1
    for i, jog := range jogadas{
        if jog.pa < 10 || jog.pb < 10{
            continue
        }
        if ind_melhor == -1 || (pontuacao(jog) < pontuacao(jogadas[ind_melhor])){
            ind_melhor = i
        }
    }
    if ind_melhor == -1 {
        fmt.Print("sem ganhador\n")
    } else {
        fmt.Print(ind_melhor, "\n")
    } 

}
