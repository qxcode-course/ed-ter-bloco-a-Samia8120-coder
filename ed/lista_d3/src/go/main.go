package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	list.size = 0
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
	l.size++
}

func (l *LList) String() string {
	if l.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	node := l.root.next
	for node != l.root {
		if node != l.root.next {
			sb.WriteString(", ")
		}
		sb.WriteString(strconv.Itoa(node.Value))
		node = node.next
	}
	sb.WriteString("]")
	return sb.String()
}

func (l *LList) Size() int {
	return l.size
}

func equals(a, b *LList) bool {
	if a.Size() != b.Size() {
		return false
	}
	na := a.root.next
	nb := b.root.next
	for na != a.root && nb != b.root {
		if na.Value != nb.Value {
			return false
		}
		na = na.next
		nb = nb.next
	}
	return true
}

func addsorted(l *LList, value int) {
	if l.Size() == 0 {
		l.PushBack(value)
		return
	}
	cur := l.root.next
	for cur != l.root && cur.Value < value {
		cur = cur.next
	}
	l.insertBefore(cur, value)
}

func reverse(l *LList) {
	if l.Size() <= 1 {
		return
	}
	cur := l.root.next
	for cur != l.root {
		cur.next, cur.prev = cur.prev, cur.next
		cur = cur.prev
	}
	l.root.next, l.root.prev = l.root.prev, l.root.next
}

func merge(a, b *LList) *LList {
	result := NewLList()
	na := a.root.next
	nb := b.root.next

	for na != a.root && nb != b.root {
		if na.Value <= nb.Value {
			result.PushBack(na.Value)
			na = na.next
		} else {
			result.PushBack(nb.Value)
			nb = nb.next
		}
	}
	for na != a.root {
		result.PushBack(na.Value)
		na = na.next
	}
	for nb != b.root {
		result.PushBack(nb.Value)
		nb = nb.next
	}
	return result
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
