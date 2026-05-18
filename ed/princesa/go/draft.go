package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    linha := scanner.Text()
    partes := strings.Fields(linha)

    if len(partes) < 2 {
        return
    }

    n, _ := strconv.Atoi(partes[0])
    e, _ := strconv.Atoi(partes[1])

    josephusRemocao(n, e)
}

func josephusRemocao(n, e int) {
    vivos := make([]int, n)
    for i := 0; i < n; i++ {
        vivos[i] = i+1
    }

    pos := e - 1
    imprimir := func() {
        fmt.Print("[ ")
        for i, val := range vivos{
            if i == pos {
                fmt.Printf("%d> ", val)
            } else {
                fmt.Printf("%d ", val)
            }
        }
        fmt.Println("]")
    }

    for len(vivos) > 1 {
        imprimir()

        proximo := (pos + 1) % len(vivos)
        vivos = append(vivos[:proximo], vivos[proximo+1:]...)

        if len(vivos) > 0 {
            pos = proximo % len(vivos)
        }
    }
    imprimir()
}
