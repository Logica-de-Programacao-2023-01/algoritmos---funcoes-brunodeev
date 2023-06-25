package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []string{"abc", "rwp", "def", "por", "ber", "erj"}
	fmt.Print(crescente(slice1))
}
func crescente(slice1 []string) (string, error) {
	sort.Strings(slice1)
	if len(slice1) == 0 {
		return "", fmt.Errorf("Erro")
	}
	string1 := ""
	for _, s1 := range slice1 {
		string1 += s1 + " "
	}

	return string1, nil
}
