  package main

  import (
	  "fmt"
  )

func main() {
	// defer is executed at the last of the function execution
	defer fmt.Println("Defer one")
	defer fmt.Println("Defer two")
	defer fmt.Println("Defer three")
	fmt.Println("Defer sample text")
	sample()
}

func sample (){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}