package main
import "fmt"

func print(valores []int, espadaIdx int) {
    fmt.Print("[ ")

    for i, v := range valores {
        if i > 0 {
            fmt.Print(" ")
        }
        if i == espadaIdx {
            if v > 0 {
                fmt.Printf("%d>", v)
            } else {
                fmt.Printf("<%d", v)
            }
        } else {
            fmt.Printf("%d", v)
        }
    }
    fmt.Println(" ]")
}

func main() {
    var N, E, F int
    fmt.Scan(&N, &E, &F)

    valores := make([]int, N)
    for i := 1; i <= N; i++ {
        sinal := F
        if i % 2 == 0 {
            sinal = -F
        }
        valores[i-1] = i * sinal
    }

    espadaIdx := E - 1
    print(valores, espadaIdx)

    for len(valores) > 1 {
        n := len(valores)
        atual := valores[espadaIdx]

        if atual > 0 {
            alvo := (espadaIdx + 1) % n
            proxDead := (alvo + 1) % n
            valores = append(valores[:alvo], valores[alvo+1:]...)

            if proxDead > alvo {
                espadaIdx = proxDead - 1
            } else {
                espadaIdx = proxDead
            }
        } else {
            target := (espadaIdx - 1 + n) % n
            prevDead := (target - 1 + n) % n
            valores = append(valores[:target], valores[target+1:]...)

            if prevDead > target {
                espadaIdx = prevDead - 1
            } else {
                espadaIdx = prevDead
            }
        }
        print(valores, espadaIdx)
    }
}

