package main

import "fmt"

func main() {
	slice := []int{0, 5, 8, 20, 7}
	num := 5
	fmt.Print(sameValue(slice, num))
}

func sameValue(slice []int, num int) int {
	value := -1
	for i := 0; i < len(slice); i++ {
		if slice[i] == num {
			value = i
		}
	}
	return value
}
