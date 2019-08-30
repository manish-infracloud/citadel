package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = string(i) //i simply typecasted it will print the unicode value for number 42 which is * we need to convert it using conversion package for string.
	fmt.Printf("%v, %T\n", j, j)
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
