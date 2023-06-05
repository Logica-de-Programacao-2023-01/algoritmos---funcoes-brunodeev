package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Eu", "sou", "Bruno"}
	result := concat(words)
	fmt.Println(result)
}

func concat(slice []string) string {
	return strings.Join(slice, "")
}
