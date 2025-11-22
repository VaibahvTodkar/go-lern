package main

import "fmt"

func main() {
	fmt.Println("Pointers Example")
	// actual value shoud passed on in pointers used 
	// var ptr *int
	// fmt.Println("Pointer value is:", ptr)  // nil

	myNumber := 42
	var ptr *int = &myNumber
	fmt.Println("Pointer value is:", ptr)   // pointer address
	fmt.Println("Value at Pointer is:", *ptr) //pointer actual value

	*ptr = 100
	fmt.Println("Updated Value at Pointer is:", *ptr)
	fmt.Println("Original variable value is:", myNumber)
}
