package main

import "fmt"

func main() {
	palavra := []string{"boa", "sorte", "amigos"}
	fmt.Println(pMais5(palavra))
}

func pMais5(palavra []string) ([]string, error) {
	pMais5 := []string{}

	if len(palavra) == 0 {
		return pMais5, fmt.Errorf("ERRO ")
	}

	for _, rangervermelho := range palavra {
		if len(rangervermelho) > 5 {
			pMais5 = append(pMais5, rangervermelho)
		}
	}

	return pMais5, nil
}
