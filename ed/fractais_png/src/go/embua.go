package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func embua(pen *Pen, coisa float64){
	if coisa < 10{
		return
	}
	pen.SetRGB(255, 150, 255)
	pen.SetLineWidth(coisa/30)
	pen.Walk(coisa)
	//pen.SavePNG("tree.png")
	//var dummy rune
	//fmt.Scanf("%c", &dummy)
	pen.Right(90)
	coisa *= 0.97
	embua(pen, coisa)
	pen.SetRGB(20, 20, 20)
	pen.Right(-90)
	pen.Walk(-coisa)
}

func main() {
	pen := NewPen(500, 500)   // cria um canvas de 500 de largura por 500 de altura
     // muda a cor do pincel para vermelho
	pen.SetHeading(0)
	pen.SetPosition(100, 100)
	
	coisa := 300.0
	embua(pen, coisa)
	
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
