// passa 2 vet,
package main

import (
	"fmt"
)
func fila(idenEntraram[]int, idenSairam[]int){
	var tam int
	removiveis := make([]int, 0)
	//ficar := make([]int, 0)
	for i := range idenEntraram{
			if idenEntraram[i] == idenSairam[i]{
				removiveis[tam] = idenEntraram[i]
				tam++
		}
	}
	for i := 0; i < tam; i++{
		if removiveis[i] != idenEntraram[i]{
			fmt.Print(idenEntraram)
		}
	}

}
func main() {
	var N int
	fmt.Scan(&N)
	idenEntraram := make([]int, N)
	for i := 0; i < N; i++{
		fmt.Scan(idenEntraram[i])
	}
	var M int
	fmt.Scan(&M)
	idenSairam := make([]int, M)
	for i := 0; i < M; i++{
		fmt.Scan(idenSairam[i])
	}
	fila(idenEntraram, idenSairam)
}
