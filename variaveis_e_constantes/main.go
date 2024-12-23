package main

import "fmt"

// Constantes
const name string = "Rodrigo"

func main() {
	var text string = "Hello world"
	var isValid bool = false

	// shortet declaration (declaração Curta)
	// valor padrão e tipagem na criação
	texto := "ola"

	fmt.Println(text)
	fmt.Println(texto)
	fmt.Println(isValid)

	fmt.Println(name)
}
