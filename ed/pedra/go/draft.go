package main
import "fmt"

type Jogada struct {
    peda1, peda2 int
}

func pontuacao(j Jogada) int {
    if j.peda1 > j.peda2 {
        return j.peda1 - j.peda2
    }
    return j.peda2 - j.peda1
}

func main() {
    qtd := 0
    fmt.Scan(&qtd)
    jogadas := make([]Jogada, qtd)

    for i := range qtd {
        fmt.Scan(&jogadas[i].peda1, &jogadas[i].peda2)
    }

    melhor := -1
    for i, jog := range jogadas {
        if jog.peda1 < 10 || jog.peda2 < 10 {
            continue
        }
        if melhor == -1 || (pontuacao(jog) < pontuacao(jogadas[melhor])) {
            melhor = i
        }
    }

    if melhor == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(melhor)
    }
}