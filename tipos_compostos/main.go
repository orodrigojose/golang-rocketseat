package main

import "fmt"

func main() {
	// Array
	// gavetas_array := [3]string{"copos", "colheres", "tapoer"}

	// gavetas_array[4] = append(gavetas_array, "um sapo?")

	/*
	   //slice é um array mutável no quesito de tamanhno

	   var gavetas []string

	   gavetas = append(gavetas, "copo", "colher", "pratos")

	   // slice[x:x-1]

	   fmt.Println(gavetas[1:3])
	*/

	// it's mean or like the object, but have an bigger difference
	// ["name of peaople"] = your respective value  (in this case it's age)
	pessoas := map[string]int{}

	pessoas["Rodrigo"] = 17
	pessoas["Priscila"] = 18

	if idade, ok := pessoas["Priscila"]; ok {
		fmt.Println("Pessoa existe!", idade, ok)
	} else {
		fmt.Println("Pessoa não existe!")
	}
}
