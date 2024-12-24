package main

import (
	"fmt"
	"strings"
)

func main() {
	var hello string = "Olá, mundo!"
	var question string = "Como vai?"

	// Concatenação
	meet := hello + question

	fmt.Println(meet)
	fmt.Println(strings.ToUpper(meet))
	fmt.Println(strings.Contains(meet, "mundo"))
}
