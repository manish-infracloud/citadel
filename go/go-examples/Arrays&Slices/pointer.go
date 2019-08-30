package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	fmt.Println(a)
	b := a   // b will copy all the array elements in a
	b[1] = 5 // and when we change the value it will be reflected for b only
	fmt.Println(a)
	fmt.Println(b)
	//using pointer
	fmt.Println("------------")
	c := [...]int{1, 2, 3}
	fmt.Println(c)
	d := &c  // d will point to all the array elements in c
	d[1] = 5 // and when we change the value it will be reflected for both b & c
	fmt.Println(c)
	fmt.Println(d)
	//using slices, empty []
	fmt.Println("------------")
	e := []int{1, 2, 3}
	fmt.Println(e)
	f := e   // f will point to all the array elements in e. pointers are not supported by slices. refering is thier default behavior
	f[1] = 6 // and when we change the value it will be reflected for both e & f
	fmt.Println(e)
	fmt.Println(f)
}
