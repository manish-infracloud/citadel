package main

import "fmt"

// elements in your slice start from 1 but the index start with 0
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice from 4th element to the end
	d := a[:6]  //slice first 6 elements
	e := a[3:6] //slice the 4th, 5th, and 6h elements in other words it starts from index 3 upto index 6 but exclusive 6.
	a[5] = 42   // it will reflect in each varible we have used
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println("Length is:", len(a))
	f := a[:len(a)-1]
	fmt.Println(f)
	fmt.Println("Length is:", len(a))

}
