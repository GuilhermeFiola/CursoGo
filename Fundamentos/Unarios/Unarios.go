package main

import "fmt"

func main() {
	x := 1
	y := 2

	// Apenas postfix (pós-fixada)
	x++ // x += 1 ou x = x + 1
	y-- // y -= 1 ou y = y - 1

	fmt.Println(x, y)

	//fmt.Println(x == y--) // Não permite misturar operadores com expressão
}

//PEMDAS - Tabela de precedência
// P - Parênteses, E - Expoente, M - Multiplicação, D - Divisão, A - Adição e S - Subração
