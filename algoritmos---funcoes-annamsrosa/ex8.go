package main

import (
	"errors"
	"fmt"
)

func pares(ints []int) ([]int, error) {
	var p []int

	if len(ints) == 0 {
		return nil, errors.New("O slice esta vazio!")
	}
	for _, num := range ints {
		if num%2 == 0 {
			p = append(p, num)
		}
	}
	return p, nil
}

func main() {
	var numeros = []int{10, 5, 20, 55, 25, 100, 95, 35, 40, 15, 65, 70}

	pares, err := pares(numeros)

	if err != nil {
		fmt.Print("Erro:", err)
		return
	}
	fmt.Println(numeros)
	fmt.Println("Os números pares são:", pares)
}
