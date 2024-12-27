package main

import "fmt"

func main() {
	var newVar int32 = 2

	fmt.Println("Print Main: ", newVar, &newVar)
	PrintVarByPointer(&newVar)
	PrintAsCopy(newVar)
}

func PrintAsCopy(num int32) {
	fmt.Println("Print Copy: ", num, &num)
}

func PrintVarByPointer(num *int32) {
	fmt.Println("Print pointer: ", *num, &num)
}
