package main
import "fmt"

func div(vetor[]int) []int{
    mapa := make(map[int]int)
    var f int
    for i := range vetor {
        f = vetor[i] / 2
    }
    for _, resto := range vetor{
        mapa[f] = resto % 2
    }
    vet := make([]int, len(vetor))
    for i, resto := range vetor {
        vet[i] = mapa[resto]
    } 
    return vet
}
/*
func contrario(vetor[]int) []int{
    for  {
        //contrario
    }
    return vetor
}
*/
func main() {
    var num int
    fmt.Scan(&num)
    vetor := make([]int, num)
    for i := range num{
        fmt.Scan(&vetor[i])
    }
    fmt.Print(div(vetor))
    //contrario(vetor)
}
