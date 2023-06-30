package main

import (
	"errors"
	"fmt"
	"sort"
)

func cres(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("Slice vazio ")
	}

	sortedSlice := make([]int, len(nums))
	copy(sortedSlice, nums)
	sort.Ints(sortedSlice)
	return sortedSlice, nil
}

func main() {
	var numeros = []int{4, 420, 56, 8, 160, 78, 15, 38, 90, 2}

	crescente, err := cres(numeros)

	if err != nil {
		fmt.Print("Erro:", err)
	}
	fmt.Println(numeros)
	fmt.Println(crescente)
}
