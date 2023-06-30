package main

import (
	"fmt"
	"unicode"
)

func main() {
	palavras := []string{"Água", "Terra", "ar", "Fogo"}
	fmt.Print(M(palavras))
}
func M(palavras []string) ([]string, error) {
	if palavras == nil {
		return palavras, fmt.Errorf("Erro ")
	}
	palavrasmaisculas := []string{}

	for i := 0; i < len(palavras); i++ {
		if unicode.IsUpper(rune(palavras[i][0])) {
			palavrasmaisculas = append(palavrasmaisculas, palavras[i])
		}
	}
	return palavrasmaisculas, nil
}
