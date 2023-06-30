package main

import "fmt"

func segundo(inteiros []int) int {
	M := inteiros[0]
	mM := inteiros[1]

	if mM > M {
		M, mM = mM, M
	}

	for i := 2; i < len(inteiros); i++ {
		if inteiros[i] > M {
			mM = M
			M = inteiros[i]
		} else if inteiros[i] > mM && inteiros[i] != M {
			mM = inteiros[i]
		}
	}
	return mM
}
func main() {
	var n = []int{10, 5, 55, 25, 100, 95, 35, 15, 65}
	fmt.Print("O segundo maior valor do slice é: ", segundo(n))
}
