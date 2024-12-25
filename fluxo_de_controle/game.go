package main

import "fmt"

func game() {
	players := map[string]int{
		"rodrigo": 65,
	}

	if value, ok := players["rodrigo"]; ok {
		fmt.Println("pontos: ", value, ok)
	}
}
