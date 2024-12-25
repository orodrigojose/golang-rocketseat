package meet // Meet is the package global, an box of functions (i think this logic for it)

import "fmt"

var X = 10

// SayHello() globally fn
// sayHello() locally fn
func sayHello() {
	fmt.Println("Hello world")
}
