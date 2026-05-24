package main

import (
    "bufio"
	"fmt"
    "os"
    "strconv"
)

var k int64
var memo map[int64]int64

func contador(n int64) int64 {
    if n <= k {
        return 1
    }
    if v, ok := memo[n]; ok {
        return v
    }

    resultado := contador((n+1)/2) + contador(n/2)
    memo[n] = resultado
    return resultado
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()

    n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
    scanner.Scan()
    k, _ = strconv.ParseInt(scanner.Text(), 10, 64)
    
    memo = make(map[int64]int64)
    fmt.Println(contador(n))
}
