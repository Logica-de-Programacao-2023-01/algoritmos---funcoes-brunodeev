package main

import "fmt"

func main() {
	slice := []int{2, 4, 6}

	fmt.Println("\n------------------------------------------------------------")
	fmt.Println("A média aritmética da slice é: ", media(slice))
	fmt.Println("------------------------------------------------------------")
}

func media(slice []int) float64 {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return float64(sum) / float64(len(slice))
}
