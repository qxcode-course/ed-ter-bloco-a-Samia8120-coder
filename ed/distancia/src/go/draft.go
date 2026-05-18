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
    ans := make([]byte, n)
    copy(ans, s)

    lastPos := make([]int, 10)
    for i := 0; i < 10; i++ {
        lastPos[i] = -1000
    }

    var dfs func(int) bool
    dfs = func(idx int) bool{
        if idx == n {
            return true
        }
        if ans[idx] != '.' {
            d := int(ans[idx] - '0')
            if idx-lastPos[d] < L {
                return false
            }

            old := lastPos[d]
            lastPos[d] = idx
            if dfs(idx + 1){
                return true
            }

            lastPos[d] = old
            return false
        }

        for d := 0; d <= L; d++ {
            if idx-lastPos[d] < L {
                continue
            }

            ok := true
            for j := idx + 1; j < n && j-idx < L; j++ {
                if ans[j] != '.' && int(ans[j]-'0') == d {
                    ok = false
                    break
                }
            }

            if !ok {
                continue
            }
            ans[idx] = byte('0' + d)
            old := lastPos[d]
            lastPos[d] = idx

            if dfs(idx + 1) {
                return true
            }
            ans[idx] = '.'
            lastPos[d] = old
        }
        return false
    }
    dfs(0)
    fmt.Println(string(ans))
}
