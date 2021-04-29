package main

import "fmt"

// Não tem operador ternário
func obterResultado(nota float64) string {

	// Modo em Go apenas com If
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

	// Outras linguagens
	// return nota >= 6 ? "Aprovado" : "Reprovado"
}

func main() {
	fmt.Println(obterResultado(7.0))
}
