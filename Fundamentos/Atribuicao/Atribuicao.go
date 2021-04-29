package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // Inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 3 // i = i % 3

	fmt.Println(i)

	x, y := 1, 2 // Múltiplas atribuições

	fmt.Println(x, y)

	x, y = y, x // Troca de valores

	fmt.Println(x, y)
}
