package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":  1234.12,
			"Gustavo Moreira": 3210.32,
		},
		"J": {
			"José Abrão": 546.2,
		},
		"P": {
			"Pedro Junior":    4567.81,
			"Priscila Santos": 7894.22,
		},
	}
	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "P")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
