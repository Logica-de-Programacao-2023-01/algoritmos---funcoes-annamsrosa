package main

import "fmt"

func concatenação(s []string) string {
	var soma string
	for _, num := range s {
		soma += num
	}
	return soma
}

func main() {
	var ss = []string{"Água", "Terra", "Ar", "Fogo"}
	fmt.Print("A concatenação é : ", concatenação(ss))

}
