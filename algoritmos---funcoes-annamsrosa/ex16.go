package main

import (
	"fmt"
	"strings"
)

func main() {
	st1 := "beatriz"
	fmt.Print(v(st1))
}
func v(st1 string) string {

	stringfinal := strings.NewReplacer("a", "*", "e", "*", "i", "*", "o", "*", "u", "*")
	st1 = stringfinal.Replace(st1)
	return st1
}
