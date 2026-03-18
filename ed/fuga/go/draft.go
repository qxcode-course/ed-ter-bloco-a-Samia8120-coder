package main

import "fmt"

func caminho(h int, p int, f int, d int) {
	if d == 1 {
		for f <= 16 {
			f++
			if f == 16 {
				f = 0
			}
			if f == p {
				fmt.Print("N\n")
				break
			} else if f == h {
				fmt.Print("S\n")
				break
			}
		}
	} else if d == -1 {
		for f < 16 {
			f--
			if f == -1 {
				f = 15
			}
			if f == p {
				fmt.Print("N\n")
				break
			} else if f == h {
				fmt.Print("S\n")
				break
			}
		}
	}
}
func main() {
	var H, P, F, D int
	fmt.Scan(&H, &P, &F, &D)

	caminho(H, P, F, D)
}
