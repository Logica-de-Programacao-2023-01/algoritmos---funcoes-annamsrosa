package main

import "fmt"

func main() {
	sl1 := []int{2, 77, 43, 8}
	n := 8
	fmt.Print(iguais(sl1, n))
}
func iguais(sl1 []int, n int) bool {

	for _, sn := range sl1 {
		if sn == n {
			return true
		}
	}
	return false
}
