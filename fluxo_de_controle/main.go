package main

import (
	"fmt"
)

/*  SWITCH CASE
func main() {
	fmt.Println("Quando é sábado?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("É hoje")
	case today + 1:
		fmt.Println("é amanhã")
	case today + 2:
		fmt.Println("é em dois dias")
	default:
		fmt.Println("Ta longe ainda..")
	}
}
*/

func main() {
	sum := 1

	// WHILE
	/*
	   for true or condition {
	     // content
	   }
	*/

	for i := 0; i < 11; i++ {
		fmt.Println(i)
		sum += i
	}

	// fmt.Println(sum)
	pessoas := map[string]string{
		"nome":  "Priscila",
		"idade": "19",
	}

	for key, data := range pessoas {
		fmt.Println(key, ":", data)
	}
}
