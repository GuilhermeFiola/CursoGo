package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i // Pegando o endereço de i e coloca no ponteiro p
	*p++   // Desreferenciando (Pegando o valor)
	i++

	//p++ // Go não tem aritmética de ponteiros

	fmt.Println(p, *p, i, &i)
}
