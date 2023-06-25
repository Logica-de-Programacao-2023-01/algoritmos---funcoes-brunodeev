package main

import "fmt"

func main() {
	slice1 := []int{10, 5, 8, 20, 7}
	n := 5
	fmt.Print(contemnosdoisslices(slice1, n))
}
func contemnosdoisslices(slice1 []int, n int) bool {

	for _, sn := range slice1 {
		if sn == n {
			return true
		}
	}
	return false
}
