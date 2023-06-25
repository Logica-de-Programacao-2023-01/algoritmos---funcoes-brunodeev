package main

import (
	"fmt"
)

func main() {
	slice1 := []int{10, 5, 8, 20, 7}
	slice2 := []int{5, 6, 20, 3}
	fmt.Print(contemnosdoisslices(slice1, slice2))
}
func contemnosdoisslices(slice1 []int, slice2 []int) ([]int, error) {
	slice3 := []int{}
	if len(slice1) == 0 || len(slice2) == 0 {
		return slice3, fmt.Errorf("Erro")
	}
	for _, s1 := range slice1 {
		for _, s2 := range slice2 {
			if s1 == s2 {
				slice3 = append(slice3, s1)
			}
		}
	}
	return slice3, nil
}
