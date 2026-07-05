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
	next *Node
	prev *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	ll := &LList{}
	root := &Node{}
	root.next = root
	root.prev = root
	ll.root = root
	ll.size = 0
	return ll
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) Clear() {
	root := ll.root
	root.next = root
	root.prev = root
	ll.size = 0
}

func (ll *LList) PushFront(value int) {
	newNode := &Node{Value: value}
	root := ll.root
	first := root.next
	root.next = newNode
	newNode.prev = root
	newNode.next = first
	first.prev = newNode
	ll.size++
}

func (ll *LList) PushBack(value int) {
	newNode := &Node{Value: value}
	root := ll.root
	last := root.prev
	last.next = newNode
	newNode.prev = last
	newNode.next = root
	root.prev = newNode
	ll.size++
}

func (ll *LList) PopFront() {
	if ll.size == 0 {
		return
	}
	root := ll.root
	first := root.next
	second := first.next
	root.next = second
	second.prev = root
	first.next = nil
	first.prev = nil
	ll.size--
}

func (ll *LList) PopBack() {
	if ll.size == 0 {
		return
	}
	root := ll.root
	last := root.prev
	prev := last.prev
	prev.next = root
	root.prev = prev
	last.next = nil
	last.prev = nil
	ll.size--
}

func (ll *LList) String() string {
	if ll.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	
	node := ll.root.next
	for node != ll.root {
		if node != ll.root.next {
			sb.WriteString(", ")
		}
		sb.WriteString(strconv.Itoa(node.Value))
		node = node.next
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
