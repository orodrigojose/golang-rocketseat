package main

import (
	"errors"
	"fmt"
)

/*
func main() {
	nota := 75

	if nota >= 90 {
		fmt.Println("Aprovado!")
	} else if nota >= 70 {
		fmt.Println("Aluno passou, porém com nota mediana!")
	} else {
		fmt.Println("Reprovado")
	}
}
*/

func ifs() {
	/*
	  err := thisAnError()

	  if err != nil {
	    fmt.Println(err.Error())
	  }
	  It's like that we saw at bellow */

	// Only If statement can use the err variable.
	//     var decalration     -   condition
	if err := thisAnError(); err != nil {
		fmt.Println(err.Error())
	}
}

func thisAnError() error {
	return errors.New("[!] It's an possible error!")
}
