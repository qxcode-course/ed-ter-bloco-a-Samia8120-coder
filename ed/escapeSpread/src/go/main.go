package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]int) int {
	const INF = 1_000_000_000_000
	const HIGH = 1_000_000_000

	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	fire := make([][]int, m)
	for i := range fire {
		fire[i] = make([]int, n)
		for j := range fire[i] {
			fire[i][j] = INF
		}
	}

	dirs := [][2]int{{-1,0}, {1,0}, {0,-1}, {0,1}}
	queue := make([][2]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				fire[i][j] = 0
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	head := 0
	for head < len(queue) {
		r, c := queue[head][0], queue[head][1]
		head++

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 ||nc >= n {
				continue
			}
			if grid[nr][nc] == 2 {
				continue
			}
			if fire[nr][nc] > fire[r][c]+1 {
				fire[nr][nc] = fire[r][c] + 1
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
	possible := func(t int) bool {
		if fire[0][0] <= t {
			return false
		}

		visited := make([][]bool, m)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		type node struct {
			r, c, time int
		}

		q := []node{{0, 0, t}}
		visited[0][0] = true
		head := 0

		for head < len(q) {
			cur := q[head]
			head++

			if cur.r == m-1 && cur.c == n-1 {
				return true
			}

			for _, d := range dirs {
				nr, nc := cur.r+d[0], cur.c+d[1]

				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					continue
				}
				if grid[nr][nc] == 2 {
					continue
				}
				if visited[nr][nc] {
					continue
				}

				nxt := cur.time + 1
				if nr == m-1 && nc == n-1 {
					if fire[nr][nc] < nxt {
						continue
					}
				} else {
					if fire[nr][nc] <= nxt {
						continue
					}
				}

				visited[nr][nc] = true
				q = append(q, node{nr, nc, nxt})
			}
		}

		return false
	}

	if !possible(0) {
		return -1
	}

	high := HIGH
	if fire[0][0] < INF {
		high = fire[0][0] - 1
	}
	if possible(high) {
		return high
	}

	low := 0
	for low < high {
		mid := (low + high + 1) / 2
		if possible(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	m, _ := strconv.Atoi(parts[0])
	n, _ := strconv.Atoi(parts[1])

	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		tokens := strings.Fields(scanner.Text())
		row := make([]int, n)
		for j := 0; j < n && j < len(tokens); j++ {
			row[j], _ = strconv.Atoi(tokens[j])
		}
		grid[i] = row
	}
	fmt.Println(numIslands(grid))
	/*scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
	*/
}
