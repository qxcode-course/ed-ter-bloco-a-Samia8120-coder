package main
import "fmt"

func print(vals []int, swordIdx int) {
    fmt.Print("[ ")

    for i, v := range vals {
        if i > 0 {
            fmt.Print(" ")
        }
        if i == swordIdx {
            if v > 0 {
                fmt.Printf("%d>", v)
            } else {
                fmt.Printf("<%d", v)
            }
        } else {
            fmt.Printf("%d", v)
        }
    }
    fmt.Println(" ]")
}

func main() {
    var N, E, F int
    fmt.Scan(&N, &E, &F)

    vals := make([]int, N)
    for i := 1; i <= N; i++ {
        sinal := F
        if i % 2 == 0 {
            sinal = -F
        }
        vals[i-1] = i * sinal
    }

    swordIdx := E - 1
    print(vals, swordIdx)

    for len(vals) > 1 {
        n := len(vals)
        atual := vals[swordIdx]

        if atual > 0 {
            target := (swordIdx + 1) % n
            nextDead := (target + 1) % n
            vals = append(vals[:target], vals[target+1:]...)

            if nextDead > target {
                swordIdx = nextDead - 1
            } else {
                swordIdx = nextDead
            }
        } else {
            target := (swordIdx - 1 + n) % n
            prevDead := (target - 1 + n) % n
            vals = append(vals[:target], vals[target+1:]...)

            if prevDead > target {
                swordIdx = prevDead - 1
            } else {
                swordIdx = prevDead
            }
        }
        print(vals, swordIdx)
    }
}

