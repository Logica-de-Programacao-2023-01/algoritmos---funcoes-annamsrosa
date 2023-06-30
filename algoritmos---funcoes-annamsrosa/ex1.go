package main

import "fmt"

func média(sl []int) float64 {
	var média float64 = 0
	for _, i := range sl {
		média += float64(i)
	}
	média /= float64(len(sl))
	return média
}
func main() {
	var slice = []int{5, 10, 15, 20}
	res := média(slice)
	fmt.Print(res)

}
