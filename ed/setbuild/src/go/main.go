package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Set struct {
	data []int
}

func NewSet(capacity int) *Set {
	return &Set{data: make([]int, 0, capacity)}
}

func (s *Set) Insert(values ...int) {
	for _, val := range values {
		idx := sort.Search(len(s.data), func(i int) bool { return s.data[i] >= val })
		if idx < len(s.data) && s.data[idx] == val {
			continue
		}
		s.data = append(s.data, 0)
		copy(s.data[idx+1:], s.data[idx:])
		s.data[idx] = val
	}
}

func (s *Set) Contains(val int) bool {
	idx := sort.Search(len(s.data), func(i int) bool { return s.data[i] >= val })
	return idx < len(s.data) && s.data[idx] == val
}

func (s *Set) Erase(val int) bool {
	idx := sort.Search(len(s.data), func(i int) bool { return s.data[i] >= val })
	if idx < len(s.data) && s.data[idx] == val {
		s.data = append(s.data[:idx], s.data[idx+1:]...)
		return true
	}
	return false
}

func (s *Set) Clear() {
	s.data = make([]int, 0, cap(s.data))
}

func (s *Set) Show() {
	if len(s.data) == 0 {
		fmt.Println("[]")
		return
	}
	var b strings.Builder
	b.WriteByte('[')
	for i, v := range s.data {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprint(&b, v)
	}
	b.WriteByte(']')
	fmt.Println(b.String())
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	var v *Set
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			capacity, _ := strconv.Atoi(parts[1])
			v = NewSet(capacity)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			v.Show()
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
			v.Clear()
		default:
			fmt.Println("value not found")
		}
	}
}
