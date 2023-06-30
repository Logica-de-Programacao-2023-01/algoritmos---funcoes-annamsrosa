package main

import "fmt"

func main() {
	sl1 := []int{14, 5, 67, 8}
	sl2 := []int{7, 89, 12, 6, 14}
	fmt.Print(contemnosdoisslices(sl1, sl2))
}
func contemnosdoisslices(sl1 []int, sl2 []int) ([]int, error) {
	sl3 := []int{}
	if len(sl1) == 0 || len(sl2) == 0 {
		return sl3, fmt.Errorf("Erro")
	}
	for _, s1 := range sl1 {
		for _, s2 := range sl2 {
			if s1 == s2 {
				sl3 = append(sl3, s1)
			}
		}
	}
	return sl3, nil
}
