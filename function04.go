package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 5, 8, 20, 7}
	fmt.Print(second(slice))
}

func second(slice []int) int {
	var biggest int
	var secondBiggest int
	for i := 0; i < len(slice); i++ {
		if slice[i] > biggest {
			secondBiggest = biggest
			biggest = slice[i]
		}
	}
	return secondBiggest
}
