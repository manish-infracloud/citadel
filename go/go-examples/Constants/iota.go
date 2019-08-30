package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota // you can just declare a once and b and c only, it will understand
)

func main() {
	fmt.Println(a, b, c)
}
