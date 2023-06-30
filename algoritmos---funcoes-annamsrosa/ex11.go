package main

import (
	"fmt"
)

func main() {
	n := 23
	fmt.Print(primo(n))
}
func primo(n int) (bool, error) {
	primo := true
	if n < 0 {
		return false, fmt.Errorf("Erro")
	}

	if n < 2 {
		primo = false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			primo = false
		}
	}
	return primo, nil

}
