package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "macacos"
	fmt.Print(vogais(string1))
}
func vogais(string1 string) string {

	stringfinal := strings.NewReplacer("a", "*", "e", "*", "i", "*", "o", "*", "u", "*")
	string1 = stringfinal.Replace(string1)
	return string1
}
