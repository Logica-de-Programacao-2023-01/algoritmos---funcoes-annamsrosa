package main

import (
	"fmt"
	"sort"
)

func main() {
	sl1 := []string{"Bruno", "Débora", "Bianca", "Rafael", "Gabriela", "Gabriel", "Anna"}
	fmt.Print(crescente(sl1))
}
func crescente(sl1 []string) (string, error) {
	sort.Strings(sl1)
	if len(sl1) == 0 {
		return "", fmt.Errorf("Erro")
	}
	string1 := ""
	for _, s1 := range sl1 {
		string1 += s1 + " "
	}

	return string1, nil
}
