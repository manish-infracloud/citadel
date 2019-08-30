package main

import (
	"fmt"
)

func main() {
	var n = 1 == 1
	var m = 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)
}
