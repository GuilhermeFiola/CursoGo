package main

import (
	"fmt"
	m "math" // Pode utilizar um alias para o nome do pacote
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // Tipo float64 inferido pelo compilador

	// Forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area, "!")

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "Oi!"

	fmt.Println(g, h, i)
}
