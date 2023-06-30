package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1438
	fmt.Print(somaalgarismos(n))
}

func somaalgarismos(n int) (int, error) {

	if n < 0 {
		return 0, fmt.Errorf("Erro ")
	}

	numString := strconv.Itoa(n)
	soma := 0

	StringNum := 0

	for i := 0; i < len(numString); i++ {
		StringNum, _ = (strconv.Atoi(string(numString[i])))
		soma += StringNum
	}
	return soma, nil
}
