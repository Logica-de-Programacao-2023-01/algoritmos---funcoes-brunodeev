package main

import (
	"fmt"
	"strings"
)

func main() {

	word := "computador"

	fmt.Println("\n------------------------------------------------------------")
	fmt.Printf("Essa string contém %d vogais", vogals(word))
	fmt.Println("\n------------------------------------------------------------")
}

func vogals(word string) int {
	theVogals := "aeiouAEIOU"
	sum := 0
	for _, letter := range word {
		if strings.ContainsRune(theVogals, letter) {
			sum++
		}
	}
	return sum
}
