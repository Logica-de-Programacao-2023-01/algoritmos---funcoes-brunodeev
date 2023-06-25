package main

import (
	"fmt"
)

func main() {
	numeros := []int{2312, 32, 77, 10, 3456, 11, 1231, 3334}
	fmt.Print(pares(numeros))
}

func pares(numeros []int) ([]int, error) {
	pares := []int{}

	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			pares = append(pares, numeros[i])
		}
	}
	if len(pares) == 0 {
		return pares, fmt.Errorf("Erro")
	}

	return pares, nil
}
