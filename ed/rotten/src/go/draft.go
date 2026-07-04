package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func LaranjasPodres(grid [][]int) int{
    if len(grid) == 0 {
        return 0
    }
    m, n := len(grid), len(grid[0])

    queue := make([][2]int, 0)
    fresh := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 2 {
                queue = append(queue, [2]int{i, j})
            } else if grid[i][j] == 1 {
                fresh++
            }
        }
    }
    if fresh == 0 {
        return 0
    }

    minutes := 0
    dirs := [][2]int{{-1,0}, {1,0}, {0,-1}, {0,1}}
    for len(queue) > 0 && fresh > 0 {
        size := len(queue)
        for k := 0; k < size; k++ {
            pos := queue[0]
            queue = queue[1:]

            for _, d := range dirs {
                ni, nj := pos[0]+d[0], pos[1]+d[1]
                if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == 1 {
                    grid[ni][nj] = 2
                    fresh--
                    queue = append(queue, [2]int{ni, nj})
                }
            }
        }
        minutes++
    }
    if fresh > 0 {
        return -1
    }
    return minutes
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var tokens []string

    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        tokens = append(tokens, fields...)
    }
    if len(tokens) == 0 {
        return 
    }

    nums := make([]int, len(tokens))
    for i, t := range tokens {
        val, _ := strconv.Atoi(t)
        nums[i] = val
    }

    idx := 0
    m, n := nums[idx], nums[idx+1]
    idx += 2

    grid := make([][]int, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]int, n)
        for j := 0; j < n; j++ {
            grid[i][j] = nums[idx]
            idx++
        } 
    }
    fmt.Println(LaranjasPodres(grid))
}