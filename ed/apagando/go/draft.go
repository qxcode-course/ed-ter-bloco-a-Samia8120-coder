package main
import "fmt"
func main() {
	var num_ini, num_deixa int
	fmt.Scan(&num_ini)
	fila := make([]int, num_ini)
	for i := range num_ini{
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&num_deixa)
	deixaram := make(map[int]bool)
	var ident_deixa int
	for i := 0; i < num_deixa; i++{
		fmt.Scan(&ident_deixa)
		deixaram[ident_deixa] = true
	}

	for i := 0; i < num_ini; i++{
		final := fila[i]
		if !(deixaram[final]){
			fmt.Printf("%d ", final)
		}
	}
	fmt.Println()
}
