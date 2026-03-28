package main
import "fmt"

type Soldado struct{
    vivo, espada int
}

func matar(){
    //mata ai
}
func cacar_vivo(){
    //caça vivo e da a espada
}
func main() {
    var tamanho, posicao int
    fmt.Scan(&tamanho, &posicao)
    vet := make([]Soldado, tamanho)
    for i := range tamanho {
        fmt.Scan(vet[i])
    }

    ax := 0
    for ax > 1{
        //coisos
        for {
            
        }
    }

}
