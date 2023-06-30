package main

import "fmt"

func funcSlice(x []int, y func(int) (int, error)) (int, error) {

	funcX := 0
	somaFuncX := 0

	if len(x) == 0 {
		return 0, fmt.Errorf("Slice vazio ")
	}

	for _, X := range x {
		funcX, _ = y(X)
		somaFuncX += funcX
	}

	return funcX, nil

}
