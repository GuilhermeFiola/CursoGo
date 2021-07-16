package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":     12333.12,
		"Mariana Silva": 6789.10,
		"Pedro Junior":  1500.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Printf("Nome: %s | Salário: %.2f\n", nome, salario)
	}
}
