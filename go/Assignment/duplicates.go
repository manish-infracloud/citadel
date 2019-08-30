package main

import (
	"fmt"
)

func printslice(slice []string) {
	fmt.Println("slice = ", slice)
}

func dupCount(list []string) map[string]int {

	duplicateFrequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicateFrequency map

		_, exist := duplicateFrequency[item]

		if exist {
			duplicateFrequency[item]++ // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[item] = 1 // else start counting from 1
		}
	}
	return duplicateFrequency
}

func main() {
	duplicate := []string{"Hello", "World", "GoodBye", "World", "We", "Love", "Love", "You"}

	printslice(duplicate)

	dupMap := dupCount(duplicate)

	//fmt.Println(dupMap)

	for k, v := range dupMap {
		fmt.Printf("Item : %s , Count : %d\n", k, v)
	}

}
