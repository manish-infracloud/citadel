package main

import "fmt"

var i = 42

func main() {
	var j float32
	j = float32(i)
	fmt.Printf("%v, %T", j, j)
}
