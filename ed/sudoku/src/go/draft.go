package main
import(
    "fmt"
    "math"
)
func main() {
    var n int
    fmt.Scan(&n)

    grid := make([][]rune, n)
    for i := 0; i < n; i++ {
        var line string
        fmt.Scan(&line)
        grid[i] = []rune(line)
    }

    solve(grid, 0, n)
    printGrid(grid, n)
}

func solve(grid [][]rune, index int, n int) bool{
    if index == n*n {
        return true
    }
    lin := index / n
    col := index % n

    if grid[lin][col] != '.'{
        return solve(grid, index+1, n)
    }

    subTam := int(math.Sqrt(float64(n)))

    for num := 1; num <= n; num++ {
        digit := rune('0' + num)
        if isValid(grid, lin, col, digit, n, subTam) {
            grid[lin][col] = digit
            if solve(grid, index+1, n) {
                return true
            }
            grid[lin][col] = '.'
        }
    }
    return false
}

func isValid(grid [][]rune, lin, col int, digit rune, n int, subTam int) bool {
    for c := 0; c < n; c++ {
        if grid[lin][c] == digit {
            return false
        }
    }
    
    for r := 0; r < n; r++ {
        if grid[r][col] == digit {
            return false
        }
    }

    startLin := (lin / subTam) * subTam
    startCol := (col / subTam) * subTam

    for l := startLin; l < startLin + subTam; l++ {
        for c := startCol; c < startCol + subTam; c++ {
            if grid[l][c] == digit {
                return false
            }
        }
    }
    return true
}

func printGrid(grid [][]rune, n int) {
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Printf("%c", grid[i][j])
        }
        fmt.Println()
    }
}