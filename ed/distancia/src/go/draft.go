package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    s := scanner.Text()
    n := len(s)
    scanner.Scan()

    L, _ := strconv.Atoi(scanner.Text())
    ans := []byte(s)

    for i := 0; i < L && i < n; i++ {
        if ans[i] == '.' {
            ans[i] = '0'
        }
    }

    padrao := s[:L]
    for i := 0; i < n; i++ {
        if s[i] != '.' {
            ans[i] = s[i]
        } else {
            ans[i] = padrao[i%L]
        }
    }
    fmt.Println(string(ans))
}
