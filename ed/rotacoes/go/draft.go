package main
import (
    "fmt"
    "strings"
)

func rotacao(r int, slice []int) [] int{
    n := len(slice)
    if n == 0 {
            return slice
        }
        r = r % n
        return append(slice[n-r:], slice[:n-r]...)
    }

func main() {
    var t, r int
    fmt.Scan(&t, &r)
    slice := make([]int, t)
    for i := 0; i < t; i++{
        //fmt.Scan(&x)
        //slice = append(slice, x)
        fmt.Scan(&slice[i])
    }
    //---
    str := strings.Trim(fmt.Sprint(rotacao(r, slice)), "[]")
    fmt.Println("[ " + str + " ]")

}
