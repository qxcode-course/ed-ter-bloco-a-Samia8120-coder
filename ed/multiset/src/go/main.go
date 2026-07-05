package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data []int
	size int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	if capacity < 0 {
		capacity = 0
	}
	return &MultiSet {
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	if ms.capacity == 0 {
		ms.capacity = 1
	} else {
		ms.capacity *= 2
	}
	newData := make([]int, ms.capacity)
	copy(newData, ms.data[:ms.size])
	ms.data = newData
}

func (ms *MultiSet) search(value int) (bool, int) {
	low, high := 0, ms.size-1
	last := -1

	for low <= high {
		mid := (low + high) / 2
		if ms.data[mid] == value {
			last = mid
			low = mid + 1
		} else if ms.data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if last != -1 {
		return true, last
	}
	return false, low
}

func (ms *MultiSet) insertAt(index int , value int) {
	if ms.size == ms.capacity {
		ms.expand()
	}
	copy(ms.data[index+1:], ms.data[index:ms.size])
	ms.data[index] = value
	ms.size++
}

func (ms *MultiSet) Insert(value int) {
	found, idx := ms.search(value)
	if found {
		idx++
	}
	ms.insertAt(idx, value)
}

func (ms *MultiSet) eraseAt(index int) {
	if index < 0 || index >= ms.size {
		return
	}
	copy(ms.data[index:], ms.data[index+1:ms.size])
	ms.size--
}

func (ms *MultiSet) Erase(value int) error {
	found, idx := ms.search(value)
	if !found {
		return errors.New("value npyt found")
	}
	ms.eraseAt(idx)
	return nil
}

func (ms *MultiSet) Contains(value int) bool {
	found, _ := ms.search(value)
	return found
}

func (ms *MultiSet) Count(value int) int {
	found, idx := ms.search(value)
	if !found {
		return 0
	}
	count := 0
	for i := idx; i >= 0 && ms.data[i] == value; i-- {
		count++
	}
	return count
}

func (ms* MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}
	count := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			count++
		}
	}
	return count
}

func (ms *MultiSet) Clear() {
	ms.size = 0
}

func (ms *MultiSet) String() string {
	if ms.size == 0 {
		return "[]"
	}
	return "[" + Join(ms.data[:ms.size], ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	var ms *MultiSet

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			if len(args) < 2 {
				fmt.Println("fail: init requer capacidade")
				continue
			}
			cap, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(cap)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			if len(args) < 2 {
				fmt.Println("fail: erase requer valor")
				continue
			}
			value, _ := strconv.Atoi(args[1])
			if err := ms.Erase(value); err != nil {
				fmt.Println("value not found")
			}
		case "contains":
			if len(args) < 2 {
				fmt.Println("fail. contains requer valor")
				continue
			}
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			if ms == nil {
				fmt.Println("0")
				continue
			}
			if len(args) < 2 {
				fmt.Println("fail: count requer valor")
				continue
			}
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			if ms == nil {
				fmt.Println("0")
			} else {
				fmt.Println(ms.Unique())
			}
		case "clear":
			if ms != nil {
				ms.Clear()
			}
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
