package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 { // if deve ter a chave e a expressão não tem parênteses (ao menos quando não tem comparação de expressões)
		fmt.Println("Aprovado com nota ", nota)
	} else {
		fmt.Println("Reprovado com nota ", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
