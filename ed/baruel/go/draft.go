package main
import "fmt"
func main() {
	var qtd_album, qtd_possui int
	fmt.Scan(&qtd_album, &qtd_possui)
	figuras := make([]int, qtd_possui)
	for i := range figuras{
		fmt.Scan(&figuras[i])
	}

	unicos := make(map[int]bool)
	repetidos := make([]int, 0, qtd_possui)
	for _, figura := range figuras{
		if unicos[figura] {
			repetidos = append(repetidos, figura)
		} else {
			unicos[figura] = true
		}
	}
	saida := fmt.Sprintf("%v", repetidos)
	if saida == "[]" {
		fmt.Println("N")
	} else {
	fmt.Println(saida[1: len(saida)-1])
	}
	faltantes := make([]int, 0, qtd_album)
	for i := 1; i <= qtd_album; i++{
		if !unicos[i]{
			faltantes = append(faltantes, i)
		}
	}
	saida = fmt.Sprintf("%v", faltantes)
	if saida == "[]" {
		fmt.Println("N")
	} else {
	fmt.Println(saida[1: len(saida)-1])
	}
}
