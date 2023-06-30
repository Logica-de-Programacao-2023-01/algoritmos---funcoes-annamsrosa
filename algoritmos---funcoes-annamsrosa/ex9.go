package main

import (
	"errors"
	"fmt"
	"strings"
)

func palavras(input string) ([]string, error) {
	if input == "" {
		return nil, errors.New("String vazia")
	}

	words := strings.Fields(input)
	return words, nil
}

func main() {
	input := "Olá, boa noite à você, leitor "
	words, err := palavras(input)
	if err != nil {
		panic(err)
	}

	for _, palavra := range words {
		fmt.Println(palavra)
	}
}
