package main

import (
	"fmt"
)

func main() {
	fmt.Println("Fuctions from myfunctions package")
	sayHello()
	add(5, 7)
	adderExample(1, 2, 3, 4, 5)
}

func sayHello() {
	fmt.Println("Hello from sayHello function!")
}

func add(a int, b int) {
	fmt.Println("A + B =", a+b)
}


// variadic parameters.
// The function can accept any number of values of that type.
func adderExample(values ...int ) {     // any number of int arguments
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println("Sum:", sum)
}
