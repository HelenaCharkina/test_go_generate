//go:generate ./generator int
package main

import (
	"fmt"
)

func main() {
	var a, b = 1, 2
	stack := NewintStack()
	stack.Add(a)
	stack.Add(b)
	fmt.Println(stack.Pop())
	return
}
