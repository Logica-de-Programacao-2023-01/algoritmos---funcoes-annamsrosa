package main

import "fmt"

func mt3(numeros int) int {
	numeros *= 3
	return numeros
}

func att(nums []int, multiplicador int) []int {
	resultado := make([]int, len(nums))

	for i, num := range nums {
		resultado[i] = mt3(num)
	}
	return resultado
}

func main() {
	var valores = []int{10, 5, 55, 25, 100, 95, 35, 15, 65}
	fmt.Println(valores)
	multi := 3
	fmt.Println(att(valores, multi))
}
