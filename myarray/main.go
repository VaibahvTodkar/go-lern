package main

import "fmt"

func main() {
	fmt.Println("Hello, World in Array!")
	var myArray [5]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40
	myArray[4] = 50
	fmt.Println("Array elements are:", myArray)

	for i := 0; i < len(myArray); i++ {
		fmt.Printf("Element at index %d: %d\n", i, myArray[i])
	}
}
