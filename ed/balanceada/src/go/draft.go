package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        s := scanner.Text()
        
        if Balanceado(s) {
            fmt.Println("balanceado")
        } else {
            fmt.Println("nao balanceado")
        }
    }
}

func Balanceado(s string) bool{
    stack := []rune{}
    for _, cha := range s {
        switch cha {
        case '(', '[':
            stack = append(stack, cha)
        case ')':
            if len(stack) == 0 || stack[len(stack)-1] != '(' {
                return false
            }
            stack = stack[:len(stack)-1]
        case ']':
            if len(stack) == 0 || stack[len(stack)-1] != '[' {
                return false
            }
            stack = stack[:len(stack)-1]
        default:
            return false
        }
    }
    return len(stack) == 0
}