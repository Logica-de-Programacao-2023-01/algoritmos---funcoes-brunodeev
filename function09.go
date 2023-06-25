package main

import (
	"fmt"
	"strings"
)

func main() {
	palavra := "Oh Shit!"
	fmt.Println(sliceWord(palavra))
}

func sliceWord(palavra string) ([]string, error) {
	palavras := []string{}
	if palavra == "" {
		return palavras, fmt.Errorf("ERRO ")
	}

	separador := strings.Split(palavra, " ")

	return separador, nil
}
