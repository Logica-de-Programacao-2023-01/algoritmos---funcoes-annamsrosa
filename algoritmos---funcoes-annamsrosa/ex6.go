package main

import (
	"errors"
	"fmt"
)

func conca(s []string) (string, error) {
	var soma string
	if len(s) == 0 {
		return "", errors.New("Slice vazio!")
	}
	soma = s[0]
	for _, num := range s[1:] {
		soma += "," + num
	}
	return soma, nil
}

func main() {
	var ss = []string{"Água", "Terra", "Ar", "Fogo"}

	concate, err := conca(ss)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	fmt.Print("A concatenação do slice ", ss, " é: ", concate)
}
