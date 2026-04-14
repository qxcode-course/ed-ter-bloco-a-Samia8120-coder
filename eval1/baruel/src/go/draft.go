package main
import "fmt"

func fig(totalfigalbum int, totalfigbaruel int){
    tam := 0
    album := make([]int, totalfigalbum)
    baruel := make([]int, totalfigbaruel)
    iguais := make([]int, tam)
    var figanterior int
    var fig int
    //as q ja tem ->figs repetidas -> figs faltando
    for i := 0; i < totalfigbaruel; i++{
        fmt.Scan(baruel[i])
    }
    for i := 0; i < totalfigbaruel; i++{
        figanterior = baruel[i-1]
        fig = baruel[i]
        if figanterior == fig{
            iguais[tam] = fig
            tam++
        }
    fmt.Print(iguais)
    }
    for 
    
}

func main() {
    var totalfigalbum, totalfigbaruel int
    fmt.Scan(&totalfigalbum)
    fmt.Scan(&totalfigbaruel)
    fig(totalfigalbum, totalfigbaruel)
}
