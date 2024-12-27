package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

// this - class/struct
// So i can say this is same "this by slected class name"
func (p Pessoa) Apresentar() {
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}

func main() {
	p1 := Pessoa{Nome: "Rodrigo", Idade: 34}

	p1.Apresentar()
}
