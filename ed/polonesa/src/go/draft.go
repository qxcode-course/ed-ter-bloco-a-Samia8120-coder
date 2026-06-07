package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func precedencia(op string) int{
    switch op {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    case "^":
        return 3
    default:
        return 0
    }
}

func associativo(op string) bool {
    return op != "^"
}

func notacao(tokens []string) []string {
    var output []string
    var stack []string

    for _, token := range tokens {
        if precedencia(token) == 0 {
            output = append(output, token)
            continue
        }

        for len(stack) > 0 {
            cima := stack[len(stack)-1]
            cimaPrec := precedencia(cima)
            agrPrec := precedencia(token)

            if (associativo(token) && agrPrec <= cimaPrec) || (!associativo(token) && agrPrec < cimaPrec) {
                output = append(output, cima)
                stack = stack[:len(stack)-1]
            } else {
                break
            }
        }
        stack = append(stack, token)
    }

    for len(stack) > 0 {
            output = append(output, stack[len(stack)-1])
            stack = stack[:len(stack)-1]
    }
    return output
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    if !scanner.Scan() {
        return
    }

    linha := scanner.Text()
    tokens := strings.Fields(linha)
    
    if len(tokens) == 0 {
        return
    }

    npr := notacao(tokens)
    fmt.Println(strings.Join(npr, " "))
}