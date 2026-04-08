package main
import "fmt"

func div(num int){
    if num == 0 {
        return
    }
    quo := num / 2
    resto := num % 2

    div(quo)
    fmt.Println(quo, resto)
}

func main() {
    var num int
    fmt.Scan(&num)
    div(num)
}
