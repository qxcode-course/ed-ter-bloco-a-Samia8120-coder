package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func DetonacaoMaxima(bombs [][]int) int{
    n := len(bombs)
    if n == 0 {
        return 0
    }

    adj := make([][]int, n)
    for i := 0; i < n; i++ {
        xi, yi, ri := bombs[i][0], bombs[i][1], bombs[i][2]

        for j := 0; j < n; j++ {
            if i == j {
                continue
            }
            
            xj, yj := bombs[j][0], bombs[j][1]
            dx := xi - xj
            dy := yi - yj
            dist2 := dx*dx + dy*dy
            if dist2 <= ri*ri {
                adj[i] = append(adj[i], j)
            }
        }
    }
    maxCount := 0

    for start := 0; start < n; start++ {
        visited := make([]bool, n)
        stack := []int{start}
        visited[start] = true
        count := 0

        for len(stack) > 0 {
            u := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            count++

            for _, v := range adj[u] {
                if !visited[v] {
                    visited[v] = true
                    stack = append(stack, v)
                }
            }
        }
        if count > maxCount {
            maxCount = count
        }
    }
    return maxCount
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    tokens := strings.Fields(line)
    for scanner.Scan(){
        line = scanner.Text()
        more := strings.Fields(line)
        tokens = append(tokens, more...)
    }

    if len(tokens) == 0 {
        return
    }
    
    nums := make([]int, len(tokens))
    for i, s := range tokens {
        val, _ := strconv.Atoi(s)
        nums[i] = val
    }

    if len(nums) < 2 {
        return
    }
    n := nums[0]
    idx := 2
    bombs := make([][]int, n)

    for i := 0; i < n; i++ {
        if idx+3 > len(nums) {
            break
        }
        bombs[i] = []int{nums[idx], nums[idx+1], nums[idx+2]}
        idx += 3
    }
    fmt.Println(DetonacaoMaxima(bombs))
}