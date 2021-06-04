package main

import (
	"fmt"
	"time"
)

func main() {

	// Bloco com inicialização separada
	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	// Bloco com inicialização de variável
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {
		fmt.Println("Loop infinito")
		time.Sleep(time.Second)
	}
}
