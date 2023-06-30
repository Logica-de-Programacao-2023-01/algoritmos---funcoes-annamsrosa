package main

import (
	"fmt"
	"strings"
)

func vogais(frase string) int {
	vogais := "aeiouAEIOU"
	cont := 0
	for _, i := range frase {
		if strings.ContainsRune(vogais, i) {
			cont++
		}
	}
	return cont
}
func main() {
	s := "Beatriz"
	resto := vogais(s)
	fmt.Print(resto)
}
